package main

import (
	"github.com/go-courier/packagesx"
	"github.com/go-courier/statuserror/generator"
	"github.com/spf13/cobra"
)

func init() {
	cmdGen.AddCommand(cmdGenStatusError)
}

var cmdGenStatusError = &cobra.Command{
	Use:     "status-error",
	Aliases: []string{"error"},
	Short:   "generate interfaces of status error",
	Run: func(cmd *cobra.Command, args []string) {
		runGenerator(func(pkg *packagesx.Package) Generator {
			g := generator.NewStatusErrorGenerator(pkg)
			g.Scan(args...)
			return g
		})
	},
}
