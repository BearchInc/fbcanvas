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
    "google.golang.org/appengine/mail"
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
	config := NewConfig(c)
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
		ExperienceProfileID: config.ExperienceProfileId,
		Payer: &paypal.Payer{
			PaymentMethod: "paypal",
		},
		RedirectURLs: &paypal.RedirectURLs{
			ReturnURL: config.BaseURL + "/paypal/success",
			CancelURL: config.BaseURL,
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
    w.Write([]byte(approveUrl))
//	http.Redirect(w, r, approveUrl, 200)
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
	config := NewConfig(c)
    http.Redirect(w, r, config.BaseURL + "/thanks?n=" + name + "&v=" + amount, 301)
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
    msg := &mail.Message{
        Sender:  "ju.ferrari.doar@gmail.com",
        To:      []string{email},
        Subject: "Agradecimento",
        Body:    "Muito obrigada, farás muitos rostinhos sorrirem!",
    }
    if err := mail.Send(c, msg); err != nil {
        log.Errorf(c, "Couldn't send email to %s: %v", email, err)
    } else {
        log.Infof(c, "Sent email to: %s", email, err)
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
	config := NewConfig(c)
	clientID := config.PaypalClientId
	secret := config.PaypalSecret
	client := paypal.NewClient(clientID, secret, config.PaypalBase)

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
