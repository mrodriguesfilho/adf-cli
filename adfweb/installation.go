package adfweb

import (
	"adf-cli/internal"
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

type ADFWebVersion struct {
	Version string
	Used    bool
}

var ProgressBarOptions = []progressbar.Option{
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

func InstallADFWeb(
	repositoryServerAddress string,
	version string,
) error {
	err := downloadADFWeb(repositoryServerAddress, version)
	if err != nil {
		return err
	}

	downloadFilePath := fmt.Sprintf("adfweb-%s", version)

	return extractZipFile(downloadFilePath+".zip", downloadFilePath)
}

func InstallJVM() error {

	jvmPreferenceData := internal.LoadedPreferences.Services["jvm:"+runtime.GOOS]

	fmt.Printf("Download JVM default version (%v)? y/n \n", jvmPreferenceData.Version)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "n" {
		availableVersions, _ := listAvailableJvmVersions()

		fmt.Println("Pick from the available versions:")
		for _, v := range availableVersions {
			fmt.Println("-", v)
		}

		fmt.Println("Select version:")
		versionInput, _ := reader.ReadString('\n')
		versionInput = strings.TrimSpace(versionInput)
	}

	err := downloadJVM(jvmPreferenceData)
	if err != nil {
		return err
	}

	destinationFilePath := internal.AdfDirectory + "/jvm/" + jvmPreferenceData.Version
	saveFilePath := internal.AdfDirectory + "/jvm/" + jvmPreferenceData.Version + "/" + jvmPreferenceData.FileName

	err = extractFile(saveFilePath, destinationFilePath)

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

func downloadADFWeb(
	repositoryServerAddress string,
	version string,
) error {
	res, err := http.Get(
		fmt.Sprintf(
			"http://%s/static/adfweb/adfweb-%s.zip",
			repositoryServerAddress, version,
		),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request to download ADF failed with status %d", res.StatusCode)
	}

	f, _ := os.OpenFile(fmt.Sprintf("adfweb-%s.zip", version), os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.NewOptions64(res.ContentLength, ProgressBarOptions...)
	bar.Describe("Baixando ADF Web " + version)

	_, err = io.Copy(io.MultiWriter(f, bar), res.Body)
	return err
}

func listAvailableJvmVersions() ([]string, error) {
	return []string{"8.0, 11.0, 12.0"}, nil
}

func downloadJVM(jvmPreferencesData internal.ServiceData) error {

	res, err := http.Get(jvmPreferencesData.DownloadUrl)

	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request to download JVM failed with status %d", res.StatusCode)
	}

	saveFilePath := internal.AdfDirectory + "/jvm/" + jvmPreferencesData.Version

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

	bar := progressbar.NewOptions64(res.ContentLength, ProgressBarOptions...)
	bar.Describe("Baixando a JVM versão: " + jvmPreferencesData.Version)

	_, err = io.Copy(io.MultiWriter(file, bar), res.Body)
	return err
}
