package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mtavano/buda-go"
)

func main() {
	key := os.Getenv("BUDA_API_KEY")
	secret := os.Getenv("BUDA_API_KEY_SECRET")

	b := buda.New(key, secret, &http.Client{})
	trades, err := b.GetTrades(buda.PairETHCLP)
	check(err)

	fmt.Println(trades)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
