package main

/*
@Time    : 2021/9/7 22:21
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 4.link.go
@Software: GoLand
*/

import "fmt"

func main() {
	fmt.Println(1)
}

// 执行 go build -x 4.link.go

// 获得链接流程

/*
WORK=/var/folders/g4/40wbk_1d10g7p0g9jsz2t_z40000gn/T/go-build098557868
mkdir -p $WORK/b001/
cat >$WORK/b001/_gomod_.go << 'EOF' # internal
package main
import _ "unsafe"
//go:linkname __debug_modinfo__ runtime.modinfo
var __debug_modinfo__ = "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tcommand-line-arguments\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
        EOF

cat >$WORK/b001/importcfg << 'EOF' # internal
# import config
packagefile fmt=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/runtime.a
EOF

cd /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/Advance/gc
/Users/austsxk/Golang_dev/go/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid hYFD-e-CyMyjVbrjDaas/hYFD-e-CyMyjVbrjDaas -goversion go1.15.6 -D _/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/Advance/gc -importcfg $WORK/b001/importcfg -pack -c=4 ./4.link.go $WORK/b001/_gomod_.go
/Users/austsxk/Golang_dev/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/_pkg_.a # internal
cp $WORK/b001/_pkg_.a /Users/austsxk/Library/Caches/go-build/e7/e7959cd5b7ac661e8ccf555c65a60f1ddccfc34e134bcf8cbea35010383e151f-d # internal
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=$WORK/b001/_pkg_.a
packagefile fmt=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/runtime.a
packagefile errors=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/errors.a
packagefile internal/fmtsort=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/fmtsort.a
packagefile io=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/io.a
packagefile math=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/math.a
packagefile os=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/os.a
packagefile reflect=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/reflect.a
packagefile strconv=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/strconv.a
packagefile sync=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/sync.a
packagefile unicode/utf8=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/unicode/utf8.a
packagefile internal/bytealg=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/bytealg.a
packagefile internal/cpu=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/cpu.a
packagefile runtime/internal/atomic=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/reflectlite.a
packagefile sort=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/sort.a
packagefile math/bits=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/math/bits.a
packagefile internal/oserror=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/oserror.a
packagefile internal/poll=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/poll.a
packagefile internal/syscall/execenv=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/syscall/unix.a
packagefile internal/testlog=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/testlog.a
packagefile sync/atomic=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/sync/atomic.a
packagefile syscall=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/syscall.a
packagefile time=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/time.a
packagefile internal/unsafeheader=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/unsafeheader.a
packagefile unicode=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/unicode.a
packagefile internal/race=/Users/austsxk/Golang_dev/go/pkg/darwin_amd64/internal/race.a
EOF
mkdir -p $WORK/b001/exe/
cd .
/Users/austsxk/Golang_dev/go/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=v9TCq4O_nK6ypmb8CkGc/hYFD-e-CyMyjVbrjDaas/JYmyum_HdB2N6zQmQ2Wt/v9TCq4O_nK6ypmb8CkGc -extld=clang $WORK/b001/_pkg_.a
/Users/austsxk/Golang_dev/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out 4.link
rm -r $WORK/b001/

*/

// readelf dumpobj 查询 ELF文件 可执行 可链接格式文件
