package main

/*
	KEMOMIMI ONLINE JUDGE SYSTEM - Runner / Judge
	|C| 2021 - 2022 Tatsuto Yamamoto
	This code is licensed under MIT license.
*/
import (
	"encoding/json"
	"fmt"
	"kojs3-worker/run"
	"kojs3-worker/status"
	"os"
	"os/exec"
)

func main() {
	arg := os.Args
	_, _ = exec.Command("rm", "-f", "a.out").Output()
	if len(arg) != 3 {
		os.Exit(1)
		return
	}

	// コンパイルの必要がある場合
	if arg[1] == "c-gcc" || arg[1] == "c-clang" || arg[1] == "cxx-gxx" || arg[1] == "cxx-clang" {
		res := run.Compile(os.Args[1])
		if res == 1 {
			// コンパイル失敗 ToDo: TestID出力するように
			fmt.Printf(`{"TaskID":"%s","Status":[{"TestID":"","ExitStatus":0,"Duration":0,"Status":"CE"},{"TestID":"","ExitStatus":0,"Duration":0,"Status":"CE"},{"TestID":"","ExitStatus":0,"Duration":0,"Status":"CE"}]}`, arg[2])
			return
		}
	}
	var point int
	var runs []status.TestStatus
	for i := 0; i < 3; i++ {
		p, s := run.Run(i+1, arg[2])
		point += p
		runs = append(runs, s)
	}
	P := status.RunningStatus{
		TaskID: arg[2],
		Status: runs,
	}

	Jsondata, err := json.Marshal(P)
	if err != nil {
		return
	}

	fmt.Printf("%s", string(Jsondata))
	return

}
