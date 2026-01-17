/*
Copyright © 2026 Shreehari Acharya shreehari.acharya.06@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/Shreehari-Acharya/go-learnings/goproc/internals"
	"github.com/spf13/cobra"
)

var source string

var rootCmd = &cobra.Command{
	Use:   "goproc",
	Short: "Get system process and memory information",
	Long: `goproc is a CLI tool written in Go that provides detailed information
about system processes and memory usage and outputs in json format.`,

	Run: func(cmd *cobra.Command, args []string) { 

		source, err := cmd.Flags().GetString("source")
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] : %v\n", err)
			return
		}

		systemInfo, err := internals.GetSystemInfo(source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] : %v\n", err)
			return
		}

		output, err := json.MarshalIndent(systemInfo, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] : %v\n", err)
			return
		}
		
		fmt.Println(string(output))

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
		"all",
		"Source of memory information [meminfo | loadavg | cpuinfo | all ]",
	)

}


