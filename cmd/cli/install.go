/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/adfweb"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var installVersion string

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

	installCmd.Flags().StringVarP(&installVersion, "version", "v", "latest", "Versão da aplicação de design do ADF")
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
	err := adfweb.InstallJVM()
	if err != nil {
		log.Print(err)
		fmt.Printf(
			"Não foi possível fazer a instação da JVM especificada. Erro: %v\n", err,
		)
		return
	}

	err = adfweb.InstallADFWeb(RepositoryServerAddress, installVersion)
	if err != nil {
		log.Print(err)
		fmt.Printf(
			"Não foi possível fazer instalar a versão especificada. Erro: %v\n", err,
		)
		return
	}
	installedVersions = append(installedVersions, installVersion)
	fmt.Printf("Versão %s do ADF Web instalada com sucesso\n", installVersion)
}
