/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strconv"
	"time"
)

const DATAFILE = "./records.json"

type Record struct {
	Name       string
	Address    string
	Mobile     string
	LastAccess string
}

type Records []Record

var records = make(Records, 0, 10)

var phoneBookIndex = make(map[string]int)

func serialize(data interface{}, to io.Writer) error {
	encoder := json.NewEncoder(to)
	return encoder.Encode(data)
}

func deserialize(data interface{}, from io.Reader) error {
	decoder := json.NewDecoder(from)
	return decoder.Decode(data)
}

func readJSONFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return deserialize(&records, file)
}

func writeJSONFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return serialize(records, file)
}

func createIndex() {
	for index, value := range records {
		phoneBookIndex[value.Mobile] = index
	}
}

func createRecord(n, a, t string) *Record {
	lastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Record{n, a, t, lastAccess}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-tut-json",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := writeJSONFile(DATAFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readJSONFile(DATAFILE)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	createIndex()

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
