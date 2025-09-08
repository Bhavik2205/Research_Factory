package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Bhavik2205/Research_Factory.git/internal/parser"
)

func main() {
	f, err := os.Open(`D:\Research\03272019.NASDAQ_ITCH50`)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sbr := parser.NewSoupBinReader(f)

	for {
		payload, err := sbr.NextPayload()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("framing error: %v", err)
		}
		if payload == nil {
			continue // heartbeat
		}

		dec := parser.NewDecoder(bytes.NewReader(payload))
		for {
			env, err := dec.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("decode error: %v", err)
			}
			// ✅ Filter here
			// switch env.Type {
			// case 'A', 'E', 'P': // only these message types
			// 	fmt.Printf("Type=%c -> %+v\n", env.Type, env.Msg)
			// default:
			// 	// skip everything else
			// }
			switch m := env.Msg.(type) {
			case *parser.AddOrderNoMPIDAttribution: // message type 'A'
				side := "Buy"
				if m.BuySellIndicator == 'S' {
					side = "Sell"
				}
				price := float64(m.Price) / 10000.0 // convert to decimal price
				fmt.Printf("[%d] %s %d shares of %s @ %.4f (OrderID=%d)\n",
					m.Header.Timestamp,
					side,
					m.Shares,
					m.Stock,
					price,
					m.OrderReferenceNumber,
				)
			// case *parser.OrderExecuted:
			// case *parser.TradeMessageNonCross:
			//     add similar formatting here if needed
			default:
				// skip other message types
			}
			// fmt.Printf("Type=%c -> %+v\n", env.Type, env.Msg)
		}
	}
}
