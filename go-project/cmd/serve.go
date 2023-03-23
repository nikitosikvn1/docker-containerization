/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fizzbuzz/lib"
	"fmt"
	"github.com/spf13/cobra"
	"html/template"
	"net/http"
	"strconv"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run an http server to anser fizzbuzz queries",
	Run: func(cmd *cobra.Command, args []string) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		type QueryContext struct {
			QueryRaw string
			Query    *int
			Result   *string
			Error    error
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				tmpl.Execute(w, nil)
				return
			}

			context := QueryContext{
				QueryRaw: r.FormValue("query"),
				Query:    nil,
				Result:   nil,
				Error:    nil,
			}
			query, err := strconv.Atoi(context.QueryRaw)
			if err == nil {
				context.Query = &query
				result := lib.FizzBuzz(*context.Query)
				context.Result = &result
			} else {
				context.Error = err
			}

			tmpl.Execute(w, context)
			return
		})
		port, _ := cmd.Flags().GetInt("port")
		fmt.Printf("Listening on http://0.0.0.0:%d\n", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", 8080, "Port to listen to clients connections")
}
