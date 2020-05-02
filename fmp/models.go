package fmp


type Sector struct {
    Name string           `json:"sector"`
    ChangePercent string    `json:"changesPercentage"`

}

type Sectors struct {
    SectorPerformance []Sector   `json:"sectorPerformance"`
}
