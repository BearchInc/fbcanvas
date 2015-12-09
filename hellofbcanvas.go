package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
	"github.com/migore/paypal"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func paypalWebhook(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var objmap map[string]*json.RawMessage
	err := decoder.Decode(&objmap)
	if err != nil {
		log.Infof(c, "Error: %+v", err)
		w.Write([]byte("Error!"))
		return
	}

	log.Infof(c, "Map keys: %+v", objmap)

	var resource json.RawMessage
	err = json.Unmarshal(*objmap["resource"], &resource)

	log.Infof(c, "resources: %+v", string(resource))
	w.Write([]byte(string(resource)))
}

func donate(w http.ResponseWriter, r *http.Request) {
	clientID := "AUGtRDBDZek5V-TWQZ4GCALZNfRTbObh5UjxVthXScB90X9W3iDrez2VEVZSFG4qFKDfMsnqPmx7tBze"
	secret := "EKLTvvNjEHZHvcrH2vmdMjNBHg4BO_8S4YBr2MFMSCfFFy9rz-TdFvk9lMe595Xd-y1UMJErjudYhiRP"
	client := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)

	c := appengine.NewContext(r)
	client.Client = urlfetch.Client(c)

	tokenResp, err := client.GetAccessToken()

	if err != nil {
		log.Infof(c, "Couldn't get access token: %+v", err)
		return
	}

	client.Token = tokenResp

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

	clientID := "AUGtRDBDZek5V-TWQZ4GCALZNfRTbObh5UjxVthXScB90X9W3iDrez2VEVZSFG4qFKDfMsnqPmx7tBze"
	secret := "EKLTvvNjEHZHvcrH2vmdMjNBHg4BO_8S4YBr2MFMSCfFFy9rz-TdFvk9lMe595Xd-y1UMJErjudYhiRP"
	client := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)

	c := appengine.NewContext(r)
	client.Client = urlfetch.Client(c)

	tokenResp, err := client.GetAccessToken()

	if err != nil {
		log.Infof(c, "Couldn't create payment: %+v", err)
		return
	}

	client.Token = tokenResp

	_, err = client.ExecutePayment(paymentID, payerID, nil)

	if err != nil {
		log.Infof(c, "Couldn't execute payment: %+v", err)
		return
	}

	w.Write([]byte("Success"))
}

func init() {
	http.HandleFunc("/webhook", paypalWebhook)
	http.HandleFunc("/paypal", donate)
	http.HandleFunc("/paypal/success", successPaypal)
	http.HandleFunc("/", handler)
}
