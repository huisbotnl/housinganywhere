package main

var typeMap = map[string]map[string]interface{}{
	"Private room in house": {
		"key": "private_room",
	},
	"Shared room in house": {
		"key": "shared_room",
	},
	"Entire apartment": {
		"key": "apartment",
	},
	"Studio": {
		"key": "studio",
	},
	"Entire house": {
		"key": "studio",
	},
}

func setupTypesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range typeMap {
		DB.Raw("SELECT * FROM `property_types`  WHERE (`key` = ?) ORDER BY `property_types`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		typeMap[name] = data
	}
}
