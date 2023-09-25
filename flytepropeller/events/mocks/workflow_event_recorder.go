// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	config "github.com/flyteorg/flyte/flytepropeller/pkg/controller/config"

	event "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/event"

	mock "github.com/stretchr/testify/mock"
)

// WorkflowEventRecorder is an autogenerated mock type for the WorkflowEventRecorder type
type WorkflowEventRecorder struct {
	mock.Mock
}

type WorkflowEventRecorder_RecordWorkflowEvent struct {
	*mock.Call
}

func (_m WorkflowEventRecorder_RecordWorkflowEvent) Return(_a0 error) *WorkflowEventRecorder_RecordWorkflowEvent {
	return &WorkflowEventRecorder_RecordWorkflowEvent{Call: _m.Call.Return(_a0)}
}

func (_m *WorkflowEventRecorder) OnRecordWorkflowEvent(ctx context.Context, _a1 *event.WorkflowExecutionEvent, eventConfig *config.EventConfig) *WorkflowEventRecorder_RecordWorkflowEvent {
	c_call := _m.On("RecordWorkflowEvent", ctx, _a1, eventConfig)
	return &WorkflowEventRecorder_RecordWorkflowEvent{Call: c_call}
}

func (_m *WorkflowEventRecorder) OnRecordWorkflowEventMatch(matchers ...interface{}) *WorkflowEventRecorder_RecordWorkflowEvent {
	c_call := _m.On("RecordWorkflowEvent", matchers...)
	return &WorkflowEventRecorder_RecordWorkflowEvent{Call: c_call}
}

// RecordWorkflowEvent provides a mock function with given fields: ctx, _a1, eventConfig
func (_m *WorkflowEventRecorder) RecordWorkflowEvent(ctx context.Context, _a1 *event.WorkflowExecutionEvent, eventConfig *config.EventConfig) error {
	ret := _m.Called(ctx, _a1, eventConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *event.WorkflowExecutionEvent, *config.EventConfig) error); ok {
		r0 = rf(ctx, _a1, eventConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
