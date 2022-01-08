package main

var bedroomsMap = map[string]map[string]interface{}{
	"Studio": {
		"key": "1",
	},
	"1 bedroom": {
		"key": "1",
	},
	"2 bedrooms": {
		"key": "2",
	},
	"3 bedrooms": {
		"key": "3",
	},
	"4 bedrooms": {
		"key": "4",
	},
	"5 bedrooms": {
		"key": "5",
	},
	"6 bedrooms": {
		"key": "6",
	},
}

func setupBedroomsMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range bedroomsMap {
		DB.Raw("SELECT * FROM `bedrooms`  WHERE (`key` = ?) ORDER BY `bedrooms`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		bedroomsMap[name] = data
	}
}
