/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		mobile, _ := cmd.Flags().GetString("mobile")
		if mobile == "" {
			fmt.Println("Invalid mobile:", mobile)
			return
		}

		// use command line argument
		//if len(args) == 0 {
		//	fmt.Println("Please provide a mobile")
		//	return
		//}
		// fmt.Println(args[0])

		err := deleteEntry(mobile)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func deleteEntry(mobile string) error {
	index, ok := phoneBookIndex[mobile]
	if !ok {
		return fmt.Errorf("%s does not exists", mobile)
	}

	records = append(records[:index], records[index+1:]...)

	err := writeJSONFile(DATAFILE)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("mobile", "m", "", "Mobile")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
