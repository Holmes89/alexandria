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
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

// getLinkCmd represents the getLink command
var getLinkCmd = &cobra.Command{
	Use:        "link",
	Short:      "List all links",
	Args:       cobra.MaximumNArgs(1),
	ArgAliases: []string{"id"},
	Aliases:    []string{"links"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			results, err := app.FindLinks()
			if err != nil {
				if debug {
					errString := fmt.Errorf("error: %w", err)
					fmt.Fprintln(out, errString.Error())
				}
				return errors.New("unable to fetch links")
			}
			tw := getTabWriter()
			fmt.Fprintf(tw, "\n %s\t%s\t", "ID", "Name")
			for _, r := range results {
				fmt.Fprintf(tw, "\n %s\t%s\t", r.ID, r.DisplayName)
			}
			fmt.Fprintf(tw, "\n\n")
			tw.Flush()
		} else {
			results, err := app.FindLinkByID(args[0])
			if err != nil {
				if debug {
					errString := fmt.Errorf("error: %w", err)
					fmt.Fprintln(out, errString.Error())
				}
				return errors.New("unable to fetch link")
			}
			b, _ := yaml.Marshal(results)
			fmt.Fprintln(out, string(b))
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getLinkCmd)
}
