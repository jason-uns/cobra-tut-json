/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func insert(n, a, m string) error {

	_, ok := phoneBookIndex[m]
	if ok {
		return fmt.Errorf("%s already exists", m)
	}

	record := createRecord(n, a, m)

	records = append(records, *record)

	err := writeJSONFile(DATAFILE)
	if err != nil {
		return err
	}

	return nil

}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Not a valid name:", name)
			return
		}

		address, _ := cmd.Flags().GetString("address")
		if address == "" {
			fmt.Println("Not a valid address:", address)
			return
		}

		mobile, _ := cmd.Flags().GetString("mobile")
		if mobile == "" {
			fmt.Println("Not a valid mobile:", mobile)
			return
		}

		mobile = strings.ReplaceAll(mobile, "-", "")
		if !validateMobile(mobile) {
			fmt.Println("Not a valid mobile:", mobile)
			return
		}

		err := insert(name, address, mobile)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func validateMobile(mobile string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.Match([]byte(mobile))
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "Name")
	addCmd.Flags().StringP("address", "a", "", "Address")
	addCmd.Flags().StringP("mobile", "m", "", "Mobile")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
