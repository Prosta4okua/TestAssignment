package swg

import (
	"crypto/tls"
	"fmt"
	gmail "github.com/go-mail/mail"
	"net/http"
	"os"
	"strconv"
)

func SendEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	CsvShow()

	m := gmail.NewMessage()
	m.SetHeader("From", "currency@currency.com")
	m.SetHeader("To", CsvRead()...)
	m.SetHeader("Subject", "Сьогоднішній курс біткоїна до гривні")
	m.SetBody("text/plain", "Поточний курс біткоїна:"+strconv.Itoa(Get_BTC_to_UAH())+"грн")

	d := gmail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	_, err := w.Write([]byte("E-mailʼи відправлено"))
	if err != nil {
		return
	}
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	fmt.Println("Your email: " + r.FormValue("email"))
	newEmail := r.FormValue("email")

	if CheckCsvValue(newEmail) == false {
		CsvAdd(newEmail)
	} else {
		http.Error(w, "Ця електронна скринька вже наявна в нашій файловій базі даних!", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("E-mail додано"))
	if err != nil {
		return
	}
}
