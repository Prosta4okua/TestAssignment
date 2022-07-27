package swg

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
)

func Rate(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(w, GetBtcToUah())
}

type ResponseBTC2UAH struct {
	Motd struct {
		Msg string `json:"msg"`
		Url string `json:"url"`
	} `json:"motd"`
	Success bool `json:"success"`
	Query   struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	} `json:"query"`
	Info struct {
		Rate float64 `json:"rate"`
	} `json:"info"`
	Historical bool    `json:"historical"`
	Date       string  `json:"date"`
	Result     float64 `json:"result"`
}

type Rates struct {
	BTC float64
}

func GetBtcToUah() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Wow, error happened!")
		return ""
	}
	btcToUahString := fmt.Sprintf("%f", makeRequest())
	return btcToUahString
}

func makeRequest() float64 {
	var newEvent ResponseBTC2UAH
	url := "http://api.exchangerate.host/convert?from=BTC&to=UAH"

	body := requestURL(url)
	err := json.Unmarshal(body, &newEvent)
	if err != nil {
		panic(err)
		return 0
	}

	return newEvent.Info.Rate

}

func requestURL(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)

	if err != nil {
		fmt.Print("Помилка" + err.Error())
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	return body
}
