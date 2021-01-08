# dash-binance
Command line tool to query the Binance API.

## Simple mode

- Run the `dash-binance` cli in a terminal window

### Build the `dash-binance` cli
- `cd cli`
- `go build dash-binance.go` 
- `watch -n 30 -c "./dash-binance"` to have the dashboard display in ticker mode, updating with interval as defined with `-n xx` seconds.

### WIP
- remove hard coded symbol for trading pair BNBEUR
- add arguments, options, flags handling for naming trading pairs dynamically
- add API key management