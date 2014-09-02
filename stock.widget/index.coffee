command: 'stock.widget/scripts/uber-stock.sh'

refreshFrequency: 60000

style: """
  bottom: 400px
  right: 10px
  color: #fff
  font-family: Helvetica Neue

  table
    border-collapse: collapse
    table-layout: fixed
  td
    text-align: center
    padding: 4px 6px
    text-shadow: 0 0 1px rgba(#000, 0.5)

  thead tr
    &:first-child td
      font-size: 24px
      font-weight: 200

    &:last-child td
      font-size: 11px
      padding-bottom: 4px
      font-weight: 500


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
"""

updateHeader: (rows, table) ->
  thead = table.find("thead")
  thead.empty()

  thead.append "<tr><td colspan='4'>#{rows[0]}</td></tr>"
  tableRow = $("<tr></tr>").appendTo(thead)
  column_names = rows[1].split(/\s+/)

  for column_name in column_names
    tableRow.append "<td>#{column_name}</td>"

updateBody: (rows, table) ->
  tbody = table.find("tbody")
  tbody.empty()
  rows.splice 0, 2
  today = rows.pop().split(/\s+/)[2]

  for ticker, i in rows
    items = ticker.split(/\s+/).filter (item) -> item.length > 0
    tableRow = $("<tr></tr>").appendTo(tbody)


    for item,col_id in items
      if col_id<1
        text_align = "text-align:left;"
      else text_align="text-align:center;"

      if (item.indexOf("+") != -1)
        cell = $("<td style='background:green;#{text_align}'>#{item}</td>").appendTo(tableRow)
      else
        if ((item.indexOf("-") != -1) & item != "-")
          cell = $("<td style='background:red;#{text_align}'>#{item}</td>").appendTo(tableRow)
        else cell = $("<td style='#{text_align}'>#{item}</td>").appendTo(tableRow)


update: (output, domEl) ->
  rows = output.split("\n")
  table = $(domEl).find("table")

  @updateHeader rows, table
  @updateBody rows, table
