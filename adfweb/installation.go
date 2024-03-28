package adfweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
