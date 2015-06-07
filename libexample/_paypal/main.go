package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"github.com/k0kubun/pp"
	"github.com/unrolled/render"
	"github.com/bitly/go-simplejson"
	"bytes"
	"io/ioutil"
)

var authToken = ""

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	ctx := context.Background()
	ctx = NewContext(ctx)

	ctx = WithAuthCallbackFunc(ctx, func(ctx context.Context, w http.ResponseWriter, r *http.Request, token *oauth2.Token) {
		fmt.Println(token.AccessToken)
	})
	ctx = WithAuthErrorFunc(ctx, func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	})

	kami.Context = ctx

	kami.Get("/login/paypal", LoginPayPal)
	kami.Get("/auth/paypal/callback", AuthPayPalCallback)
	kami.Get("/payment/create", PaymentCreate)
	kami.Get("/payment/done", PaymentDone)
	kami.Get("/paypal/payment/execute", PayPalPaymentExecute)
	kami.Get("/paypal/payment/cancel", PayPalPaymentCancel)

	log.Println("Starting server...")
	log.Println("GOMAXPROCS: ", cpus)
	kami.Serve()
}

var createData = `
{
  "intent":"sale",
  "redirect_urls":{
    "return_url":"http://localhost:8000/paypal/payment/execute",
    "cancel_url":"http://localhost:8000//paypal/payment/cancel"
  },
  "payer":{
    "payment_method":"paypal"
  },
  "transactions":[
    {
      "amount":{
        "total":"9.99",
        "currency":"USD"
      },
      "description":"This is the payment transaction description."
    }
  ]
}
`

var renderer = render.New()

func PaymentCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	buf := bytes.NewBufferString(createData)
	req, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/payments/payment", buf)
	if err != nil {
		log.Fatalln(err)
		renderer.JSON(w, 400, err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Println("Bearer %s", authToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		renderer.JSON(w, 400, err.Error())
		return
	}
	defer res.Body.Close()

	sj, err := simplejson.NewFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
		renderer.JSON(w, 400, err.Error())
		return
	}

	pp.Println(sj)

	approvalURL := ""
	for _, link := range sj.Get("links").MustArray() {
		l := link.(map[string]interface{})
		if l["rel"].(string) != "approval_url" {
			continue
		}

		approvalURL = l["href"].(string)
		break
	}

	fmt.Println("approvalURL", approvalURL)

	http.Redirect(w, r, approvalURL, 302)
}

func PayPalPaymentExecute(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pp.Println(r.Header)
	pp.Println(r.Cookies())

	buf := bytes.NewBufferString(fmt.Sprintf("{ \"payer_id\" : \"%s\"}", r.FormValue("PayerID")))
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.sandbox.paypal.com/v1/payments/payment/%s/execute/", r.FormValue("paymentId")), buf)
	if err != nil {
		log.Fatalln(err)
		renderer.JSON(w, 400, err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Println("Bearer %s", authToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		renderer.JSON(w, 400, err.Error())
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		renderer.JSON(w, 400, err.Error())
		return
	}

	fmt.Println(string(data))

	http.Redirect(w, r, "/payment/done", 302)
}

func PayPalPaymentCancel(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pp.Println(r.Form)
	pp.Println(r.Header)
	pp.Println(r.Cookies())
}

func PaymentDone(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pp.Println(r.Form)
	pp.Println(r.Header)
	pp.Println(r.Cookies())

	w.Write([]byte("OK"))
}