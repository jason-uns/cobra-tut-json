/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		mobile, _ := cmd.Flags().GetString("mobile")
		if mobile == "" {
			fmt.Println("Invalid mobile", mobile)
			return
		}

		result := search(mobile)
		if result == nil {
			fmt.Println("Mobile not found", mobile)
			return
		}
		fmt.Println(*result)
	},
}

func search(mobile string) *Record {
	index, ok := phoneBookIndex[mobile]
	if !ok {
		return nil
	}
	return &records[index]
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("mobile", "m", "", "Mobile")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
