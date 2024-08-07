package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRemoteVersionsForDownload(
	repositoryServerAddress string,
) ([]string, error) {
	remoteVersions := []string{}

	resp, err := http.Get(
		fmt.Sprintf(
			"http://%s/api/adfweb-version/",
			repositoryServerAddress,
		),
	)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&remoteVersions)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return remoteVersions, nil
}
