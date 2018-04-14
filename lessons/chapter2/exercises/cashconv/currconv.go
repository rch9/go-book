// converts Bitcoin <-> USD
package currconv

type btc float64
type usd float64

const (
    BTCCostInUSD usd = 8100
)

func BTCtoUSD(btc b) usd {
    return usd(b * BTCCostInUSD)
}

func USDtoBTC(usd u) btc {
    return btc(u / BTCCostInUSD)
}
