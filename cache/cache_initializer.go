package cache

import (
	"log"
	"open-rts/ruleset"
	rulesetResource "open-rts/ruleset/resource"
)

func Initialize() {
	resources, err := rulesetResource.ParseResources("ruleset/resource/resources_ruleset.json")
	if err != nil {
		panic(err)
	}

	err = ruleset.ValidateJSON(resources, "ruleset/resource/resource_schema.json")
	if err != nil {
		panic(err)
	}
	log.Println("Resources validated successfully")

	cache := GetInstance()
	cache.SetResources(resources)
	log.Println("Resources set in cache")
	log.Println("Cache initialized successfully")
}
