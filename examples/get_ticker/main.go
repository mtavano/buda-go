package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blue-factory/cryptobot/internal/httpclient"
	"github.com/blue-factory/cryptobot/pkg/exchange"
	"github.com/blue-factory/cryptobot/pkg/exchange/buda"
)

func main() {
	//

	key := os.Getenv("BUDA_API_KEY")
	secret := os.Getenv("BUDA_API_KEY_SECRET")

	log.Println("BUDA_API_KEY", key)
	log.Println("BUDA_API_KEY_SECRET", secret)

	b := buda.New(key, secret, &httpclient.HTTPClient{
		Client:   &http.Client{},
		MaxRetry: 8,
		Delay:    1,
		Validate: validate,
	})

	// lt, t, _ := b.GetTickers()
	// fmt.Printf("%+v\n", t)
	bbs, _ := b.GetBalances()
	fmt.Println(bbs)

	or := &exchange.Order{
		Pair: exchange.PairBTCCLP,

		Side:   "sell",
		Type:   "limit",
		Amount: 0.001,
		Price:  20000001,
	}
	o, err := b.CreateOrder(or)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">> order created \n%+v\n", o)

	oo, err := b.GetOrder(o.ID)

	if err != nil {
		log.Fatal(err)
		log.Fatal(err.Error())
	}
	fmt.Printf(">> order fetched \n%+v\n", oo)
	time.Sleep(time.Second * 5)

	oc, err := b.CancelOrder(oo.ID)
	if err != nil {
		log.Fatal(err)
		log.Fatal(err.Error())
	}
	fmt.Printf(">> order canceled \n%+v\n", oc)

}

func validate(res *http.Response) (*http.Response, error) {
	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return res, nil
	case http.StatusTooManyRequests:
		return res, errors.New("Max rate limit")
	case http.StatusForbidden:
		return res, errors.New("Forbidden request")
	case http.StatusUnauthorized:
		return res, errors.New("Unauthorized request")
	case http.StatusNotFound:
		return res, errors.New("Not Found")
	default:
		return res, fmt.Errorf("Unknown http status code: %d", res.StatusCode)
	}
}
