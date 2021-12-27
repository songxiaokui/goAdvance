package image

import (
	"github.com/spf13/cobra"
	"images/internal/image"
)

/*
@Time    : 2021/12/27 14:47
@Author  : austsxk
@Email   : austsxk@163.com
@File    : crop.go
@Software: GoLand
*/

var cropCmd = &cobra.Command{
	Use:   "crop",
	Short: "图片裁剪",
	Long:  "支持多种格式的图片裁剪",
	Run: func(cmd *cobra.Command, args []string) {
		image.CropHandler(inputFile, outputFile, x1, y1, x2, y2)
	},
}

var (
	inputFile  string
	outputFile string
	x1         int
	y1         int
	x2         int
	y2         int
)

func init() {

	cropCmd.Flags().StringVarP(&inputFile, "input_file", "i", "", "输入文件路径")
	cropCmd.Flags().StringVarP(&outputFile, "output_file", "o", "", "输出文件路径")
	cropCmd.Flags().IntVarP(&x1, "x1", "a", 0, "第一个点x轴坐标")
	cropCmd.Flags().IntVarP(&y1, "y1", "b", 0, "第一个点y轴坐标")
	cropCmd.Flags().IntVarP(&x2, "x2", "x", 10, "第二个点x轴坐标")
	cropCmd.Flags().IntVarP(&y2, "y2", "y", 10, "第二个点y轴坐标")
}
