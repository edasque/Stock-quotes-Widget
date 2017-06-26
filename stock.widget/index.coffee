command: '/usr/local/bin/node Stock-quotes-Widget/stock.widget/src/get_stocks_info.js'

refreshFrequency: '1m'

style: """
  bottom: 10px
  right: 5px
  color: #fff
  font-family: Monaco
  table
    border-collapse: collapse
    table-layout: fixed
    background: rgba(#334477)
  td
    text-align: center
    padding: 3px 6px
    text-shadow: 0 0 1px rgba(#000, 0.5)

  thead tr
    background: rgba(#336699,0.2)
    &:first-child td
      font-weight: 400
	  font-size: 9px
    &:last-child td
      padding-bottom: 4px
      font-weight: 500
  a
    color:white
    text-decoration: none
  .updated
    text-align: center
    font-size:.4em
    padding-top: 1em

  tbody td
    font-size: 12px

  .today
    font-weight: bold
    background: rgba(#fff, 0.2)
    border-radius: 50%
"""

render: -> """
  <table>
    <thead>
    </thead>
    <tbody>
    </tbody>
  </table>
  <div class="updated">Time</div>

"""

updateHeader: (table) ->
  thead = table.find("thead")
  thead.empty()

  thead.append "<tr><td colspan='5'>Stock Market</td></tr>"
  tableRow = $("<tr></tr>").appendTo(thead)
  column_names = ["Ticker","Value","Change","%"] # ,"Spark"]

  for column_name in column_names
    tableRow.append "<td>#{column_name}</td>"
  $(".updated").html(Date().toLocaleString())

updateBody: (data, table) ->
  tbody = table.find("tbody")
  tbody.empty()

  darkred = "background:#600000;"
  midred = "background:#900000;"
  lightred = "background:#FF0000;"
  darkgreen = "background:#006000;"
  midgreen = "background:#009000;"
  lightgreen = "background:#00B000;"

  for property, ticker of data

    tableRow = $("<tr></tr>").appendTo(tbody)
    currency = "$"
    # if ticker.symbol.indexOf("^") is -1 then currencey = "er#"

    ticker.changeInPercent = (ticker.changeInPercent*100).toFixed(2)

    colour = lightgreen

    colour = switch
      when ticker.changeInPercent == 0 then ''
      when ticker.changeInPercent < -4 then lightred
      when ticker.changeInPercent < -1 then midred
      when ticker.changeInPercent < 0 then darkred
      when ticker.changeInPercent > 4 then lightgreen
      when ticker.changeInPercent > 1 then midgreen
      when ticker.changeInPercent > 0 then darkgreen

    if (ticker.change>0) then ticker.change = "+" + ticker.change

    ticker.history_csv = ""

    for item,index in ticker.historical
      if (item.close) then ticker.history_csv += item.close+","

    ticker.history_csv = (ticker.history_csv).substring(0, (ticker.history_csv).length - 1)

    ticker.change = if (ticker.change) then ticker.change else "&nbsp;&nbsp;&nbsp;"
    $("<td style='text-align:left;background:#336699;'><a href='https://www.google.com/finance?client=ob&q=#{ticker.symbol}'>#{ticker.symbol}</a></td>").appendTo(tableRow)
    $("<td style='text-align:center;#{colour}'>#{currency}#{ticker.lastTradePriceOnly}</td>").appendTo(tableRow)
    $("<td style='text-align:center;#{colour}'>#{ticker.change}</td>").appendTo(tableRow)
    $("<td style='text-align:center;#{colour}'>#{ticker.changeInPercent+"%"}</td>").appendTo(tableRow)
    # $("<td style='text-align:center;background:#FFF;'><span class='inlinesparkline'>#{ticker.history_csv}</span></td>").appendTo(tableRow)


  # $.fn.sparkline.defaults.common.chartRangeMin = 0;
  # $.fn.sparkline.defaults.common.type = 'line';
  # $('.inlinesparkline').sparkline(); 


update: (output, domEl) ->
  rows = JSON.parse(output)
  table = $(domEl).find("table")

  @updateHeader table

  # $.getScript './stock.widget/libs/jquery.sparkline.min.js.lib',
  @updateBody rows, table
