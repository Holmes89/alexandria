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
	"github.com/spf13/cobra"
)

// tagLinkCmd represents the tagLink command
var tagLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Add a tag to a link",
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"id", "tag"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.TagLink(args[0], args[1])
	},
}

func init() {
	tagCmd.AddCommand(tagLinkCmd)
}
