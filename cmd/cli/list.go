/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Exibe as versões disponíveis dos serviços relacinados ao ADF",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bundleArr := getBundles()
		for _, bundle := range bundleArr {
			fmt.Println("-------")
			fmt.Printf("Service Data Version: %s \n", bundle.Version)
			fmt.Printf("In Use: %v \n", bundle.InUse)
			for key, serviceEntry := range bundle.Services {
				fmt.Printf("Service: %s | Version: %s \n", key, serviceEntry.Version)
			}
			fmt.Println("-------")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getBundles() []models.Bundle {

	models.AdfDirectory = filepath.Join(models.AdfDirectory, models.AdfPreferencesFileName)
	_, err := os.Stat(models.AdfDirectory)
	if err != nil {
		cobra.CheckErr(err)
		return nil
	}

	serviceCollectionJsonData, err := os.ReadFile(models.AdfDirectory)
	if err != nil {
		cobra.CheckErr(err)
		return nil
	}

	var preferences models.Preferences
	if err := json.Unmarshal(serviceCollectionJsonData, &preferences); err != nil {
		cobra.CheckErr(err)
		return nil
	}

	return preferences.Bundles
}
