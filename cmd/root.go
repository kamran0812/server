/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"server/cmd/helpers"

	"github.com/spf13/cobra"
)

var (
	Port string
	Path string
)
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Simple server",
	Long:  `A simple server that serves static files`,
	Run: func(cmd *cobra.Command, args []string) {

		helpers.Serve(Path, Port)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "80", "PORT you want to serve")
	rootCmd.PersistentFlags().StringVarP(&Path, "path", "d", "", "Path to your static folder (Required)")
	rootCmd.MarkFlagRequired("path")
}
