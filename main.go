package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "hello",
		Short: "this is prints the hellow world",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hellow World")

		},
	}
	fmt.Println("Hellow world")
	cobra.CheckErr(cmd.Execute())

}
