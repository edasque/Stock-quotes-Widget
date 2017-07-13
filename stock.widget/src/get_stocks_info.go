package main

// http://www.nasdaq.com/earnings/report/amzn
// $("#left-column-div h2").eq(0)
// <h2>Earnings announcement* for AMZN: Feb 04, 2016</h2>
// <h2>Earnings announcement* for BOBX:</h2>

import (
	"encoding/json"
	"fmt"
	// "log"
	"os"
	"sort"
	// "strings"
	// "sync"
	"io/ioutil"
	// "log"
	"net/http"
	"time"
)

const RobinHood_API_URL_pre = "https://api.robinhood.com/quotes/historicals/" // "NVDA,GOOG"
const RobinHood_API_URL_post = "/?span=year&interval=day&bounds=regular"      // "NVDA,GOOG"

var myClient = &http.Client{Timeout: 10 * time.Second}

func getCSVfromTickersArray(tickers []string) string {
	list := ""
	for key, element := range tickers {
		list += element
		// if (key!=(listObj.length-1)) quotes += ','
		if key != (len(tickers) - 1) {
			list += ","
		}

	}
	return (list)
}

type rhHistoricaQuotesResult struct {
	// Begins_at   string
	Close_price string
}
type rhHistoricaQuotesResponse struct {
	Historicals []rhHistoricaQuotesResult
	Symbol      string
}

func getQuotesFromRHYearWorthOfDays(ticker string) []string {
	var URLToQuery = RobinHood_API_URL_pre + ticker + RobinHood_API_URL_post
	// fmt.Println("URLToQuery: ", URLToQuery)
	response, err := myClient.Get(URLToQuery)
	if err != nil {
		// return err
		fmt.Println("Error:", err)
		if response != nil {
			fmt.Println("Status:", response.Status)
			fmt.Println("StatusCode:", response.StatusCode)
		}
		return nil

	} else {
		var quoteResponse rhHistoricaQuotesResponse

		buf, _ := ioutil.ReadAll(response.Body)

		jsonErr := json.Unmarshal(buf, &quoteResponse)

		if jsonErr != nil {
			// return err
			fmt.Println("jsonErr:", jsonErr)
		}
		const X = 10
		var numberOfDataPoints = len(quoteResponse.Historicals)
		var lastX = append(quoteResponse.Historicals[numberOfDataPoints-X : numberOfDataPoints])
		// fmt.Println("Historicals length:", numberOfDataPoints)

		// fmt.Println("Historicals:", lastX)

		var closingPrices = make([]string, X, 2*X)
		for index, value := range lastX {
			// fmt.Println("Historicals:", value.Close_price)
			closingPrices[index] = value.Close_price
		}
		// fmt.Println("Historicals last X:", closingPrices)

		defer response.Body.Close()

		return closingPrices

		// foo1 := new(RH_Quotes_Response)
		// json.NewDecoder(r.Body).Decode(foo1)
		// fmt.Println(foo1)

	}

}

// var request = require('request');

// var quotes = ""

// for_tickers.forEach(function(value, key, listObj) {

//   quotes += value
//   if (key!=(listObj.length-1)) quotes += ','

// })

// request(RobinHood_API_URL+quotes, function (error, response, body) {
//   if (error) {
//    console.log('error:', error);
//    console.log('statusCode:', response && response.statusCode);
//   }
//   else{
//   callback(JSON.parse(body).results);
// }

// });

// }

func parseDate(timeString string) string {

	// The standard time used in the layouts is:
	// Mon Jan 2 15:04:05 MST 2006 (MST is GMT-0700)

	// dateFormat := "Jan 2, 2006"
	dateFormat := "2/1/2006"

	t, err := time.Parse(dateFormat, timeString)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return t.Format("Jan 2")

}

// func getEarningsInfoFromLambda(tickers string) map[string]interface{} {

// }

// var wg sync.WaitGroup
type Historical_info struct {
	Symbol string
	// Begins_at   string
	Close_price []string
}

type JSONOutput struct {
	Historicals []Historical_info
}

func main() {

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	// fmt.Printf("config.json: %s\n", string(file))

	type Config struct {
		Tickers []string `json:"tickers"`
	}

	conf := Config{}
	json.Unmarshal(file, &conf)
	sort.Strings(conf.Tickers)

	// g_tickers := [...]string{"AAPL", "EIGI", "DIS", "HAS", "AMZN", "NFLX", "AXP", "GOOG", "FB", "TSLA", "MSFT"}
	// fmt.Println(conf.Tickers)

	// CSVstring := getCSVfromTickersArray(conf.Tickers)
	// fmt.Println(CSVstring)
	var outputStructure JSONOutput
	outputStructure.Historicals = make([]Historical_info, len(conf.Tickers), len(conf.Tickers)*2)

	for index, element := range conf.Tickers {
		var Quotes = getQuotesFromRHYearWorthOfDays(element)
		var bob = Historical_info{element, Quotes}
		outputStructure.Historicals[index] = bob
	}
	fmt.Println(outputStructure)

	b, _ := json.Marshal(outputStructure)
	fmt.Println(string(b))

	// wg.Wait()

}
