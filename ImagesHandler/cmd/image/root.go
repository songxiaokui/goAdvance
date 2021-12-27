package image

import (
	"github.com/spf13/cobra"
)

/*
@Time    : 2021/12/27 14:53
@Author  : austsxk
@Email   : austsxk@163.com
@File    : root.go
@Software: GoLand
*/

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cropCmd)
	rootCmd.AddCommand(alphaCmd)
}
