// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/city-mobil/gobuns/redis (interfaces: Redis)

// Package mock_redis is a generated GoMock package.
package mock_redis

import (
	context "context"
	reflect "reflect"
	time "time"

	redis "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockRedis is a mock of Redis interface.
type MockRedis struct {
	ctrl     *gomock.Controller
	recorder *MockRedisMockRecorder
}

// MockRedisMockRecorder is the mock recorder for MockRedis.
type MockRedisMockRecorder struct {
	mock *MockRedis
}

// NewMockRedis creates a new mock instance.
func NewMockRedis(ctrl *gomock.Controller) *MockRedis {
	mock := &MockRedis{ctrl: ctrl}
	mock.recorder = &MockRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedis) EXPECT() *MockRedisMockRecorder {
	return m.recorder
}

// Decr mocks base method.
func (m *MockRedis) Decr(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decr indicates an expected call of Decr.
func (mr *MockRedisMockRecorder) Decr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockRedis)(nil).Decr), arg0, arg1)
}

// DecrBy mocks base method.
func (m *MockRedis) DecrBy(arg0 context.Context, arg1 string, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockRedisMockRecorder) DecrBy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockRedis)(nil).DecrBy), arg0, arg1, arg2)
}

// Del mocks base method.
func (m *MockRedis) Del(arg0 context.Context, arg1 ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del.
func (mr *MockRedisMockRecorder) Del(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedis)(nil).Del), varargs...)
}

// Eval mocks base method.
func (m *MockRedis) Eval(arg0 context.Context, arg1 string, arg2 []string, arg3 []interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eval", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eval indicates an expected call of Eval.
func (mr *MockRedisMockRecorder) Eval(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockRedis)(nil).Eval), arg0, arg1, arg2, arg3)
}

// EvalSha mocks base method.
func (m *MockRedis) EvalSha(arg0 context.Context, arg1 string, arg2 []string, arg3 []interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvalSha", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvalSha indicates an expected call of EvalSha.
func (mr *MockRedisMockRecorder) EvalSha(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalSha", reflect.TypeOf((*MockRedis)(nil).EvalSha), arg0, arg1, arg2, arg3)
}

// Exists mocks base method.
func (m *MockRedis) Exists(arg0 context.Context, arg1 ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRedisMockRecorder) Exists(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRedis)(nil).Exists), varargs...)
}

// Expire mocks base method.
func (m *MockRedis) Expire(arg0 context.Context, arg1 string, arg2 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expire indicates an expected call of Expire.
func (mr *MockRedisMockRecorder) Expire(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockRedis)(nil).Expire), arg0, arg1, arg2)
}

// ExpireAt mocks base method.
func (m *MockRedis) ExpireAt(arg0 context.Context, arg1 string, arg2 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpireAt indicates an expected call of ExpireAt.
func (mr *MockRedisMockRecorder) ExpireAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireAt", reflect.TypeOf((*MockRedis)(nil).ExpireAt), arg0, arg1, arg2)
}

// GeoAdd mocks base method.
func (m *MockRedis) GeoAdd(arg0 context.Context, arg1 string, arg2 ...*redis.GeoLocation) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoAdd", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeoAdd indicates an expected call of GeoAdd.
func (mr *MockRedisMockRecorder) GeoAdd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoAdd", reflect.TypeOf((*MockRedis)(nil).GeoAdd), varargs...)
}

// GeoRadius mocks base method.
func (m *MockRedis) GeoRadius(arg0 context.Context, arg1 string, arg2, arg3 float64, arg4 *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadius", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]redis.GeoLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeoRadius indicates an expected call of GeoRadius.
func (mr *MockRedisMockRecorder) GeoRadius(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadius", reflect.TypeOf((*MockRedis)(nil).GeoRadius), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method.
func (m *MockRedis) Get(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRedisMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedis)(nil).Get), arg0, arg1)
}

// HDel mocks base method.
func (m *MockRedis) HDel(arg0 context.Context, arg1 string, arg2 ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HDel", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HDel indicates an expected call of HDel.
func (mr *MockRedisMockRecorder) HDel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockRedis)(nil).HDel), varargs...)
}

// HGet mocks base method.
func (m *MockRedis) HGet(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HGet indicates an expected call of HGet.
func (mr *MockRedisMockRecorder) HGet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockRedis)(nil).HGet), arg0, arg1, arg2)
}

// HKeys mocks base method.
func (m *MockRedis) HKeys(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HKeys", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HKeys indicates an expected call of HKeys.
func (mr *MockRedisMockRecorder) HKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HKeys", reflect.TypeOf((*MockRedis)(nil).HKeys), arg0, arg1)
}

// HLen mocks base method.
func (m *MockRedis) HLen(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HLen", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HLen indicates an expected call of HLen.
func (mr *MockRedisMockRecorder) HLen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HLen", reflect.TypeOf((*MockRedis)(nil).HLen), arg0, arg1)
}

// HMSet mocks base method.
func (m *MockRedis) HMSet(arg0 context.Context, arg1 string, arg2 []interface{}, arg3 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HMSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HMSet indicates an expected call of HMSet.
func (mr *MockRedisMockRecorder) HMSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSet", reflect.TypeOf((*MockRedis)(nil).HMSet), arg0, arg1, arg2, arg3)
}

// HSet mocks base method.
func (m *MockRedis) HSet(arg0 context.Context, arg1, arg2 string, arg3 interface{}, arg4 time.Duration) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSet", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HSet indicates an expected call of HSet.
func (mr *MockRedisMockRecorder) HSet(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockRedis)(nil).HSet), arg0, arg1, arg2, arg3, arg4)
}

// HSetNX mocks base method.
func (m *MockRedis) HSetNX(arg0 context.Context, arg1, arg2 string, arg3 interface{}, arg4 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSetNX", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HSetNX indicates an expected call of HSetNX.
func (mr *MockRedisMockRecorder) HSetNX(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSetNX", reflect.TypeOf((*MockRedis)(nil).HSetNX), arg0, arg1, arg2, arg3, arg4)
}

// Incr mocks base method.
func (m *MockRedis) Incr(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Incr indicates an expected call of Incr.
func (mr *MockRedisMockRecorder) Incr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockRedis)(nil).Incr), arg0, arg1)
}

// IncrBy mocks base method.
func (m *MockRedis) IncrBy(arg0 context.Context, arg1 string, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockRedisMockRecorder) IncrBy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockRedis)(nil).IncrBy), arg0, arg1, arg2)
}

// MGet mocks base method.
func (m *MockRedis) MGet(arg0 context.Context, arg1 ...string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGet", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGet indicates an expected call of MGet.
func (mr *MockRedisMockRecorder) MGet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockRedis)(nil).MGet), varargs...)
}

// MSet mocks base method.
func (m *MockRedis) MSet(arg0 context.Context, arg1 []interface{}, arg2 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MSet indicates an expected call of MSet.
func (mr *MockRedisMockRecorder) MSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSet", reflect.TypeOf((*MockRedis)(nil).MSet), arg0, arg1, arg2)
}

// SAdd mocks base method.
func (m *MockRedis) SAdd(arg0 context.Context, arg1 string, arg2 []interface{}, arg3 time.Duration) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SAdd", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SAdd indicates an expected call of SAdd.
func (mr *MockRedisMockRecorder) SAdd(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockRedis)(nil).SAdd), arg0, arg1, arg2, arg3)
}

// SCard mocks base method.
func (m *MockRedis) SCard(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SCard", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SCard indicates an expected call of SCard.
func (mr *MockRedisMockRecorder) SCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SCard", reflect.TypeOf((*MockRedis)(nil).SCard), arg0, arg1)
}

// SIsMember mocks base method.
func (m *MockRedis) SIsMember(arg0 context.Context, arg1 string, arg2 interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIsMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SIsMember indicates an expected call of SIsMember.
func (mr *MockRedisMockRecorder) SIsMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIsMember", reflect.TypeOf((*MockRedis)(nil).SIsMember), arg0, arg1, arg2)
}

// SMembers mocks base method.
func (m *MockRedis) SMembers(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SMembers indicates an expected call of SMembers.
func (mr *MockRedisMockRecorder) SMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockRedis)(nil).SMembers), arg0, arg1)
}

// SRem mocks base method.
func (m *MockRedis) SRem(arg0 context.Context, arg1 string, arg2 ...interface{}) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SRem indicates an expected call of SRem.
func (mr *MockRedisMockRecorder) SRem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockRedis)(nil).SRem), varargs...)
}

// Set mocks base method.
func (m *MockRedis) Set(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockRedisMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedis)(nil).Set), arg0, arg1, arg2, arg3)
}

// SetNX mocks base method.
func (m *MockRedis) SetNX(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX.
func (mr *MockRedisMockRecorder) SetNX(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockRedis)(nil).SetNX), arg0, arg1, arg2, arg3)
}

// StrLen mocks base method.
func (m *MockRedis) StrLen(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StrLen", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StrLen indicates an expected call of StrLen.
func (mr *MockRedisMockRecorder) StrLen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrLen", reflect.TypeOf((*MockRedis)(nil).StrLen), arg0, arg1)
}

// ZAdd mocks base method.
func (m *MockRedis) ZAdd(arg0 context.Context, arg1 string, arg2 []*redis.Z, arg3 time.Duration) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAdd", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockRedisMockRecorder) ZAdd(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockRedis)(nil).ZAdd), arg0, arg1, arg2, arg3)
}

// ZCard mocks base method.
func (m *MockRedis) ZCard(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCard", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZCard indicates an expected call of ZCard.
func (mr *MockRedisMockRecorder) ZCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCard", reflect.TypeOf((*MockRedis)(nil).ZCard), arg0, arg1)
}

// ZRangeByScore mocks base method.
func (m *MockRedis) ZRangeByScore(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockRedisMockRecorder) ZRangeByScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockRedis)(nil).ZRangeByScore), arg0, arg1, arg2)
}

// ZRem mocks base method.
func (m *MockRedis) ZRem(arg0 context.Context, arg1 string, arg2 ...interface{}) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZRem", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRem indicates an expected call of ZRem.
func (mr *MockRedisMockRecorder) ZRem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockRedis)(nil).ZRem), varargs...)
}

// ZRemRangeByScore mocks base method.
func (m *MockRedis) ZRemRangeByScore(arg0 context.Context, arg1, arg2, arg3 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByScore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRemRangeByScore indicates an expected call of ZRemRangeByScore.
func (mr *MockRedisMockRecorder) ZRemRangeByScore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByScore", reflect.TypeOf((*MockRedis)(nil).ZRemRangeByScore), arg0, arg1, arg2, arg3)
}