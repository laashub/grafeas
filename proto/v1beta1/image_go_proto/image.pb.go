// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package image_go_proto

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

// Instructions from Dockerfile.
type Layer_Directive int32

const (
	// Default value for unsupported/missing directive.
	Layer_DIRECTIVE_UNSPECIFIED Layer_Directive = 0
	// https://docs.docker.com/engine/reference/builder/
	Layer_MAINTAINER Layer_Directive = 1
	// https://docs.docker.com/engine/reference/builder/
	Layer_RUN Layer_Directive = 2
	// https://docs.docker.com/engine/reference/builder/
	Layer_CMD Layer_Directive = 3
	// https://docs.docker.com/engine/reference/builder/
	Layer_LABEL Layer_Directive = 4
	// https://docs.docker.com/engine/reference/builder/
	Layer_EXPOSE Layer_Directive = 5
	// https://docs.docker.com/engine/reference/builder/
	Layer_ENV Layer_Directive = 6
	// https://docs.docker.com/engine/reference/builder/
	Layer_ADD Layer_Directive = 7
	// https://docs.docker.com/engine/reference/builder/
	Layer_COPY Layer_Directive = 8
	// https://docs.docker.com/engine/reference/builder/
	Layer_ENTRYPOINT Layer_Directive = 9
	// https://docs.docker.com/engine/reference/builder/
	Layer_VOLUME Layer_Directive = 10
	// https://docs.docker.com/engine/reference/builder/
	Layer_USER Layer_Directive = 11
	// https://docs.docker.com/engine/reference/builder/
	Layer_WORKDIR Layer_Directive = 12
	// https://docs.docker.com/engine/reference/builder/
	Layer_ARG Layer_Directive = 13
	// https://docs.docker.com/engine/reference/builder/
	Layer_ONBUILD Layer_Directive = 14
	// https://docs.docker.com/engine/reference/builder/
	Layer_STOPSIGNAL Layer_Directive = 15
	// https://docs.docker.com/engine/reference/builder/
	Layer_HEALTHCHECK Layer_Directive = 16
	// https://docs.docker.com/engine/reference/builder/
	Layer_SHELL Layer_Directive = 17
)

var Layer_Directive_name = map[int32]string{
	0:  "DIRECTIVE_UNSPECIFIED",
	1:  "MAINTAINER",
	2:  "RUN",
	3:  "CMD",
	4:  "LABEL",
	5:  "EXPOSE",
	6:  "ENV",
	7:  "ADD",
	8:  "COPY",
	9:  "ENTRYPOINT",
	10: "VOLUME",
	11: "USER",
	12: "WORKDIR",
	13: "ARG",
	14: "ONBUILD",
	15: "STOPSIGNAL",
	16: "HEALTHCHECK",
	17: "SHELL",
}

var Layer_Directive_value = map[string]int32{
	"DIRECTIVE_UNSPECIFIED": 0,
	"MAINTAINER":            1,
	"RUN":                   2,
	"CMD":                   3,
	"LABEL":                 4,
	"EXPOSE":                5,
	"ENV":                   6,
	"ADD":                   7,
	"COPY":                  8,
	"ENTRYPOINT":            9,
	"VOLUME":                10,
	"USER":                  11,
	"WORKDIR":               12,
	"ARG":                   13,
	"ONBUILD":               14,
	"STOPSIGNAL":            15,
	"HEALTHCHECK":           16,
	"SHELL":                 17,
}

func (x Layer_Directive) String() string {
	return proto.EnumName(Layer_Directive_name, int32(x))
}

func (Layer_Directive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{0, 0}
}

// Layer holds metadata specific to a layer of a Docker image.
type Layer struct {
	// Required. The recovered Dockerfile directive used to construct this layer.
	Directive Layer_Directive `protobuf:"varint,1,opt,name=directive,proto3,enum=grafeas.v1beta1.image.Layer_Directive" json:"directive,omitempty"`
	// The recovered arguments to the Dockerfile directive.
	Arguments            string   `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer) Reset()         { *m = Layer{} }
func (m *Layer) String() string { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()    {}
func (*Layer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{0}
}

func (m *Layer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer.Unmarshal(m, b)
}
func (m *Layer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer.Marshal(b, m, deterministic)
}
func (m *Layer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer.Merge(m, src)
}
func (m *Layer) XXX_Size() int {
	return xxx_messageInfo_Layer.Size(m)
}
func (m *Layer) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer.DiscardUnknown(m)
}

var xxx_messageInfo_Layer proto.InternalMessageInfo

func (m *Layer) GetDirective() Layer_Directive {
	if m != nil {
		return m.Directive
	}
	return Layer_DIRECTIVE_UNSPECIFIED
}

func (m *Layer) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

// A set of properties that uniquely identify a given Docker image.
type Fingerprint struct {
	// Required. The layer ID of the final layer in the Docker image's v1
	// representation.
	V1Name string `protobuf:"bytes,1,opt,name=v1_name,json=v1Name,proto3" json:"v1_name,omitempty"`
	// Required. The ordered list of v2 blobs that represent a given image.
	V2Blob []string `protobuf:"bytes,2,rep,name=v2_blob,json=v2Blob,proto3" json:"v2_blob,omitempty"`
	// Output only. The name of the image's v2 blobs computed via:
	//   [bottom] := v2_blob[bottom]
	//   [N] := sha256(v2_blob[N] + " " + v2_name[N+1])
	// Only the name of the final blob is kept.
	V2Name               string   `protobuf:"bytes,3,opt,name=v2_name,json=v2Name,proto3" json:"v2_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fingerprint) Reset()         { *m = Fingerprint{} }
func (m *Fingerprint) String() string { return proto.CompactTextString(m) }
func (*Fingerprint) ProtoMessage()    {}
func (*Fingerprint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{1}
}

func (m *Fingerprint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fingerprint.Unmarshal(m, b)
}
func (m *Fingerprint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fingerprint.Marshal(b, m, deterministic)
}
func (m *Fingerprint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fingerprint.Merge(m, src)
}
func (m *Fingerprint) XXX_Size() int {
	return xxx_messageInfo_Fingerprint.Size(m)
}
func (m *Fingerprint) XXX_DiscardUnknown() {
	xxx_messageInfo_Fingerprint.DiscardUnknown(m)
}

var xxx_messageInfo_Fingerprint proto.InternalMessageInfo

func (m *Fingerprint) GetV1Name() string {
	if m != nil {
		return m.V1Name
	}
	return ""
}

func (m *Fingerprint) GetV2Blob() []string {
	if m != nil {
		return m.V2Blob
	}
	return nil
}

func (m *Fingerprint) GetV2Name() string {
	if m != nil {
		return m.V2Name
	}
	return ""
}

// Basis describes the base image portion (Note) of the DockerImage
// relationship. Linked occurrences are derived from this or an
// equivalent image via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g. a tag of the resource_url.
type Basis struct {
	// Required. Immutable. The resource_url for the resource representing the
	// basis of associated occurrence images.
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// Required. Immutable. The fingerprint of the base image.
	Fingerprint          *Fingerprint `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Basis) Reset()         { *m = Basis{} }
func (m *Basis) String() string { return proto.CompactTextString(m) }
func (*Basis) ProtoMessage()    {}
func (*Basis) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{2}
}

func (m *Basis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Basis.Unmarshal(m, b)
}
func (m *Basis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Basis.Marshal(b, m, deterministic)
}
func (m *Basis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basis.Merge(m, src)
}
func (m *Basis) XXX_Size() int {
	return xxx_messageInfo_Basis.Size(m)
}
func (m *Basis) XXX_DiscardUnknown() {
	xxx_messageInfo_Basis.DiscardUnknown(m)
}

var xxx_messageInfo_Basis proto.InternalMessageInfo

func (m *Basis) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *Basis) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

// Details of an image occurrence.
type Details struct {
	// Required. Immutable. The child image derived from the base image.
	DerivedImage         *Derived `protobuf:"bytes,1,opt,name=derived_image,json=derivedImage,proto3" json:"derived_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{3}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetDerivedImage() *Derived {
	if m != nil {
		return m.DerivedImage
	}
	return nil
}

// Derived describes the derived image portion (Occurrence) of the DockerImage
// relationship. This image would be produced from a Dockerfile with FROM
// <DockerImage.Basis in attached Note>.
type Derived struct {
	// Required. The fingerprint of the derived image.
	Fingerprint *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Output only. The number of layers by which this image differs from the
	// associated image basis.
	Distance int32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	// This contains layer-specific metadata, if populated it has length
	// "distance" and is ordered with [distance] being the layer immediately
	// following the base image and [1] being the final layer.
	LayerInfo []*Layer `protobuf:"bytes,3,rep,name=layer_info,json=layerInfo,proto3" json:"layer_info,omitempty"`
	// Output only. This contains the base image URL for the derived image
	// occurrence.
	BaseResourceUrl      string   `protobuf:"bytes,4,opt,name=base_resource_url,json=baseResourceUrl,proto3" json:"base_resource_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Derived) Reset()         { *m = Derived{} }
func (m *Derived) String() string { return proto.CompactTextString(m) }
func (*Derived) ProtoMessage()    {}
func (*Derived) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{4}
}

func (m *Derived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Derived.Unmarshal(m, b)
}
func (m *Derived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Derived.Marshal(b, m, deterministic)
}
func (m *Derived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Derived.Merge(m, src)
}
func (m *Derived) XXX_Size() int {
	return xxx_messageInfo_Derived.Size(m)
}
func (m *Derived) XXX_DiscardUnknown() {
	xxx_messageInfo_Derived.DiscardUnknown(m)
}

var xxx_messageInfo_Derived proto.InternalMessageInfo

func (m *Derived) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *Derived) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Derived) GetLayerInfo() []*Layer {
	if m != nil {
		return m.LayerInfo
	}
	return nil
}

func (m *Derived) GetBaseResourceUrl() string {
	if m != nil {
		return m.BaseResourceUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.image.Layer_Directive", Layer_Directive_name, Layer_Directive_value)
	proto.RegisterType((*Layer)(nil), "grafeas.v1beta1.image.Layer")
	proto.RegisterType((*Fingerprint)(nil), "grafeas.v1beta1.image.Fingerprint")
	proto.RegisterType((*Basis)(nil), "grafeas.v1beta1.image.Basis")
	proto.RegisterType((*Details)(nil), "grafeas.v1beta1.image.Details")
	proto.RegisterType((*Derived)(nil), "grafeas.v1beta1.image.Derived")
}

func init() { proto.RegisterFile("image.proto", fileDescriptor_9624c68e2b547544) }

var fileDescriptor_9624c68e2b547544 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0x9b, 0x4c,
	0x14, 0xc5, 0x3f, 0x4c, 0x6c, 0x87, 0x4b, 0xfe, 0x4c, 0x46, 0x8a, 0x3e, 0xb7, 0x8a, 0x2a, 0xd7,
	0x8b, 0x2a, 0xea, 0x82, 0xc8, 0x74, 0xd1, 0x45, 0x57, 0x18, 0x26, 0x31, 0x0a, 0x01, 0x6b, 0x6c,
	0xa7, 0x49, 0xbb, 0x40, 0x60, 0x8f, 0xdd, 0x91, 0x6c, 0x88, 0x06, 0x6c, 0xa9, 0xaf, 0xd3, 0x67,
	0xea, 0xba, 0xcf, 0xd0, 0x47, 0xa8, 0x66, 0x42, 0xec, 0xb4, 0x4a, 0xba, 0xe8, 0x8a, 0x99, 0x7b,
	0xee, 0xef, 0x70, 0xb8, 0x33, 0x80, 0xc9, 0x97, 0xc9, 0x9c, 0x59, 0x77, 0x22, 0x2f, 0x73, 0x7c,
	0x3c, 0x17, 0xc9, 0x8c, 0x25, 0x85, 0xb5, 0xee, 0xa6, 0xac, 0x4c, 0xba, 0x96, 0x12, 0x3b, 0x3f,
	0x6a, 0x50, 0x0f, 0x92, 0xaf, 0x4c, 0x60, 0x0f, 0x8c, 0x29, 0x17, 0x6c, 0x52, 0xf2, 0x35, 0x6b,
	0x69, 0x6d, 0xed, 0xf4, 0xc0, 0x7e, 0x63, 0x3d, 0x09, 0x59, 0x0a, 0xb0, 0xbc, 0x87, 0x6e, 0xba,
	0x05, 0xf1, 0x09, 0x18, 0x89, 0x98, 0xaf, 0x96, 0x2c, 0x2b, 0x8b, 0x56, 0xad, 0xad, 0x9d, 0x1a,
	0x74, 0x5b, 0xe8, 0xfc, 0xd4, 0xc0, 0xd8, 0x60, 0xf8, 0x05, 0x1c, 0x7b, 0x3e, 0x25, 0xee, 0xc8,
	0xbf, 0x26, 0xf1, 0x38, 0x1c, 0x0e, 0x88, 0xeb, 0x9f, 0xfb, 0xc4, 0x43, 0xff, 0xe1, 0x03, 0x80,
	0x2b, 0xc7, 0x0f, 0x47, 0x8e, 0x1f, 0x12, 0x8a, 0x34, 0xdc, 0x04, 0x9d, 0x8e, 0x43, 0x54, 0x93,
	0x0b, 0xf7, 0xca, 0x43, 0x3a, 0x36, 0xa0, 0x1e, 0x38, 0x3d, 0x12, 0xa0, 0x1d, 0x0c, 0xd0, 0x20,
	0x37, 0x83, 0x68, 0x48, 0x50, 0x5d, 0xea, 0x24, 0xbc, 0x46, 0x0d, 0xb9, 0x70, 0x3c, 0x0f, 0x35,
	0xf1, 0x2e, 0xec, 0xb8, 0xd1, 0xe0, 0x16, 0xed, 0x4a, 0x53, 0x12, 0x8e, 0xe8, 0xed, 0x20, 0xf2,
	0xc3, 0x11, 0x32, 0x24, 0x77, 0x1d, 0x05, 0xe3, 0x2b, 0x82, 0x40, 0x76, 0x8d, 0x87, 0x84, 0x22,
	0x13, 0x9b, 0xd0, 0xfc, 0x18, 0xd1, 0x4b, 0xcf, 0xa7, 0x68, 0x4f, 0xb9, 0xd0, 0x0b, 0xb4, 0x2f,
	0xab, 0x51, 0xd8, 0x1b, 0xfb, 0x81, 0x87, 0x0e, 0xa4, 0xd1, 0x70, 0x14, 0x0d, 0x86, 0xfe, 0x45,
	0xe8, 0x04, 0xe8, 0x10, 0x1f, 0x82, 0xd9, 0x27, 0x4e, 0x30, 0xea, 0xbb, 0x7d, 0xe2, 0x5e, 0x22,
	0x24, 0xc3, 0x0d, 0xfb, 0x24, 0x08, 0xd0, 0x51, 0xe7, 0x06, 0xcc, 0x73, 0x9e, 0xcd, 0x99, 0xb8,
	0x13, 0x3c, 0x2b, 0xf1, 0xff, 0xd0, 0x5c, 0x77, 0xe3, 0x2c, 0x59, 0xde, 0xcf, 0xd8, 0xa0, 0x8d,
	0x75, 0x37, 0x4c, 0x96, 0x4c, 0x09, 0x76, 0x9c, 0x2e, 0xf2, 0xb4, 0x55, 0x6b, 0xeb, 0x4a, 0xb0,
	0x7b, 0x8b, 0x3c, 0xad, 0x04, 0x45, 0xe8, 0x15, 0x61, 0x4b, 0xa2, 0x73, 0x07, 0xf5, 0x5e, 0x52,
	0xf0, 0x02, 0xbf, 0x86, 0x3d, 0xc1, 0x8a, 0x7c, 0x25, 0x26, 0x2c, 0x5e, 0x89, 0x45, 0x65, 0x6c,
	0x3e, 0xd4, 0xc6, 0x62, 0x81, 0x3d, 0x30, 0x67, 0xdb, 0x14, 0xea, 0x60, 0x4c, 0xbb, 0xf3, 0xcc,
	0xf1, 0x3e, 0xca, 0x4b, 0x1f, 0x63, 0x9d, 0x10, 0x9a, 0x1e, 0x2b, 0x13, 0xbe, 0x28, 0xb0, 0x0b,
	0xfb, 0x53, 0x26, 0xf8, 0x9a, 0x4d, 0x63, 0x05, 0xa9, 0x97, 0x9a, 0xf6, 0xab, 0x67, 0x2c, 0xbd,
	0xfb, 0x5e, 0xba, 0x57, 0x41, 0xbe, 0xba, 0x7c, 0xdf, 0x35, 0x69, 0xa8, 0x0a, 0x7f, 0x26, 0xd4,
	0xfe, 0x29, 0x21, 0x7e, 0x09, 0xbb, 0x53, 0x5e, 0x94, 0x49, 0x36, 0x61, 0xea, 0x23, 0xeb, 0x74,
	0xb3, 0xc7, 0x1f, 0x00, 0x16, 0xf2, 0xe2, 0xc6, 0x3c, 0x9b, 0xe5, 0x2d, 0xbd, 0xad, 0x9f, 0x9a,
	0xf6, 0xc9, 0xdf, 0x6e, 0x38, 0x35, 0x54, 0xbf, 0x9f, 0xcd, 0x72, 0xfc, 0x16, 0x8e, 0xd2, 0xa4,
	0x60, 0xf1, 0x6f, 0x83, 0xde, 0x51, 0x83, 0x3e, 0x94, 0x02, 0xdd, 0x0e, 0xbb, 0xf7, 0x19, 0x5a,
	0x3c, 0x7f, 0xda, 0x78, 0xa0, 0x7d, 0x7a, 0x3f, 0xe7, 0xe5, 0x97, 0x55, 0x6a, 0x4d, 0xf2, 0xe5,
	0x59, 0xd5, 0xb3, 0x79, 0xaa, 0x5f, 0xf5, 0xac, 0x22, 0xce, 0x14, 0x11, 0xcf, 0xf3, 0x58, 0x95,
	0xbf, 0xd5, 0xf4, 0x0b, 0xea, 0xa4, 0x0d, 0xb5, 0x79, 0xf7, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc7,
	0x23, 0x29, 0x22, 0xdd, 0x03, 0x00, 0x00,
}
