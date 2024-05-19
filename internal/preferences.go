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
	Installed   bool
}

var LoadedPreferences Preferences
var AdfDirectory string

const (
	AdfVersion         = "0.0.1"
	DefaultWebPort int = 5263

	RepositoryServerAddress       string = "localhost:5263"
	ServiceDatacCollectionAddress string = "https://raw.githubusercontent.com/mrodriguesfilho/adf-cli/main/preferences.json"
	AdfDefaultDir                        = ".adf"
	AdfPreferencesFileName               = "preferences.json"
)

var staticServiceDataArr = map[string]ServiceData{
	"jvm:darwin":  {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_aarch64_mac_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_aarch64_mac_hotspot_21.0.3_9.tar.gz", false},
	"jvm:linux":   {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", false},
	"jvm:windows": {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_windows_hotspot_21.0.3_9.zip", "OpenJDK21U-jre_x64_windows_hotspot_21.0.3_9.zip", false},
}

const PreferencesBuiltInVersion = "0.0.1"

var staticPreferences = Preferences{
	Services: staticServiceDataArr,
	Version:  PreferencesBuiltInVersion,
}

func GetStaticServiceDataAsJson() (string, error) {
	jsonData, err := json.MarshalIndent(staticPreferences, "", "    ")

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
