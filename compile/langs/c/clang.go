package c

import (
	"os/exec"
)

func Execclang(arg string) int{
	switch arg {
	case "c-clang":
		res := compileCClang()
		if res != 0 {
			return 1
		}
		break
	case "cxx-clang":
		res := compileCppClang()
		if res != 0 {
			return 1
		}
		break
	}
	return 0
}

func compileCClang () int{
	CompileArgs := "clang -w -lm -std=c11 -o /work/a.out -O2 main.c"
	_, err := exec.Command("sh", "-c", CompileArgs).Output()
	if err != nil {return 1}
	return 0
}

func compileCppClang () int{
	CompileArgs := "clang++ -w -lm -std=c++11 -o /work/a.out -O2 main.cpp"
	_, err := exec.Command("sh", "-c", CompileArgs).Output()
	if err != nil {return 1}
	return 0
}
