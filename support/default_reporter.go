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

func (inst *MockReporter) UpdateStatus(s task.Status) {
	vlog.Info("task.Status=", s)
}

func (inst *MockReporter) UpdateState(s task.State) {
	vlog.Info("task.State=", s)
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
