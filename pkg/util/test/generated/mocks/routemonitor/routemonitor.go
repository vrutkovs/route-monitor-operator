// Code generated by MockGen. DO NOT EDIT.
// Source: routemonitor_interface.go

// Package routemonitor is a generated GoMock package.
package routemonitor

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/route/v1"
	v1alpha1 "github.com/openshift/route-monitor-operator/api/v1alpha1"
	blackbox "github.com/openshift/route-monitor-operator/pkg/const/blackbox"
	reconcile "github.com/openshift/route-monitor-operator/pkg/util/reconcile"
	reflect "reflect"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

// MockRouteMonitorSupplement is a mock of RouteMonitorSupplement interface
type MockRouteMonitorSupplement struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMonitorSupplementMockRecorder
}

// MockRouteMonitorSupplementMockRecorder is the mock recorder for MockRouteMonitorSupplement
type MockRouteMonitorSupplementMockRecorder struct {
	mock *MockRouteMonitorSupplement
}

// NewMockRouteMonitorSupplement creates a new mock instance
func NewMockRouteMonitorSupplement(ctrl *gomock.Controller) *MockRouteMonitorSupplement {
	mock := &MockRouteMonitorSupplement{ctrl: ctrl}
	mock.recorder = &MockRouteMonitorSupplementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteMonitorSupplement) EXPECT() *MockRouteMonitorSupplementMockRecorder {
	return m.recorder
}

// GetRouteMonitor mocks base method
func (m *MockRouteMonitorSupplement) GetRouteMonitor(ctx context.Context, req controllerruntime.Request) (v1alpha1.RouteMonitor, reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteMonitor", ctx, req)
	ret0, _ := ret[0].(v1alpha1.RouteMonitor)
	ret1, _ := ret[1].(reconcile.Result)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRouteMonitor indicates an expected call of GetRouteMonitor
func (mr *MockRouteMonitorSupplementMockRecorder) GetRouteMonitor(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteMonitor", reflect.TypeOf((*MockRouteMonitorSupplement)(nil).GetRouteMonitor), ctx, req)
}

// GetRoute mocks base method
func (m *MockRouteMonitorSupplement) GetRoute(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (v1.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoute", ctx, routeMonitor)
	ret0, _ := ret[0].(v1.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoute indicates an expected call of GetRoute
func (mr *MockRouteMonitorSupplementMockRecorder) GetRoute(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoute", reflect.TypeOf((*MockRouteMonitorSupplement)(nil).GetRoute), ctx, routeMonitor)
}

// EnsureRouteURLExists mocks base method
func (m *MockRouteMonitorSupplement) EnsureRouteURLExists(ctx context.Context, route v1.Route, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureRouteURLExists", ctx, route, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureRouteURLExists indicates an expected call of EnsureRouteURLExists
func (mr *MockRouteMonitorSupplementMockRecorder) EnsureRouteURLExists(ctx, route, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureRouteURLExists", reflect.TypeOf((*MockRouteMonitorSupplement)(nil).EnsureRouteURLExists), ctx, route, routeMonitor)
}

// EnsureFinalizerAbsent mocks base method
func (m *MockRouteMonitorSupplement) EnsureFinalizerAbsent(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFinalizerAbsent", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureFinalizerAbsent indicates an expected call of EnsureFinalizerAbsent
func (mr *MockRouteMonitorSupplementMockRecorder) EnsureFinalizerAbsent(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFinalizerAbsent", reflect.TypeOf((*MockRouteMonitorSupplement)(nil).EnsureFinalizerAbsent), ctx, routeMonitor)
}

// MockRouteMonitorDeleter is a mock of RouteMonitorDeleter interface
type MockRouteMonitorDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMonitorDeleterMockRecorder
}

// MockRouteMonitorDeleterMockRecorder is the mock recorder for MockRouteMonitorDeleter
type MockRouteMonitorDeleterMockRecorder struct {
	mock *MockRouteMonitorDeleter
}

// NewMockRouteMonitorDeleter creates a new mock instance
func NewMockRouteMonitorDeleter(ctrl *gomock.Controller) *MockRouteMonitorDeleter {
	mock := &MockRouteMonitorDeleter{ctrl: ctrl}
	mock.recorder = &MockRouteMonitorDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteMonitorDeleter) EXPECT() *MockRouteMonitorDeleterMockRecorder {
	return m.recorder
}

// ShouldDeleteBlackBoxExporterResources mocks base method
func (m *MockRouteMonitorDeleter) ShouldDeleteBlackBoxExporterResources(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (blackbox.ShouldDeleteBlackBoxExporter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldDeleteBlackBoxExporterResources", ctx, routeMonitor)
	ret0, _ := ret[0].(blackbox.ShouldDeleteBlackBoxExporter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldDeleteBlackBoxExporterResources indicates an expected call of ShouldDeleteBlackBoxExporterResources
func (mr *MockRouteMonitorDeleterMockRecorder) ShouldDeleteBlackBoxExporterResources(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldDeleteBlackBoxExporterResources", reflect.TypeOf((*MockRouteMonitorDeleter)(nil).ShouldDeleteBlackBoxExporterResources), ctx, routeMonitor)
}

// EnsureBlackBoxExporterDeploymentAbsent mocks base method
func (m *MockRouteMonitorDeleter) EnsureBlackBoxExporterDeploymentAbsent(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBlackBoxExporterDeploymentAbsent", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureBlackBoxExporterDeploymentAbsent indicates an expected call of EnsureBlackBoxExporterDeploymentAbsent
func (mr *MockRouteMonitorDeleterMockRecorder) EnsureBlackBoxExporterDeploymentAbsent(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBlackBoxExporterDeploymentAbsent", reflect.TypeOf((*MockRouteMonitorDeleter)(nil).EnsureBlackBoxExporterDeploymentAbsent), ctx)
}

// EnsureBlackBoxExporterServiceAbsent mocks base method
func (m *MockRouteMonitorDeleter) EnsureBlackBoxExporterServiceAbsent(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBlackBoxExporterServiceAbsent", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureBlackBoxExporterServiceAbsent indicates an expected call of EnsureBlackBoxExporterServiceAbsent
func (mr *MockRouteMonitorDeleterMockRecorder) EnsureBlackBoxExporterServiceAbsent(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBlackBoxExporterServiceAbsent", reflect.TypeOf((*MockRouteMonitorDeleter)(nil).EnsureBlackBoxExporterServiceAbsent), ctx)
}

// EnsureServiceMonitorResourceAbsent mocks base method
func (m *MockRouteMonitorDeleter) EnsureServiceMonitorResourceAbsent(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureServiceMonitorResourceAbsent", ctx, routeMonitor)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureServiceMonitorResourceAbsent indicates an expected call of EnsureServiceMonitorResourceAbsent
func (mr *MockRouteMonitorDeleterMockRecorder) EnsureServiceMonitorResourceAbsent(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureServiceMonitorResourceAbsent", reflect.TypeOf((*MockRouteMonitorDeleter)(nil).EnsureServiceMonitorResourceAbsent), ctx, routeMonitor)
}

// MockRouteMonitorAdder is a mock of RouteMonitorAdder interface
type MockRouteMonitorAdder struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMonitorAdderMockRecorder
}

// MockRouteMonitorAdderMockRecorder is the mock recorder for MockRouteMonitorAdder
type MockRouteMonitorAdderMockRecorder struct {
	mock *MockRouteMonitorAdder
}

// NewMockRouteMonitorAdder creates a new mock instance
func NewMockRouteMonitorAdder(ctrl *gomock.Controller) *MockRouteMonitorAdder {
	mock := &MockRouteMonitorAdder{ctrl: ctrl}
	mock.recorder = &MockRouteMonitorAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteMonitorAdder) EXPECT() *MockRouteMonitorAdderMockRecorder {
	return m.recorder
}

// EnsureBlackBoxExporterDeploymentExists mocks base method
func (m *MockRouteMonitorAdder) EnsureBlackBoxExporterDeploymentExists(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBlackBoxExporterDeploymentExists", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureBlackBoxExporterDeploymentExists indicates an expected call of EnsureBlackBoxExporterDeploymentExists
func (mr *MockRouteMonitorAdderMockRecorder) EnsureBlackBoxExporterDeploymentExists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBlackBoxExporterDeploymentExists", reflect.TypeOf((*MockRouteMonitorAdder)(nil).EnsureBlackBoxExporterDeploymentExists), ctx)
}

// EnsureBlackBoxExporterServiceExists mocks base method
func (m *MockRouteMonitorAdder) EnsureBlackBoxExporterServiceExists(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBlackBoxExporterServiceExists", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureBlackBoxExporterServiceExists indicates an expected call of EnsureBlackBoxExporterServiceExists
func (mr *MockRouteMonitorAdderMockRecorder) EnsureBlackBoxExporterServiceExists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBlackBoxExporterServiceExists", reflect.TypeOf((*MockRouteMonitorAdder)(nil).EnsureBlackBoxExporterServiceExists), ctx)
}

// EnsureServiceMonitorResourceExists mocks base method
func (m *MockRouteMonitorAdder) EnsureServiceMonitorResourceExists(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureServiceMonitorResourceExists", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureServiceMonitorResourceExists indicates an expected call of EnsureServiceMonitorResourceExists
func (mr *MockRouteMonitorAdderMockRecorder) EnsureServiceMonitorResourceExists(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureServiceMonitorResourceExists", reflect.TypeOf((*MockRouteMonitorAdder)(nil).EnsureServiceMonitorResourceExists), ctx, routeMonitor)
}
