'use strict'
const tickers = require("./config.json").tickers.sort()
const yahooFinance = require('yahoo-finance');
const moment = require('moment');
const fields_we_need = ['s', 'l1', 'c1', 'p', 'p2']
const now = moment()
  .format();
const a_month_ago = moment()
  .subtract(1, 'months')
  .format();
var historical_quotes;
var current_quotes;
var consolidated_quotes = {}

const request  = require('request')


const RobinHood_API_URL="https://api.robinhood.com/quotes/?symbols=" // "NVDA,GOOG"

function getHistoricalQuotes(from_date, to_date, for_tickers) {
  yahooFinance.historical({
    symbols: for_tickers,
    from: from_date,
    to: to_date,
  }, (err, result) => {
    if (err) console.error(err)
    else {
      historical_quotes = result
      getQuotes(tickers);
    }
  })
}
function getQuotes_from_RH(for_tickers,callback) {

  var request = require('request');

  var quotes = ""

  for_tickers.forEach(function(value, key, listObj) {

    quotes += value
    if (key!=(listObj.length-1)) quotes += ','

  })


  request(RobinHood_API_URL+quotes, function (error, response, body) {
    if (error) {
     console.log('error:', error); 
     console.log('statusCode:', response && response.statusCode); 
    }
    else{
    callback(JSON.parse(body).results); 
  }
    
  });

}

function printResults (data) {
  // console.log(JSON.stringify(data));
  var output_array=[]
  data.forEach(function(value, key, listObj) {

    const lastTradePriceOnly=value.last_trade_price
    const previousClose = value.previous_close
    const change = (lastTradePriceOnly - previousClose).toFixed(2)
    const changeInPercent = (lastTradePriceOnly  / previousClose)-1
    // console.log(lastTradePriceOnly)
    // console.log(previousClose)
    // console.log(change)
    // console.log(changeInPercent)

    output_array.push({"symbol":value.symbol,"previousClose":previousClose,"lastTradePriceOnly":lastTradePriceOnly,"change":change,"changeInPercent":changeInPercent})


  })
  var output = {"quotes":output_array}
  console.log(JSON.stringify(output))
}


function getQuotes(for_tickers) {
  yahooFinance.snapshot({
    symbols: for_tickers,
    fields: fields_we_need
  }, (err, snapshot) => {
    if (err) console.error(err)
    else {
      current_quotes = snapshot
      processResults()
    }
  })
}

function processResults() {
  current_quotes.forEach((entry, index) => {
    consolidated_quotes[entry.symbol] = entry
  })
  tickers.forEach((ticker, index) => {
    consolidated_quotes[ticker].historical = historical_quotes[ticker]
  })
  console.log(JSON.stringify(consolidated_quotes))
}
// getHistoricalQuotes(a_month_ago, now, tickers);
getQuotes_from_RH(tickers,printResults)
