package c

import (
	"context"
	"kojs3-worker/run/langs/lib"
	runningstatus "kojs3-worker/status"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

func RunC(i int, TaskID string) (int, runningstatus.TestStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Todo: Dockerのとき
	arg := "/a.out < /test-case/" + TaskID + "/test_" + strconv.Itoa(i) + ".txt"
	//arg := "cd ./work && ../a.out < ../test-case/" + TaskID + "/test_" + strconv.Itoa(i) + ".txt"
	startTime := time.Now()
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", arg)
	out, err := cmd.Output()
	if err != nil {
		return 0, runningstatus.TestStatus{}
	}
	usage, _ := cmd.ProcessState.SysUsage().(*syscall.Rusage)
	memory := usage.Maxrss / 4 // キロバイト単位(小数点以下切り捨て)
	exitStatus := cmd.ProcessState.ExitCode()
	duration := int(time.Since(startTime).Milliseconds())
	pts, status := lib.CheckExecResult(out, exitStatus, duration, TaskID, i, memory)

	return pts, status
}
