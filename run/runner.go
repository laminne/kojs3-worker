package run

import (
	"kojs3-worker/run/langs/c"
	"kojs3-worker/run/langs/nodejs"
	"kojs3-worker/run/langs/ruby"
	"kojs3-worker/status"
	"os"
)

func Run(i int, TaskID string) (int, status.TestStatus) {
	if os.Args[1] == "node-js" {
		pts, s := nodejs.RunJS(i, TaskID)
		return pts, s
	} else if os.Args[1] == "ruby" {
		pts, s := ruby.RunRuby(i, TaskID)
		return pts, s
	} else {
		pts, s := c.RunC(i, TaskID)
		return pts, s
	}
}
