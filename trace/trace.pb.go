// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/api/trace/trace.proto

package trace // import "github.com/TheThingsNetwork/api/trace"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Trace information
type Trace struct {
	// Generated ID
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Time in Unix nanoseconds
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	// The ID of the component
	ServiceID string `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// The name of the component (router/broker/handler)
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Short event name
	Event string `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	// metadata for the event
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Parents of the event
	Parents              []*Trace `protobuf:"bytes,11,rep,name=parents" json:"parents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trace) Reset()      { *m = Trace{} }
func (*Trace) ProtoMessage() {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_1a0e1745152703b8, []int{0}
}
func (m *Trace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(dst, src)
}
func (m *Trace) XXX_Size() int {
	return m.Size()
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Trace) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Trace) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *Trace) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Trace) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Trace) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Trace) GetParents() []*Trace {
	if m != nil {
		return m.Parents
	}
	return nil
}

func init() {
	proto.RegisterType((*Trace)(nil), "trace.Trace")
	golang_proto.RegisterType((*Trace)(nil), "trace.Trace")
	proto.RegisterMapType((map[string]string)(nil), "trace.Trace.MetadataEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "trace.Trace.MetadataEntry")
}
func (this *Trace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Trace)
	if !ok {
		that2, ok := that.(Trace)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.Time != that1.Time {
		return false
	}
	if this.ServiceID != that1.ServiceID {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.Event != that1.Event {
		return false
	}
	if len(this.Metadata) != len(that1.Metadata) {
		return false
	}
	for i := range this.Metadata {
		if this.Metadata[i] != that1.Metadata[i] {
			return false
		}
	}
	if len(this.Parents) != len(that1.Parents) {
		return false
	}
	for i := range this.Parents {
		if !this.Parents[i].Equal(that1.Parents[i]) {
			return false
		}
	}
	return true
}
func (m *Trace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trace) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTrace(dAtA, i, uint64(m.Time))
	}
	if len(m.ServiceID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.ServiceID)))
		i += copy(dAtA[i:], m.ServiceID)
	}
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.Event) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.Event)))
		i += copy(dAtA[i:], m.Event)
	}
	if len(m.Metadata) > 0 {
		for k, _ := range m.Metadata {
			dAtA[i] = 0x32
			i++
			v := m.Metadata[k]
			mapSize := 1 + len(k) + sovTrace(uint64(len(k))) + 1 + len(v) + sovTrace(uint64(len(v)))
			i = encodeVarintTrace(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrace(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTrace(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Parents) > 0 {
		for _, msg := range m.Parents {
			dAtA[i] = 0x5a
			i++
			i = encodeVarintTrace(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintTrace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTrace(r randyTrace, easy bool) *Trace {
	this := &Trace{}
	this.ID = string(randStringTrace(r))
	this.Time = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Time *= -1
	}
	this.ServiceID = string(randStringTrace(r))
	this.ServiceName = string(randStringTrace(r))
	this.Event = string(randStringTrace(r))
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.Metadata = make(map[string]string)
		for i := 0; i < v1; i++ {
			this.Metadata[randStringTrace(r)] = randStringTrace(r)
		}
	}
	if r.Intn(10) == 0 {
		v2 := r.Intn(5)
		this.Parents = make([]*Trace, v2)
		for i := 0; i < v2; i++ {
			this.Parents[i] = NewPopulatedTrace(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTrace interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTrace(r randyTrace) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTrace(r randyTrace) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTrace(r)
	}
	return string(tmps)
}
func randUnrecognizedTrace(r randyTrace, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTrace(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTrace(dAtA []byte, r randyTrace, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTrace(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTrace(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTrace(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTrace(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTrace(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTrace(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTrace(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Trace) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovTrace(uint64(m.Time))
	}
	l = len(m.ServiceID)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	l = len(m.Event)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTrace(uint64(len(k))) + 1 + len(v) + sovTrace(uint64(len(v)))
			n += mapEntrySize + 1 + sovTrace(uint64(mapEntrySize))
		}
	}
	if len(m.Parents) > 0 {
		for _, e := range m.Parents {
			l = e.Size()
			n += 1 + l + sovTrace(uint64(l))
		}
	}
	return n
}

func sovTrace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTrace(x uint64) (n int) {
	return sovTrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Trace) String() string {
	if this == nil {
		return "nil"
	}
	keysForMetadata := make([]string, 0, len(this.Metadata))
	for k, _ := range this.Metadata {
		keysForMetadata = append(keysForMetadata, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMetadata)
	mapStringForMetadata := "map[string]string{"
	for _, k := range keysForMetadata {
		mapStringForMetadata += fmt.Sprintf("%v: %v,", k, this.Metadata[k])
	}
	mapStringForMetadata += "}"
	s := strings.Join([]string{`&Trace{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Time:` + fmt.Sprintf("%v", this.Time) + `,`,
		`ServiceID:` + fmt.Sprintf("%v", this.ServiceID) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`Event:` + fmt.Sprintf("%v", this.Event) + `,`,
		`Metadata:` + mapStringForMetadata + `,`,
		`Parents:` + strings.Replace(fmt.Sprintf("%v", this.Parents), "Trace", "Trace", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTrace(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Trace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: Trace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Event = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTrace
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTrace
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTrace
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTrace
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTrace
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTrace(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTrace
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parents = append(m.Parents, &Trace{})
			if err := m.Parents[len(m.Parents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func skipTrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrace
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
					return 0, ErrIntOverflowTrace
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
					return 0, ErrIntOverflowTrace
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
				return 0, ErrInvalidLengthTrace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrace
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
				next, err := skipTrace(dAtA[start:])
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
	ErrInvalidLengthTrace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrace   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/api/trace/trace.proto", fileDescriptor_trace_1a0e1745152703b8)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/api/trace/trace.proto", fileDescriptor_trace_1a0e1745152703b8)
}

var fileDescriptor_trace_1a0e1745152703b8 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x6c, 0x13, 0x31,
	0x18, 0xf5, 0x77, 0x69, 0x02, 0x71, 0x5a, 0x09, 0x59, 0x08, 0x9d, 0x32, 0x7c, 0x0d, 0x48, 0xa0,
	0x0c, 0x70, 0x11, 0x20, 0x21, 0x04, 0x13, 0x51, 0x19, 0x32, 0x50, 0x45, 0x47, 0x26, 0x16, 0xe4,
	0xe4, 0xcc, 0xc5, 0x2a, 0xf7, 0xa3, 0x8b, 0x13, 0xd4, 0xad, 0x63, 0x47, 0x46, 0x46, 0xc6, 0x8a,
	0xa9, 0x1b, 0x1d, 0x3b, 0x76, 0xec, 0xd8, 0xa9, 0xea, 0xd9, 0x12, 0xea, 0xd8, 0xb1, 0x23, 0x3a,
	0x3b, 0x45, 0x81, 0x89, 0xe5, 0xf4, 0x3d, 0xbf, 0xf7, 0x4e, 0x7e, 0x9f, 0x1f, 0x7d, 0x1a, 0x4b,
	0x35, 0x9d, 0x8f, 0x83, 0x49, 0x96, 0xf4, 0x46, 0x53, 0x31, 0x9a, 0xca, 0x34, 0x9e, 0x6d, 0x0b,
	0xf5, 0x25, 0x2b, 0x76, 0x7a, 0x3c, 0x97, 0x3d, 0x55, 0xf0, 0x89, 0x70, 0xdf, 0x20, 0x2f, 0x32,
	0x95, 0xb1, 0xba, 0x05, 0xed, 0x27, 0x2b, 0xce, 0x38, 0x8b, 0xb3, 0x9e, 0x65, 0xc7, 0xf3, 0x4f,
	0x16, 0x59, 0x60, 0x27, 0xe7, 0x7a, 0xf0, 0xd3, 0xa3, 0xf5, 0x51, 0x65, 0x64, 0xf7, 0xa8, 0x27,
	0x23, 0x1f, 0x3a, 0xd0, 0x6d, 0xf6, 0x1b, 0xfa, 0x7c, 0xd3, 0x1b, 0x6c, 0x85, 0x9e, 0x8c, 0x18,
	0xa3, 0x6b, 0x4a, 0x26, 0xc2, 0xf7, 0x3a, 0xd0, 0xad, 0x85, 0x76, 0x66, 0x8f, 0x29, 0x9d, 0x89,
	0x62, 0x21, 0x27, 0xe2, 0xa3, 0x8c, 0xfc, 0x9a, 0xf5, 0x6c, 0xe8, 0xf3, 0xcd, 0xe6, 0x7b, 0x77,
	0x3a, 0xd8, 0x0a, 0x9b, 0x4b, 0xc1, 0x20, 0x62, 0xf7, 0xe9, 0xfa, 0x8d, 0x3a, 0xe5, 0x89, 0xf0,
	0xd7, 0x2a, 0x7d, 0xd8, 0x5a, 0x9e, 0x6d, 0xf3, 0x44, 0xb0, 0xbb, 0xb4, 0x2e, 0x16, 0x22, 0x55,
	0x7e, 0xdd, 0x72, 0x0e, 0xb0, 0x17, 0xf4, 0x76, 0x22, 0x14, 0x8f, 0xb8, 0xe2, 0x7e, 0xa3, 0x53,
	0xeb, 0xb6, 0x9e, 0xb5, 0x03, 0x17, 0xd9, 0x5e, 0x39, 0x78, 0xb7, 0x24, 0xdf, 0xa6, 0xaa, 0xd8,
	0x0d, 0xff, 0x68, 0xd9, 0x23, 0x7a, 0x2b, 0xe7, 0x85, 0x48, 0xd5, 0xcc, 0x6f, 0x59, 0xdb, 0xfa,
	0xaa, 0x2d, 0xbc, 0x21, 0xdb, 0xaf, 0xe9, 0xc6, 0x5f, 0xbf, 0x60, 0x77, 0x68, 0x6d, 0x47, 0xec,
	0xba, 0x25, 0x84, 0xd5, 0x58, 0x5d, 0x6c, 0xc1, 0x3f, 0xcf, 0x5d, 0xfc, 0x66, 0xe8, 0xc0, 0x2b,
	0xef, 0x25, 0xf4, 0x7f, 0xc1, 0x49, 0x89, 0x70, 0x5a, 0x22, 0x9c, 0x95, 0x48, 0x2e, 0x4a, 0x24,
	0x97, 0x25, 0x92, 0xab, 0x12, 0xc9, 0x75, 0x89, 0xb0, 0xa7, 0x11, 0xf6, 0x35, 0x92, 0x03, 0x8d,
	0x70, 0xa8, 0x91, 0x1c, 0x69, 0x24, 0xc7, 0x1a, 0xc9, 0x89, 0x46, 0x38, 0xd5, 0x08, 0x67, 0x1a,
	0xc9, 0x85, 0x46, 0xb8, 0xd4, 0x48, 0xae, 0x34, 0xc2, 0xb5, 0x46, 0xb2, 0x67, 0x90, 0xec, 0x1b,
	0x84, 0xaf, 0x06, 0xc9, 0x37, 0x83, 0xf0, 0xdd, 0x20, 0x39, 0x30, 0x48, 0x0e, 0x0d, 0xc2, 0x91,
	0x41, 0x38, 0x36, 0x08, 0x14, 0xb3, 0x22, 0x0e, 0xd4, 0x54, 0x28, 0x5b, 0x89, 0xd4, 0x55, 0x22,
	0xe0, 0xb9, 0x74, 0x11, 0xfb, 0xd4, 0x66, 0x1c, 0x56, 0x8f, 0x3b, 0x84, 0x0f, 0x0f, 0xff, 0xab,
	0x47, 0x3f, 0xbc, 0xf6, 0xbf, 0x64, 0xf0, 0x66, 0x38, 0x70, 0xeb, 0x1a, 0x37, 0x6c, 0x53, 0x9e,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x50, 0x75, 0xed, 0x94, 0x02, 0x00, 0x00,
}
