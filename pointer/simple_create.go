// pointer package provides some pointer shortcuts. See:
// 1: https://stackoverflow.com/questions/50830676/set-int-pointer-to-int-value-golang
// 2: https://github.com/openlyinc/pointy
package pointer

func Int64(num int64) *int64 {
	return &num
}
