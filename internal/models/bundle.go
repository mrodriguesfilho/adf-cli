package models

import "strings"

type Bundle struct {
	Services map[string]ServiceData
	Version  string
	InUse    bool
}

type ServiceData struct {
	Version     string
	DownloadUrl string
	FileName    string
	Installed   bool
}

func (b Bundle) Validate() bool {
	for _, service := range b.Services {

		if service.DownloadUrl == "" || !strings.HasPrefix(service.DownloadUrl, "http://") || !strings.HasPrefix(service.DownloadUrl, "https://") {
			return false
		}

	}

	return true
}
