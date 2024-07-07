package services

import (
	"adf-cli/internal/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/schollz/progressbar/v3"
)

var JvmProgressBarOptions = []progressbar.Option{
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

func InstallJVM(installDir string, installVersion string, bundleToUse models.Bundle) error {

	jvmPreferenceData := bundleToUse.Services["jvm:"+runtime.GOOS]
	err := downloadJVM(installDir, jvmPreferenceData)

	if err != nil {
		return err
	}

	destinationFilePath := installDir + "/jvm/" + jvmPreferenceData.Version
	saveFilePath := installDir + "/jvm/" + jvmPreferenceData.Version + "/" + jvmPreferenceData.FileName

	err = extractFile(saveFilePath, destinationFilePath, JvmProgressBarOptions)

	if err != nil {
		return err
	}

	// var javaPath string
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	javaPath = destinationFilePath + "/Contents/Home/bin/java"
	// case "linux":
	// 	javaPath = destinationFilePath + "/bin/java"
	// case "windows":
	// 	javaPath = destinationFilePath + "/bin/java.exe"
	// default:
	// 	fmt.Printf("Operating System: %s\n not found", os)
	// }

	// cmd := exec.Command(javaPath, "-version")

	// output, err := cmd.CombinedOutput()

	// if err != nil {
	// 	return err
	// }

	// fmt.Println(string(output))
	return nil
}

func downloadJVM(installDir string, jvmPreferencesData models.ServiceData) error {

	res, err := http.Get(jvmPreferencesData.DownloadUrl)

	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request to download JVM failed with status %d", res.StatusCode)
	}

	saveFilePath := installDir + "/jvm/" + jvmPreferencesData.Version

	err = os.MkdirAll(saveFilePath, os.ModePerm)
	if err != nil {
		return err
	}

	saveFilePath += "/" + jvmPreferencesData.FileName

	file, err := os.Create(saveFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	bar := progressbar.NewOptions64(res.ContentLength, JvmProgressBarOptions...)
	bar.Describe("Baixando a JVM versão: " + jvmPreferencesData.Version)

	_, err = io.Copy(io.MultiWriter(file, bar), res.Body)
	return err
}
