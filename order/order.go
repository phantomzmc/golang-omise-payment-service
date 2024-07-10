package order

type Order struct {
    ID     int     `json:"id"`
    Name   string  `json:"name"`
    Price  float64 `json:"price"`
}