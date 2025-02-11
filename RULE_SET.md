## Building Ruleset

The building ruleset defines the properties and behavior of different buildings in the game. Each building entry contains basic information like ID, name, description, and maximum level, along with more complex attributes that determine its costs, production, and modifiers.

### Metadata Modes

Buildings can operate in three different metadata modes that determine how their attributes are calculated at different levels:

#### 1. FORMULA Mode
When `metadataMode` is set to "FORMULA", building attributes are calculated using mathematical formulas. This mode uses:
- Base values (like `baseBuildingTime`, `baseCosts`, `baseProduction`)
- Growth factor
- Various formulas (e.g., `buildingTimeFormula`, `costFormula`, `productionFormula`)

Example from Farm building:

```json
{
"buildingTime": {
"baseBuildingTime": 10,
"buildingTimeFormula": "baseBuildingTime (buildingLevel)^growthFactor (1 - buildTimeModifier)"
}
}
```


#### 2. CONCRETE Mode
When `metadataMode` is set to "CONCRETE", all values are explicitly defined for each level in the metadata object. No calculations are performed - the values are used exactly as specified.

Example from Iron Mine building:
```json
{
"metadata": {
"1": {
"baseBuildingTime": 10,
"baseCost": [...],
"baseProduction": [...],
"baseModifier": [...]
},
"2": {
"baseBuildingTime": 20,
"baseCost": [...],
"baseProduction": [...],
"baseModifier": [...]
}
}
}
```


#### 3. MIXED Mode
When `metadataMode` is set to "MIXED", the building uses a combination of formula-based and concrete values. It primarily works like FORMULA mode, but for levels that are explicitly defined in the metadata object, those concrete values take precedence over calculated values.

For example, in the Stone Quarry building:
- Most levels use the formula calculations
- Level 5 uses explicitly defined values from the metadata object
- No calculations are attempted for level 5; all values come directly from the metadata configuration

Example:

```json
{
"metadataMode": "MIXED",
"buildingTime": {
"baseBuildingTime": 10,
"buildingTimeFormula": "baseBuildingTime (buildingLevel)^growthFactor (1 - buildTimeModifier)"
},
"metadata": {
"5": {
"baseBuildingTime": 50,
"baseCost": [...],
"baseProduction": [...],
"baseModifier": [...]
}
}
}
```

In this example, level 5 uses the explicit values from metadata, while all other levels use formula calculations.

### Common Attributes

Each building definition includes:
- Basic information (id, name, description, maxLevel)
- Building time configuration
- Resource costs
- Production rates
- Modifiers (both local and global effects)
- Growth factor (for formula calculations)

### Growth Factor

The growth factor is a multiplier that determines how the building's attributes grow as the level increases. It is used in formula calculations to ensure that the building's attributes increase at a reasonable rate.

