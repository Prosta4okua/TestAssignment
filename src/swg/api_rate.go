package swg

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var BtcToUah int = 0

func Rate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(w, strconv.Itoa(Get_BTC_to_UAH()))
}

type ResponseUSD2UAH struct {
	Success bool    `json:"success"`
	Date    string  `json:"date"`
	Result  float64 `json:"result"`
}

type ResponseBTC2USD struct {
	Success bool  `json:"success"`
	Rates   Rates `json:"rates"`
}

type Rates struct {
	BTC float64
}

func Get_BTC_to_UAH() int {
	if BtcToUah <= 0 {
		btcToUsd := BTS_to_USD().Rates.BTC
		usdToUah := USD_to_UAH().Result
		BtcToUah = int(btcToUsd * usdToUah)
	}
	return BtcToUah
}

func USD_to_UAH() ResponseUSD2UAH {
	var newEvent ResponseUSD2UAH

	url := "https://api.apilayer.com/fixer/convert?to=UAH&from=USD&amount=1"

	err := godotenv.Load(".env")
	if err != nil {
		return newEvent
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("USDtoUAH"))

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &newEvent)

	return newEvent

}

func BTS_to_USD() ResponseBTC2USD {
	var newEvent ResponseBTC2USD

	url := "http://api.coinlayer.com/api/live?access_key=" + os.Getenv("BTCtoUSD") + "&symbols=BTC"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &newEvent)

	return newEvent

}
