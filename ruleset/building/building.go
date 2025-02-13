package building

import (
	"github.com/Knetic/govaluate"
)

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
	Type            string      `json:"type"`
	Scope           string      `json:"scope"`
	ResourceID      string      `json:"resourceId,omitempty"`
	Factor          float64     `json:"factor"`
	ModifierFormula string      `json:"modifierFormula,omitempty"`
	CompiledFormula FormulaFunc `json:"-"`
}

// FormulaFunc is a type that represents an executable formula
type FormulaFunc func(params map[string]float64) float64

// BuildingTime represents the formula or base time needed for building/upgrade
type BuildingTime struct {
	BaseTime            float64     `json:"baseBuildingTime"`
	BuildingTimeFormula string      `json:"buildingTimeFormula"`
	CompiledFormula     FormulaFunc `json:"-"` // Won't be marshaled to JSON
}

// Costs structure encapsulates base costs and cost formulas
type Costs struct {
	BaseCosts       []Cost      `json:"baseCosts"`
	CostFormula     string      `json:"costFormula"`
	CompiledFormula FormulaFunc `json:"-"`
}

// ProductionSpec defines the base production and formulas
type ProductionSpec struct {
	BaseProduction    []Production `json:"baseProduction"`
	ProductionFormula string       `json:"productionFormula"`
	CompiledFormula   FormulaFunc  `json:"-"`
}

// ModifierSpec encapsulates base modifiers and formulas
type ModifierSpec struct {
	BaseModifiers   []Modifier  `json:"baseModifier"`
	ModifierFormula string      `json:"modifierFormula"`
	CompiledFormula FormulaFunc `json:"-"`
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

func CompileFormula(formula string) (FormulaFunc, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return nil, err
	}

	return func(params map[string]float64) float64 {
		// Convert float64 params to interface{} for govaluate
		parameters := make(map[string]interface{}, len(params))
		for k, v := range params {
			parameters[k] = v
		}

		result, err := expression.Evaluate(parameters)
		if err != nil {
			// In production, you might want to handle this differently
			return 0
		}
		return result.(float64)
	}, nil
}

// InitializeFormulas compiles all formulas in the building
func (b *Building) InitializeFormulas() error {
	if b.MetadataMode == "FORMULA" || b.MetadataMode == "MIXED" {
		if b.BuildingTime.BuildingTimeFormula != "" {
			formula, err := CompileFormula(b.BuildingTime.BuildingTimeFormula)
			if err != nil {
				return err
			}
			b.BuildingTime.CompiledFormula = formula
		}
		if b.Costs.CostFormula != "" {
			formula, err := CompileFormula(b.Costs.CostFormula)
			if err != nil {
				return err
			}
			b.Costs.CompiledFormula = formula
		}
		if b.Production.ProductionFormula != "" {
			formula, err := CompileFormula(b.Production.ProductionFormula)
			if err != nil {
				return err
			}
			b.Production.CompiledFormula = formula
		}
		if b.Modifiers.ModifierFormula != "" {
			formula, err := CompileFormula(b.Modifiers.ModifierFormula)
			if err != nil {
				return err
			}
			b.Modifiers.CompiledFormula = formula
		}
		for i := range b.Modifiers.BaseModifiers {
			if b.Modifiers.BaseModifiers[i].ModifierFormula != "" {
				formula, err := CompileFormula(b.Modifiers.BaseModifiers[i].ModifierFormula)
				if err != nil {
					return err
				}
				b.Modifiers.BaseModifiers[i].CompiledFormula = formula
			}
		}
	}

	return nil
}
