// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/kyl2016/tusd/s3store (interfaces: S3API)

package s3store

import (
	s3 "github.com/aws/aws-sdk-go/service/s3"
	gomock "github.com/golang/mock/gomock"
)

// Mock of S3API interface
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *_MockS3APIRecorder
}

// Recorder for MockS3API (not exported)
type _MockS3APIRecorder struct {
	mock *MockS3API
}

func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &_MockS3APIRecorder{mock}
	return mock
}

func (_m *MockS3API) EXPECT() *_MockS3APIRecorder {
	return _m.recorder
}

func (_m *MockS3API) AbortMultipartUpload(_param0 *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	ret := _m.ctrl.Call(_m, "AbortMultipartUpload", _param0)
	ret0, _ := ret[0].(*s3.AbortMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) AbortMultipartUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AbortMultipartUpload", arg0)
}

func (_m *MockS3API) CompleteMultipartUpload(_param0 *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	ret := _m.ctrl.Call(_m, "CompleteMultipartUpload", _param0)
	ret0, _ := ret[0].(*s3.CompleteMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) CompleteMultipartUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CompleteMultipartUpload", arg0)
}

func (_m *MockS3API) CreateMultipartUpload(_param0 *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateMultipartUpload", _param0)
	ret0, _ := ret[0].(*s3.CreateMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) CreateMultipartUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateMultipartUpload", arg0)
}

func (_m *MockS3API) DeleteObject(_param0 *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteObject", _param0)
	ret0, _ := ret[0].(*s3.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteObject", arg0)
}

func (_m *MockS3API) DeleteObjects(_param0 *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteObjects", _param0)
	ret0, _ := ret[0].(*s3.DeleteObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) DeleteObjects(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteObjects", arg0)
}

func (_m *MockS3API) GetObject(_param0 *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	ret := _m.ctrl.Call(_m, "GetObject", _param0)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) GetObject(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetObject", arg0)
}

func (_m *MockS3API) ListParts(_param0 *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	ret := _m.ctrl.Call(_m, "ListParts", _param0)
	ret0, _ := ret[0].(*s3.ListPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) ListParts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListParts", arg0)
}

func (_m *MockS3API) PutObject(_param0 *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	ret := _m.ctrl.Call(_m, "PutObject", _param0)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) PutObject(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutObject", arg0)
}

func (_m *MockS3API) UploadPart(_param0 *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	ret := _m.ctrl.Call(_m, "UploadPart", _param0)
	ret0, _ := ret[0].(*s3.UploadPartOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) UploadPart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UploadPart", arg0)
}

func (_m *MockS3API) UploadPartCopy(_param0 *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	ret := _m.ctrl.Call(_m, "UploadPartCopy", _param0)
	ret0, _ := ret[0].(*s3.UploadPartCopyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockS3APIRecorder) UploadPartCopy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UploadPartCopy", arg0)
}
