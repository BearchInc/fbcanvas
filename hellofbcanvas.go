package main

import (
	"encoding/json"
	"fmt"
	"github.com/migore/paypal"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"net/smtp"
    "html/template"
    "strconv"
    "net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
	// Set up authentication information.
}

type CartItem struct {
	Product  *Product
	Quantity int
}

func (this CartItem) toPaypalItem() *paypal.Item {
	return &paypal.Item{
		Quantity: this.Quantity,
		Name:     this.Product.Name,
		Price:    this.Product.Price,
		Currency: "BRL",
		SKU:      this.Product.Id,
	}
}

type ShoppingCart struct {
	Items []CartItem
}

func (this ShoppingCart) toPaypalItemList() []paypal.Item {
	var items []paypal.Item
	for _, item := range this.Items {
		items = append(items, *item.toPaypalItem())
	}
	return items
}

func (this ShoppingCart) total() string {
	total := float64(0)
	for _, item := range this.Items {
		//handle this error
		price, _ := strconv.ParseFloat(item.Product.Price, 64)
		total += price * float64(item.Quantity)
	}

	return strconv.FormatFloat(total, 'f', 2, 64)
}

func donate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	c := appengine.NewContext(r)
	client, err := newPaypalClient(c)

	if err != nil {
		return // error was logged inside newPaypalClient function
	}

	shoppingCart := []CartItem{}

	for i := 1; i <= len(products); i++ {
		itemQuantity := r.FormValue(fmt.Sprintf("quantity-%d", i))
		log.Infof(c, "Item %d quantity %s", i, itemQuantity)

		//handle this error
		integerQuantity, _ := strconv.Atoi(itemQuantity)
		if integerQuantity > 0 {
			cartItem := CartItem{
				Quantity: integerQuantity,
				Product:  getProduct(strconv.Itoa(i)),
			}
			shoppingCart = append(shoppingCart, cartItem)
		}
	}

	cart := ShoppingCart{shoppingCart}
	total := cart.total()

	log.Infof(c, "Total: %s", total)

	payment := paypal.Payment{
		Intent:              "sale",
		ExperienceProfileID: "XP-3L6B-V2T3-RGFL-2JBZ",
		Payer: &paypal.Payer{
			PaymentMethod: "paypal",
		},
		RedirectURLs: &paypal.RedirectURLs{
			ReturnURL: "https://fb-canvas-dot-staging-api-getunseen.appspot.com/paypal/success",
			CancelURL: "https://apps.facebook.com/bearchcanvas/",
		},
		Transactions: []paypal.Transaction{
			paypal.Transaction{
				ItemList: &paypal.ItemList{
					Items: cart.toPaypalItemList(),
				},
				Amount: &paypal.Amount{
					Currency: "BRL",
					Total:    total,
				},
			},
		},
	}

	paymentResp, err := client.CreatePayment(payment)

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
	Id       string `datastore:"-" goon:"id"`
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
		item := &Item{Id: i.SKU}
		_ = g.Get(item)

		item.Quantity += i.Quantity

		g.Put(item)
	}

	if err != nil {
		log.Infof(c, "Couldn't execute payment: %+v", err)
		return
	}

	sendEmail(executePaymentResponse.Payer.PayerInfo.Email, c)

    name := url.QueryEscape(executePaymentResponse.Payer.PayerInfo.FirstName)
    amount := url.QueryEscape(executePaymentResponse.Transactions[0].Amount.Total)
    http.Redirect(w, r, "http://apps.facebook.com/bearchcanvas/thanks?n="+name+"&v="+amount, 301)

}


func thanks(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    p := struct {
        Name string
        Amount string
    }{
        Name: r.URL.Query().Get("n"),
        Amount: r.URL.Query().Get("v"),
    }

    t, err := template.ParseFiles("templates/thanks.html")
    if err != nil {
        log.Infof(c, "Couldn't say thanks: %+v", err)
        return
    } else {
        t.Execute(w, p)
    }
}


func sendEmail(email string, c context.Context) {

	auth := smtp.PlainAuth(
		"",
		"ju.ferrari.doar.com",
		"ygorbruxel",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		[]string{email},
		[]byte("This is the email body."),
	)

	if err != nil {
		log.Infof(c, "Couldn't e-mail to %s: %+v", email, err)
	} else {
		log.Infof(c, "E-mail sent to %s", email)
	}
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
    http.HandleFunc("/thanks", thanks)
	http.HandleFunc("/paypal/success", successPaypal)
	http.HandleFunc("/", handler)
}
