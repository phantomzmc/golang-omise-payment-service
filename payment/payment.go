package payment

import (
	"github.com/biskitsx/go-fiber-api/omise"
	"github.com/biskitsx/go-fiber-api/order"
)



type Payment struct {
	ID int `json:"id"`
	TransactionID int `json:"transactionId"`
	Order order.Order `json:"order"`
}

func (p *Payment) GetPayment() Payment {
	var order  = order.Order{
		ID: 1,
		Name: "Product 1",
		Price: 10.99,
	}
	return Payment{
		ID: 1,
		TransactionID: 1,
		Order: order,
	}
}

func (p *Payment) GetPaymentById() interface{} {
	charge := omise.GetChargeList()
	return charge
}

func (p *Payment) CreatePayment() interface{} {
	source := omise.CreateSource()
	charge := omise.CreateChargeBySource(source.Base.ID)
	return charge
}

func (p *Payment) UpdatePayment() {

}

func (p *Payment) DeletePayment() {

}
