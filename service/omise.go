package omiseSevice

import (
	"log"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

const (
	// Read these from environment variables or configuration files!
	OmisePublicKey = "pkey_test_5xurq3wypybxwpbp4nq"
	OmiseSecretKey = "skey_test_5xurq3xvsiikzgmthaq"
)

const (
	PROMPTPAY = "promptpay"
	PAYMENT_TESCO_LOTUS = "bill_payment_tesco_lotus"
)

type OmiseService struct {

}

func CreateSource(amount int, currency string, paymentType string) *omise.Source {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Source{}
	log.Println(amount)
	log.Println(currency)
	log.Println(paymentType)
	client.Do(result, &operations.CreateSource{
		Amount:   int64(amount) * 100,
		Currency: currency,
		Type:     paymentType,
	})
	return result
}

func CreateToken() *omise.Card {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Card{}

	err := client.Do(result, &operations.CreateToken{
		Name:            "Somchai Prasert",
		Number:          "4242424242424242",
		ExpirationMonth: 10,
		ExpirationYear:  2025,

		City:         "Bangkok",
		PostalCode:   "10320",
		SecurityCode: "123",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	return result
}

func CreateChargeBySource(source string, amount int, currency string) *omise.Charge {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Charge{}
	err := client.Do(result, &operations.CreateCharge{
		Amount:   int64(amount),
		Currency: currency,
		Source:  source,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	return result
}

func CreateChargeByToken(tokenId string, amount int, currency string) *omise.Charge {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Charge{}
	err := client.Do(result, &operations.CreateCharge{
		Amount:   int64(amount),
		Currency: currency,
		Card: tokenId,
		ReturnURI: "http://localhost:8080/sync",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	return result
}

func GetChargeById(changeId string) *omise.Charge {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Charge{}
	err := client.Do(result, &operations.RetrieveCharge{
		ChargeID: changeId,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	return result
}

func GetChargeList() *omise.ChargeList {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	
	result := &omise.ChargeList{}
	err := client.Do(result, &operations.ListCharges{
		operations.List{
			Limit: 100,
			From:  time.Now().Add(-1 * time.Hour),
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	return result
}