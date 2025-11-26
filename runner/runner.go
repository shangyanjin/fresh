package runner

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func run() bool {
	runnerLog("Running...")

	args := []string{}
	appArgs := os.Getenv("RUNNER_APP_ARGS")
	if appArgs != "" {
		args = strings.Fields(appArgs)
		runnerLog("Passing args to app: %v", args)
	} else {
		runnerLog("No args passed to app")
	}

	cmd := exec.Command(buildPath(), args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	go io.Copy(appLogWriter{}, stderr)
	go io.Copy(appLogWriter{}, stdout)

	go func() {
		<-stopChannel
		pid := cmd.Process.Pid
		runnerLog("Killing PID %d", pid)
		cmd.Process.Kill()
	}()

	return true
}
