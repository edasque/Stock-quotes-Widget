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
getHistoricalQuotes(a_month_ago, now, tickers);
