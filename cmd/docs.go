/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var pathToGenerateDocs string

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Gera documentação em Markdown para essa ferramenta de CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if directory does not exist
		_, err := os.Stat(pathToGenerateDocs)
		if errors.Is(err, os.ErrNotExist) {
			// If directory does not exist, create it.
			err = os.Mkdir(pathToGenerateDocs, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = doc.GenMarkdownTree(rootCmd, pathToGenerateDocs)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	docsCmd.Flags().StringVarP(
		&pathToGenerateDocs,
		"path",
		"p",
		"",
		"Caminho do diretório em que serão criados os arquivos de documentação",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
