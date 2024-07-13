package payment

import (
	"log"

	"github.com/phantomzmc/omise-go-payment-service/order"
	omiseSevice "github.com/phantomzmc/omise-go-payment-service/service"
)

type RequestPayment struct {
	Amount int `json:"amount"`
	Currency string `json:"currency"`
	PaymentType string `json:"paymentType"`
}

type RequestChargeBySource struct {
	SourceId string `json:"sourceId"`
	Amount int `json:"amount"`
	Currency string `json:"currency"`
}

type Payment struct {
	ID int `json:"id"`
	TransactionID int `json:"transactionId"`
	Order order.Order `json:"order"`
}

type ChargeBySource struct {
	Status string `json:"status"`
	Amount int `json:"amount"`
	Currency string `json:"currency"`
	Source Source `json:"source"`
}

type Source struct {
	Type string `json:"type"`
	DownloadUrl string `json:"dowloadUrl"`
	FileName string `json:"fileName"`
	Id string `json:"id"`
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

func (p *Payment) GetPaymentById(chargeId string) interface{} {
	charge := omiseSevice.GetChargeById(chargeId)
	return charge
}

func (p *Payment) GetPaymentList() interface{} {
	charge := omiseSevice.GetChargeList()
	return charge
}

func (p *Payment) CreatePayment(req *RequestPayment) interface{} {
	source := omiseSevice.CreateSource(req.Amount, req.Currency, req.PaymentType)
	log.Println(source)
	charge := omiseSevice.CreateChargeBySource(source.Base.ID, req.Amount, req.Currency)
	return charge
}

func (p *Payment) ChargeBySource(req *RequestChargeBySource) ChargeBySource {
	log.Println(req.SourceId)
	charge := omiseSevice.CreateChargeBySource(req.SourceId, req.Amount, req.Currency)
	result := ChargeBySource{
		Status: string(charge.Status),
		Amount: int(charge.Amount),
		Currency: charge.Currency,
		Source: Source{
			Type: charge.Source.Type,
			DownloadUrl: charge.Source.ScannableCode.Image.DownloadURI,
			FileName: charge.Source.ScannableCode.Image.Filename,
			Id: charge.Source.ID,
		},
	}
	return result
}

func (p *Payment) UpdatePayment() {

}

func (p *Payment) DeletePayment() {

}
