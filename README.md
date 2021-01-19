# dash-binance
Command line tool to query the Binance API.

### Build the dash-binance cli
- `git clone https://github.com/gmzabos/dash-binance.git`
- `cd cli`
- `go build dash-binance.go`

## Simple mode
- Run the `dash-binance` cli in a terminal window using [`watch`](https://linux.die.net/man/1/watch) 
- Add flag to specify trade pair (required!)
- Example: `watch -n 30 -c "./dash-binance -tp BNBEUR"` updating with interval as defined with `-n xx` seconds, requesting `-tp BNBEUR` trade pair

### WIP
- [ ] add more flags
- [ ] add API key management

### DONE
- [x] remove hard coded symbol for calling trade pair BNBEUR
- [x] add arguments, options, flags handling for naming trading pairs dynamically