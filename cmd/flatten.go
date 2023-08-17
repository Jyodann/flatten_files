/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// flattenCmd represents the flatten command
var flattenCmd = &cobra.Command{
	Use:   "flatten",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		basepath, _ := cmd.Flags().GetString("path")
		seperator, _ := cmd.Flags().GetString("seperator")
		var totalFileNum = 0
		var currentFileNum = 0
		if basepath == "" {
			fmt.Print("Flatten called on Current working directory.")
		} else {
			fmt.Printf("Flatten called on %v", basepath)
		}

		if seperator == "" {
			seperator = " "
		}

		err := filepath.Walk(
			basepath,
			func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					totalFileNum++
				}

				return nil
			})

		fmt.Printf("%v Files Discovered\n", totalFileNum)
		if err != nil {
			log.Fatal(err)
		}

		err = filepath.Walk(
			basepath,
			func(path string, info os.FileInfo, err error) error {

				if err != nil {
					fmt.Println(err)
					return nil
				}

				if !info.IsDir() {
					bytesRead, err := os.ReadFile(path)
					rel, err := filepath.Rel(basepath, path)

					if err != nil {
						fmt.Println("Error reading file")
					}

					convertedPath := filepath.ToSlash(rel)
					//fmt.Println(convertedPath)
					finalPath := strings.ReplaceAll(convertedPath, "/", seperator)

					finalPath = filepath.Join(basepath, finalPath)

					err = os.WriteFile(finalPath, bytesRead, 0755)
					currentFileNum++
					fmt.Printf("\r")
					fmt.Printf("Wrote File %s (FileNo. %v / %v)\n", finalPath, currentFileNum, totalFileNum)
				}

				return nil
			})
	},
}

func init() {
	rootCmd.AddCommand(flattenCmd)

	flattenCmd.PersistentFlags().String("path", "", "The path to flatten, defaults to current working directory.")
	flattenCmd.PersistentFlags().String("seperator", "", "The seperator used when flattening the structure. Defaults to Space")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flattenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flattenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
