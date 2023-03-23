package cmd

import (
	"fizzbuzz/lib"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query if the number is fizz/buzz or fizzbuzz",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			var queried_num, err = strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			if queried_num < 0 {
				fmt.Printf("Error: negative numbers are not supported yet")
				os.Exit(1)
			}
			fmt.Printf("%d is %s\n", queried_num, lib.FizzBuzz(queried_num))
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
