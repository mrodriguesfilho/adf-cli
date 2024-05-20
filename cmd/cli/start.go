/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/models"
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicializa a versão em uso do ADF Web",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		hapifhirFolder := models.AdfDirectory + "/hapifhir/hapifhir.war"

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
