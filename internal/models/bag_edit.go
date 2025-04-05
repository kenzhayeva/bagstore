package models

type BagEdit struct {
	Title    string  `json:"title" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Color    string  `json:"color" binding:"required"`
	Price    float64 `json:"price" binding:"required,gte=0"`
	Size     string  `json:"size" binding:"required"`
}
