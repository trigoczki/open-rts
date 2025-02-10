package cache

import (
	ruleset "open-rts/ruleset/resource"
)

func Initialize() {
	resources, err := ruleset.ParseResources("ruleset/resource/resources_ruleset.json")
	if err != nil {
		panic(err)
	}

	cache := GetInstance()
	cache.SetResources(resources)
}
