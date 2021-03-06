/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"

	"github.com/SethCurry/magicman/pkg/ent"
	"github.com/SethCurry/magicman/pkg/magicman"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// cacheImagesCmd represents the cacheImages command
var cacheImagesCmd = &cobra.Command{
	Use:   "cacheImages",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cacheImages called")
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		db, err := ent.Open("sqlite3", "file:db.sqlite3?cache=shared&_fk=1")
		if err != nil {
			panic(err)
		}

		if err = db.Schema.Create(context.Background()); err != nil {
			panic(err)
		}

		err = magicman.CacheImages(context.Background(), db, "./backups/images", logger)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cacheImagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheImagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheImagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
