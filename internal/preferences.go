package internal

import (
	"encoding/json"
)

type Preferences struct {
	Services []ServiceData
	Version  string
}

type ServiceData struct {
	Name        string
	Version     string
	DownloadUrl string
}

var staticServiceDataArr = []ServiceData{
	{"adfweb", "0.0.1", "http://localhost:5000/downloadadf"},
	{"jvm", "12.0.0", "http://localhost:5000/downloadjvm"},
}

const PreferencesVersion = "0.0.1"

var preference = Preferences{
	Services: staticServiceDataArr,
	Version:  PreferencesVersion,
}

func GetStaticServiceDataAsJson() (string, error) {
	jsonData, err := json.MarshalIndent(preference, "", "    ")

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
