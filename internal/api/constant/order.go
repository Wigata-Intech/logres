package constant

type OrderStatus string

const (
	OrderStatusWaitingPayment   OrderStatus = "WAITING_PAYMENT"
	OrderStatusPaymentConfirmed OrderStatus = "PAYMENT_CONFIRMED"
	OrderStatusShipped          OrderStatus = "SHIPPED"
	OrderStatusOnDelivery       OrderStatus = "ON_DELIVERY"
	OrderStatusDelivered        OrderStatus = "DELIVERED"
	OrderStatusConfirmed        OrderStatus = "CONFIRMED"
	OrderStatusExpired          OrderStatus = "EXPIRED"
	OrderStatusCancelled        OrderStatus = "CANCELLED"
)
