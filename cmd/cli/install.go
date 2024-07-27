/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/internal/models"
	"adf-cli/internal/services"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installVersion string
var installDir string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Instala uma versão do ADF Web",
	Long:  ``,
	Example: `
	$ adf install --version 0.0.1	
	Instalando versão 0.0.1 do ADF Web...
	Versão 0.0.1 do ADF Web instalada com sucesso
	`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&installVersion, "version", "v", "", "Versão da aplicação de design do ADF")
	installCmd.Flags().StringVarP(&installDir, "output", "o", "", "Diretório personalizado de instalação")

	// installCmd.Flags().StringVarP(&installVersion, "", "", "latest", "Versão da aplicação de design do ADF")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func execute() {

	if installVersion == "" {
		installVersion = models.PreferencesLatestVersion
	}

	var bundle = GetBundleVersion(installVersion)

	if installDir == "" {
		installDir = models.AdfDirectory
	}

	installDir = filepath.Join(models.AdfDirectory, installVersion)

	installStrategies := []services.InstallStrategy{
		services.JVM{},
		services.HAPIFHIR{},
		services.HapifhirValidator{},
	}

	for _, strategy := range installStrategies {
		err := strategy.Install(installDir, installVersion, bundle)
		if err != nil {
			fmt.Printf("Não foi possível fazer a instalação %v. Erro: %v\n", strategy.ServiceName(), err)
			return
		}
	}

	err := WriteToReferences(installVersion, installDir)
	if err != nil {
		fmt.Printf("falha ao escrever o novo arquivo de referencias. Erro: %v", err)
	}
}

func GetBundleVersion(installVersion string) models.Bundle {

	for i := 0; i < len(models.LoadedPreferences.Bundles); i++ {
		if installVersion == models.LoadedPreferences.Bundles[i].Version {
			return models.LoadedPreferences.Bundles[i]
		}
	}

	return models.Bundle{}
}

func WriteToReferences(installVersion, installedDir string) error {

	referencesData, err := os.ReadFile(models.AdfDirectory)

	var references models.References
	if err != nil {
		newFile, err := os.OpenFile(filepath.Join(models.AdfDirectory, models.ReferenceFileName), os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return err
		}

		defer newFile.Close()

		references = models.NewReference(installVersion, installedDir)
	} else {
		references, err = UpdateReferenceFile(referencesData, installVersion, installedDir)

		if err != nil {
			return err
		}
	}

	newReferenceData, err := json.Marshal(references)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(models.AdfDirectory, models.ReferenceFileName), newReferenceData, 0644)

	if err != nil {
		return err
	}

	return nil
}

func UpdateReferenceFile(referenceData []byte, installedVersion, installedDir string) (models.References, error) {

	var references models.References
	if err := json.Unmarshal(referenceData, &references); err != nil {
		return models.References{}, err
	}

	if alreadyExits, index := ReferenceAlreadyExists(references.InstalledBundles); alreadyExits {
		references.InstalledBundles[index].DirectoryPath = installedDir
	} else {
		newInstallation := models.BundleInstalled{Version: installedVersion, DirectoryPath: installedDir}
		references.InstalledBundles = append(references.InstalledBundles, newInstallation)
	}

	return references, nil
}

func ReferenceAlreadyExists(installedBundles []models.BundleInstalled) (bool, int) {

	for i := 0; i < len(installedBundles); i++ {
		if installVersion == installedBundles[i].Version {
			return true, i
		}
	}

	return false, 0
}
