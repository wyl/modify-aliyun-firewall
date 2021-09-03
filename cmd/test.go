/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/viper"
	"modify-aliyun-firewall/src"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"test", "verify"},
	Short:   "展示获取到的变量信息",
	Long: `展示获取到的变量信息. For example:

变量信息以 json 的格式进行输出.`,
	Run: func(cmd *cobra.Command, args []string) {
		ic := src.NewIndexClient(viper.Get("key_id").(string), viper.Get("key_secret").(string))
		ic.Show()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	//ic := src.NewIndexClient()
	//ic.Show()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
