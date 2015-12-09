package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"github.com/migore/paypal"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"github.com/mjibson/goon"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func donate(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client, err := newPaypalClient(c)

	if err != nil {
		return // error was logged inside newPaypalClient function
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
				ItemList: &paypal.ItemList{
					Items: []paypal.Item{
						paypal.Item{
							Quantity: 3,
							Name: "Boneco da Peppa",
							Price: "10.10",
							Currency: "BRL",
							SKU: "1",
						},
						paypal.Item{
							Quantity: 1,
							Name: "Cachorro quente",
							Price: "120.20",
							Currency: "BRL",
							SKU: "2",
						},
					},
				},
				Amount: &paypal.Amount{
					Currency: "BRL",
					Total: "150.50",
				},
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

	http.Redirect(w, r, approveUrl, http.StatusFound)
}

type Item struct {
	Id string `datastore:"-" goon:"id"`
	Quantity int
}

func successPaypal(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client, err := newPaypalClient(c)

	if err != nil {
		return // error was logged inside newPaypalClient function
	}

	paymentID := r.URL.Query().Get("paymentId")
	payerID := r.URL.Query().Get("PayerID")
	executePaymentResponse, err := client.ExecutePayment(paymentID, payerID, nil)

	g := goon.NewGoon(r)
	for _, i := range executePaymentResponse.Transactions[0].ItemList.Items {
		item := &Item{ Id: i.SKU }
		_ = g.Get(item)

		item.Quantity += i.Quantity

		g.Put(item)
	}

	if err != nil {
		log.Infof(c, "Couldn't execute payment: %+v", err)
		return
	}

	w.Write([]byte("Successfully payed!"))
}

func data(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	productsBytes, err := json.Marshal(products)
	if err != nil {
		log.Errorf(c, "Couldn't execute payment: %+v", err)
		return
	}
	w.Write(productsBytes)
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
	http.HandleFunc("/items", data)
	http.HandleFunc("/paypal", donate)
	http.HandleFunc("/paypal/success", successPaypal)
	http.HandleFunc("/", handler)
}