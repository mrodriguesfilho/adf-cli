/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listRemoteCmd represents the listRemote command
var listRemoteCmd = &cobra.Command{
	Use:   "listRemote",
	Short: "Lista as versões disponíveis do ADF Web para instalação",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(listRemoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listRemoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listRemoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
