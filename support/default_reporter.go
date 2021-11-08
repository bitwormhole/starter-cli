package support

import (
	"github.com/bitwormhole/starter/task"
	"github.com/bitwormhole/starter/vlog"
)

type MockReporter struct {
}

func (inst *MockReporter) _Impl() task.ProgressReporter {
	return inst
}

func (inst *MockReporter) Report(p *task.Progress) {
	vlog.Debug("task.Report: progress")
}

func (inst *MockReporter) Update(p *task.Progress) {
	vlog.Info("task.update State:", p.State, " Status:", p.Status)
}

func (inst *MockReporter) HandleCancel(f task.ProgressControlHandlerFn) {
	vlog.Warn("NOP for:HandleCancel")
}

func (inst *MockReporter) HandlePause(f task.ProgressControlHandlerFn) {
	vlog.Warn("NOP for:HandlePause")
}

func (inst *MockReporter) HandleResume(f task.ProgressControlHandlerFn) {
	vlog.Warn("NOP for:HandleResume")
}
