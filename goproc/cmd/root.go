/*
Copyright © 2026 Shreehari Acharya shreehari.acharya.06@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var source string
var output string


var rootCmd = &cobra.Command{
	Use:   "goproc",
	Short: "Get system process and memory information",
	Long: `goproc is a CLI tool written in Go that provides detailed information
about system processes and memory usage and outputs in json or prometheus format.`,

	Run: func(cmd *cobra.Command, args []string) { 

		sourceFlag, _ := cmd.Flags().GetString("source")
		outputFlag, _ := cmd.Flags().GetString("output")

		fmt.Println("source file : ", sourceFlag)
		fmt.Println("output format : ", outputFlag)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVarP(
		&source,
		"source",
		"s",
		"meminfo",
		"Source of memory information [meminfo | loadavg | cpuinfo | all ]",
	)

	rootCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"json",
		"Output format [json, prom] prom - prometheus format",
	)

}


