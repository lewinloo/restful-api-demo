package cmd

import (
	"fmt"

	"github.com/lewinloo/restful-api-demo/version"
	"github.com/spf13/cobra"
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo-api 后端API",
	Long:  "demo-api 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print demo-api version")
}
