package main

import (
	"fmt"
	"unsafe"
)

/*
@Time    : 2021/10/15 09:31
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 23.sliceCapacityExpansion.go
@Software: GoLand
*/

func main() {
	s1 := []int{1, 2, 3, 5, 4}
	fmt.Println(len(s1), cap(s1)) // cap=5
	s1 = append(s1, []int{1, 2, 3}...)
	fmt.Println(len(s1), cap(s1)) // cap=10
	// 扩容源码 runtime/slice.go
	/*
		1 newcap := old.cap
		2 if newcap+newcap < cap {
		3     newcap = cap
		4 } else {
		5     for {
		6         if old.len < 1024 {
		7             newcap += newcap
		8         } else {
		9             newcap += newcap / 4
		10         }
		11         if newcap >= cap {
		12             break
		13         }
		14     }
		15 }
	*/
	// 第1行：old.cap = 5
	// 第2行：cap：扩容前容量加上扩容的元素数量，对于此例，就是 5+3
	// 第2行：5+5=10 > 8
	// 第6行：old.len = 5
	// 第7行：newcap += newcap newcap=10
	// 第11行：newcap=10 cap=8 10 >= 8
	// 第12行：break
	// 最后结果就是 newcap=10 符合
	// newcap 只是预估容量，并不是最终的容量，要计算最终的容量，还需要参考另一个维度，也就是内存分配。

	/*
		内存管理模块的代码，在 runtime/sizeclasses.go
		// class  bytes/obj  bytes/span  objects  tail waste  max waste
		//     1          8        8192     1024           0     87.50%
		//     2         16        8192      512           0     43.75%
		//     3         32        8192      256           0     46.88%
		//     4         48        8192      170          32     31.52%
		//     5         64        8192      128           0     23.44%
		//     6         80        8192      102          32     19.07%
		//     7         96        8192       85          32     15.95%
		//     8        112        8192       73          16     13.56%
		//     9        128        8192       64           0     11.72%
		//    10        144        8192       56         128     11.82%
		//    11        160        8192       51          32      9.73%
		//    12        176        8192       46          96      9.59%
		//    13        192        8192       42         128      9.25%
		//    14        208        8192       39          80      8.12%
		//    15        224        8192       36         128      8.15%
		//    16        240        8192       34          32      6.62%
		//    17        256        8192       32           0      5.86%
		//    18        288        8192       28         128     12.16%
		//    19        320        8192       25         192     11.80%
		//    20        352        8192       23          96      9.88%
		//    21        384        8192       21         128      9.51%
		//    22        416        8192       19         288     10.71%
		//    23        448        8192       18         128      8.37%
		//    24        480        8192       17          32      6.82%
		//    25        512        8192       16           0      6.05%
		//    26        576        8192       14         128     12.33%
		//    27        640        8192       12         512     15.48%
		//    28        704        8192       11         448     13.93%
		//    29        768        8192       10         512     13.94%
		//    30        896        8192        9         128     15.52%
		//    31       1024        8192        8           0     12.40%
		//    32       1152        8192        7         128     12.41%
		//    33       1280        8192        6         512     15.55%
		//    34       1408       16384       11         896     14.00%
		//    35       1536        8192        5         512     14.00%
		//    36       1792       16384        9         256     15.57%
		//    37       2048        8192        4           0     12.45%
		//    38       2304       16384        7         256     12.46%
		//    39       2688        8192        3         128     15.59%
		//    40       3072       24576        8           0     12.47%
		//    41       3200       16384        5         384      6.22%
		//    42       3456       24576        7         384      8.83%
		//    43       4096        8192        2           0     15.60%
		//    44       4864       24576        5         256     16.65%
		//    45       5376       16384        3         256     10.92%
		//    46       6144       24576        4           0     12.48%
		//    47       6528       32768        5         128      6.23%
		//    48       6784       40960        6         256      4.36%
		//    49       6912       49152        7         768      3.37%
		//    50       8192        8192        1           0     15.61%
		//    51       9472       57344        6         512     14.28%

		// int占8字节 使用sizeof可以查看值占用空间
		// 新分配的空间为10 * 8 = 80
		// 对应第6类 内存分配的规则 分配80字节，可以计算真实容量为80/8 = 10
		// 所以实际分配的内存容量为： 10

	*/
	var a int
	var b string
	var c float32
	var d float64
	var e bool
	var f struct{}
	var h chan int
	var m interface{}
	fmt.Println("int 占用字节数:", unsafe.Sizeof(a))
	fmt.Println("string 占用字节数:", unsafe.Sizeof(b))
	fmt.Println("float32 占用字节数:", unsafe.Sizeof(c))
	fmt.Println("float64 占用字节数:", unsafe.Sizeof(d))
	fmt.Println("bool 占用字节数:", unsafe.Sizeof(e))
	fmt.Println("struct{} 占用字节数:", unsafe.Sizeof(f))
	fmt.Println("chan int 占用字节数:", unsafe.Sizeof(h))
	fmt.Println("interface{} 占用字节数:", unsafe.Sizeof(m))

}
