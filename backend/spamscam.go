package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"time"
)

type Webhistory struct {
	Timestamp time.Time `json:"timestamp"`
	Duration  float64   `json:"duration"`
	Data      struct {
		URL       string `json:"url"`

	} `json:"data"`
}


func Requestperson() {
	url := "https://randomuser.me/api/"

  res, _ := http.Get(url)

  body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

  //var data []Webhistory
	var data interface{}

  json.Unmarshal([]byte(body), &data)
  //fmt.Println(data[0])
	fmt.Printf("%+v\n", (data))

	//fmt.Println(data)
}

	//fmt.Println(data)
func Requestcurrency() {
	url := "https://api.currencyfreaks.com/latest?apikey=4f7765e522f84a37ab63672b887ed45e"

  res, _ := http.Get(url)

  body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

  //var data []Webhistory
	var data interface{}

  json.Unmarshal([]byte(body), &data)
  fmt.Println(data)
	//fmt.Println(data)
}


func Requestweb() {
	url := "http://localhost:5600/api/0/buckets/aw-watcher-web-chrome/events?limit=1"

  res, err := http.Get(url)
  if err != nil {
      fmt.Println(err)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      fmt.Println(err)
  }

  var data []Webhistory

  err = json.Unmarshal(body, &data)
  if err != nil {
      panic(err)
  }

  for _, values := range data {
		datetime := values.Timestamp.Format("Mon, 02 Jan 2006 15:04:05") 

		//fmt.Println(values.Timestamp.String)
		fmt.Println(`Hello, My name is TITLE FIRSTNAME LASTNAME and I am located
								in CITY, COUNTRY. My company is COMPANY. On ` + datetime +
								`you visited the url` + values.Data.URL +  `. You must send 35
								dollars via this link: example.com within 7 days of opening
								this email. If you do not I will tell your friends and family
								that you visited the site. Be kind, TITLE FIRSTNAME LASTNAME`)

	}
	//fmt.Println(data)
}
