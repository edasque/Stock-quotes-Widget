Stock-quotes-Widget
===========
Übersicht Stock quotes Widget **Version 2.0**

Rewritten in GoLang (versus JavaScript)

Displays stock quotes in a table format with color-coding of wins/loss and sparklines for the last 20d

## Setup & customization

Install by moving the folder stock.widget to the Übersicht widget folder.

Adjust bottom and right in the style section in index.coffee if you want it placed somewhere else on the screen

Adjust and amend stock list in src/config.json

<img width="310" alt="screenshot_11_7_17__10_33_am" src="https://user-images.githubusercontent.com/219187/32501860-2a0d9fca-c3a7-11e7-9234-6f734637a390.png">

## Important

Since it seems that the Yahoo Finance API has been deprecated as had been the Google Finance API, I switched to the Robinhood Private API, [documented here](https://github.com/sanko/Robinhood), which *for now* allows querying real time stock quotes and get good historicals to plot the sparklines.
