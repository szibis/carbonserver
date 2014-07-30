// Code generated by protoc-gen-gogo.
// source: carbonserver.proto
// DO NOT EDIT!

/*
Package carbonserverpb is a generated protocol buffer package.

It is generated from these files:
	carbonserver.proto

It has these top-level messages:
	FetchResponse
	GlobMatch
	GlobResponse
	Retention
	InfoResponse
*/
package carbonserverpb

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type FetchResponse struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	StartTime        *int32    `protobuf:"varint,2,req,name=startTime" json:"startTime,omitempty"`
	StopTime         *int32    `protobuf:"varint,3,req,name=stopTime" json:"stopTime,omitempty"`
	StepTime         *int32    `protobuf:"varint,4,req,name=stepTime" json:"stepTime,omitempty"`
	Values           []float64 `protobuf:"fixed64,5,rep,name=values" json:"values,omitempty"`
	IsAbsent         []bool    `protobuf:"varint,6,rep,name=isAbsent" json:"isAbsent,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}

func (m *FetchResponse) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FetchResponse) GetStartTime() int32 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *FetchResponse) GetStopTime() int32 {
	if m != nil && m.StopTime != nil {
		return *m.StopTime
	}
	return 0
}

func (m *FetchResponse) GetStepTime() int32 {
	if m != nil && m.StepTime != nil {
		return *m.StepTime
	}
	return 0
}

func (m *FetchResponse) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *FetchResponse) GetIsAbsent() []bool {
	if m != nil {
		return m.IsAbsent
	}
	return nil
}

type GlobMatch struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	IsLeaf           *bool   `protobuf:"varint,2,req,name=isLeaf" json:"isLeaf,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GlobMatch) Reset()         { *m = GlobMatch{} }
func (m *GlobMatch) String() string { return proto.CompactTextString(m) }
func (*GlobMatch) ProtoMessage()    {}

func (m *GlobMatch) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *GlobMatch) GetIsLeaf() bool {
	if m != nil && m.IsLeaf != nil {
		return *m.IsLeaf
	}
	return false
}

type GlobResponse struct {
	Name             *string      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Matches          []*GlobMatch `protobuf:"bytes,2,rep,name=matches" json:"matches,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GlobResponse) Reset()         { *m = GlobResponse{} }
func (m *GlobResponse) String() string { return proto.CompactTextString(m) }
func (*GlobResponse) ProtoMessage()    {}

func (m *GlobResponse) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GlobResponse) GetMatches() []*GlobMatch {
	if m != nil {
		return m.Matches
	}
	return nil
}

type Retention struct {
	SecondsPerPoint  *int32 `protobuf:"varint,1,req,name=secondsPerPoint" json:"secondsPerPoint,omitempty"`
	NumberOfPoints   *int32 `protobuf:"varint,2,req,name=numberOfPoints" json:"numberOfPoints,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Retention) Reset()         { *m = Retention{} }
func (m *Retention) String() string { return proto.CompactTextString(m) }
func (*Retention) ProtoMessage()    {}

func (m *Retention) GetSecondsPerPoint() int32 {
	if m != nil && m.SecondsPerPoint != nil {
		return *m.SecondsPerPoint
	}
	return 0
}

func (m *Retention) GetNumberOfPoints() int32 {
	if m != nil && m.NumberOfPoints != nil {
		return *m.NumberOfPoints
	}
	return 0
}

type InfoResponse struct {
	Name              *string      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	AggregationMethod *string      `protobuf:"bytes,2,req,name=aggregationMethod" json:"aggregationMethod,omitempty"`
	MaxRetention      *int32       `protobuf:"varint,3,req,name=maxRetention" json:"maxRetention,omitempty"`
	XFilesFactor      *float32     `protobuf:"fixed32,4,req,name=xFilesFactor" json:"xFilesFactor,omitempty"`
	Retentions        []*Retention `protobuf:"bytes,5,rep,name=retentions" json:"retentions,omitempty"`
	XXX_unrecognized  []byte       `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}

func (m *InfoResponse) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *InfoResponse) GetAggregationMethod() string {
	if m != nil && m.AggregationMethod != nil {
		return *m.AggregationMethod
	}
	return ""
}

func (m *InfoResponse) GetMaxRetention() int32 {
	if m != nil && m.MaxRetention != nil {
		return *m.MaxRetention
	}
	return 0
}

func (m *InfoResponse) GetXFilesFactor() float32 {
	if m != nil && m.XFilesFactor != nil {
		return *m.XFilesFactor
	}
	return 0
}

func (m *InfoResponse) GetRetentions() []*Retention {
	if m != nil {
		return m.Retentions
	}
	return nil
}

func init() {
}
