// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mockusecase

import (
	"context"
	"ryusuke410/golang-cli-tryout/internal/usecase"

	mock "github.com/stretchr/testify/mock"
)

// NewIMathUseCase creates a new instance of IMathUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMathUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMathUseCase {
	mock := &IMathUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// IMathUseCase is an autogenerated mock type for the IMathUseCase type
type IMathUseCase struct {
	mock.Mock
}

type IMathUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *IMathUseCase) EXPECT() *IMathUseCase_Expecter {
	return &IMathUseCase_Expecter{mock: &_m.Mock}
}

// Add provides a mock function for the type IMathUseCase
func (_mock *IMathUseCase) Add(ctx context.Context, input *usecase.MathAddInput) (*usecase.MathAddOutput, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 *usecase.MathAddOutput
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.MathAddInput) (*usecase.MathAddOutput, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.MathAddInput) *usecase.MathAddOutput); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.MathAddOutput)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *usecase.MathAddInput) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// IMathUseCase_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type IMathUseCase_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *IMathUseCase_Expecter) Add(ctx interface{}, input interface{}) *IMathUseCase_Add_Call {
	return &IMathUseCase_Add_Call{Call: _e.mock.On("Add", ctx, input)}
}

func (_c *IMathUseCase_Add_Call) Run(run func(ctx context.Context, input *usecase.MathAddInput)) *IMathUseCase_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*usecase.MathAddInput))
	})
	return _c
}

func (_c *IMathUseCase_Add_Call) Return(mathAddOutput *usecase.MathAddOutput, err error) *IMathUseCase_Add_Call {
	_c.Call.Return(mathAddOutput, err)
	return _c
}

func (_c *IMathUseCase_Add_Call) RunAndReturn(run func(ctx context.Context, input *usecase.MathAddInput) (*usecase.MathAddOutput, error)) *IMathUseCase_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Multiply provides a mock function for the type IMathUseCase
func (_mock *IMathUseCase) Multiply(ctx context.Context, input *usecase.MathMultiplyInput) (*usecase.MathMultiplyOutput, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Multiply")
	}

	var r0 *usecase.MathMultiplyOutput
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.MathMultiplyInput) (*usecase.MathMultiplyOutput, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.MathMultiplyInput) *usecase.MathMultiplyOutput); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.MathMultiplyOutput)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *usecase.MathMultiplyInput) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// IMathUseCase_Multiply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Multiply'
type IMathUseCase_Multiply_Call struct {
	*mock.Call
}

// Multiply is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *IMathUseCase_Expecter) Multiply(ctx interface{}, input interface{}) *IMathUseCase_Multiply_Call {
	return &IMathUseCase_Multiply_Call{Call: _e.mock.On("Multiply", ctx, input)}
}

func (_c *IMathUseCase_Multiply_Call) Run(run func(ctx context.Context, input *usecase.MathMultiplyInput)) *IMathUseCase_Multiply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*usecase.MathMultiplyInput))
	})
	return _c
}

func (_c *IMathUseCase_Multiply_Call) Return(mathMultiplyOutput *usecase.MathMultiplyOutput, err error) *IMathUseCase_Multiply_Call {
	_c.Call.Return(mathMultiplyOutput, err)
	return _c
}

func (_c *IMathUseCase_Multiply_Call) RunAndReturn(run func(ctx context.Context, input *usecase.MathMultiplyInput) (*usecase.MathMultiplyOutput, error)) *IMathUseCase_Multiply_Call {
	_c.Call.Return(run)
	return _c
}

// NewITransformUseCase creates a new instance of ITransformUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITransformUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITransformUseCase {
	mock := &ITransformUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// ITransformUseCase is an autogenerated mock type for the ITransformUseCase type
type ITransformUseCase struct {
	mock.Mock
}

type ITransformUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *ITransformUseCase) EXPECT() *ITransformUseCase_Expecter {
	return &ITransformUseCase_Expecter{mock: &_m.Mock}
}

// Base64Encode provides a mock function for the type ITransformUseCase
func (_mock *ITransformUseCase) Base64Encode(ctx context.Context, input *usecase.TransformBase64Input) (*usecase.TransformBase64Output, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Base64Encode")
	}

	var r0 *usecase.TransformBase64Output
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.TransformBase64Input) (*usecase.TransformBase64Output, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.TransformBase64Input) *usecase.TransformBase64Output); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.TransformBase64Output)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *usecase.TransformBase64Input) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ITransformUseCase_Base64Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Base64Encode'
type ITransformUseCase_Base64Encode_Call struct {
	*mock.Call
}

// Base64Encode is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *ITransformUseCase_Expecter) Base64Encode(ctx interface{}, input interface{}) *ITransformUseCase_Base64Encode_Call {
	return &ITransformUseCase_Base64Encode_Call{Call: _e.mock.On("Base64Encode", ctx, input)}
}

func (_c *ITransformUseCase_Base64Encode_Call) Run(run func(ctx context.Context, input *usecase.TransformBase64Input)) *ITransformUseCase_Base64Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*usecase.TransformBase64Input))
	})
	return _c
}

func (_c *ITransformUseCase_Base64Encode_Call) Return(transformBase64Output *usecase.TransformBase64Output, err error) *ITransformUseCase_Base64Encode_Call {
	_c.Call.Return(transformBase64Output, err)
	return _c
}

func (_c *ITransformUseCase_Base64Encode_Call) RunAndReturn(run func(ctx context.Context, input *usecase.TransformBase64Input) (*usecase.TransformBase64Output, error)) *ITransformUseCase_Base64Encode_Call {
	_c.Call.Return(run)
	return _c
}

// Escape provides a mock function for the type ITransformUseCase
func (_mock *ITransformUseCase) Escape(ctx context.Context, input *usecase.TransformEscapeInput) (*usecase.TransformEscapeOutput, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Escape")
	}

	var r0 *usecase.TransformEscapeOutput
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.TransformEscapeInput) (*usecase.TransformEscapeOutput, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.TransformEscapeInput) *usecase.TransformEscapeOutput); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.TransformEscapeOutput)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *usecase.TransformEscapeInput) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ITransformUseCase_Escape_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Escape'
type ITransformUseCase_Escape_Call struct {
	*mock.Call
}

// Escape is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *ITransformUseCase_Expecter) Escape(ctx interface{}, input interface{}) *ITransformUseCase_Escape_Call {
	return &ITransformUseCase_Escape_Call{Call: _e.mock.On("Escape", ctx, input)}
}

func (_c *ITransformUseCase_Escape_Call) Run(run func(ctx context.Context, input *usecase.TransformEscapeInput)) *ITransformUseCase_Escape_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*usecase.TransformEscapeInput))
	})
	return _c
}

func (_c *ITransformUseCase_Escape_Call) Return(transformEscapeOutput *usecase.TransformEscapeOutput, err error) *ITransformUseCase_Escape_Call {
	_c.Call.Return(transformEscapeOutput, err)
	return _c
}

func (_c *ITransformUseCase_Escape_Call) RunAndReturn(run func(ctx context.Context, input *usecase.TransformEscapeInput) (*usecase.TransformEscapeOutput, error)) *ITransformUseCase_Escape_Call {
	_c.Call.Return(run)
	return _c
}

// UrlEncode provides a mock function for the type ITransformUseCase
func (_mock *ITransformUseCase) UrlEncode(ctx context.Context, input *usecase.TransformUrlEncodeInput) (*usecase.TransformUrlEncodeOutput, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for UrlEncode")
	}

	var r0 *usecase.TransformUrlEncodeOutput
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.TransformUrlEncodeInput) (*usecase.TransformUrlEncodeOutput, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *usecase.TransformUrlEncodeInput) *usecase.TransformUrlEncodeOutput); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.TransformUrlEncodeOutput)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *usecase.TransformUrlEncodeInput) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ITransformUseCase_UrlEncode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UrlEncode'
type ITransformUseCase_UrlEncode_Call struct {
	*mock.Call
}

// UrlEncode is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *ITransformUseCase_Expecter) UrlEncode(ctx interface{}, input interface{}) *ITransformUseCase_UrlEncode_Call {
	return &ITransformUseCase_UrlEncode_Call{Call: _e.mock.On("UrlEncode", ctx, input)}
}

func (_c *ITransformUseCase_UrlEncode_Call) Run(run func(ctx context.Context, input *usecase.TransformUrlEncodeInput)) *ITransformUseCase_UrlEncode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*usecase.TransformUrlEncodeInput))
	})
	return _c
}

func (_c *ITransformUseCase_UrlEncode_Call) Return(transformUrlEncodeOutput *usecase.TransformUrlEncodeOutput, err error) *ITransformUseCase_UrlEncode_Call {
	_c.Call.Return(transformUrlEncodeOutput, err)
	return _c
}

func (_c *ITransformUseCase_UrlEncode_Call) RunAndReturn(run func(ctx context.Context, input *usecase.TransformUrlEncodeInput) (*usecase.TransformUrlEncodeOutput, error)) *ITransformUseCase_UrlEncode_Call {
	_c.Call.Return(run)
	return _c
}
