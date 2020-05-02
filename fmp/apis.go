package fmp

import (
    "fmt"
    "encoding/json"
)

const FMPBaseUrl = "https://financialmodelingprep.com/api/v3/"

func SectorPerformance() []Sector {
    subpath := "stock/sectors-performance"
    url := FMPBaseUrl + subpath
    fmt.Println("SectorPerformance URL: ", url)

    res := GetRequest(url)
    var sectors Sectors
    json.Unmarshal([]byte(res), &sectors)
    fmt.Println(sectors.SectorPerformance)
    return sectors.SectorPerformance

}

func Profile(symbol string) ProfileDetail {
    profilePath := "company/profile/"

    url := FMPBaseUrl + profilePath + symbol
    fmt.Println("Profile URL: ", url)

    res := GetRequest(url)
    var profile CompanyProfile
    json.Unmarshal([]byte(res), &profile)
    fmt.Println(profile)
    return profile.ProfileDetail
}

func SymbolsList() []SymbolDetail {
    subpath := "company/stock/list"
    url := FMPBaseUrl + subpath
    fmt.Println("SymbolsList URL: ", url)

    res := GetRequest(url)
    var symbols Symbols
    json.Unmarshal([]byte(res), &symbols)
    fmt.Println(symbols)
    return symbols.SymbolList

}

