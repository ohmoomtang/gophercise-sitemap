/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"oot.me/sitemap/utils"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "Sitemap generator",
	Long:  `Sitemap genetator for specific URL and output as an XML format file`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			fmt.Println("Error occured! :", err)
			os.Exit(1)
		}
		urlValid, err := utils.CheckURL(url)
		if err != nil && !urlValid {
			fmt.Println("Error occured! :", err)
			os.Exit(1)
		}
		htmlReader, err := utils.URLtoReader(url)
		if err != nil {
			fmt.Println("Error occured! :", err)
			os.Exit(1)
		}
		defer htmlReader.Close()
		links, err := utils.ParseLink(htmlReader)
		if err != nil {
			fmt.Println("Error occured! :", err)
			os.Exit(1)
		}
		utils.WriteToXML(links, "out.xml")
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sitemap.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("url", "u", "", "Specify URL for genearting sitemap")
	rootCmd.MarkFlagRequired("url")
}
