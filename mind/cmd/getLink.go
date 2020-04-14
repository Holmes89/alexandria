/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// getLinkCmd represents the getLink command
var getLinkCmd = &cobra.Command{
	Use:   "links",
	Short: "List all links",
	RunE: func(cmd *cobra.Command, args []string) error {
		results, err := app.FindLinks()
		if err != nil {
			if debug {
				errString := fmt.Errorf("error: %w", err)
				fmt.Fprintln(out, errString.Error())
			}
			return errors.New("unable to fetch links")
		}
		tw := getTabWriter()
		fmt.Fprintf(tw, "\n %s\t%s\t",  "ID", "Name")
		for _, r := range results {
			fmt.Fprintf(tw, "\n %s\t%s\t", r.ID, r.DisplayName)
		}
		fmt.Fprintf(tw, "\n\n")
		tw.Flush()
		return nil
	},
}

func init() {
	getCmd.AddCommand(getLinkCmd)
}
