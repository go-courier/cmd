package main

import (
	"github.com/go-courier/enumeration/generator"
	"github.com/go-courier/packagesx"
	"github.com/spf13/cobra"
)

func init() {
	cmdGen.AddCommand(cmdGenEnum)
}

var cmdGenEnum = &cobra.Command{
	Use:   "enum",
	Short: "generate interfaces of enumeration",
	Run: func(cmd *cobra.Command, args []string) {
		runGenerator(func(p *packagesx.Package) Generator {
			g := generator.NewEnumGenerator(p)
			g.Scan(args...)
			return g
		})
	},
}
