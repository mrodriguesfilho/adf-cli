package services

import (
	"adf-cli/internal/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

var hapiFhirValidatorProgressBarOptions = []progressbar.Option{
	progressbar.OptionSetWriter(os.Stderr),
	progressbar.OptionShowBytes(true),
	progressbar.OptionSetWidth(10),
	progressbar.OptionThrottle(65 * time.Millisecond),
	progressbar.OptionShowCount(),
	progressbar.OptionOnCompletion(func() {
		fmt.Fprint(os.Stderr, "\n")
	}),
	progressbar.OptionSpinnerType(14),
	progressbar.OptionFullWidth(),
	progressbar.OptionSetRenderBlankState(true),
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}),
}

type HapifhirValidator struct{}

func (h HapifhirValidator) Install(installDir string, installVersion string, bundleToUse models.Bundle) error {

	hapifhirServiceData := bundleToUse.Services["hapifhir-validator"]
	err := downloadHapifhirValidator(installDir, hapifhirServiceData)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (h HapifhirValidator) ServiceName() string {
	return "HapiFhirValidator"
}

func downloadHapifhirValidator(installDir string, hapifhirePreferencesData models.ServiceData) error {

	req, _ := http.NewRequest("GET", hapifhirePreferencesData.DownloadUrl, nil)
	req.Header.Add("Accept", "application/octet-stream")
	res, err := http.Get(hapifhirePreferencesData.DownloadUrl)

	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request to download HAPIFHIR-VALIDTOR failed with status %d", res.StatusCode)
	}

	saveFilePath := installDir + "/hapifhir-validator/"

	err = os.MkdirAll(saveFilePath, os.ModePerm)
	if err != nil {
		return err
	}

	saveFilePath += hapifhirePreferencesData.FileName

	file, err := os.Create(saveFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	bar := progressbar.NewOptions64(res.ContentLength, hapiFhirValidatorProgressBarOptions...)
	bar.Describe("Baixando o HAPIFHIR-Validator versão: " + hapifhirePreferencesData.Version)

	_, err = io.Copy(io.MultiWriter(file, bar), res.Body)
	return err
}
