package main

import (
    Fmp "stock-analysis/fmp"
    "fmt"
    "encoding/json"
)

func main() {
    //Fmp.Profile("AAPL")
    //Fmp.SectorPerformance()
    //Fmp.SymbolsList()
    //Fmp.Ratios("AAPL")
    //categories := Fmp.GetCategoryMap()
    //categoriesJson, _ := json.MarshalIndent(categories, "", "   ")
    //fmt.Println(string(categoriesJson))

    vComparison := Fmp.VerticalBasicValuation("AAPL")
    vComparisonJson, _ := json.MarshalIndent(vComparison, "", "   ")
    fmt.Println(string(vComparisonJson))

    hComparison := Fmp.HorizontalBasicValuation([]string{"AAPL", "GOOG"})
    hComparisonJson, _ := json.MarshalIndent(hComparison, "", "   ")
    fmt.Println(string(hComparisonJson))
}
