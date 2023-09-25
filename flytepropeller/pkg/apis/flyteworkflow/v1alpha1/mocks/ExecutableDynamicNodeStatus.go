// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// ExecutableDynamicNodeStatus is an autogenerated mock type for the ExecutableDynamicNodeStatus type
type ExecutableDynamicNodeStatus struct {
	mock.Mock
}

type ExecutableDynamicNodeStatus_GetDynamicNodePhase struct {
	*mock.Call
}

func (_m ExecutableDynamicNodeStatus_GetDynamicNodePhase) Return(_a0 v1alpha1.DynamicNodePhase) *ExecutableDynamicNodeStatus_GetDynamicNodePhase {
	return &ExecutableDynamicNodeStatus_GetDynamicNodePhase{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableDynamicNodeStatus) OnGetDynamicNodePhase() *ExecutableDynamicNodeStatus_GetDynamicNodePhase {
	c_call := _m.On("GetDynamicNodePhase")
	return &ExecutableDynamicNodeStatus_GetDynamicNodePhase{Call: c_call}
}

func (_m *ExecutableDynamicNodeStatus) OnGetDynamicNodePhaseMatch(matchers ...interface{}) *ExecutableDynamicNodeStatus_GetDynamicNodePhase {
	c_call := _m.On("GetDynamicNodePhase", matchers...)
	return &ExecutableDynamicNodeStatus_GetDynamicNodePhase{Call: c_call}
}

// GetDynamicNodePhase provides a mock function with given fields:
func (_m *ExecutableDynamicNodeStatus) GetDynamicNodePhase() v1alpha1.DynamicNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.DynamicNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.DynamicNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.DynamicNodePhase)
	}

	return r0
}

type ExecutableDynamicNodeStatus_GetDynamicNodeReason struct {
	*mock.Call
}

func (_m ExecutableDynamicNodeStatus_GetDynamicNodeReason) Return(_a0 string) *ExecutableDynamicNodeStatus_GetDynamicNodeReason {
	return &ExecutableDynamicNodeStatus_GetDynamicNodeReason{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableDynamicNodeStatus) OnGetDynamicNodeReason() *ExecutableDynamicNodeStatus_GetDynamicNodeReason {
	c_call := _m.On("GetDynamicNodeReason")
	return &ExecutableDynamicNodeStatus_GetDynamicNodeReason{Call: c_call}
}

func (_m *ExecutableDynamicNodeStatus) OnGetDynamicNodeReasonMatch(matchers ...interface{}) *ExecutableDynamicNodeStatus_GetDynamicNodeReason {
	c_call := _m.On("GetDynamicNodeReason", matchers...)
	return &ExecutableDynamicNodeStatus_GetDynamicNodeReason{Call: c_call}
}

// GetDynamicNodeReason provides a mock function with given fields:
func (_m *ExecutableDynamicNodeStatus) GetDynamicNodeReason() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutableDynamicNodeStatus_GetExecutionError struct {
	*mock.Call
}

func (_m ExecutableDynamicNodeStatus_GetExecutionError) Return(_a0 *core.ExecutionError) *ExecutableDynamicNodeStatus_GetExecutionError {
	return &ExecutableDynamicNodeStatus_GetExecutionError{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableDynamicNodeStatus) OnGetExecutionError() *ExecutableDynamicNodeStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError")
	return &ExecutableDynamicNodeStatus_GetExecutionError{Call: c_call}
}

func (_m *ExecutableDynamicNodeStatus) OnGetExecutionErrorMatch(matchers ...interface{}) *ExecutableDynamicNodeStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError", matchers...)
	return &ExecutableDynamicNodeStatus_GetExecutionError{Call: c_call}
}

// GetExecutionError provides a mock function with given fields:
func (_m *ExecutableDynamicNodeStatus) GetExecutionError() *core.ExecutionError {
	ret := _m.Called()

	var r0 *core.ExecutionError
	if rf, ok := ret.Get(0).(func() *core.ExecutionError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionError)
		}
	}

	return r0
}

type ExecutableDynamicNodeStatus_GetIsFailurePermanent struct {
	*mock.Call
}

func (_m ExecutableDynamicNodeStatus_GetIsFailurePermanent) Return(_a0 bool) *ExecutableDynamicNodeStatus_GetIsFailurePermanent {
	return &ExecutableDynamicNodeStatus_GetIsFailurePermanent{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableDynamicNodeStatus) OnGetIsFailurePermanent() *ExecutableDynamicNodeStatus_GetIsFailurePermanent {
	c_call := _m.On("GetIsFailurePermanent")
	return &ExecutableDynamicNodeStatus_GetIsFailurePermanent{Call: c_call}
}

func (_m *ExecutableDynamicNodeStatus) OnGetIsFailurePermanentMatch(matchers ...interface{}) *ExecutableDynamicNodeStatus_GetIsFailurePermanent {
	c_call := _m.On("GetIsFailurePermanent", matchers...)
	return &ExecutableDynamicNodeStatus_GetIsFailurePermanent{Call: c_call}
}

// GetIsFailurePermanent provides a mock function with given fields:
func (_m *ExecutableDynamicNodeStatus) GetIsFailurePermanent() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
