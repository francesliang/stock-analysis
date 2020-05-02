package fmp

import (
    "fmt"
    "strconv"
)

type BasicValuation struct {
    Symbol string
    Beta float64
    PERatio float64
    PBRatio float64
    PEGRatio float64
    DividendYield float64
}

func GetCategoryMap() map[string]map[string][]string {
    categories := make(map[string]map[string][]string)
    symbols := SymbolsList()
    for _, symbol := range symbols[:5] {
        profile := Profile(symbol.Symbol)
        if categories[profile.Sector] == nil {
            categories[profile.Sector] = map[string][]string{}
        }
        var symbolList []string
        symbolList = categories[profile.Sector][profile.Industry]
        symbolList = append(symbolList, symbol.Symbol)
        categories[profile.Sector][profile.Industry] = symbolList
    }
    fmt.Println(categories)
    return categories
}

func HorizontalBasicValuation(symbols []string) map[string][]BasicValuation {
    comparison := make(map[string][]BasicValuation)

    for _, symbol := range symbols {
        basicValuation := VerticalBasicValuation(symbol)
        for key, value := range basicValuation {
            var valuationList []BasicValuation
            valuationList = comparison[key]
            valuationList = append(valuationList, value)
            comparison[key] = valuationList
        }
    }
    return comparison
}

func VerticalBasicValuation(symbol string) map[string]BasicValuation {
    //beta - lower the value, the less volatile the stock is
    //PE ratio - high P/E is ok if the growth is fast, otherwise expensive; low P/E but fast growth is worth watching
    //PB ratio - an accurate low P/B ratio generally protect you
    //PEG ratio - the lower PEG ratio, the better deal for stock's future estimated earnings; 1 is breaking even
    // Dividend yield - like the interest, vary by industries 
    comparison := make(map[string]BasicValuation)
    ratios := Ratios(symbol)
    beta, _ := strconv.ParseFloat(Profile(symbol).Beta, 64)
    for _, data := range ratios {
        peRatio, _ := strconv.ParseFloat(data.InvestmentValuationRatios.PriceEarningsRatio, 64)
        pbRatio, _ := strconv.ParseFloat(data.InvestmentValuationRatios.PriceToBookRatio, 64)
        pegRatio, _ := strconv.ParseFloat(data.InvestmentValuationRatios.PriceEarningsToGrowthRatio, 64)
        dividendYield, _ := strconv.ParseFloat(data.InvestmentValuationRatios.DividendYield, 64)

        comparison[data.Date] = BasicValuation{
            Symbol: symbol,
            Beta: beta,
            PERatio: peRatio,
            PBRatio: pbRatio,
            PEGRatio: pegRatio,
            DividendYield: dividendYield,
        }
    }
    return comparison
}
