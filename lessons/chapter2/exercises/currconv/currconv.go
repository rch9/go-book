// converts Bitcoin <-> USD
package currconv

import "fmt"

type Btc float64
type Usd float64

const (
    BTCCostInUSD Usd = 8100
)

func (b Btc) String() string {
    return fmt.Sprintf("%g BTC", b)
}

func (u Usd) String() string {
    return fmt.Sprintf("%g USD", u)
}

// func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32)*5/9) }
