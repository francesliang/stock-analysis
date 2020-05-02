package fmp


type Sector struct {
    Name string           `json:"sector"`
    ChangePercent string    `json:"changesPercentage"`

}

type Sectors struct {
    SectorPerformance []Sector   `json:"sectorPerformance"`
}

type SymbolDetail struct {
    Symbol string    `json:"symbol"`
    Name string      `json:"name"`
    Price float32    `json:"price"`
}

type Symbols struct {
    SymbolList []SymbolDetail    `json:"symbolsList"`
}


type CompanyProfile struct {
    Symbol string       `json:"symbol"`
    ProfileDetail ProfileDetail    `json:"profile"`
}

type ProfileDetail struct {
    Price float32    `json:"price"`
    Beta string      `json:"beta"`
    VolumeAvg string `json:"volAvg"`
    MarketCap string `json:"mktCap"`
    LastDiv string   `json:"lastDiv"`
    Range string     `json:"range"`
    Changes float32  `json:"changes"`
    ChangesPercent string `json:"changesPercentage"`
    CompanyName string    `json:"companyName"`
    Exchange string       `json:"exchange"`
    Industry string       `json:"industry"`
    Website string        `json:"website"`
    Description string    `json:"description"`
    CEO string            `json:"ceo"`
    Sector string         `json:"sector"`
    Image string          `json:"image"`
}



