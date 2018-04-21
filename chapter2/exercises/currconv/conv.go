package currconv

func BTCtoUSD(b Btc) Usd {
    return Usd(b * Btc(BTCCostInUSD))
}

func USDtoBTC(u Usd) Btc {
    return Btc(u / BTCCostInUSD)
}
