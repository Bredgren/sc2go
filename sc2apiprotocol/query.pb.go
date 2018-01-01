// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2clientprotocol/query.proto

/*
Package sc2apiprotocol is a generated protocol buffer package.

It is generated from these files:
	s2clientprotocol/query.proto

It has these top-level messages:
	RequestQuery
	ResponseQuery
	RequestQueryPathing
	ResponseQueryPathing
	RequestQueryAvailableAbilities
	ResponseQueryAvailableAbilities
	RequestQueryBuildingPlacement
	ResponseQueryBuildingPlacement
*/
package sc2apiprotocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestQuery struct {
	Pathing                    []*RequestQueryPathing            `protobuf:"bytes,1,rep,name=pathing" json:"pathing,omitempty"`
	Abilities                  []*RequestQueryAvailableAbilities `protobuf:"bytes,2,rep,name=abilities" json:"abilities,omitempty"`
	Placements                 []*RequestQueryBuildingPlacement  `protobuf:"bytes,3,rep,name=placements" json:"placements,omitempty"`
	IgnoreResourceRequirements *bool                             `protobuf:"varint,4,opt,name=ignore_resource_requirements,json=ignoreResourceRequirements" json:"ignore_resource_requirements,omitempty"`
	XXX_unrecognized           []byte                            `json:"-"`
}

func (m *RequestQuery) Reset()                    { *m = RequestQuery{} }
func (m *RequestQuery) String() string            { return proto.CompactTextString(m) }
func (*RequestQuery) ProtoMessage()               {}
func (*RequestQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{0} }

func (m *RequestQuery) GetPathing() []*RequestQueryPathing {
	if m != nil {
		return m.Pathing
	}
	return nil
}

func (m *RequestQuery) GetAbilities() []*RequestQueryAvailableAbilities {
	if m != nil {
		return m.Abilities
	}
	return nil
}

func (m *RequestQuery) GetPlacements() []*RequestQueryBuildingPlacement {
	if m != nil {
		return m.Placements
	}
	return nil
}

func (m *RequestQuery) GetIgnoreResourceRequirements() bool {
	if m != nil && m.IgnoreResourceRequirements != nil {
		return *m.IgnoreResourceRequirements
	}
	return false
}

type ResponseQuery struct {
	Pathing          []*ResponseQueryPathing            `protobuf:"bytes,1,rep,name=pathing" json:"pathing,omitempty"`
	Abilities        []*ResponseQueryAvailableAbilities `protobuf:"bytes,2,rep,name=abilities" json:"abilities,omitempty"`
	Placements       []*ResponseQueryBuildingPlacement  `protobuf:"bytes,3,rep,name=placements" json:"placements,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *ResponseQuery) Reset()                    { *m = ResponseQuery{} }
func (m *ResponseQuery) String() string            { return proto.CompactTextString(m) }
func (*ResponseQuery) ProtoMessage()               {}
func (*ResponseQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{1} }

func (m *ResponseQuery) GetPathing() []*ResponseQueryPathing {
	if m != nil {
		return m.Pathing
	}
	return nil
}

func (m *ResponseQuery) GetAbilities() []*ResponseQueryAvailableAbilities {
	if m != nil {
		return m.Abilities
	}
	return nil
}

func (m *ResponseQuery) GetPlacements() []*ResponseQueryBuildingPlacement {
	if m != nil {
		return m.Placements
	}
	return nil
}

// --------------------------------------------------------------------------------------------------
type RequestQueryPathing struct {
	// Types that are valid to be assigned to Start:
	//	*RequestQueryPathing_StartPos
	//	*RequestQueryPathing_UnitTag
	Start            isRequestQueryPathing_Start `protobuf_oneof:"start"`
	EndPos           *Point2D                    `protobuf:"bytes,3,opt,name=end_pos,json=endPos" json:"end_pos,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *RequestQueryPathing) Reset()                    { *m = RequestQueryPathing{} }
func (m *RequestQueryPathing) String() string            { return proto.CompactTextString(m) }
func (*RequestQueryPathing) ProtoMessage()               {}
func (*RequestQueryPathing) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{2} }

type isRequestQueryPathing_Start interface {
	isRequestQueryPathing_Start()
}

type RequestQueryPathing_StartPos struct {
	StartPos *Point2D `protobuf:"bytes,1,opt,name=start_pos,json=startPos,oneof"`
}
type RequestQueryPathing_UnitTag struct {
	UnitTag uint64 `protobuf:"varint,2,opt,name=unit_tag,json=unitTag,oneof"`
}

func (*RequestQueryPathing_StartPos) isRequestQueryPathing_Start() {}
func (*RequestQueryPathing_UnitTag) isRequestQueryPathing_Start()  {}

func (m *RequestQueryPathing) GetStart() isRequestQueryPathing_Start {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *RequestQueryPathing) GetStartPos() *Point2D {
	if x, ok := m.GetStart().(*RequestQueryPathing_StartPos); ok {
		return x.StartPos
	}
	return nil
}

func (m *RequestQueryPathing) GetUnitTag() uint64 {
	if x, ok := m.GetStart().(*RequestQueryPathing_UnitTag); ok {
		return x.UnitTag
	}
	return 0
}

func (m *RequestQueryPathing) GetEndPos() *Point2D {
	if m != nil {
		return m.EndPos
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RequestQueryPathing) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RequestQueryPathing_OneofMarshaler, _RequestQueryPathing_OneofUnmarshaler, _RequestQueryPathing_OneofSizer, []interface{}{
		(*RequestQueryPathing_StartPos)(nil),
		(*RequestQueryPathing_UnitTag)(nil),
	}
}

func _RequestQueryPathing_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RequestQueryPathing)
	// start
	switch x := m.Start.(type) {
	case *RequestQueryPathing_StartPos:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StartPos); err != nil {
			return err
		}
	case *RequestQueryPathing_UnitTag:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.UnitTag))
	case nil:
	default:
		return fmt.Errorf("RequestQueryPathing.Start has unexpected type %T", x)
	}
	return nil
}

func _RequestQueryPathing_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RequestQueryPathing)
	switch tag {
	case 1: // start.start_pos
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Point2D)
		err := b.DecodeMessage(msg)
		m.Start = &RequestQueryPathing_StartPos{msg}
		return true, err
	case 2: // start.unit_tag
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Start = &RequestQueryPathing_UnitTag{x}
		return true, err
	default:
		return false, nil
	}
}

func _RequestQueryPathing_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RequestQueryPathing)
	// start
	switch x := m.Start.(type) {
	case *RequestQueryPathing_StartPos:
		s := proto.Size(x.StartPos)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestQueryPathing_UnitTag:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.UnitTag))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResponseQueryPathing struct {
	Distance         *float32 `protobuf:"fixed32,1,opt,name=distance" json:"distance,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ResponseQueryPathing) Reset()                    { *m = ResponseQueryPathing{} }
func (m *ResponseQueryPathing) String() string            { return proto.CompactTextString(m) }
func (*ResponseQueryPathing) ProtoMessage()               {}
func (*ResponseQueryPathing) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{3} }

func (m *ResponseQueryPathing) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

// --------------------------------------------------------------------------------------------------
type RequestQueryAvailableAbilities struct {
	UnitTag          *uint64 `protobuf:"varint,1,opt,name=unit_tag,json=unitTag" json:"unit_tag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestQueryAvailableAbilities) Reset()         { *m = RequestQueryAvailableAbilities{} }
func (m *RequestQueryAvailableAbilities) String() string { return proto.CompactTextString(m) }
func (*RequestQueryAvailableAbilities) ProtoMessage()    {}
func (*RequestQueryAvailableAbilities) Descriptor() ([]byte, []int) {
	return fileDescriptorQuery, []int{4}
}

func (m *RequestQueryAvailableAbilities) GetUnitTag() uint64 {
	if m != nil && m.UnitTag != nil {
		return *m.UnitTag
	}
	return 0
}

type ResponseQueryAvailableAbilities struct {
	Abilities        []*AvailableAbility `protobuf:"bytes,1,rep,name=abilities" json:"abilities,omitempty"`
	UnitTag          *uint64             `protobuf:"varint,2,opt,name=unit_tag,json=unitTag" json:"unit_tag,omitempty"`
	UnitTypeId       *uint32             `protobuf:"varint,3,opt,name=unit_type_id,json=unitTypeId" json:"unit_type_id,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ResponseQueryAvailableAbilities) Reset()         { *m = ResponseQueryAvailableAbilities{} }
func (m *ResponseQueryAvailableAbilities) String() string { return proto.CompactTextString(m) }
func (*ResponseQueryAvailableAbilities) ProtoMessage()    {}
func (*ResponseQueryAvailableAbilities) Descriptor() ([]byte, []int) {
	return fileDescriptorQuery, []int{5}
}

func (m *ResponseQueryAvailableAbilities) GetAbilities() []*AvailableAbility {
	if m != nil {
		return m.Abilities
	}
	return nil
}

func (m *ResponseQueryAvailableAbilities) GetUnitTag() uint64 {
	if m != nil && m.UnitTag != nil {
		return *m.UnitTag
	}
	return 0
}

func (m *ResponseQueryAvailableAbilities) GetUnitTypeId() uint32 {
	if m != nil && m.UnitTypeId != nil {
		return *m.UnitTypeId
	}
	return 0
}

// --------------------------------------------------------------------------------------------------
type RequestQueryBuildingPlacement struct {
	AbilityId        *int32   `protobuf:"varint,1,opt,name=ability_id,json=abilityId" json:"ability_id,omitempty"`
	TargetPos        *Point2D `protobuf:"bytes,2,opt,name=target_pos,json=targetPos" json:"target_pos,omitempty"`
	PlacingUnitTag   *uint64  `protobuf:"varint,3,opt,name=placing_unit_tag,json=placingUnitTag" json:"placing_unit_tag,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RequestQueryBuildingPlacement) Reset()         { *m = RequestQueryBuildingPlacement{} }
func (m *RequestQueryBuildingPlacement) String() string { return proto.CompactTextString(m) }
func (*RequestQueryBuildingPlacement) ProtoMessage()    {}
func (*RequestQueryBuildingPlacement) Descriptor() ([]byte, []int) {
	return fileDescriptorQuery, []int{6}
}

func (m *RequestQueryBuildingPlacement) GetAbilityId() int32 {
	if m != nil && m.AbilityId != nil {
		return *m.AbilityId
	}
	return 0
}

func (m *RequestQueryBuildingPlacement) GetTargetPos() *Point2D {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

func (m *RequestQueryBuildingPlacement) GetPlacingUnitTag() uint64 {
	if m != nil && m.PlacingUnitTag != nil {
		return *m.PlacingUnitTag
	}
	return 0
}

type ResponseQueryBuildingPlacement struct {
	Result           *ActionResult `protobuf:"varint,1,opt,name=result,enum=ActionResult" json:"result,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ResponseQueryBuildingPlacement) Reset()         { *m = ResponseQueryBuildingPlacement{} }
func (m *ResponseQueryBuildingPlacement) String() string { return proto.CompactTextString(m) }
func (*ResponseQueryBuildingPlacement) ProtoMessage()    {}
func (*ResponseQueryBuildingPlacement) Descriptor() ([]byte, []int) {
	return fileDescriptorQuery, []int{7}
}

func (m *ResponseQueryBuildingPlacement) GetResult() ActionResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return ActionResult_Success
}

func init() {
	proto.RegisterType((*RequestQuery)(nil), "RequestQuery")
	proto.RegisterType((*ResponseQuery)(nil), "ResponseQuery")
	proto.RegisterType((*RequestQueryPathing)(nil), "RequestQueryPathing")
	proto.RegisterType((*ResponseQueryPathing)(nil), "ResponseQueryPathing")
	proto.RegisterType((*RequestQueryAvailableAbilities)(nil), "RequestQueryAvailableAbilities")
	proto.RegisterType((*ResponseQueryAvailableAbilities)(nil), "ResponseQueryAvailableAbilities")
	proto.RegisterType((*RequestQueryBuildingPlacement)(nil), "RequestQueryBuildingPlacement")
	proto.RegisterType((*ResponseQueryBuildingPlacement)(nil), "ResponseQueryBuildingPlacement")
}

func init() { proto.RegisterFile("s2clientprotocol/query.proto", fileDescriptorQuery) }

var fileDescriptorQuery = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xda, 0x6e, 0x6d, 0xef, 0xb6, 0x0a, 0x19, 0x24, 0xba, 0xd2, 0x42, 0x14, 0x78, 0xe8,
	0x0b, 0x1d, 0x8a, 0xd0, 0x5e, 0x10, 0x13, 0x1d, 0x3c, 0xac, 0x12, 0x43, 0xc1, 0x7c, 0xbc, 0x56,
	0x59, 0x62, 0x05, 0x4b, 0xa9, 0x9d, 0xda, 0x0e, 0x52, 0xfe, 0x0c, 0xe2, 0x85, 0x37, 0x7e, 0x1b,
	0xbf, 0x01, 0xc5, 0x4e, 0xbb, 0xd0, 0x54, 0xe6, 0x2d, 0xf6, 0x3d, 0xf7, 0xdc, 0xe3, 0x73, 0x6e,
	0x60, 0x2c, 0xfd, 0x28, 0xa5, 0x84, 0xa9, 0x4c, 0x70, 0xc5, 0x23, 0x9e, 0x9e, 0xaf, 0x73, 0x22,
	0x8a, 0x99, 0x3e, 0xa2, 0xc1, 0xa7, 0xb7, 0xfe, 0x3c, 0x58, 0x04, 0x55, 0x6d, 0x34, 0x69, 0xa0,
	0x23, 0xbe, 0x5a, 0x71, 0x66, 0xe0, 0xa3, 0x26, 0x19, 0x11, 0x82, 0x0b, 0x53, 0xf5, 0x7e, 0xb7,
	0xe0, 0x04, 0x93, 0x75, 0x4e, 0xa4, 0xfa, 0x58, 0xce, 0x40, 0xaf, 0xa1, 0x9b, 0x85, 0xea, 0x1b,
	0x65, 0xc9, 0xd0, 0x71, 0xdb, 0xd3, 0x63, 0xff, 0xe9, 0xec, 0xdf, 0x79, 0xb3, 0x3a, 0x3c, 0x30,
	0x50, 0xbc, 0xe9, 0x41, 0xef, 0xa1, 0x1f, 0xde, 0xd2, 0x94, 0x2a, 0x4a, 0xe4, 0xb0, 0xa5, 0x09,
	0x66, 0x36, 0x82, 0xf9, 0xf7, 0x90, 0xa6, 0xe1, 0x6d, 0x4a, 0xe6, 0x9b, 0x2e, 0x7c, 0x47, 0x80,
	0x6e, 0x00, 0xb2, 0x34, 0x8c, 0xc8, 0x8a, 0x30, 0x25, 0x87, 0x6d, 0x4d, 0xf7, 0xdc, 0x46, 0x77,
	0x95, 0xd3, 0x34, 0xa6, 0x2c, 0x09, 0x36, 0x5d, 0xb8, 0x46, 0x80, 0xde, 0xc0, 0x98, 0x26, 0x8c,
	0x0b, 0xb2, 0x14, 0x44, 0xf2, 0x5c, 0x44, 0xe5, 0xc7, 0x3a, 0xa7, 0xa2, 0x1a, 0xd0, 0x71, 0x9d,
	0x69, 0x0f, 0x8f, 0x0c, 0x06, 0x57, 0x10, 0x5c, 0x43, 0x78, 0x7f, 0x1c, 0x38, 0xc5, 0x44, 0x66,
	0x9c, 0x49, 0x62, 0xfc, 0xba, 0xdc, 0xf5, 0xeb, 0x59, 0x53, 0x5f, 0x0d, 0xdf, 0x30, 0xec, 0xa6,
	0x69, 0xd8, 0xb9, 0x95, 0xc1, 0xee, 0xd8, 0x87, 0x3d, 0x8e, 0xcd, 0xac, 0x7c, 0x56, 0xcb, 0xbc,
	0x5f, 0x0e, 0xdc, 0xdf, 0x13, 0x38, 0xba, 0x80, 0xbe, 0x54, 0xa1, 0x50, 0xcb, 0x8c, 0xcb, 0xa1,
	0xe3, 0x3a, 0xd3, 0x63, 0xff, 0xe1, 0xee, 0x98, 0x80, 0x53, 0xa6, 0xfc, 0x77, 0xd7, 0x07, 0xb8,
	0xa7, 0xb1, 0x01, 0x97, 0xe8, 0x11, 0xf4, 0x72, 0x46, 0xd5, 0x52, 0x85, 0xc9, 0xb0, 0xe5, 0x3a,
	0xd3, 0xce, 0xf5, 0x01, 0xee, 0x96, 0x37, 0x9f, 0xc3, 0x04, 0xbd, 0x80, 0x2e, 0x61, 0xb1, 0xa6,
	0x6c, 0x5b, 0x29, 0xf1, 0x11, 0x61, 0x71, 0xc0, 0xe5, 0x55, 0x17, 0x0e, 0x35, 0xb5, 0xe7, 0xc3,
	0x83, 0x7d, 0x3e, 0xa3, 0x11, 0xf4, 0x62, 0x2a, 0x55, 0xc8, 0x22, 0xa2, 0x65, 0xb6, 0xf0, 0xf6,
	0xec, 0xbd, 0x82, 0xc7, 0xf6, 0x55, 0x44, 0x67, 0x35, 0xb5, 0x65, 0x77, 0x67, 0xab, 0xd5, 0xfb,
	0xe1, 0xc0, 0x93, 0xff, 0xe4, 0x82, 0x2e, 0xeb, 0xd9, 0x9a, 0xed, 0x70, 0x77, 0x5f, 0xb4, 0xd3,
	0x56, 0xd4, 0xc3, 0x3c, 0xdb, 0x35, 0xeb, 0xce, 0x2a, 0x17, 0x4e, 0x4c, 0xa9, 0xc8, 0xc8, 0x92,
	0xc6, 0xda, 0xaf, 0x53, 0x0c, 0xba, 0x5c, 0x64, 0x64, 0x11, 0x7b, 0x3f, 0x1d, 0x98, 0x58, 0x7f,
	0x0d, 0x34, 0x01, 0x30, 0xb3, 0x8a, 0x92, 0xa1, 0x7c, 0xdf, 0xe1, 0x66, 0x7a, 0xb1, 0x88, 0xd1,
	0x05, 0x80, 0x0a, 0x45, 0x42, 0x4c, 0xc6, 0x2d, 0x7b, 0x20, 0x7d, 0x03, 0x2d, 0x23, 0x9e, 0xc2,
	0xbd, 0x72, 0x81, 0x28, 0x4b, 0x96, 0x5b, 0xf5, 0x6d, 0xad, 0x7e, 0x50, 0xdd, 0x7f, 0xa9, 0x3c,
	0xfc, 0x5a, 0x06, 0x60, 0x5b, 0x45, 0xf4, 0x12, 0x8e, 0x04, 0x91, 0x79, 0xaa, 0xb4, 0xbc, 0x81,
	0x3f, 0x6e, 0xd8, 0x17, 0x29, 0xca, 0x19, 0xd6, 0x18, 0x5c, 0x61, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x23, 0xa4, 0x79, 0xe9, 0x40, 0x05, 0x00, 0x00,
}