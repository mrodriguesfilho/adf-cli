package models

import (
	"encoding/json"
)

type Preferences struct {
	Bundles []Bundle
}

var LoadedPreferences Preferences
var LoadedBundle Bundle
var AdfDirectory string

const (
	AdfReleaseVersion             = "0.0.6"
	PreferencesBuiltInVersion     = "0.0.1"
	HapiFhirDefaultPort           = ":8080"
	ServiceDatacCollectionAddress = "https://raw.githubusercontent.com/mrodriguesfilho/adf-cli/main/preferences.json"
	AdfPreferencesFileName        = "preferences.json"
	AdfDefaultDir                 = ".adf"
)

var staticServiceDataArr = map[string]ServiceData{
	"jvm:darwin":         {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_aarch64_mac_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_aarch64_mac_hotspot_21.0.3_9.tar.gz", false},
	"jvm:linux":          {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", "OpenJDK21U-jre_x64_linux_hotspot_21.0.3_9.tar.gz", false},
	"jvm:windows":        {"21.0.3", "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.3%2B9/OpenJDK21U-jre_x64_windows_hotspot_21.0.3_9.zip", "OpenJDK21U-jre_x64_windows_hotspot_21.0.3_9.zip", false},
	"hapifhir":           {"7.0.2", "https://github.com/mrodriguesfilho/adf-cli/releases/download/v.0.2.0-beta/hapifhir.war", "hapifhir.war", false},
	"hapifhir-validator": {"6.13.16", "https://github.com/hapifhir/org.hl7.fhir.core/releases/latest/download/validator_cli.jar", "validator_cli.jar", false},
}

var staticPreferences = Preferences{
	Bundles: []Bundle{
		{Services: staticServiceDataArr, Version: PreferencesBuiltInVersion},
	},
}

func GetStaticServiceDataAsJson() (string, error) {
	jsonData, err := json.MarshalIndent(staticPreferences, "", "    ")

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
