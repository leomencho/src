package invoiceitem

import "time"

type Model struct {
	ID              uint
	invoiceheaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdateAt        time.Time
}
