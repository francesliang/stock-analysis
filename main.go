package main

import (
    Fmp "stock-analysis/fmp"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "flag"
    "strings"
)

type CLIFlags struct {
    getcategory bool
    sector string
    industry string
    symbols string
}


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

    var flags CLIFlags

    flag.BoolVar(&flags.getcategory, "getcategory", false, "Bool flag to indicate if to generate category mapping")
    flag.StringVar(&flags.sector, "sector", "", "Sector of the stock")
    flag.StringVar(&flags.industry, "industry", "", "Industry of the stock")
    flag.StringVar(&flags.symbols, "symbols", "", "Symbols of stocks, separated by comma")
    flag.Parse()

    var categoryJsonFile = "stock-categories.json"
    if flags.getcategory { 
        generateCategoryMap(categoryJsonFile)
    }

    if (flags.sector != "" && flags.industry != "") {
        CategoryValuation(categoryJsonFile, flags.sector, flags.industry)
    }

    if flags.symbols != "" {
        for _, symbol := range strings.Split(flags.symbols, ",") {
            vComparison := Fmp.VerticalBasicValuation(symbol)
            vComparisonJson, _ := json.MarshalIndent(vComparison, "", "   ")
            fmt.Println(string(vComparisonJson))
        }
    }

}
