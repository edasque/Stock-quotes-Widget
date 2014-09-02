/bin/echo "Stocks"
/bin/echo "Ticker Value Change %"
/bin/echo -n "NASDAQ "
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=^IXIC&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "ConstantContact $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=CTCT&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "Apple $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=AAPL&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "MUTF:FINSX $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=FINSX&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "NFLX $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=NFLX&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "AMZN $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=AMZN&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "GE $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=GE&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "Exxon_Mobil $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=XOM&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "Astrazeneca $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=AZN&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "Amex $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=AXP&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
/bin/echo -n "Google $"
curl -s 'http://download.finance.yahoo.com/d/quotes.csv?s=GOOG&f=l1c' | tr '"' ' ' | tr ',' ' ' | sed s/-[[:space:]]//
