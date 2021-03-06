// Code generated by protoc-gen-godin. DO NOT EDIT.
// source: example.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// EncodeFooRequest is responsible of encoding a domain-layer FooRequest into a protobuf message
func EncodeFooRequest(request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil FooRequest")
	}
	req := pbRequest.(*FooRequest)
}

// EncodeBarRequest is responsible of encoding a domain-layer BarRequest into a protobuf message
func EncodeBarRequest(request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil BarRequest")
	}
	req := pbRequest.(*BarRequest)
}
