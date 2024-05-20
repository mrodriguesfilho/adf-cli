/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adf-cli/models"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Pausa uma versão do ADF Web",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if runtime.GOOS == "windows" {
			err = RunWindows()
		} else {
			err = RunUnix()
		}

		cobra.CheckErr(err)

		fmt.Printf("Servidor HAPIFHIR da porta %v foi parado com sucesso", models.HapiFhirDefaultPort)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunUnix() error {
	bashCmd := exec.Command("lsof", "-i", models.HapiFhirDefaultPort)

	stdout, err := bashCmd.Output()
	cobra.CheckErr(err)
	lines := strings.Split(string(stdout), "\n")

	if len(lines) < 2 {
		return fmt.Errorf("nenhum processo foi encontrado na porta %v", models.HapiFhirDefaultPort)
	}

	parts := strings.Fields(lines[1])
	if len(parts) < 2 {
		return fmt.Errorf("output incompatível com a CLI")
	}

	pid := parts[1]

	bashCmd = exec.Command("kill", pid)
	output, err := bashCmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}

func RunWindows() error {
	windowsCmd := exec.Command("netstat", "-ano")

	output, err := windowsCmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if strings.Contains(line, models.HapiFhirDefaultPort) {
			fields := strings.Fields(line)
			if len(fields) < 5 {
				return fmt.Errorf("output incompatível com a CLI")
			}

			pid := fields[len(fields)-1]

			windowsCmd = exec.Command("taskill", "/pid", pid, "/F")
			err = windowsCmd.Run()
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("nenhum processo com porta %s foi encontrado para ser finalizado", models.HapiFhirDefaultPort)
}
