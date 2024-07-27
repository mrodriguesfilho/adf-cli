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
	"adf-cli/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adf",
	Short: "Ferramenta administrativa do Ambiente de Design FHIR (ADF)",
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
	preferences, err := readLocalPreferencesFile()

	if err != nil {
		createPreferencesFile()
	}

	models.LoadedPreferences = preferences

	validateLatestPreferences(preferences)
}

func createApplicationFolder() {

	if os.Getenv("ADF_HOME") != "" {
		models.AdfDirectory = os.Getenv("ADF_HOME")
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		models.AdfDirectory = filepath.Join(home, models.AdfDefaultDir)
	}

	_, err := os.Stat(models.AdfDirectory)

	if os.IsNotExist(err) {
		err := os.Mkdir(models.AdfDirectory, 0700)
		cobra.CheckErr(err)
	}
}

func readLocalPreferencesFile() (models.Preferences, error) {
	viper.AddConfigPath(models.AdfDirectory)
	viper.SetConfigName("preferences")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		return models.Preferences{}, err
	}

	var preferences models.Preferences
	if err := viper.Unmarshal(&preferences); err != nil {
		return models.Preferences{}, err
	}

	err = ValidateBundles(preferences)
	if err != nil {
		return models.Preferences{}, err
	}

	return preferences, nil
}

func createPreferencesFile() {

	adfPreferencesFilePath := filepath.Join(models.AdfDirectory, models.AdfPreferencesFileName)

	if _, err := os.Stat(adfPreferencesFilePath); err == nil {
		if err := os.Remove(adfPreferencesFilePath); err != nil {
			cobra.CheckErr(fmt.Errorf("failed to delete existing preferences file: %w", err))
			return
		}
	}

	serviceDataArr, err := downloadLatestServiceDataFile()

	if err != nil {
		fmt.Println("Não foi possível baixar a lista das versões mais atualizadas do serviço.")
		fmt.Printf("Utilizando as versões built-in da versão %v do ADF \n", models.PreferencesLatestVersion)
		serviceDataArr, _ = models.GetStaticServiceDataAsJson()
	}

	err = os.WriteFile(adfPreferencesFilePath, []byte(serviceDataArr), 0644)

	if err != nil {
		cobra.CheckErr(err)
		fmt.Printf("Não foi possível salvar o arquivo JSON com os dados de serviço. %v\n", err)
	}

	if _, err = readLocalPreferencesFile(); err != nil {
		fmt.Println("Não foi possível ler o arquivo preferences")
	}
}

func downloadLatestServiceDataFile() (string, error) {

	httpClient := &http.Client{}

	res, err := httpClient.Get(models.ServiceDatacCollectionAddress)

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

func ValidateBundles(preferences models.Preferences) error {

	for _, bundle := range preferences.Bundles {
		if !bundle.Validate() {
			return errors.New("bundle has invalid data in json file")
		}
	}

	return nil
}

func validateLatestPreferences(preferences models.Preferences) {

	serviceDataArr, err := downloadLatestServiceDataFile()

	if err != nil {
		return
	}

	var latestPreferences models.Preferences
	json.Unmarshal([]byte(serviceDataArr), &latestPreferences)

	if versionGreaterThan(latestPreferences.LatestVersion, preferences.LatestVersion) {
		createPreferencesFile()
		models.LoadedPreferences = latestPreferences
	}
}

func versionGreaterThan(v1, v2 string) bool {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	if len(v1Parts) != 3 || len(v2Parts) != 3 {
		return true
	}

	for i := 0; i < 3; i++ {
		v1Int, err1 := strconv.Atoi(v1Parts[i])
		v2Int, err2 := strconv.Atoi(v2Parts[i])

		if err1 != nil || err2 != nil {
			return false
		}

		if v1Int > v2Int {
			return true
		} else if v1Int < v2Int {
			return false
		}
	}

	return false
}
