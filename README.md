# dash-binance
Command line tool to query the Binance API.

## Simple mode
- Run the `dash-binance` cli in a terminal window
- -tp <string> , e.g. BNBEUR for the trade pair (required!)

### Build the dash-binance cli
- `cd cli`
- `go build dash-binance.go` 
- `watch -n 30 -c "./dash-binance -tp BNBEUR"` ticker mode, updating with interval as defined with `-n xx` seconds, requesting `-tp BNBEUR` trade pair

### WIP
- add API key management

### DONE
- (done) remove hard coded symbol for calling trade pair BNBEUR
- (done) add arguments, options, flags handling for naming trading pairs dynamically