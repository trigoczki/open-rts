package cache

import (
	"log"
	"open-rts/ruleset/building"
	rulesetResource "open-rts/ruleset/resource"
)

func Initialize() {
	resources, err := rulesetResource.ParseResources("ruleset/resource/resources_ruleset.json")
	if err != nil {
		panic(err)
	}

	buildings, err := building.ParseBuildings("ruleset/building/buildings_ruleset.json")
	if err != nil {
		panic(err)
	}

	cache := GetInstance()
	cache.SetResources(resources)
	log.Println("Resources set in cache")
	cache.SetBuildings(buildings)
	log.Println("Buildings set in cache")
	log.Println("Cache initialized successfully")
}
