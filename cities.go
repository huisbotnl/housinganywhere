package main

import "github.com/jinzhu/gorm"

var citiesUrl = map[string]string{
	"Almelo":                 "/s/Almelo--Netherlands?sorting=mostRecent",
	"Alphen-aan-den-Rijn":    "/s/Alphen-aan-den-Rijn--Netherlands?sorting=mostRecent",
	"Amstelveen":             "/s/Amstelveen--Netherlands?sorting=mostRecent",
	"Amsterdam":              "/s/Amsterdam--Netherlands?sorting=mostRecent",
	"Baarn":                  "/s/Baarn--Netherlands?sorting=mostRecent",
	"Bussum":                 "/s/Bussum--Netherlands?sorting=mostRecent",
	"Capelle-aan-den-IJssel": "/s/Capelle-aan-den-IJssel--Netherlands?sorting=mostRecent",
	"Culemborg":              "/s/Culemborg--Netherlands?sorting=mostRecent",
	"De-Bilt":                "/s/De-Bilt--Netherlands?sorting=mostRecent",
	"Delft":                  "/s/Delft--Netherlands?sorting=mostRecent",
	"Diemen":                 "/s/Diemen--Netherlands?sorting=mostRecent",
	"Dordrecht":              "/s/Dordrecht--Netherlands?sorting=mostRecent",
	"Enschede":               "/s/Enschede--Netherlands?sorting=mostRecent",
	"Gouda":                  "/s/Gouda--Netherlands?sorting=mostRecent",
	"Groningen":              "/s/Groningen--Netherlands?sorting=mostRecent",
	"Haaksbergen":            "/s/Haaksbergen--Netherlands?sorting=mostRecent",
	"Heemskerk":              "/s/Heemskerk--Netherlands?sorting=mostRecent",
	"Heemstede":              "/s/Heemstede--Netherlands?sorting=mostRecent",
	"Hengelo":                "/s/Hengelo--Netherlands?sorting=mostRecent",
	"Hilversum":              "/s/Hilversum--Netherlands?sorting=mostRecent",
	"Houten":                 "/s/Houten--Netherlands?sorting=mostRecent",
	"IJsselstein":            "/s/IJsselstein--Netherlands?sorting=mostRecent",
	"Kerkrade":               "/s/IJsselstein--Netherlands?sorting=mostRecent",
	"Krimpen-aan-den-IJssel": "/s/Krimpen-aan-den-IJssel--Netherlands?sorting=mostRecent",
	"Leiden":                 "/s/Leiden--Netherlands?sorting=mostRecent",
	"Leiderdorp":             "/s/Leiderdorp--Netherlands?sorting=mostRecent",
	"Leusden":                "/s/Leusden--Netherlands?sorting=mostRecent",
	"Maarssen":               "/s/Maarssen--Netherlands?sorting=mostRecent",
	"Maassluis":              "/s/Maassluis--Netherlands?sorting=mostRecent",
	"Mijdrecht":              "/s/Mijdrecht--Netherlands?sorting=mostRecent",
	"Nieuwegein":             "/s/Nieuwegein--Netherlands?sorting=mostRecent",
	"Noordwijk~Binnen":       "/s/Noordwijk~Binnen--Netherlands?sorting=mostRecent",
	"Oldenzaal":              "/s/Oldenzaal--Netherlands?sorting=mostRecent",
	"Oud~Beijerland":         "/s/Oud~Beijerland--Netherlands?sorting=mostRecent",
	"Papendrecht":            "/s/Papendrecht--Netherlands?sorting=mostRecent",
	"Purmerend":              "/s/Purmerend--Netherlands?sorting=mostRecent",
	"Ridderkerk":             "/s/Ridderkerk--Netherlands?sorting=mostRecent",
	"Rijswijk":               "/s/Rijswijk--Netherlands?sorting=mostRecent",
	"Rotterdam":              "/s/Rotterdam--Netherlands?sorting=mostRecent",
	"Scheveningen":           "/s/Scheveningen--Netherlands?sorting=mostRecent",
	"Schiedam":               "/s/Schiedam--Netherlands?sorting=mostRecent",
	"Soest":                  "/s/Soest--Netherlands?sorting=mostRecent",
	"Spijkenisse":            "/s/Spijkenisse--Netherlands?sorting=mostRecent",
	"The-Hague":              "/s/The-Hague--Netherlands?sorting=mostRecent",
	"Uithoorn":               "/s/Uithoorn--Netherlands?sorting=mostRecent",
	"Utrecht":                "/s/Utrecht--Netherlands?sorting=mostRecent",
	"Velsen":                 "/s/Velsen--Netherlands?sorting=mostRecent",
	"Vlaardingen":            "/s/Vlaardingen--Netherlands?sorting=mostRecent",
	"Voorburg":               "/s/Voorburg--Netherlands?sorting=mostRecent",
	"Waddinxveen":            "/s/Waddinxveen--Netherlands?sorting=mostRecent",
	"Wassenaar":              "/s/Wassenaar--Netherlands?sorting=mostRecent",
	"Wijk-bij-Duurstede":     "/s/Wijk-bij-Duurstede--Netherlands?sorting=mostRecent",
	"Woerden":                "/s/Woerden--Netherlands?sorting=mostRecent",
	"Zaandam":                "/s/Zaandam--Netherlands?sorting=mostRecent",
	"Zaanstad":               "/s/Zaanstad--Netherlands?sorting=mostRecent",
	"Zeist":                  "/s/Zeist--Netherlands?sorting=mostRecent",
}

var citiesMap = map[string]map[string]interface{}{
	"Alkmaar": {
		"key": "alkmaar",
	},
	"Almere-Stad": {
		"key": "almere_stad",
	},
	"Alphen aan den Rijn": {
		"key": "alphen_aan_den_rijn",
	},
	"Amersfoort": {
		"key": "amersfoort",
	},
	"Amstelveen": {
		"key": "amstelveen",
	},
	"Amsterdam": {
		"key": "amsterdam",
	},
	"Bleiswijk": {
		"key": "bleiswijk",
	},
	"Capelle aan den IJssel": {
		"key": "capelle_aan_den_i_jssel",
	},
	"Delft": {
		"key": "delft",
	},
	"Diemen": {
		"key": "diemen",
	},
	"Driebergen-Rijsenburg": {
		"key": "driebergen_rijsenburg",
	},
	"Egmond aan den Hoef": {
		"key": "egmond_aan_den_hoef",
	},
	"Enschede": {
		"key": "enschede",
	},
	"Groningen": {
		"key": "groningen",
	},
	"Haarlem": {
		"key": "haarlem",
	},
	"Haren": {
		"key": "haren",
	},
	"Hazerswoude-Rijndijk": {
		"key": "hazerswoude_rijndijk",
	},
	"Hengelo": {
		"key": "hengelo",
	},
	"Hillegom": {
		"key": "hillegom",
	},
	"Hilversum": {
		"key": "hilversum",
	},
	"Hoofddorp": {
		"key": "hoofddorp",
	},
	"Leiden": {
		"key": "leiden",
	},
	"Maarssen": {
		"key": "maarssen",
	},
	"Nieuw-Vennep": {
		"key": "nieuw_vennep",
	},
	"Nieuwveen": {
		"key": "nieuwveen",
	},
	"Ouderkerk aan de Amstel": {
		"key": "ouderkerk_aan_de_amstel",
	},
	"Putten": {
		"key": "putten",
	},
	"Rijswijk": {
		"key": "rijswijk",
	},
	"Rotterdam": {
		"key": "rotterdam",
	},
	"Scheveningen": {
		"key": "scheveningen",
	},
	"Schiedam": {
		"key": "schiedam",
	},
	"The Hague": {
		"key": "the_hague",
	},
	"Utrecht": {
		"key": "utrecht",
	},
	"Veenendaal": {
		"key": "veenendaal",
	},
	"Voorburg": {
		"key": "voorburg",
	},
	"Voorschoten": {
		"key": "voorschoten",
	},
	"Nieuwegein": {
		"key": "nieuwegein",
	},
	"Warmond": {
		"key": "warmond",
	},
	"Zaandam": {
		"key": "zaandam",
	},
	"Zeist": {
		"key": "zeist",
	},
	"Breda": {
		"key": "breda",
	},
	"Zoetermeer": {
		"key": "zoetermeer",
	},
	"Almere": {
		"key": "almere",
	},
	"Maastricht": {
		"key": "maastricht",
	},
	"Vlaardingen": {
		"key": "vlaardingen",
	},
}

func setupCitiesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range citiesMap {
		DB.Raw("SELECT * FROM `cities`  WHERE (`key` = ?) ORDER BY `cities`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		citiesMap[name] = data
	}
}

type City struct {
	gorm.Model
	Key            string `json:"key" gorm:"type:varchar(50);index;unique"`
	Name           string `json:"name" gorm:"type:varchar(50);index;unique"`
	Neighbourhoods []Neighbourhood
}
