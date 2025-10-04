# CSV Filter Tool

This is a simple script that takes an exported CSV from Trading212 and converts it to a CSV that can be imported
into YNAB

It is meant to be used to pull card transactions rather than investment details, so it filters out:
 - Interest earned on cash
 - Cashback earned
 - Market orders (Buy and Sell)

It takes the path of a Trading212 export as parameter 1

---

## Example Usage

```bash
go run csvconvert.go trading212.csv
```

Alternately build the binary and run:

```bash
go build -o csvconvert csvconvert.go
./csvconvert trading212.csv
```