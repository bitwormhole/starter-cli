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

	p := &task.Progress{}
	p.Name = "sleep"
	p.Title = "Sleep"
	p.Unit = "ms"
	p.Value = ms - ttl
	p.ValueMin = 0
	p.ValueMax = ms

	p.State = task.StateRunning
	reporter.Update(p)

	defer func() {
		if cancelled {
			p.Status = task.StatusCancelled
			reporter.Update(p)
		} else {
			p.Status = task.StatusOK
			reporter.Update(p)
		}
		p.State = task.StateStopped
		reporter.Update(p)
	}()

	reporter.HandleCancel(func(reporter task.ProgressReporter) error {
		cancelled = true
		p.State = task.StateCancelling
		reporter.Update(p)
		return nil
	})
	reporter.HandlePause(func(reporter task.ProgressReporter) error {
		paused = true
		p.State = task.StatePaused
		reporter.Update(p)
		return nil
	})
	reporter.HandleResume(func(reporter task.ProgressReporter) error {
		paused = false
		p.State = task.StateRunning
		reporter.Update(p)
		return nil
	})

	p.ValueMin = 0
	p.ValueMax = ttl
	for ttl > 0 {
		p.Value = ttl
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
		reporter.Report(p)
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

// GetHelpInfo 取帮助信息
func (inst *Sleep) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "sleep"
	info.Title = "睡一下"
	info.Description = "睡指定的毫秒"
	info.Content = "usage: sleep [ms]"
	return info
}
