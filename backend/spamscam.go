package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"time"
)

type scammer struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City        string `json:"city"`
			State       string `json:"state"`
			Country     string `json:"country"`
			Postcode    int    `json:"postcode"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
			Timezone struct {
				Offset      string `json:"offset"`
				Description string `json:"description"`
			} `json:"timezone"`
		} `json:"location"`
		Email string `json:"email"`
		Login struct {
			UUID     string `json:"uuid"`
			Username string `json:"username"`
			Password string `json:"password"`
			Salt     string `json:"salt"`
			Md5      string `json:"md5"`
			Sha1     string `json:"sha1"`
			Sha256   string `json:"sha256"`
		} `json:"login"`
		Dob struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"dob"`
		Registered struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"registered"`
		Phone string `json:"phone"`
		Cell  string `json:"cell"`
		ID    struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"id"`
		Picture struct {
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Thumbnail string `json:"thumbnail"`
		} `json:"picture"`
		Nat string `json:"nat"`
	} `json:"results"`
	Info struct {
		Seed    string `json:"seed"`
		Results int    `json:"results"`
		Page    int    `json:"page"`
		Version string `json:"version"`
	} `json:"info"`
}

type Webhistory struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Duration  float64   `json:"duration"`
	Data      struct {
		URL       string `json:"url"`
		Title     string `json:"title"`
		Audible   bool   `json:"audible"`
		Incognito bool   `json:"incognito"`
		TabCount  int    `json:"tabCount"`
	} `json:"data"`
}


type Data struct {
Company string `json:"company"`
Currency string `json:"currency"`
}

func Requestscammer() {
	url := "https://randomuser.me/api/"

  res, err := http.Get(url)
  if err != nil {
      fmt.Println(err)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      fmt.Println(err)
  }

	var scam scammer
	json.Unmarshal([]byte(body), &scam)
	fmt.Printf(scam)

	//fmt.Println(data)
}


func Requestbusiness() {
	url := "http://localhost:8000/data"

  res, err := http.Get(url)
  if err != nil {
      fmt.Println(err)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      fmt.Println(err)
  }

	var company Data
	json.Unmarshal([]byte(body), &company)
	fmt.Printf(company.Company)
	fmt.Printf(company.Currency)

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
		fmt.Println(values.Duration)
   	fmt.Println(values.Timestamp)
		fmt.Println(values.Data.URL)
	}
	//fmt.Println(data)
}
