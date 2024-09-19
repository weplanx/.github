package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"server/bootstrap"
	"server/cmd"
	"server/common"
)

func main() {
	var config string
	root := &cobra.Command{
		Use:               "weplanx",
		Short:             "API service, based on Hertz's project",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cmd.Help()
			}
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.Background()
			var v *common.Values
			if v, err = bootstrap.LoadStaticValues(config); err != nil {
				return
			}
			cmd.SetContext(context.WithValue(ctx, "values", v))
			return
		},
	}
	root.PersistentFlags().StringVarP(&config,
		"config", "c", "config/default.values.yml",
		"The default configuration file of weplanx server values",
	)
	root.AddCommand(cmd.API)
	root.AddCommand(cmd.XAPI)
	root.AddCommand(cmd.OpenAPI)
	root.AddCommand(cmd.Setup)
	root.AddCommand(cmd.Sync)
	root.AddCommand(cmd.User())
	root.AddCommand(cmd.Values)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
