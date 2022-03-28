package lib

import (
	runningstatus "github.com/laminne/nmoj-worker/status"
	"io/ioutil"
	"strconv"
)

// CheckExecResult チェック関数
func CheckExecResult(result []byte, exitCode int, duration int, TaskID string, i int, memory int64) (int, runningstatus.TestStatus) {
	if exitCode != 0 {
		return 0, runningstatus.TestStatus{
			TestID:     "",
			ExitStatus: exitCode,
			Duration:   duration,
			Status:     "RE",
			Memory: memory,
		}
	}

	fileName := "./test-case/" + TaskID + "/ans_" + strconv.Itoa(i) + ".txt"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		// ファイルが読めないのはIE
		return 0, runningstatus.TestStatus{
			TestID:     "",
			ExitStatus: 0,
			Duration:   duration,
			Memory: memory,
			Status:     "IE",
		}
	}

	// stdoutとテストケースを照合させて判定
	if string(result) == string(bytes) {
		return 1, runningstatus.TestStatus{
			TestID:     "",
			ExitStatus: 0,
			Duration:   duration,
			Memory: memory,
			Status:     "AC",
		}
	} else {
		return 0, runningstatus.TestStatus{
			TestID:     "",
			ExitStatus: 0,
			Duration:   duration,
			Memory: memory,
			Status:     "WA",
		}
	}
}
