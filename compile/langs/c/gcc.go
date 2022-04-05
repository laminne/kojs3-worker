package c

import (
	"os/exec"
)

func Execgcc(arg string) int {
	switch arg {
	case "c-gcc":
		res := compileC()
		if res != 0 {
			return 1
		}
		break
	case "cxx-gxx":
		res := compileCpp()
		if res != 0 {
			return 1
		}
		break
	}
	return 0
}

func compileC() int {
	CompileArgs := "gcc -w -lm -std=gnu11 -o a.out -O2 main.c"
	_, err := exec.Command("/bin/sh", "-c", CompileArgs).Output()
	if err != nil {
		return 1
	}
	return 0
}

func compileCpp() int {
	CompileArgs := "g++ -w -lm -std=gnu++11 -o a.out -O2 main.cpp"
	_, err := exec.Command("/bin/sh", "-c", CompileArgs).Output()
	if err != nil {
		return 1
	}
	return 0
}
