package adfweb

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	repositoryServerPort int,
	version string,
) error {
	err := downloadADFWeb(repositoryServerAddress, repositoryServerPort, version)
	if err != nil {
		return err
	}

	downloadFilePath := fmt.Sprintf("adfweb-%s", version)

	return extractZipFile(downloadFilePath+".zip", downloadFilePath)
}

func InstallJVM(
	repositoryServerAddress string,
	repositoryServerPort int,
	version string,
) error {
	fmt.Sprintf("Download JVM default version (%s)? y/n", version)
	fmt.Println("Download JVM default version (%s)? y/n", version)
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
		version = versionInput
	}

	err := downloadJVM(repositoryServerAddress, repositoryServerPort, version)
	if err != nil {
		return err
	}

	downloadFilePath := fmt.Sprintf("jvm/%s/", version)

	return extractZipFile(downloadFilePath+".zip", downloadFilePath)
}

func downloadADFWeb(
	repositoryServerAddress string,
	repositoryServerPort int,
	version string,
) error {
	res, err := http.Get(
		fmt.Sprintf(
			"http://%s:%d/static/adfweb/adfweb-%s.zip",
			repositoryServerAddress, repositoryServerPort, version,
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

func downloadJVM(
	repositoryServerAddress string,
	repositoryServerPort int,
	versionJVM string,
) error {
	res, err := http.Get(
		fmt.Sprintf(
			"http://%s:%d/static/jvm/jvm-%s.zip",
			repositoryServerAddress, repositoryServerPort, versionJVM,
		),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request to download JVM failed with status %d", res.StatusCode)
	}

	f, _ := os.OpenFile(fmt.Sprintf("jvm-%s.zip", versionJVM), os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.NewOptions64(res.ContentLength, ProgressBarOptions...)
	bar.Describe("Baixando a JVM versão: " + versionJVM)

	_, err = io.Copy(io.MultiWriter(f, bar), res.Body)
	return err
}
