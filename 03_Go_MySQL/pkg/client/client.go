package client

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint
	Name      string
	CreateAt  time.Time
	UpadateAt time.Time
}
