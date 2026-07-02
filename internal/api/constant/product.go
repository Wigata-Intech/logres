package constant

type ProductStatus uint8

const (
	ProductStatusDraft ProductStatus = iota
	ProductStatusPublished
	ProductStatusArchived
)
