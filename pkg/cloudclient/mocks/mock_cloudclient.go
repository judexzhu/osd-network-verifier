// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloudclient/cloudclient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	output "github.com/openshift/osd-network-verifier/pkg/output"
	proxy "github.com/openshift/osd-network-verifier/pkg/proxy"
)

// MockCloudClient is a mock of CloudClient interface.
type MockCloudClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudClientMockRecorder
}

// MockCloudClientMockRecorder is the mock recorder for MockCloudClient.
type MockCloudClientMockRecorder struct {
	mock *MockCloudClient
}

// NewMockCloudClient creates a new mock instance.
func NewMockCloudClient(ctrl *gomock.Controller) *MockCloudClient {
	mock := &MockCloudClient{ctrl: ctrl}
	mock.recorder = &MockCloudClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudClient) EXPECT() *MockCloudClientMockRecorder {
	return m.recorder
}

// ByoVPCValidator mocks base method.
func (m *MockCloudClient) ByoVPCValidator(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByoVPCValidator", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ByoVPCValidator indicates an expected call of ByoVPCValidator.
func (mr *MockCloudClientMockRecorder) ByoVPCValidator(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByoVPCValidator", reflect.TypeOf((*MockCloudClient)(nil).ByoVPCValidator), ctx)
}

// ValidateEgress mocks base method.
func (m *MockCloudClient) ValidateEgress(ctx context.Context, vpcSubnetID, cloudImageID, kmsKeyID string, timeout time.Duration, proxy proxy.ProxyConfig) *output.Output {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateEgress", ctx, vpcSubnetID, cloudImageID, kmsKeyID, timeout, proxy)
	ret0, _ := ret[0].(*output.Output)
	return ret0
}

// ValidateEgress indicates an expected call of ValidateEgress.
func (mr *MockCloudClientMockRecorder) ValidateEgress(ctx, vpcSubnetID, cloudImageID, kmsKeyID, timeout, proxy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateEgress", reflect.TypeOf((*MockCloudClient)(nil).ValidateEgress), ctx, vpcSubnetID, cloudImageID, kmsKeyID, timeout, proxy)
}

// VerifyDns mocks base method.
func (m *MockCloudClient) VerifyDns(ctx context.Context, vpcID string) *output.Output {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyDns", ctx, vpcID)
	ret0, _ := ret[0].(*output.Output)
	return ret0
}

// VerifyDns indicates an expected call of VerifyDns.
func (mr *MockCloudClientMockRecorder) VerifyDns(ctx, vpcID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyDns", reflect.TypeOf((*MockCloudClient)(nil).VerifyDns), ctx, vpcID)
}
