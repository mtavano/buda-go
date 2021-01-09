package balance

import (
	"log"
	"os"
	"testing"

	"github.com/blue-factory/cryptobot/pkg/exchange/buda"
)

func GetBudaBalance(t *testing.T) {
	testName := "BudaBalance"
	log.Println("balance")

	budaAPIKey := os.Getenv("BUDA_API_KEY")
	if budaAPIKey == "" {
		log.Fatalln("missing env variable BUDA_API_KEY")
	}
	budaAPIkeySecret := os.Getenv("BUDA_API_KEY_SECRET")
	if budaAPIkeySecret == "" {
		log.Fatalln("missing env variable BUDA_API_KEY_SECRET")
	}

	secondary := buda.New(budaAPIKey, budaAPIkeySecret)

	balance, err := secondary.GetBalances()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(testName, "balance", balance)

	for _, b := range balance {
		log.Println(b)
	}
}
