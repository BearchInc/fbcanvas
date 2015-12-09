package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
	"github.com/migore/paypal"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func donate(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client, err := newPaypalClient(c)

	if err != nil {
		return
	}

	payment := paypal.Payment{
		Intent: "sale",
		Payer: &    paypal.Payer{
			PaymentMethod: "paypal",
		},
		RedirectURLs: &paypal.RedirectURLs{
			ReturnURL:"https://fb-canvas-dot-staging-api-getunseen.appspot.com/success",
			CancelURL:"https://fb-canvas-dot-staging-api-getunseen.appspot.com/cancel",
		},
		Transactions: []paypal.Transaction{
			paypal.Transaction{
				Amount: &paypal.Amount{
					Currency: "BRL",
					Total: "12.23",
				},
				Description: "Ajudando a Ju Ferrari!",
			},
		},
	}
	paymentResp, err := client.CreatePayment(payment);

	if err != nil {
		log.Infof(c, "Couldn't create payment: %+v", err)
		return
	}

	var approveUrl string
	for _, link := range paymentResp.Links {
		if link.Rel == "approval_url" {
			approveUrl = link.Href
		}
	}

	w.Write([]byte(approveUrl))
}

func successPaypal(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Query().Get("paymentId")
	payerID := r.URL.Query().Get("PayerID")
	c := appengine.NewContext(r)
	client, err := newPaypalClient(c)

	if err != nil {
		return
	}

	_, err = client.ExecutePayment(paymentID, payerID, nil)

	if err != nil {
		log.Infof(c, "Couldn't execute payment: %+v", err)
		return
	}

	w.Write([]byte("Success"))
}

func newPaypalClient(c context.Context) (*paypal.Client, error) {
	clientID := "AUGtRDBDZek5V-TWQZ4GCALZNfRTbObh5UjxVthXScB90X9W3iDrez2VEVZSFG4qFKDfMsnqPmx7tBze"
	secret := "EKLTvvNjEHZHvcrH2vmdMjNBHg4BO_8S4YBr2MFMSCfFFy9rz-TdFvk9lMe595Xd-y1UMJErjudYhiRP"
	client := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)


	client.Client = urlfetch.Client(c)

	tokenResp, err := client.GetAccessToken()

	if err != nil {
		log.Infof(c, "Couldn't create access token: %+v", err)
		return nil, err
	}

	client.Token = tokenResp

	return client, nil
}

func init() {
	http.HandleFunc("/paypal", donate)
	http.HandleFunc("/paypal/success", successPaypal)
	http.HandleFunc("/", handler)
}
