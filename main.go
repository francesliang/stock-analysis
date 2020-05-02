package main

import (
    Fmp "stock-analysis/fmp"
    "fmt"
    "encoding/json"
    "io/ioutil"
)


func generateCategoryMap(categoryJsonFile string) {
    categories := Fmp.GetCategoryMap(categoryJsonFile)
    categoriesJson, _ := json.MarshalIndent(categories, "", "   ")
    ioutil.WriteFile(categoryJsonFile, categoriesJson, 0644)
    fmt.Println(string(categoriesJson))
}

func CategoryValuation(categoryJsonFile string, sector string, industry string) {
    var categories = make(map[string]map[string][]string)
    data, _ := ioutil.ReadFile(categoryJsonFile)
    json.Unmarshal(data, &categories)

    symbols := categories[sector][industry]

    hComparison := Fmp.HorizontalBasicValuation(symbols)
    hComparisonJson, _ := json.MarshalIndent(hComparison, "", "   ")
    fmt.Println(string(hComparisonJson))

}

func main() {
    var categoryJsonFile = "stock-categories.json"
    //generateCategoryMap(categoryJsonFile)

    var sector = "Technology"
    var industry = "Application Software"
    CategoryValuation(categoryJsonFile, sector, industry)

    //vComparison := Fmp.VerticalBasicValuation("AAPL")
    //vComparisonJson, _ := json.MarshalIndent(vComparison, "", "   ")
    //fmt.Println(string(vComparisonJson))

}
