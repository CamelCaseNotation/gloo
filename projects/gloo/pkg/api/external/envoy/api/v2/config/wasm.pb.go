// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/api/v2/config/wasm/wasm.proto

package config

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for a Wasm VM.
// [#next-free-field: 6]
type VmConfig struct {
	// An ID which will be used along with a hash of the wasm code (or null_vm_id) to determine which
	// VM will be used for the plugin. All plugins which use the same vm_id and code will use the same
	// VM. May be left blank.
	VmId string `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// The Wasm runtime type (see source/extensions/commmon/wasm/well_known_names.h).
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The Wasm code that Envoy will execute.
	Code *core.AsyncDataSource `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The Wasm configuration string used on initialization of a new VM (proxy_onStart).
	Configuration string `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Allow the wasm file to include pre-compiled code on VMs which support it.
	AllowPrecompiled     bool     `protobuf:"varint,5,opt,name=allow_precompiled,json=allowPrecompiled,proto3" json:"allow_precompiled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmConfig) Reset()         { *m = VmConfig{} }
func (m *VmConfig) String() string { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()    {}
func (*VmConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9483ad3f200a8, []int{0}
}
func (m *VmConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmConfig.Unmarshal(m, b)
}
func (m *VmConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmConfig.Marshal(b, m, deterministic)
}
func (m *VmConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmConfig.Merge(m, src)
}
func (m *VmConfig) XXX_Size() int {
	return xxx_messageInfo_VmConfig.Size(m)
}
func (m *VmConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VmConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VmConfig proto.InternalMessageInfo

func (m *VmConfig) GetVmId() string {
	if m != nil {
		return m.VmId
	}
	return ""
}

func (m *VmConfig) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *VmConfig) GetCode() *core.AsyncDataSource {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *VmConfig) GetConfiguration() string {
	if m != nil {
		return m.Configuration
	}
	return ""
}

func (m *VmConfig) GetAllowPrecompiled() bool {
	if m != nil {
		return m.AllowPrecompiled
	}
	return false
}

// Base Configuration for Wasm Plugins, e.g. filters and services.
type PluginConfig struct {
	// A unique name for a filters/services in a VM for use in identifiying the filter/service if
	// multiple filters/services are handled by the same vm_id and root_id and for logging/debugging.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a set of filters/services in a VM which will share a RootContext and Contexts
	// if applicable (e.g. an Wasm HttpFilter and an Wasm AccessLog). If left blank, all
	// filters/services with a blank root_id with the same vm_id will share Context(s).
	RootId string `protobuf:"bytes,2,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// Configuration for finding or starting VM.
	VmConfig *VmConfig `protobuf:"bytes,3,opt,name=vm_config,json=vmConfig,proto3" json:"vm_config,omitempty"`
	// Filter/service configuration string e.g. a serialized protobuf which will be the
	// argument to the proxy_onConfigure() call.
	Configuration        string   `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginConfig) Reset()         { *m = PluginConfig{} }
func (m *PluginConfig) String() string { return proto.CompactTextString(m) }
func (*PluginConfig) ProtoMessage()    {}
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9483ad3f200a8, []int{1}
}
func (m *PluginConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginConfig.Unmarshal(m, b)
}
func (m *PluginConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginConfig.Marshal(b, m, deterministic)
}
func (m *PluginConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginConfig.Merge(m, src)
}
func (m *PluginConfig) XXX_Size() int {
	return xxx_messageInfo_PluginConfig.Size(m)
}
func (m *PluginConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PluginConfig proto.InternalMessageInfo

func (m *PluginConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginConfig) GetRootId() string {
	if m != nil {
		return m.RootId
	}
	return ""
}

func (m *PluginConfig) GetVmConfig() *VmConfig {
	if m != nil {
		return m.VmConfig
	}
	return nil
}

func (m *PluginConfig) GetConfiguration() string {
	if m != nil {
		return m.Configuration
	}
	return ""
}

// WasmService is configured as a built-in *envoy.wasm_service* `ServiceConfig
// (envoy_api_msg_config.wasm.v2.ServiceConfig)`. This opaque configuration will be used to
// create a Wasm Service.
type WasmService struct {
	// General plugin configuration.
	Config *PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// If true, create a single VM rather than creating one VM per silo. Such a singleton can
	// not be used with filters.
	Singleton bool `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
	// If set add 'stat_prefix' as a prefix to all stats.
	StatPrefix           string   `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WasmService) Reset()         { *m = WasmService{} }
func (m *WasmService) String() string { return proto.CompactTextString(m) }
func (*WasmService) ProtoMessage()    {}
func (*WasmService) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9483ad3f200a8, []int{2}
}
func (m *WasmService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmService.Unmarshal(m, b)
}
func (m *WasmService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmService.Marshal(b, m, deterministic)
}
func (m *WasmService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmService.Merge(m, src)
}
func (m *WasmService) XXX_Size() int {
	return xxx_messageInfo_WasmService.Size(m)
}
func (m *WasmService) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmService.DiscardUnknown(m)
}

var xxx_messageInfo_WasmService proto.InternalMessageInfo

func (m *WasmService) GetConfig() *PluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *WasmService) GetSingleton() bool {
	if m != nil {
		return m.Singleton
	}
	return false
}

func (m *WasmService) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*VmConfig)(nil), "envoy.config.wasm.v2.VmConfig")
	proto.RegisterType((*PluginConfig)(nil), "envoy.config.wasm.v2.PluginConfig")
	proto.RegisterType((*WasmService)(nil), "envoy.config.wasm.v2.WasmService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/api/v2/config/wasm/wasm.proto", fileDescriptor_56d9483ad3f200a8)
}

var fileDescriptor_56d9483ad3f200a8 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x89, 0x6e, 0xb7, 0x3b, 0x59, 0x05, 0x8d, 0x42, 0x87, 0x52, 0x74, 0x19, 0x3c, 0x2c,
	0x88, 0x19, 0x58, 0xc1, 0x83, 0x9e, 0xac, 0x5e, 0xea, 0x69, 0x98, 0x42, 0x05, 0x2f, 0x4b, 0x76,
	0x26, 0x1d, 0xa3, 0x49, 0xde, 0x90, 0xc9, 0xa4, 0xbb, 0x47, 0x6f, 0x7e, 0x0a, 0x3f, 0x87, 0x1f,
	0xc0, 0xaf, 0xe4, 0x5d, 0x92, 0xcc, 0x50, 0x5b, 0x0a, 0xed, 0x65, 0x78, 0xef, 0xff, 0xe6, 0xe5,
	0xbd, 0xdf, 0x9f, 0x87, 0xcf, 0x1a, 0x61, 0xbf, 0xf6, 0x1b, 0x5a, 0x81, 0xca, 0x3b, 0x90, 0xf0,
	0x4a, 0x40, 0xde, 0x48, 0x80, 0xbc, 0x35, 0xf0, 0x8d, 0x57, 0xb6, 0x8b, 0x19, 0x6b, 0x45, 0xce,
	0xb7, 0x96, 0x1b, 0xcd, 0x64, 0xce, 0xb5, 0x83, 0x5d, 0x90, 0xdc, 0x2a, 0xaf, 0x40, 0x9f, 0x8b,
	0x26, 0xbf, 0x60, 0x9d, 0x0a, 0x1f, 0xda, 0x1a, 0xb0, 0x40, 0x9e, 0x86, 0x9f, 0x68, 0xac, 0xd2,
	0x50, 0x70, 0xab, 0xc3, 0xa3, 0x6b, 0xad, 0x86, 0xe7, 0x1b, 0xd6, 0xf1, 0xd8, 0x73, 0x78, 0xe0,
	0x98, 0x14, 0x35, 0xb3, 0x3c, 0x1f, 0x83, 0xa1, 0x40, 0xf8, 0xd6, 0x86, 0xc8, 0xef, 0x10, 0xb5,
	0xec, 0x0f, 0xc2, 0xb3, 0x33, 0xf5, 0x21, 0xbc, 0x4f, 0x9e, 0xe0, 0x3d, 0xa7, 0xd6, 0xa2, 0x4e,
	0xd1, 0x02, 0x2d, 0x93, 0x72, 0xe2, 0xd4, 0x49, 0x4d, 0x52, 0xbc, 0x6f, 0x7a, 0x6d, 0x85, 0xe2,
	0xe9, 0xbd, 0x20, 0x8f, 0x29, 0x79, 0x83, 0x27, 0x15, 0xd4, 0x3c, 0xbd, 0xbf, 0x40, 0xcb, 0xf9,
	0x2a, 0xa3, 0x71, 0x57, 0xd6, 0x0a, 0xea, 0x56, 0xd4, 0x6f, 0x45, 0xdf, 0x77, 0x3b, 0x5d, 0x7d,
	0x64, 0x96, 0x9d, 0x42, 0x6f, 0x2a, 0x5e, 0x86, 0xff, 0xc9, 0x0b, 0xfc, 0x30, 0x02, 0xf5, 0x86,
	0x59, 0x01, 0x3a, 0x9d, 0x84, 0x77, 0xaf, 0x8a, 0xe4, 0x25, 0x7e, 0xcc, 0xa4, 0x84, 0x8b, 0x75,
	0x6b, 0x78, 0x05, 0xaa, 0x15, 0x92, 0xd7, 0xe9, 0xde, 0x02, 0x2d, 0x67, 0xe5, 0xa3, 0x50, 0x28,
	0x2e, 0xf5, 0xec, 0x17, 0xc2, 0x0f, 0x0a, 0xd9, 0x37, 0x42, 0x0f, 0x28, 0x04, 0x4f, 0x34, 0x53,
	0x7c, 0x24, 0xf1, 0x31, 0x39, 0xc0, 0xfb, 0x06, 0xc0, 0x7a, 0xc0, 0x48, 0x32, 0xf5, 0xe9, 0x49,
	0x4d, 0xde, 0xe1, 0xc4, 0xa9, 0x75, 0x1c, 0x3f, 0xd0, 0x3c, 0xa3, 0x37, 0x39, 0x4f, 0x47, 0xab,
	0xca, 0x99, 0x1b, 0x4d, 0xbb, 0x13, 0x4d, 0xf6, 0x13, 0xe1, 0xf9, 0x67, 0xd6, 0xa9, 0x53, 0x6e,
	0x9c, 0xa8, 0x38, 0x79, 0x8b, 0xa7, 0xc3, 0x3c, 0x74, 0xc5, 0xbd, 0x6b, 0xf3, 0xfe, 0x67, 0x2a,
	0x87, 0x0e, 0x72, 0x84, 0x93, 0x4e, 0xe8, 0x46, 0x72, 0x0b, 0x3a, 0x90, 0xcc, 0xca, 0x4b, 0x81,
	0x3c, 0xc7, 0xf3, 0xce, 0x32, 0xeb, 0x6d, 0x3b, 0x17, 0xdb, 0x80, 0x93, 0x94, 0xd8, 0x4b, 0x45,
	0x50, 0x8e, 0x7f, 0xa0, 0xdf, 0x7f, 0x27, 0x08, 0x67, 0x02, 0xe2, 0xcc, 0xd6, 0xc0, 0x76, 0x77,
	0xe3, 0xf8, 0xe3, 0xc4, 0xaf, 0x5c, 0xf8, 0x43, 0x29, 0xd0, 0x97, 0x4f, 0x77, 0xbb, 0xf1, 0xf6,
	0x7b, 0x73, 0xeb, 0x9d, 0x6f, 0xa6, 0xe1, 0xfa, 0x5e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0xa5, 0x01, 0x18, 0x38, 0x03, 0x00, 0x00,
}
