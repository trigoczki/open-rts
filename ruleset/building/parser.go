package building

import (
	"encoding/json"
	"os"
)

func ParseBuildings(path string) ([]*Building, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	var buildings []*Building
	err = json.NewDecoder(jsonFile).Decode(&buildings)
	if err != nil {
		return nil, err
	}
	for _, building := range buildings {
		if err := building.InitializeFormulas(); err != nil {
			return nil, err
		}
	}

	return buildings, nil
}
