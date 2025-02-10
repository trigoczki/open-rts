package ruleset

import (
	"encoding/json"
	"os"
)

func ParseResources(path string) ([]*Resource, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var resources []*Resource
	err = json.NewDecoder(jsonFile).Decode(&resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
