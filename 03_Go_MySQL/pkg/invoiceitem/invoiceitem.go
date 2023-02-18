package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID        uint
	ClienteID uint
	ProductID uint
	CreateAt  time.Time
	UpadateAt time.Time
}
