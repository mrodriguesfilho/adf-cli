/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"adf-cli/internal"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Default config parameters values
const (
	adfVersion         = "0.0.1"
	defaultWebPort int = 5263

	RepositoryServerAddress       string = "localhost:5263"
	ServiceDatacCollectionAddress string = "https://raw.githubusercontent.com/mrodriguesfilho/adf-cli/main/preferences.json"
	adfDefaultDir                        = ".adf"
	adfPreferencesFileName               = "preferences.json"
)

var (
	installedVersions []string
	usedVersion       string
)

// Config parameters
var webPort int = defaultWebPort
var adfDirectory string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adf",
	Short: "Ferramenta de linha de comando para funcionalidades administrativas do Ambiente de Design FIHR (ADF)",
	Long:  ``,
	Example: `
	$ adf install --version 0.0.1
	Instalando versão 0.0.1 do ADF Web...
	Versão 0.0.1 do ADF Web instalada com sucesso

	$ adf list
	Versões instaladas:
	ADF Web 0.0.1

	$ adf use 0.0.1
	Definida a versão 0.0.1 do ADF Web a ser utilizada

	$ adf list
	ADF Web 0.0.1 - em uso`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	createApplicationFolder()
	err := readPreferencesFile()
	if err != nil {
		createPreferencesFile()
	}
}

func createApplicationFolder() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	adfDirectory = filepath.Join(home, adfDefaultDir)
	_, err = os.Stat(adfDirectory)

	if os.IsNotExist(err) {
		err := os.Mkdir(adfDirectory, 0700)
		cobra.CheckErr(err)
	}
}

func readPreferencesFile() error {
	viper.SetConfigName("preferences")
	viper.AddConfigPath(adfDefaultDir)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	if err := viper.Unmarshal(&internal.LoadedPreferences); err != nil {
		return err
	}

	return nil
}

func createPreferencesFile() {

	adfPreferencesFilePath := filepath.Join(adfDirectory, adfPreferencesFileName)

	if _, err := os.Stat(adfPreferencesFilePath); err == nil {
		cobra.CheckErr(err)
		return
	}

	serviceDataArr, err := downloadLatestServiceDataFile()

	if err != nil {
		fmt.Println("Não foi possível baixar a lista das versões mais atualizadas do serviço.")
		fmt.Println("Utilizando as versões built-in")
		serviceDataArr, _ = internal.GetStaticServiceDataAsJson()
	}

	err = os.WriteFile(adfPreferencesFilePath, []byte(serviceDataArr), 0644)

	if err != nil {
		cobra.CheckErr(err)
		fmt.Printf("Não foi possível salvar o arquivo JSON com os dados de serviço. %v\n", err)
	}
}

func downloadLatestServiceDataFile() (string, error) {

	httpClient := &http.Client{}

	res, err := httpClient.Get(ServiceDatacCollectionAddress)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
