package fmp

import (
    "fmt"
    "encoding/json"
)

const FMPBaseUrl = "https://financialmodelingprep.com/api/v3/"

func SectorPerformance() {
    subpath := "stock/sectors-performance"
    url := FMPBaseUrl + subpath
    fmt.Println("SectorPerformance URL: ", url)

    res := GetRequest(url)
    fmt.Println(res)
    var sectors Sectors
    json.Unmarshal([]byte(res), &sectors)
    fmt.Println(sectors.SectorPerformance)

}

func Profile(symbol string) {
    profilePath := "company/profile/"

    url := FMPBaseUrl + profilePath + symbol
    fmt.Println("Profile URL: ", url)

    res := GetRequest(url)
    fmt.Println(res)
}

