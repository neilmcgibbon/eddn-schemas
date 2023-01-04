package v1

import (
	"time"
)

type Schema struct {
	AbsoluteMagnitude     float64                 `json:"AbsoluteMagnitude"`     // Events: Scan
	Age_MY                int32                   `json:"Age_MY"`                // Events: Scan
	AscendingNode         float64                 `json:"AscendingNode"`         // Events: Scan
	Atmosphere            string                  `json:"Atmosphere"`            // Events: Scan
	AtmosphereComposition []AtmosphereComposition `json:"AtmosphereComposition"` // Events: Scan
	AtmosphereType        string                  `json:"AtmosphereType"`        // Events: Scan
	AxialTilt             float64                 `json:"AxialTilt"`             // Events: Scan
	Body                  string                  `json:"Body"`                  // Events: FSDJump
	BodyID                int                     `json:"BodyID"`                // Events: FSDJump, Scan
	BodyName              string                  `json:"BodyName"`              // Events: Scan
	BodyType              string                  `json:"BodyType"`              // Events: FSDJump
	Composition           Composition             `json:"Composition"`           // Events: Scan
	Conflicts             []Conflict              `json:"Conflicts"`             // Events: FSDJump
	Eccentricity          float64                 `json:"Eccentricity"`          // Events: Scan
	Event                 string                  `json:"event"`                 // Events: all
	Factions              []Faction               `json:"Factions"`              // Events: FSDJump
	IsHorizons            bool                    `json:"horizons"`              // Events: all
	IsOdyssey             bool                    `json:"odyssey"`               // Events: all
	Landable              bool                    `json:"Landable"`              // Events: Scan
	LandingPads           LandingPads             `json:"LandingPads"`           // Events: Docked
	Luminosity            string                  `json:"Luminosity"`            // Events: Scan
	MarketID              int64                   `json:"MarketID"`              // Events: Docked
	MassEM                float64                 `json:"MassEM"`                // Events: Scan
	Materials             []Material              `json:"Materials"`             // Events: Scan
	MeanAnomaly           float64                 `json:"MeanAnomaly"`           // Events: Scan
	OrbitalInclination    float64                 `json:"OrbitalInclination"`    // Events: Scan
	OrbitalPeriod         float64                 `json:"OrbitalPeriod"`         // Events: Scan
	Parents               []map[string]int        `json:"Parents"`               // Events: Scan
	Periapsis             float64                 `json:"Periapsis"`             // Events: Scan
	PlanetClass           string                  `json:"PlanetClass"`           // Events: Scan
	Population            int64                   `json:"Population"`            // Events: FSDJump
	Radius                float64                 `json:"Radius"`                // Events: Scan
	ReserveLevel          string                  `json:"ReserveLevel"`          // Events: Scan
	Rings                 []Ring                  `json:"Rings"`                 // Events: Scan
	RotationPeriod        float64                 `json:"RotationPeriod"`        // Events: Scan
	ScanType              string                  `json:"ScanType"`              // Events: Scan
	SemiMajorAxis         float64                 `json:"SemiMajorAxis"`         // Events: Scan
	StarPosition          []float64               `json:"StarPos"`               // Events: FSDJump
	StarSystem            string                  `json:"StarSystem"`            // Events: Docked, FSDJump, Scan
	StarType              string                  `json:"StarType"`              // Events: Scan
	StationAllegiance     string                  `json:"StationAllegiance"`     // Events: Docked
	StationEconomies      []StationEconomy        `json:"StationEconomies"`      // Events: Docked
	StationEconomy        string                  `json:"StationEconomy"`        // Events: Docked
	StationFaction        StationFaction          `json:"StationFaction"`        // Events: Docked
	StationGovernment     string                  `json:"StationGovernment"`     // Events: Docked
	StationName           string                  `json:"StationName"`           // Events: Docked
	StationServices       []string                `json:"StationServices"`       // Events: Docked
	StationType           string                  `json:"StationType"`           // Events: Docked
	StellarMass           float64                 `json:"StellarMass"`           // Events: Scan
	Subclass              int                     `json:"Subclass"`              // Events: Scan
	SurfaceGravity        float64                 `json:"SurfaceGravity"`        // Events: Scan
	SurfacePressure       float64                 `json:"SurfacePressure"`       // Events: Scan
	SurfaceTemperature    float64                 `json:"SurfaceTemperature"`    // Events: Scan
	SystemAddress         int64                   `json:"SystemAddress"`         // Events: Docked, FSDJump, Scan
	SystemAllegiance      string                  `json:"SystemAllegiance"`      // Events: FSDJump
	SystemEconomy         string                  `json:"SystemEconomy"`         // Events: FSDJump
	SystemFaction         SystemFaction           `json:"SystemFaction"`         // Events: FSDJump
	SystemGovernment      string                  `json:"SystemGovernment"`      // Events: FSDJump
	SystemSecondEconomy   string                  `json:"SystemSecondEconomy"`   // Events: FSDJump
	SystemSecurity        string                  `json:"SystemSecurity"`        // Events: FSDJump
	TerraformState        string                  `json:"TerraformState"`        // Events: Scan
	TidalLock             bool                    `json:"TidalLock"`             // Events: Scan
	Timestamp             time.Time               `json:"timestamp"`             // Events: all
	Volcanism             string                  `json:"Volcanism"`             // Events: Scan

	// Currently ignored fields:
	// - Taxi (not sure what this is)
	// - Multicrew
	// - JumpDist
	// - FuelUsed
	// - FuelLevel
	// - DistanceFromArrivalLS
	// - WasDiscovered
	// - WasMapped
	// - DistFromStarLS
	// - BoostUsed
	// - ActiveFine
	// - Wanted
}

type Conflict struct {
	WarType  string          `json:"WarType"`
	Status   string          `json:"Status"`
	Faction1 ConflictFaction `json:"Faction1"`
	Faction2 ConflictFaction `json:"Faction2"`
}

type ConflictFaction struct {
	Name    string `json:"Name"`
	Stake   string `json:"Stake"`
	WonDays int    `json:"WonDays"`
}

type Composition struct {
	Ice   float64 `json:"Ice"`
	Rock  float64 `json:"Rock"`
	Metal float64 `json:"Metal"`
}

type Material struct {
	Name    string  `json:"Name"`
	Percent float64 `json:"Percent"`
}

type Ring struct {
	Name      string  `json:"Name"`
	RingClass string  `json:"RingClass"`
	MassMT    float64 `json:"MassMT"`
	InnerRad  float64 `json:"InnerRad"`
	OuterRad  float64 `json:"OuterRad"`
}

type AtmosphereComposition struct {
	Name    string  `json:"Name"`
	Percent float64 `json:"Percent"`
}

type Faction struct {
	ActiveStates     []FactionState `json:"ActiveStates"`
	Allegiance       string         `json:"Allegiance"`
	FactionState     string         `json:"FactionState"`
	Government       string         `json:"Government"`
	Happiness        string         `json:"Happiness"`
	Influence        float64        `json:"Influence"`
	Name             string         `json:"Name"`
	RecoveringStates []FactionState `json:"RecoveringStates"`
	PendingStates    []FactionState `json:"PendingStates"`
}

type StationEconomy struct {
	Name       string  `json:"Name"`
	Proportion float64 `json:"Proportion"`
}

type FactionState struct {
	State string `json:"State"`
	Trend int    `json:"Trend"`
}

type LandingPads struct {
	Small  int `json:"Small"`
	Medium int `json:"Medium"`
	Large  int `json:"Large"`
}

type SystemFaction struct {
	Name         string `json:"Name"`
	FactionState string `json:"FactionState"`
}

type StationFaction struct {
	Name string `json:"Name"`
}

func (s *Schema) PostDecode() {
	// do nothing
}

func (s *Schema) GetSchema() string {
	return "https://eddn.edcd.io/schemas/journal/1"
}
