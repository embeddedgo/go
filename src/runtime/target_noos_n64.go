//go:build noos && n64

package runtime

import "github.com/clktmr/n64/machine"

//go:nowritebarrierrec
//go:nosplit
func targetDefaultWrite(fd int, p []byte) int {
	return machine.DefaultWrite(fd, p)
}
