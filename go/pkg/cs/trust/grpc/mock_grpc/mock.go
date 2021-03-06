// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/pkg/cs/trust/grpc (interfaces: ChainBuilder,RenewalRequestVerifier,Signer)

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	x509 "crypto/x509"
	gomock "github.com/golang/mock/gomock"
	control_plane "github.com/scionproto/scion/go/pkg/proto/control_plane"
	crypto "github.com/scionproto/scion/go/pkg/proto/crypto"
	reflect "reflect"
)

// MockChainBuilder is a mock of ChainBuilder interface
type MockChainBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockChainBuilderMockRecorder
}

// MockChainBuilderMockRecorder is the mock recorder for MockChainBuilder
type MockChainBuilderMockRecorder struct {
	mock *MockChainBuilder
}

// NewMockChainBuilder creates a new mock instance
func NewMockChainBuilder(ctrl *gomock.Controller) *MockChainBuilder {
	mock := &MockChainBuilder{ctrl: ctrl}
	mock.recorder = &MockChainBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainBuilder) EXPECT() *MockChainBuilderMockRecorder {
	return m.recorder
}

// CreateChain mocks base method
func (m *MockChainBuilder) CreateChain(arg0 context.Context, arg1 *x509.CertificateRequest) ([]*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChain", arg0, arg1)
	ret0, _ := ret[0].([]*x509.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChain indicates an expected call of CreateChain
func (mr *MockChainBuilderMockRecorder) CreateChain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChain", reflect.TypeOf((*MockChainBuilder)(nil).CreateChain), arg0, arg1)
}

// MockRenewalRequestVerifier is a mock of RenewalRequestVerifier interface
type MockRenewalRequestVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockRenewalRequestVerifierMockRecorder
}

// MockRenewalRequestVerifierMockRecorder is the mock recorder for MockRenewalRequestVerifier
type MockRenewalRequestVerifierMockRecorder struct {
	mock *MockRenewalRequestVerifier
}

// NewMockRenewalRequestVerifier creates a new mock instance
func NewMockRenewalRequestVerifier(ctrl *gomock.Controller) *MockRenewalRequestVerifier {
	mock := &MockRenewalRequestVerifier{ctrl: ctrl}
	mock.recorder = &MockRenewalRequestVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRenewalRequestVerifier) EXPECT() *MockRenewalRequestVerifierMockRecorder {
	return m.recorder
}

// VerifyChainRenewalRequest mocks base method
func (m *MockRenewalRequestVerifier) VerifyChainRenewalRequest(arg0 *control_plane.ChainRenewalRequest, arg1 [][]*x509.Certificate) (*x509.CertificateRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyChainRenewalRequest", arg0, arg1)
	ret0, _ := ret[0].(*x509.CertificateRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyChainRenewalRequest indicates an expected call of VerifyChainRenewalRequest
func (mr *MockRenewalRequestVerifierMockRecorder) VerifyChainRenewalRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyChainRenewalRequest", reflect.TypeOf((*MockRenewalRequestVerifier)(nil).VerifyChainRenewalRequest), arg0, arg1)
}

// MockSigner is a mock of Signer interface
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// Sign mocks base method
func (m *MockSigner) Sign(arg0 context.Context, arg1 []byte, arg2 ...[]byte) (*crypto.SignedMessage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sign", varargs...)
	ret0, _ := ret[0].(*crypto.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockSignerMockRecorder) Sign(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSigner)(nil).Sign), varargs...)
}
