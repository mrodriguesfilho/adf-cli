/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Default config parameters valus
const (
	defaultWebPort int = 8050

	RepositoryServerAddress string = "localhost"
	RepositoryServerPort    int    = 8001
)

var (
	installedVersions []string
	usedVersion       string
)

// Config parameters
var webPort int = defaultWebPort
var webWorkDir string
var dataServerAddress string
var dataServerPort int

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adf",
	Short: "Ferramenta de linha de comando para funcionalidades administrativas do Ambiente de Design FIHR (ADF)",
	Long:  ``,
	Example: `
	$ adf install --version 0.0.1
	Instalando versão 0.0.1 do ADF Web...
	Versão 0.0.1 do ADF Web instalada com sucesso

	$ adf list
	Versões instaladas:
	ADF Web 0.0.1

	$ adf use 0.0.1
	Definida a versão 0.0.1 do ADF Web a ser utilizada

	$ adf list
	ADF Web 0.0.1 - em uso`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.adf.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".adf" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".adf")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault("web.port", defaultWebPort)

	rootCmd.Flags().Int("webPort", viper.GetInt("web.port"), "Número de porta TCP do ADF Web")
	rootCmd.Flags().String("webWorkDir", viper.GetString("web.workdir"), "Diretório de trabalho do ADF Web")
	rootCmd.Flags().String("dataServerAddress", viper.GetString("dataserver.address"), "Endereço do servidor de dados")
	rootCmd.Flags().Int("dataServerPort", viper.GetInt("dataserver.port"), "Número de porta do servidor de dados")

	viper.BindPFlag("web.port", rootCmd.Flags().Lookup("web.port"))
	viper.BindPFlag("web.workdir", rootCmd.Flags().Lookup("web.workdir"))
	viper.BindPFlag("dataserver.address", rootCmd.Flags().Lookup("dataserver.address"))
	viper.BindPFlag("dataserver.port", rootCmd.Flags().Lookup("dataserver.port"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Usando arquivo de configuração:", viper.ConfigFileUsed()+"\n")

		// Set config variables according to config value.
		webPort = viper.GetInt("web.port")
		webWorkDir = viper.GetString("web.workdir")
		dataServerAddress = viper.GetString("dataserver.address")
		dataServerPort = viper.GetInt("dataserver.port")

		fmt.Printf(
			"web.port: %d\nweb.workdir: %s\n"+
				"dataserver.address: %s\ndataserver.port: %d\n\n",
			webPort, webWorkDir, dataServerAddress, dataServerPort,
		)
	}
}
