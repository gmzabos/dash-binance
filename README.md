# dash-binance
Dashboard your favorite data from the Binance API.

## Simple mode

- Run the dashboard in simple mode from a terminal window ("ticker")

### Build for ticker mode
- `cd ticker`
- `go build ticker.go` 
- `watch -n 30 -c "./ticker"` to have the dashboard display in ticker mode, updating with interval as defined with `-n xx` seconds.

## Complex mode

- Complex mode is currently WIP