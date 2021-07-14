package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


type RequestBody struct {
	Text string `json:"text"`
	Language string `json:"language"`
}

type ResponseBody struct {
	D struct {
		Result []struct {
			Level   int    `json:"Level"`
			Label   string `json:"Label"`
			Value   string `json:"Value"`
			Matches []struct {
				Index       int `json:"Index"`
				Len         int `json:"Len"`
				SearchIndex int `json:"SearchIndex"`
			} `json:"Matches"`
			Type   string `json:"Type"`
			TypeID string `json:"TypeId"`
		} `json:"Result"`
		Succeeded bool        `json:"Succeeded"`
		Error     interface{} `json:"Error"`
		Message   interface{} `json:"Message"`
	} `json:"d"`
}

//No good. Not real data
func getCentrisData(location string) ResponseBody{

	var reqBody RequestBody = RequestBody{Text: location, Language: "en"}

	reqBytes, err := json.Marshal(reqBody);
	if err != nil {
		panic(err)
	}

    req, _ := http.NewRequest("POST", "https://www.centris.ca/Property/GetSearchAutoCompleteData", strings.NewReader(string(reqBytes)))
	req.Header.Add("referer", "https://www.centris.ca/");
	req.Header.Add("accept-language", "en-US,en;q=0.9");
	req.Header.Add("origin", "https://www.centris.ca");
	req.Header.Add("content-type", "application/json; charset=UTF-8");
	req.Header.Add("accept", "/*");
	req.Header.Add("authority", "www.centris.ca");
	
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    var listingResult ResponseBody
    err = json.Unmarshal(body, &listingResult)
    if err != nil {
        panic(err)
    }
	
	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected Status", res.Status)
	} else {
		fmt.Println("Success.", res.Status)

	}

	return listingResult
}