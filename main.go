package main

import (
    Fmp "stock-analysis/fmp"
)

func main() {
    Fmp.Profile("AAPL")
    Fmp.SectorPerformance()
    Fmp.SymbolsList()
}
