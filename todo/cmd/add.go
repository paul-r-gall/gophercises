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
	"strings"

	"github.com/boltdb/bolt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long:  `Add a task to the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if label == "imp" {
			fmt.Println("Invalid Label")
			return
		}
		taskName := strings.Join(args, " ")

		db, err := bolt.Open(dbFile, 0600, nil)
		defer db.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		// add taskName to the database.
		err = db.Update(func(tx *bolt.Tx) error {
			// create bucket fot the label if it doesn't already exist
			b, err := tx.CreateBucketIfNotExists([]byte(label))
			// put this task in that bucket
			err = b.Put([]byte(taskName), nil)
			if err != nil {
				return err
			}
			// if it's important, put it in the important bucket.
			if imp {
				b = tx.Bucket([]byte("imp"))
				b.Put([]byte(taskName), []byte(label))
			}
			return err
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringVarP(&label, "label", "l", "", "the label of the task to add (optional)")
	addCmd.Flags().BoolVarP(&imp, "imp", "i", false, "Is this an important task? (optional)")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
