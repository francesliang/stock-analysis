package main

import (
    Fmp "stock-analysis/fmp"
    "fmt"
    "encoding/json"
    "io/ioutil"
)

func main() {
    var categoryJsonFile = "stock-categories.json"
    categories := Fmp.GetCategoryMap(categoryJsonFile)
    categoriesJson, _ := json.MarshalIndent(categories, "", "   ")
    ioutil.WriteFile(categoryJsonFile, categoriesJson, 0644)
    fmt.Println(string(categoriesJson))

    //vComparison := Fmp.VerticalBasicValuation("AAPL")
    //vComparisonJson, _ := json.MarshalIndent(vComparison, "", "   ")
    //fmt.Println(string(vComparisonJson))

    //hComparison := Fmp.HorizontalBasicValuation([]string{"AAPL", "GOOG"})
    //hComparisonJson, _ := json.MarshalIndent(hComparison, "", "   ")
    //fmt.Println(string(hComparisonJson))
}
