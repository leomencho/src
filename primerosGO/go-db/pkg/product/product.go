package product

import "time"

//modelo de prodcuto
type Model struct {
	ID          uint
	Name        string
	Obsevations string
	Price       int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
