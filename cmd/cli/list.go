/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/internal"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
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

	serviceDataArr, err := downloadJson()

	if err != nil {
		fmt.Println("Não foi possível baixar a lista das versões mais atualizadas do serviço.")
		fmt.Println("Mostrando versões locais:")
		return internal.StaticServiceDataArr
	}

	return serviceDataArr
}

func downloadJson() ([]internal.ServiceData, error) {

	httpClient := &http.Client{}

	res, err := httpClient.Get(RepositoryServerAddress)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var serviceDataArr []internal.ServiceData

	jsonDecoder := json.NewDecoder(res.Body)
	if err := jsonDecoder.Decode(&serviceDataArr); err != nil {
		return nil, err
	}

	return serviceDataArr, nil
}
