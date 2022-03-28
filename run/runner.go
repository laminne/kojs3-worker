package run

import (
	"github.com/laminne/nmoj-worker/run/langs/c"
	"github.com/laminne/nmoj-worker/run/langs/nodejs"
	"github.com/laminne/nmoj-worker/run/langs/ruby"
	"github.com/laminne/nmoj-worker/status"
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
		return pts,s
	}
}
