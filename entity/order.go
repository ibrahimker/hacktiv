package entity

import "time"

type Order struct {
	ID           int       `json:"order_id"`
	CustomerName string    `json:"customerName`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ID          int    `json:"item_id"`
	Code        string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}
