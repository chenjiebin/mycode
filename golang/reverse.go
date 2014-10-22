package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	httpPost()
}

func httpPost() {
	resp, err := http.PostForm("http://www.xmsmjk.com/UrpOnline/Home/DoReservation",
		url.Values{
			"ORG_ID":      {"Value"},
			"SCHEDULE_ID": {"123"},
			"NUMBER_ID":   {"11"},
			"DEPT_CODE":   {"123"},
			"DOCTOR_CODE": {"123"},
			"AM_PM":       {"123"},
			"START_TIME":  {"123"},
			"CARD_NO":     {"D64272480"},
			"ID_CARD":     {"350825198907214531"},
			"NAME":        {"%E6%B1%9F%E5%BB%BA%E6%9D%BE"},
			"TELEPHONE":   {"123"}})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
