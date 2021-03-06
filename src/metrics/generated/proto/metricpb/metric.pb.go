// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/metrics/generated/proto/metricpb/metric.proto

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metricpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MetricType int32

const (
	MetricType_UNKNOWN MetricType = 0
	MetricType_COUNTER MetricType = 1
	MetricType_TIMER   MetricType = 2
	MetricType_GAUGE   MetricType = 3
)

var MetricType_name = map[int32]string{
	0: "UNKNOWN",
	1: "COUNTER",
	2: "TIMER",
	3: "GAUGE",
}
var MetricType_value = map[string]int32{
	"UNKNOWN": 0,
	"COUNTER": 1,
	"TIMER":   2,
	"GAUGE":   3,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (MetricType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetric, []int{0} }

type Counter struct {
	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Counter) Reset()                    { *m = Counter{} }
func (m *Counter) String() string            { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()               {}
func (*Counter) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{0} }

func (m *Counter) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Counter) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BatchTimer struct {
	Id     []byte    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Values []float64 `protobuf:"fixed64,2,rep,packed,name=values" json:"values,omitempty"`
}

func (m *BatchTimer) Reset()                    { *m = BatchTimer{} }
func (m *BatchTimer) String() string            { return proto.CompactTextString(m) }
func (*BatchTimer) ProtoMessage()               {}
func (*BatchTimer) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{1} }

func (m *BatchTimer) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BatchTimer) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Gauge struct {
	Id    []byte  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Gauge) Reset()                    { *m = Gauge{} }
func (m *Gauge) String() string            { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()               {}
func (*Gauge) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{2} }

func (m *Gauge) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Gauge) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TimedMetric struct {
	Type       MetricType `protobuf:"varint,1,opt,name=type,proto3,enum=metricpb.MetricType" json:"type,omitempty"`
	Id         []byte     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TimeNanos  int64      `protobuf:"varint,3,opt,name=time_nanos,json=timeNanos,proto3" json:"time_nanos,omitempty"`
	Value      float64    `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
	Annotation []byte     `protobuf:"bytes,5,opt,name=annotation,proto3" json:"annotation,omitempty"`
}

func (m *TimedMetric) Reset()                    { *m = TimedMetric{} }
func (m *TimedMetric) String() string            { return proto.CompactTextString(m) }
func (*TimedMetric) ProtoMessage()               {}
func (*TimedMetric) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{3} }

func (m *TimedMetric) GetType() MetricType {
	if m != nil {
		return m.Type
	}
	return MetricType_UNKNOWN
}

func (m *TimedMetric) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TimedMetric) GetTimeNanos() int64 {
	if m != nil {
		return m.TimeNanos
	}
	return 0
}

func (m *TimedMetric) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TimedMetric) GetAnnotation() []byte {
	if m != nil {
		return m.Annotation
	}
	return nil
}

type ForwardedMetric struct {
	Type       MetricType `protobuf:"varint,1,opt,name=type,proto3,enum=metricpb.MetricType" json:"type,omitempty"`
	Id         []byte     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TimeNanos  int64      `protobuf:"varint,3,opt,name=time_nanos,json=timeNanos,proto3" json:"time_nanos,omitempty"`
	Values     []float64  `protobuf:"fixed64,4,rep,packed,name=values" json:"values,omitempty"`
	Annotation []byte     `protobuf:"bytes,5,opt,name=annotation,proto3" json:"annotation,omitempty"`
}

func (m *ForwardedMetric) Reset()                    { *m = ForwardedMetric{} }
func (m *ForwardedMetric) String() string            { return proto.CompactTextString(m) }
func (*ForwardedMetric) ProtoMessage()               {}
func (*ForwardedMetric) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{4} }

func (m *ForwardedMetric) GetType() MetricType {
	if m != nil {
		return m.Type
	}
	return MetricType_UNKNOWN
}

func (m *ForwardedMetric) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ForwardedMetric) GetTimeNanos() int64 {
	if m != nil {
		return m.TimeNanos
	}
	return 0
}

func (m *ForwardedMetric) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ForwardedMetric) GetAnnotation() []byte {
	if m != nil {
		return m.Annotation
	}
	return nil
}

type Tag struct {
	Name  []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{5} }

func (m *Tag) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Tag) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Counter)(nil), "metricpb.Counter")
	proto.RegisterType((*BatchTimer)(nil), "metricpb.BatchTimer")
	proto.RegisterType((*Gauge)(nil), "metricpb.Gauge")
	proto.RegisterType((*TimedMetric)(nil), "metricpb.TimedMetric")
	proto.RegisterType((*ForwardedMetric)(nil), "metricpb.ForwardedMetric")
	proto.RegisterType((*Tag)(nil), "metricpb.Tag")
	proto.RegisterEnum("metricpb.MetricType", MetricType_name, MetricType_value)
}
func (m *Counter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Counter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Value != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *BatchTimer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchTimer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Values) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Values)*8))
		for _, num := range m.Values {
			f1 := math.Float64bits(float64(num))
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(f1))
			i += 8
		}
	}
	return i, nil
}

func (m *Gauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gauge) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Value != 0 {
		dAtA[i] = 0x11
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i += 8
	}
	return i, nil
}

func (m *TimedMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedMetric) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.Type))
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.TimeNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.TimeNanos))
	}
	if m.Value != 0 {
		dAtA[i] = 0x21
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i += 8
	}
	if len(m.Annotation) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Annotation)))
		i += copy(dAtA[i:], m.Annotation)
	}
	return i, nil
}

func (m *ForwardedMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardedMetric) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.Type))
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.TimeNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.TimeNanos))
	}
	if len(m.Values) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Values)*8))
		for _, num := range m.Values {
			f2 := math.Float64bits(float64(num))
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(f2))
			i += 8
		}
	}
	if len(m.Annotation) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Annotation)))
		i += copy(dAtA[i:], m.Annotation)
	}
	return i, nil
}

func (m *Tag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tag) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeVarintMetric(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Counter) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovMetric(uint64(m.Value))
	}
	return n
}

func (m *BatchTimer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if len(m.Values) > 0 {
		n += 1 + sovMetric(uint64(len(m.Values)*8)) + len(m.Values)*8
	}
	return n
}

func (m *Gauge) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *TimedMetric) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMetric(uint64(m.Type))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.TimeNanos != 0 {
		n += 1 + sovMetric(uint64(m.TimeNanos))
	}
	if m.Value != 0 {
		n += 9
	}
	l = len(m.Annotation)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func (m *ForwardedMetric) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMetric(uint64(m.Type))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.TimeNanos != 0 {
		n += 1 + sovMetric(uint64(m.TimeNanos))
	}
	if len(m.Values) > 0 {
		n += 1 + sovMetric(uint64(len(m.Values)*8)) + len(m.Values)*8
	}
	l = len(m.Annotation)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func (m *Tag) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func sovMetric(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetric(x uint64) (n int) {
	return sovMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Counter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Counter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Counter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BatchTimer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BatchTimer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchTimer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Values = append(m.Values, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMetric
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMetric
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Values = append(m.Values, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Gauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Gauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimedMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimedMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (MetricType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeNanos", wireType)
			}
			m.TimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotation", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Annotation = append(m.Annotation[:0], dAtA[iNdEx:postIndex]...)
			if m.Annotation == nil {
				m.Annotation = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ForwardedMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ForwardedMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardedMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (MetricType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeNanos", wireType)
			}
			m.TimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Values = append(m.Values, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMetric
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMetric
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Values = append(m.Values, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotation", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Annotation = append(m.Annotation[:0], dAtA[iNdEx:postIndex]...)
			if m.Annotation == nil {
				m.Annotation = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Tag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetric(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetric
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetric
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetric(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetric = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetric   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/metrics/generated/proto/metricpb/metric.proto", fileDescriptorMetric)
}

var fileDescriptorMetric = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0xed, 0xe4, 0xa7, 0xb5, 0xb7, 0xa5, 0x86, 0xa1, 0x48, 0x36, 0x86, 0x92, 0x55, 0x10, 0xcc,
	0x80, 0x15, 0xdc, 0xb8, 0xb1, 0xb5, 0x96, 0x22, 0x4d, 0x21, 0xa4, 0x08, 0x6e, 0x64, 0x92, 0x0c,
	0x69, 0xc0, 0xcc, 0x84, 0x64, 0xa2, 0xf4, 0x2d, 0x7c, 0x00, 0x57, 0x3e, 0x8d, 0x4b, 0x1f, 0x41,
	0xea, 0x8b, 0x48, 0xa6, 0xa9, 0xad, 0x20, 0x0a, 0xc2, 0xb7, 0xbb, 0xe7, 0xcc, 0xbd, 0xf7, 0x9c,
	0x7b, 0x18, 0x78, 0x99, 0xe5, 0xf2, 0xd0, 0xc4, 0x7e, 0x22, 0x0a, 0x52, 0xcc, 0xd3, 0x98, 0x14,
	0x73, 0x52, 0x57, 0x09, 0x29, 0x98, 0xac, 0xf2, 0xa4, 0x26, 0x19, 0xe3, 0xac, 0xa2, 0x92, 0xa5,
	0xa4, 0xac, 0x84, 0x14, 0x1d, 0x5f, 0xc6, 0x5d, 0xe1, 0x2b, 0x16, 0xdf, 0xbb, 0xd0, 0x2e, 0x81,
	0xc1, 0x52, 0x34, 0x5c, 0xb2, 0x0a, 0x4f, 0x40, 0xcb, 0x53, 0x1b, 0xcd, 0x90, 0x37, 0x0e, 0xb5,
	0x3c, 0xc5, 0x53, 0x30, 0x3f, 0xd0, 0xf7, 0x0d, 0xb3, 0xb5, 0x19, 0xf2, 0xf4, 0xf0, 0x0c, 0xdc,
	0xa7, 0x00, 0x0b, 0x2a, 0x93, 0x43, 0x94, 0x17, 0x7f, 0x98, 0x79, 0x00, 0x7d, 0xd5, 0x56, 0xdb,
	0xda, 0x4c, 0xf7, 0x50, 0xd8, 0x21, 0xf7, 0x31, 0x98, 0x6b, 0xda, 0x64, 0xec, 0xef, 0x22, 0xe8,
	0x22, 0xf2, 0x19, 0xc1, 0xa8, 0x15, 0x48, 0xb7, 0xca, 0x27, 0xf6, 0xc0, 0x90, 0xc7, 0x92, 0xa9,
	0xb9, 0xc9, 0x93, 0xa9, 0x7f, 0xb1, 0xef, 0x9f, 0xdf, 0xa3, 0x63, 0xc9, 0x42, 0xd5, 0xd1, 0xed,
	0xd7, 0x7e, 0xed, 0x7f, 0x08, 0x20, 0xf3, 0x82, 0xbd, 0xe3, 0x94, 0x8b, 0xda, 0xd6, 0xd5, 0x25,
	0xc3, 0x96, 0x09, 0x5a, 0xe2, 0x2a, 0x6f, 0xdc, 0xc8, 0x63, 0x07, 0x80, 0x72, 0x2e, 0x24, 0x95,
	0xb9, 0xe0, 0xb6, 0xa9, 0x96, 0xdd, 0x30, 0xee, 0x17, 0x04, 0xf7, 0x5f, 0x89, 0xea, 0x23, 0xad,
	0xd2, 0xbb, 0xb7, 0x78, 0x8d, 0xd4, 0xb8, 0x8d, 0xf4, 0x9f, 0x26, 0x09, 0xe8, 0x11, 0xcd, 0x30,
	0x06, 0x83, 0xd3, 0x82, 0x75, 0x91, 0xab, 0xfa, 0xf7, 0xd0, 0xc7, 0xdd, 0xd5, 0x8f, 0x9e, 0x03,
	0x5c, 0xbd, 0xe2, 0x11, 0x0c, 0xf6, 0xc1, 0xeb, 0x60, 0xf7, 0x26, 0xb0, 0x7a, 0x2d, 0x58, 0xee,
	0xf6, 0x41, 0xb4, 0x0a, 0x2d, 0x84, 0x87, 0x60, 0x46, 0x9b, 0xed, 0x2a, 0xb4, 0xb4, 0xb6, 0x5c,
	0xbf, 0xd8, 0xaf, 0x57, 0x96, 0xbe, 0xd8, 0x7c, 0x3d, 0x39, 0xe8, 0xdb, 0xc9, 0x41, 0xdf, 0x4f,
	0x0e, 0xfa, 0xf4, 0xc3, 0xe9, 0xbd, 0x7d, 0xf6, 0x9f, 0x5f, 0x35, 0xee, 0x2b, 0x3c, 0xff, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x82, 0xff, 0x64, 0x34, 0xec, 0x02, 0x00, 0x00,
}
