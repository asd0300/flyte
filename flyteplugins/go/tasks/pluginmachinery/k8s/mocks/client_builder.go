// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	cache "sigs.k8s.io/controller-runtime/pkg/cache"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	k8s "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/k8s"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"
)

// ClientBuilder is an autogenerated mock type for the ClientBuilder type
type ClientBuilder struct {
	mock.Mock
}

type ClientBuilder_Build struct {
	*mock.Call
}

func (_m ClientBuilder_Build) Return(_a0 client.Client, _a1 error) *ClientBuilder_Build {
	return &ClientBuilder_Build{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ClientBuilder) OnBuild(_a0 cache.Cache, config *rest.Config, options client.Options) *ClientBuilder_Build {
	c_call := _m.On("Build", _a0, config, options)
	return &ClientBuilder_Build{Call: c_call}
}

func (_m *ClientBuilder) OnBuildMatch(matchers ...interface{}) *ClientBuilder_Build {
	c_call := _m.On("Build", matchers...)
	return &ClientBuilder_Build{Call: c_call}
}

// Build provides a mock function with given fields: _a0, config, options
func (_m *ClientBuilder) Build(_a0 cache.Cache, config *rest.Config, options client.Options) (client.Client, error) {
	ret := _m.Called(_a0, config, options)

	var r0 client.Client
	if rf, ok := ret.Get(0).(func(cache.Cache, *rest.Config, client.Options) client.Client); ok {
		r0 = rf(_a0, config, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cache.Cache, *rest.Config, client.Options) error); ok {
		r1 = rf(_a0, config, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ClientBuilder_WithUncached struct {
	*mock.Call
}

func (_m ClientBuilder_WithUncached) Return(_a0 k8s.ClientBuilder) *ClientBuilder_WithUncached {
	return &ClientBuilder_WithUncached{Call: _m.Call.Return(_a0)}
}

func (_m *ClientBuilder) OnWithUncached(objs ...client.Object) *ClientBuilder_WithUncached {
	c_call := _m.On("WithUncached", objs)
	return &ClientBuilder_WithUncached{Call: c_call}
}

func (_m *ClientBuilder) OnWithUncachedMatch(matchers ...interface{}) *ClientBuilder_WithUncached {
	c_call := _m.On("WithUncached", matchers...)
	return &ClientBuilder_WithUncached{Call: c_call}
}

// WithUncached provides a mock function with given fields: objs
func (_m *ClientBuilder) WithUncached(objs ...client.Object) k8s.ClientBuilder {
	_va := make([]interface{}, len(objs))
	for _i := range objs {
		_va[_i] = objs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 k8s.ClientBuilder
	if rf, ok := ret.Get(0).(func(...client.Object) k8s.ClientBuilder); ok {
		r0 = rf(objs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(k8s.ClientBuilder)
		}
	}

	return r0
}
