package run

import (
	"github.com/laminne/nmoj-worker/compile/langs/c"
	"os"
)

func Compile(arg string) int {
	switch arg {
	case "c-gcc":
		err := c.Execgcc(arg)
		if err != 0 {
			return 1
		}
		break
	case "c-clang":
		err := c.Execclang(arg)
		if err != 0 {
			return 1
		}
		break
	case "cxx-gxx":
		err := c.Execgcc(arg)
		if err != 0 {
			return 1
		}
		break
	case "cxx-clang":
		err := c.Execclang(arg)
		if err != 0 {
			return 1
		}
		break
	default:
		os.Exit(1)
	}
	return 0
}
