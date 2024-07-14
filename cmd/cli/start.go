/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/internal/models"
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var startVersion string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicializa a versão em uso do ADF Web",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if startVersion == "" {
			startVersion = models.PreferencesBuiltInVersion
		}

		hapifhirFolder := models.AdfDirectory + "/" + startVersion + "/hapifhir/hapifhir.war"

		bashCmd := exec.Command("java", "-jar", hapifhirFolder)

		stdout, err := bashCmd.StdoutPipe()
		if err != nil {
			fmt.Printf("Error obtaining stdout pipe: %s\n", err)
			return
		}

		stderr, err := bashCmd.StderrPipe()
		if err != nil {
			fmt.Printf("Error obtaining stderr pipe: %s\n", err)
			return
		}

		if err := bashCmd.Start(); err != nil {
			fmt.Printf("Error starting command: %s\n", err)
			return
		}

		stdoutScanner := bufio.NewScanner(stdout)
		go func() {
			for stdoutScanner.Scan() {
				fmt.Printf("STDOUT: %s\n", stdoutScanner.Text())
			}
		}()

		stderrScanner := bufio.NewScanner(stderr)
		go func() {
			for stderrScanner.Scan() {
				fmt.Printf("STDERR: %s\n", stderrScanner.Text())
			}
		}()

		if err := bashCmd.Wait(); err != nil {
			fmt.Printf("Error waiting for command: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&startVersion, "version", "v", "", "Versão do Bundle ADF a ser iniciada")
}
