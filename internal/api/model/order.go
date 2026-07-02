package model

import (
	"time"

	"github.com/wigata-intech/logres/internal/api/constant"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID            uuid.UUID            `json:"id" db:"id"`
	DocumentID    uuid.UUID            `json:"document_id" db:"document_id"`
	PaymentID     uuid.NullUUID        `json:"payment_id" db:"payment_id"`
	FullName      string               `json:"full_name" db:"full_name"`
	Email         string               `json:"email" db:"email"`
	PhoneNumber   string               `json:"phone_number" db:"phone_number"`
	TotalPrice    decimal.Decimal      `json:"total_price" db:"total_price"`
	TotalDiscount decimal.Decimal      `json:"total_discount" db:"total_discount"`
	TotalTax      decimal.Decimal      `json:"total_tax" db:"total_tax"`
	BilledAmount  decimal.Decimal      `json:"billed_amount" db:"billed_amount"`
	Status        constant.OrderStatus `json:"status" db:"status"`
	CreatedAt     time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time           `json:"deleted_at,omitempty" db:"deleted_at"`
}

type OrderItem struct {
	OrderID    uuid.UUID       `json:"order_id" db:"order_id"`
	ProductID  uuid.UUID       `json:"product_id" db:"product_id"`
	Quantity   uint16          `json:"quantity" db:"quantity"`
	Price      decimal.Decimal `json:"price" db:"price"`
	TotalPrice decimal.Decimal `json:"total_price" db:"total_price"`
	Order      uint8           `json:"order" db:"order"`
}

func NewOrder(
	fullName, email, phoneNumber string,
	totalPrice, totalDiscount, totalTax, billedAmount decimal.Decimal,
) (*Order, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Order{
		ID:            id,
		FullName:      fullName,
		Email:         email,
		PhoneNumber:   phoneNumber,
		TotalPrice:    totalPrice,
		TotalDiscount: totalDiscount,
		TotalTax:      totalTax,
		BilledAmount:  billedAmount,
		Status:        constant.OrderStatusWaitingPayment,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// TODO: Add FSM for order status transitions
func (o *Order) UpdateStatus(status constant.OrderStatus) {
	o.Status = status
	o.UpdatedAt = time.Now()
}

func NewOrderItem(
	orderID uuid.UUID,
	product *Product,
	quantity uint16,
	order uint8,
) (*OrderItem, error) {
	totalPrice := product.Price.Mul(
		decimal.NewFromInt(
			int64(quantity),
		),
	)

	return &OrderItem{
		OrderID:    orderID,
		ProductID:  product.ID,
		Quantity:   quantity,
		Price:      product.Price,
		TotalPrice: totalPrice,
		Order:      order,
	}, nil
}
