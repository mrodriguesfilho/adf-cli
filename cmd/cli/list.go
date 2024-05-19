/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/internal"
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
		serviceList := getList()
		for _, serviceEntry := range serviceList {
			fmt.Printf("Service: %s | Version: %s \n", serviceEntry.Name, serviceEntry.Version)
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

func getList() []internal.ServiceData {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	adfPreferenceFilePath = filepath.Join(home, adfDefaultDir, adfPreferencesFileName)
	_, err = os.Stat(adfPreferenceFilePath)
	if err != nil {
		cobra.CheckErr(err)
		return nil
	}

	serviceCollectionJsonData, err := os.ReadFile(adfPreferenceFilePath)
	if err != nil {
		cobra.CheckErr(err)
		return nil
	}

	var preferences internal.Preferences
	if err := json.Unmarshal(serviceCollectionJsonData, &preferences); err != nil {
		cobra.CheckErr(err)
		return nil
	}

	return preferences.Services
}
