package main

var facilitiesMap = map[string]map[string]interface{}{
	"Private bathroom": {
		"key": "private_bathroom",
	},

	"Private kitchen": {
		"key": "kitchen",
	},
	"Shared kitchen": {
		"key": "kitchen",
	},

	"Private garden": {
		"key": "garden",
	},
	"Shared garden": {
		"key": "garden",
	},

	"Private balcony": {
		"key": "balcony",
	},
	"Shared balcony": {
		"key": "balcony",
	},

	"Private basement": {
		"key": "basement",
	},
	"Shared basement": {
		"key": "basement",
	},

	"Private parking": {
		"key": "parking",
	},
	"Shared parking": {
		"key": "parking",
	},

	"Pets negotiable": {
		"key": "pets_allowed",
	},
}

func setupFacilitiesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range facilitiesMap {
		DB.Raw("SELECT * FROM `facilities`  WHERE (`key` = ?) ORDER BY `facilities`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		facilitiesMap[name] = data
	}
}
