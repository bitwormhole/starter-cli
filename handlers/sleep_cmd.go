package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
)

// Sleep ...
type Sleep struct {
	markup.Component `class:"cli-handler"`
}

func (inst *Sleep) _Impl() cli.Handler {
	return inst
}

// Init ...
func (inst *Sleep) Init(service cli.Service) error {
	return service.RegisterHandler("sleep", inst)
}

// Handle ...
func (inst *Sleep) Handle(t *cli.TaskContext) error {

	console := t.Console
	task := t.CurrentTask
	ms, err := inst.getArgMillisec(task.Arguments, 1)
	if err != nil {
		return err
	}

	console.WriteString("Sleep ")
	console.WriteString(strconv.FormatInt(ms, 10))
	console.WriteString(" ms ... ")
	if ms < 3000 {
		time.Sleep(time.Millisecond * time.Duration(ms))
	} else {
		inst.doLongSleep(t, ms)
	}
	console.WriteString("[done]\n")

	return nil
}

func (inst *Sleep) doLongSleep(t *cli.TaskContext, ms int64) error {

	reporter := t.Reporter
	const step = 1000
	ttl := ms
	cancelled := false
	paused := false

	reporter.UpdateState(task.StateRunning)

	defer func() {
		if cancelled {
			reporter.UpdateStatus(task.StatusCancelled)
		} else {
			reporter.UpdateStatus(task.StatusOK)
		}
		reporter.UpdateState(task.StateStopped)
	}()

	reporter.HandleCancel(func(reporter task.ProgressReporter) error {
		cancelled = true
		reporter.UpdateState(task.StateCancelling)
		return nil
	})
	reporter.HandlePause(func(reporter task.ProgressReporter) error {
		paused = true
		reporter.UpdateState(task.StatePaused)
		return nil
	})
	reporter.HandleResume(func(reporter task.ProgressReporter) error {
		paused = false
		reporter.UpdateState(task.StateRunning)
		return nil
	})

	for ttl > 0 {
		todo := ttl
		if ttl > step {
			todo = step
		}
		time.Sleep(time.Millisecond * time.Duration(todo))
		if cancelled {
			break
		}
		if !paused {
			ttl -= todo
		}
		p := task.Progress{}
		p.Name = "sleep"
		p.Title = "Sleep"
		p.Value = ms - ttl
		p.ValueMin = 0
		p.ValueMax = ms
		reporter.Report(&p)
	}
	return nil
}

func (inst *Sleep) getArgMillisec(args []string, index int) (int64, error) {
	if index < len(args) {
		text := args[index]
		return strconv.ParseInt(text, 10, 64)
	}
	return 0, errors.New("bad argument")
}

func (inst *Sleep) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "sleep"
	info.Title = "睡一下"
	info.Description = "睡指定的毫秒"
	info.Content = "usage: sleep [ms]"
	return info
}
