
    Shares.go

- Reads current price for selected tickers from CSV file 
- Reads Dividen calculation parameters from CSV files
- Reads desired sum to invest from STDIO
- writes to output file sorted list of shares to buy

Source data are placed into data/subdirectory

CurrentPrice.txt - list of the current proces in format:

Ticker, price

Dividends.txt - list of dividends and parrameters 

Ticker, dividend sum, tax value (percent in 0..1 range), months for period

Example of usage:

go run Shares.go 2000

Output:

Invested sum: 2000.00 USD

Ticker          Amount of shares        Planned monthly income
ARR                     189                     16.07 USD/month
OKE                     60                      15.98 USD/month
STWD                    113                     15.37 USD/month

Colombina:6 - FilesWorking 1$ go run Shares.go

Invested sum: 1000.00 USD

Ticker          Amount of shares        Planned monthly income
ARR                     94                      7.99 USD/month
OKE                     30                      7.99 USD/month
STWD                    56                      7.62 USD/month
Colombina:6 - FilesWorking 1$ 

olombina:6 - FilesWorking 1$ go run Shares.go 1100

Invested sum: 1100.00 USD

Ticker          Amount of shares        Planned monthly income
OKE                     33                      8.79 USD/month
ARR                     103                     8.76 USD/month
STWD                    62                      8.43 USD/month
Colombina:6 - FilesWorking 1$ 

