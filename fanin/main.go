package main

import (
  "log"
  "time"
  "strconv"
)

type Message interface {
  Payload() string
  Attributes() map[string]string
}

func FanIn(inputs ...<-chan Message) <-chan Message {
  c := make(chan Message)

  for idx, i := range inputs {
    log.Println("Fanning in input..." + strconv.FormatInt(int64(idx), 10))
    go func(i <-chan Message) { for { c<- <-i } }(i)
  }

  return c
}

type quote struct {
  ticker string
  price int64
  currency string
}

func (q *quote) Payload() string {
  return q.currency + " " + strconv.FormatInt(q.price, 10)
}

func (q *quote) Attributes() map[string]string {
  attrs := make(map[string]string, 3)
  attrs["ticker"] = q.ticker
  attrs["price"] = string(q.price)
  attrs["currency"] = q.currency
  return attrs
}


func main() {
  const SECS = 7
  c1 := quoteSource("CRM")
  c2 := quoteSource("RHT")
  c3 := quoteSource("MSFT")

  c4 := FanIn(c1, c2, c3)

  go printQuotes(c4)

  time.Sleep(SECS * time.Second)
  log.Panic("fin")
}

func printQuotes(c <-chan Message) {
  for {
    select {
    case q := <-c:
      ticker := q.Attributes()["ticker"]
      log.Println("Quote for " + ticker + ": " + q.Payload())
    }
  }
}

func quoteSource(ticker string) <-chan Message {
  c := make(chan Message)
  go func(t string) {
    for {
      c<- &quote{ ticker: t, price: 1234, currency: "USD" }
      time.Sleep(1 * time.Second)
    }
  }(ticker)
  return c
}
