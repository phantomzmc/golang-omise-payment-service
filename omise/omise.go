package omise

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


func CreateSource() *omise.Source {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Source{}

	client.Do(result, &operations.CreateSource{
		Amount:   100000,
		Currency: "thb",
		Type:     "bill_payment_tesco_lotus",
	})
	return result
}

func CreateChargeBySource(source string) *omise.Charge {
	client, _ := omise.NewClient(
		OmisePublicKey,
		OmiseSecretKey,
	)
	result := &omise.Charge{}
	err := client.Do(result, &operations.CreateCharge{
		Amount:   100000,
		Currency: "thb",
		Source:  source,
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