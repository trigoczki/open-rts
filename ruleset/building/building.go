package building

// Cost represents the cost of resources to build/upgrade a building
type Cost struct {
	ResourceID string  `json:"resourceId"`
	Amount     float64 `json:"amount"`
}

// Production represents the resources produced by a building
type Production struct {
	ResourceID string  `json:"resourceId"`
	Amount     float64 `json:"amount"`
}

// Modifier represents a dynamic effect a building applies
type Modifier struct {
	Type            string  `json:"type"`
	Scope           string  `json:"scope"`
	ResourceID      string  `json:"resourceId,omitempty"`
	Factor          float64 `json:"factor"`
	ModifierFormula string  `json:"modifierFormula,omitempty"`
}

// BuildingTime represents the formula or base time needed for building/upgrade
type BuildingTime struct {
	BaseTime            float64 `json:"baseBuildingTime"`
	BuildingTimeFormula string  `json:"buildingTimeFormula"`
}

// Costs structure encapsulates base costs and cost formulas
type Costs struct {
	BaseCosts   []Cost `json:"baseCosts"`
	CostFormula string `json:"costFormula"`
}

// ProductionSpec defines the base production and formulas
type ProductionSpec struct {
	BaseProduction    []Production `json:"baseProduction"`
	ProductionFormula string       `json:"productionFormula"`
}

// ModifierSpec encapsulates base modifiers and formulas
type ModifierSpec struct {
	BaseModifiers   []Modifier `json:"baseModifier"`
	ModifierFormula string     `json:"modifierFormula"`
}

type Level struct {
	BaseBuildingTime float64      `json:"baseBuildingTime"`
	BaseCosts        []Cost       `json:"baseCost"`
	BaseProduction   []Production `json:"baseProduction"`
	BaseModifiers    []Modifier   `json:"baseModifier"`
}

// Levels stores level-specific data for buildings using the CONCRETE mode
type Levels map[int]Level

// Building represents the main structure for a building
type Building struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	MaxLevel     int            `json:"maxLevel"`
	MetadataMode string         `json:"metadataMode"`
	BuildingTime BuildingTime   `json:"buildingTime,omitempty"`
	Costs        Costs          `json:"costs,omitempty"`
	Production   ProductionSpec `json:"production,omitempty"`
	Modifiers    ModifierSpec   `json:"modifier,omitempty"`
	GrowthFactor float64        `json:"growthFactor,omitempty"`
	Levels       Levels         `json:"levels,omitempty"`
}
