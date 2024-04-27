/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var useVersion string

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Define uma versão do ADF Web a ser utilizada",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Definida a versão %s do ADF Web a ser utilizada\n", useVersion)
		usedVersion = useVersion
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	useCmd.Flags().StringVarP(&useVersion, "version", "v", "latest", "Versão da aplicação de design do ADF a ser utilizada")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
