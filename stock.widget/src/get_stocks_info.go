package main

// http://www.nasdaq.com/earnings/report/amzn
// $("#left-column-div h2").eq(0)
// <h2>Earnings announcement* for AMZN: Feb 04, 2016</h2>
// <h2>Earnings announcement* for BOBX:</h2>

import (
	"encoding/json"
	"fmt"
	"strconv"

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

const RobinHood_API_Historicals_pre = "https://api.robinhood.com/quotes/historicals/" // "NVDA,GOOG"
const RobinHood_API_Historicals_post = "/?span=year&interval=day&bounds=regular"      // "NVDA,GOOG"
const RobinHood_API_Intraday = "https://api.robinhood.com/quotes/?symbols="

var myClient = &http.Client{Timeout: 10 * time.Second}

/**
 * @brief      Gets a CSV string from an array of string.
 *
 * @param      tickers  The array of string
 *
 * @return     The csv as a string.
 */
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

func getQuotesFromRHYearWorthOfDays(ticker string, auth string) []string {
	var URLToQuery = RobinHood_API_Historicals_pre + ticker + RobinHood_API_Historicals_post
	// fmt.Println("URLToQuery: ", URLToQuery)
	req, err := http.NewRequest("GET", URLToQuery, nil)
	// ...

	req.Header.Add("Authorization", auth)
	response, err := myClient.Do(req)

	// response, err := myClient.Get(URLToQuery)
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
		const X = 20
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

type rhAuthResult struct {
	AccessToken string `json:"access_token"`
}

type rhIntradayResult struct {
	Symbol           string
	Last_trade_price string
	Previous_close   string
}
type rhIntradayResponse struct {
	Results []rhIntradayResult
}

func getAuthToken(user string, password string) string {

	url := "https://api.robinhood.com/oauth2/token/?username=" + user + "&password=" + password + "&grant_type=password&client_id=c82SH0WZOsabOXGP2sxqcj34FxkvfnWRZBKlBjFS"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("Postman-Token", "4e5775f8-f7cd-41e9-85c9-29aa6b2c7583")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	auth := rhAuthResult{}
	json.Unmarshal(body, &auth)

	// fmt.Println("AuthToken:", auth.AccessToken)

	return ("Bearer " + auth.AccessToken)
}
func getIntradayFromRH(tickersAsCSV string, Auth string) []rhIntradayResult {
	var URLToQuery = RobinHood_API_Intraday + tickersAsCSV

	req, err := http.NewRequest("GET", URLToQuery, nil)

	req.Header.Add("Authorization", Auth)
	response, err := myClient.Do(req)

	// response, err := myClient.Get(URLToQuery)
	if err != nil {
		// return err
		fmt.Println("Error:", err)
		if response != nil {
			fmt.Println("Status:", response.Status)
			fmt.Println("StatusCode:", response.StatusCode)
		}
		return nil
	} else {

		var quoteResponse rhIntradayResponse

		buf, _ := ioutil.ReadAll(response.Body)

		jsonErr := json.Unmarshal(buf, &quoteResponse)

		if jsonErr != nil {
			// return err
			fmt.Println("jsonErr:", jsonErr)
		}

		defer response.Body.Close()

		return quoteResponse.Results

		// foo1 := new(RH_Quotes_Response)
		// json.NewDecoder(r.Body).Decode(foo1)
		// fmt.Println(foo1)

	}

}

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
// type Historical_info struct {
// 	Close_price string
// }

type Ticker struct {
	Symbol             string   `json:"symbol"`
	LastTradePriceOnly string   `json:"lastTradePriceOnly"`
	Previous_close     string   `json:"previousClose"`
	ChangeInPercent    string   `json:"changeInPercent"`
	Change             string   `json:"change"`
	LastX              []string `json:"lastX"`
}

type JSONOutput struct {
	// Historicals []Historical_info
	Tickers []Ticker `json:"quotes"`
}

func main() {

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	type Config struct {
		Tickers  []string `json:"tickers"`
		User     string   `json:"user"`
		Password string   `json:"password"`
	}

	conf := Config{}
	json.Unmarshal(file, &conf)

	AuthString := getAuthToken(conf.User, conf.Password)

	sort.Strings(conf.Tickers)

	CSVstring := getCSVfromTickersArray(conf.Tickers)

	var intradaysResults = getIntradayFromRH(CSVstring, AuthString)
	// fmt.Println(intradaysResults)

	var outputStructure JSONOutput
	outputStructure.Tickers = make([]Ticker, len(conf.Tickers), len(conf.Tickers)*2)

	for index := range conf.Tickers {
		var itr = intradaysResults[index]
		ltp, err := strconv.ParseFloat(itr.Last_trade_price, 64)
		lpc, err2 := strconv.ParseFloat(itr.Previous_close, 64)

		if err != nil || err2 != nil {
			fmt.Printf("Float conversion error: %v, %v\n", err, err2)
		}
		var change = strconv.FormatFloat(ltp-lpc, 'f', 2, 64)
		var changePercent = strconv.FormatFloat((ltp-lpc)*100.0/ltp, 'f', 2, 64)
		var historicals = getQuotesFromRHYearWorthOfDays(itr.Symbol, AuthString)

		outputStructure.Tickers[index] = Ticker{itr.Symbol, itr.Last_trade_price, itr.Previous_close, changePercent, change, historicals}
	}

	// for index, element := range conf.Tickers {
	// 	var historicals = getQuotesFromRHYearWorthOfDays(element)
	// 	var bob = Historical_info{element, Quotes}
	// 	outputStructure.Historicals[index] = bob
	// }

	b, _ := json.Marshal(outputStructure)
	fmt.Println(string(b))

	// wg.Wait()

}
