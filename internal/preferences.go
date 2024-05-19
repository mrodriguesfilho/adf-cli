package internal

import (
	"encoding/json"
)

type Preferences struct {
	Services map[string]ServiceData
	Version  string
}

type ServiceData struct {
	Version     string
	DownloadUrl string
	FileName    string
}

var LoadedPreferences Preferences

var staticServiceDataArr = map[string]ServiceData{
	"adfweb":      {"0.0.1", "http://localhost:5000/downloadadf", ""},
	"jvm:darwin":  {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz"},
	"jvm:linux":   {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz"},
	"jvm:windows": {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz"},
}

const PreferencesVersion = "0.0.2"

var staticPreferences = Preferences{
	Services: staticServiceDataArr,
	Version:  PreferencesVersion,
}

func GetStaticServiceDataAsJson() (string, error) {
	jsonData, err := json.MarshalIndent(staticPreferences, "", "    ")

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
