// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		fmt.Println(label)

		// don't want to double-list tasks.
		taskSet := make(map[string]bool)
		// start by listing the important tasks only:
		db, err := bolt.Open(dbFile, 0600, nil)
		defer db.Close()
		if err!= nil {
			fmt.Println(err)
			return
		}
		db.View(func(tx *bolt.Tx) error {

			// first, let's go through the important ones
			b := tx.Bucket([]byte("imp"))
			fmt.Println("Important Tasks:")
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				// if label has not been provided, we will print everything.
				if label == "" {
					taskSet[string(k)] = true
					fmt.Println(string(k))
				} else if label == string(v) {
					taskSet[string(k)] = true
					fmt.Println(string(k))
				}

				if imp {
					return nil
				}
			}

			// now, let's go through the rest, and let's not double count.

			// if there is a label, only print those.
			if label != "" {
				b = tx.Bucket([]byte(label))
				b.ForEach(func(k, v []byte) error {

					if taskSet[string(k)] {
						return nil
					}
					if v != nil {
						return nil
					}
					fmt.Println(string(k))
					return nil
				})
				return nil
			}
			// otherwise, print everything.
			tx.ForEach(func(name []byte, b *bolt.Bucket) error {
				if string(name) == "imp" {
					return nil
				}
				b.ForEach(func(k, v []byte) error {
					if taskSet[string(k)] {
						return nil
					}
					if v != nil {
						return nil
					}
					fmt.Println(string(k))
					return nil
				})
				return nil
			})

			return nil
		})

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().StringVarP(&label, "label", "l", "", "list all tasks with this label")
	listCmd.Flags().BoolVarP(&imp, "imp", "i", false, "Use this flag to list only the important tasks")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
