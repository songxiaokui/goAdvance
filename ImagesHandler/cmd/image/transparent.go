package image

import "github.com/spf13/cobra"

/*
@Time    : 2021/12/27 14:48
@Author  : austsxk
@Email   : austsxk@163.com
@File    : transparent.go
@Software: GoLand
*/

var alphaCmd = &cobra.Command{
	Use:   "transparent",
	Short: "图片设置透明度",
	Long:  "支持多种格式的图片指定颜色透明度设置",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {}
