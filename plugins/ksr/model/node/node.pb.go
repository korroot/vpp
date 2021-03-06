// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

// Package node defines data model for Kubernetes Node.

package node

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// These are valid address type of node.
type NodeAddress_AddressType int32

const (
	NodeAddress_NodeUnknownAddr NodeAddress_AddressType = 0
	NodeAddress_NodeHostName    NodeAddress_AddressType = 1
	NodeAddress_NodeExternalIP  NodeAddress_AddressType = 2
	NodeAddress_NodeInternalIP  NodeAddress_AddressType = 3
	NodeAddress_NodeExternalDNS NodeAddress_AddressType = 4
	NodeAddress_NodeInternalDNS NodeAddress_AddressType = 5
)

var NodeAddress_AddressType_name = map[int32]string{
	0: "NodeUnknownAddr",
	1: "NodeHostName",
	2: "NodeExternalIP",
	3: "NodeInternalIP",
	4: "NodeExternalDNS",
	5: "NodeInternalDNS",
}

var NodeAddress_AddressType_value = map[string]int32{
	"NodeUnknownAddr": 0,
	"NodeHostName":    1,
	"NodeExternalIP":  2,
	"NodeInternalIP":  3,
	"NodeExternalDNS": 4,
	"NodeInternalDNS": 5,
}

func (x NodeAddress_AddressType) String() string {
	return proto.EnumName(NodeAddress_AddressType_name, int32(x))
}

func (NodeAddress_AddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1, 0}
}

// Pod is a collection of containers that can run on a host.
// This resource is created by clients and scheduled onto hosts.
type Node struct {
	// Name of the node unique within the cluster.
	// Cannot be updated.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// PodCIDR represents the pod IP range assigned to the node.
	// +optional
	Pod_CIDR string `protobuf:"bytes,2,opt,name=pod_CIDR,json=podCIDR,proto3" json:"pod_CIDR,omitempty"`
	// ID of the node assigned by the cloud provider in the format:
	// <ProviderName>://<ProviderSpecificNodeID>
	// +optional
	Provider_ID string `protobuf:"bytes,3,opt,name=provider_ID,json=providerID,proto3" json:"provider_ID,omitempty"`
	// List of addresses reachable to the node.
	// Queried from cloud provider, if available.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	// +optional
	Addresses []*NodeAddress `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// Set of ids/uuids to uniquely identify the node.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#info
	// +optional
	NodeInfo             *NodeSystemInfo `protobuf:"bytes,5,opt,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetPod_CIDR() string {
	if m != nil {
		return m.Pod_CIDR
	}
	return ""
}

func (m *Node) GetProvider_ID() string {
	if m != nil {
		return m.Provider_ID
	}
	return ""
}

func (m *Node) GetAddresses() []*NodeAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Node) GetNodeInfo() *NodeSystemInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

// NodeAddress contains information for the node's address.
type NodeAddress struct {
	// Node address type, one of Hostname, ExternalIP or InternalIP.
	Type NodeAddress_AddressType `protobuf:"varint,1,opt,name=type,proto3,enum=node.NodeAddress_AddressType" json:"type,omitempty"`
	// The node address.
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (m *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(m, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetType() NodeAddress_AddressType {
	if m != nil {
		return m.Type
	}
	return NodeAddress_NodeUnknownAddr
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
type NodeSystemInfo struct {
	// MachineID reported by the node. For unique machine identification
	// in the cluster this field is preferred. Learn more from man(5)
	// machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	Machine_ID string `protobuf:"bytes,1,opt,name=machine_ID,json=machineID,proto3" json:"machine_ID,omitempty"`
	// SystemUUID reported by the node. For unique machine identification
	// MachineID is preferred. This field is specific to Red Hat hosts
	// https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	System_UUID string `protobuf:"bytes,2,opt,name=system_UUID,json=systemUUID,proto3" json:"system_UUID,omitempty"`
	// Boot ID reported by the node.
	Boot_ID string `protobuf:"bytes,3,opt,name=boot_ID,json=bootID,proto3" json:"boot_ID,omitempty"`
	// Kernel Version reported by the node from 'uname -r'
	// (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `protobuf:"bytes,4,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
	// OS Image reported by the node from /etc/os-release
	// (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage string `protobuf:"bytes,5,opt,name=os_image,json=osImage,proto3" json:"os_image,omitempty"`
	// ContainerRuntime Version reported by the node through runtime remote API
	// (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `protobuf:"bytes,6,opt,name=container_runtime_version,json=containerRuntimeVersion,proto3" json:"container_runtime_version,omitempty"`
	// Kubelet Version reported by the node.
	KubeletVersion string `protobuf:"bytes,7,opt,name=kubelet_version,json=kubeletVersion,proto3" json:"kubelet_version,omitempty"`
	// KubeProxy Version reported by the node.
	KubeProxyVersion string `protobuf:"bytes,8,opt,name=KubeProxyVersion,proto3" json:"KubeProxyVersion,omitempty"`
	OperatingSystem  string `protobuf:"bytes,9,opt,name=OperatingSystem,proto3" json:"OperatingSystem,omitempty"`
	// The Architecture reported by the node
	Architecture         string   `protobuf:"bytes,10,opt,name=Architecture,proto3" json:"Architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeSystemInfo) Reset()         { *m = NodeSystemInfo{} }
func (m *NodeSystemInfo) String() string { return proto.CompactTextString(m) }
func (*NodeSystemInfo) ProtoMessage()    {}
func (*NodeSystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}
func (m *NodeSystemInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSystemInfo.Unmarshal(m, b)
}
func (m *NodeSystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSystemInfo.Marshal(b, m, deterministic)
}
func (m *NodeSystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSystemInfo.Merge(m, src)
}
func (m *NodeSystemInfo) XXX_Size() int {
	return xxx_messageInfo_NodeSystemInfo.Size(m)
}
func (m *NodeSystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSystemInfo proto.InternalMessageInfo

func (m *NodeSystemInfo) GetMachine_ID() string {
	if m != nil {
		return m.Machine_ID
	}
	return ""
}

func (m *NodeSystemInfo) GetSystem_UUID() string {
	if m != nil {
		return m.System_UUID
	}
	return ""
}

func (m *NodeSystemInfo) GetBoot_ID() string {
	if m != nil {
		return m.Boot_ID
	}
	return ""
}

func (m *NodeSystemInfo) GetKernelVersion() string {
	if m != nil {
		return m.KernelVersion
	}
	return ""
}

func (m *NodeSystemInfo) GetOsImage() string {
	if m != nil {
		return m.OsImage
	}
	return ""
}

func (m *NodeSystemInfo) GetContainerRuntimeVersion() string {
	if m != nil {
		return m.ContainerRuntimeVersion
	}
	return ""
}

func (m *NodeSystemInfo) GetKubeletVersion() string {
	if m != nil {
		return m.KubeletVersion
	}
	return ""
}

func (m *NodeSystemInfo) GetKubeProxyVersion() string {
	if m != nil {
		return m.KubeProxyVersion
	}
	return ""
}

func (m *NodeSystemInfo) GetOperatingSystem() string {
	if m != nil {
		return m.OperatingSystem
	}
	return ""
}

func (m *NodeSystemInfo) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func init() {
	proto.RegisterEnum("node.NodeAddress_AddressType", NodeAddress_AddressType_name, NodeAddress_AddressType_value)
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodeAddress)(nil), "node.NodeAddress")
	proto.RegisterType((*NodeSystemInfo)(nil), "node.NodeSystemInfo")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x80, 0x7f, 0x37, 0x6e, 0x1c, 0x8f, 0xfb, 0x27, 0x66, 0x40, 0xea, 0xf6, 0x50, 0x11, 0x59,
	0x42, 0x44, 0x1c, 0x82, 0x1a, 0x6e, 0xdc, 0x2a, 0x8c, 0xc4, 0x0a, 0x29, 0x54, 0x2e, 0xe1, 0x6a,
	0x39, 0xf1, 0xb4, 0xb5, 0x12, 0xef, 0x5a, 0xeb, 0x4d, 0x69, 0x5e, 0x80, 0x03, 0x8f, 0xc4, 0xeb,
	0xf0, 0x22, 0x68, 0xd7, 0x76, 0x52, 0xe8, 0xc9, 0x3b, 0xdf, 0x7c, 0x33, 0xeb, 0xdd, 0xb1, 0x01,
	0x84, 0xcc, 0x69, 0x5a, 0x29, 0xa9, 0x25, 0xba, 0x66, 0x1d, 0xfd, 0x72, 0xc0, 0x9d, 0xcb, 0x9c,
	0x10, 0xc1, 0x15, 0x59, 0x49, 0xcc, 0x19, 0x3b, 0x13, 0x3f, 0xb1, 0x6b, 0x3c, 0x83, 0x41, 0x25,
	0xf3, 0xf4, 0x03, 0x8f, 0x13, 0x76, 0x64, 0xb9, 0x57, 0xc9, 0xdc, 0x84, 0xf8, 0x12, 0x82, 0x4a,
	0xc9, 0xfb, 0x22, 0x27, 0x95, 0xf2, 0x98, 0xf5, 0x6c, 0x16, 0x3a, 0xc4, 0x63, 0x7c, 0x0b, 0x7e,
	0x96, 0xe7, 0x8a, 0xea, 0x9a, 0x6a, 0xe6, 0x8e, 0x7b, 0x93, 0x60, 0xf6, 0x6c, 0x6a, 0xb7, 0x37,
	0xdb, 0x5d, 0x36, 0xa9, 0xe4, 0xe0, 0xe0, 0x05, 0xf8, 0x26, 0x9d, 0x16, 0xe2, 0x46, 0xb2, 0xe3,
	0xb1, 0x33, 0x09, 0x66, 0x2f, 0x0e, 0x05, 0xd7, 0xbb, 0x5a, 0x53, 0xc9, 0xc5, 0x8d, 0x4c, 0x06,
	0x06, 0x9a, 0x55, 0xf4, 0xdb, 0x81, 0xe0, 0x51, 0x37, 0xbc, 0x00, 0x57, 0xef, 0xaa, 0xe6, 0x0c,
	0xc3, 0xd9, 0xf9, 0x93, 0xed, 0xa6, 0xed, 0xf3, 0xeb, 0xae, 0xa2, 0xc4, 0xaa, 0xc8, 0xc0, 0x6b,
	0x5f, 0xa1, 0x3b, 0x61, 0x1b, 0x46, 0x3f, 0x1c, 0x08, 0x1e, 0xf9, 0xf8, 0x1c, 0x46, 0xa6, 0xd5,
	0x42, 0xac, 0x85, 0xfc, 0x2e, 0x4c, 0x26, 0xfc, 0x0f, 0x43, 0x38, 0x31, 0xf0, 0x93, 0xac, 0xf5,
	0x3c, 0x2b, 0x29, 0x74, 0x10, 0x61, 0x68, 0xc8, 0xc7, 0x07, 0x4d, 0x4a, 0x64, 0x1b, 0x7e, 0x15,
	0x1e, 0x75, 0x8c, 0x8b, 0x3d, 0xeb, 0x75, 0xed, 0x3a, 0x2f, 0x9e, 0x5f, 0x87, 0x6e, 0x07, 0x3b,
	0xd1, 0xc0, 0xe3, 0xe8, 0x67, 0xaf, 0x29, 0x3f, 0x5c, 0x01, 0x9e, 0x03, 0x94, 0xd9, 0xea, 0xae,
	0x10, 0x64, 0x2e, 0xbf, 0x19, 0x99, 0xdf, 0x12, 0x1e, 0x9b, 0xe1, 0xd4, 0x56, 0x4e, 0x17, 0x0b,
	0x1e, 0xb7, 0x07, 0x83, 0x06, 0x19, 0x82, 0xa7, 0xe0, 0x2d, 0xa5, 0xd4, 0x87, 0xc9, 0xf5, 0x4d,
	0xc8, 0x63, 0x7c, 0x05, 0xc3, 0x35, 0x29, 0x41, 0x9b, 0xf4, 0x9e, 0x54, 0x5d, 0x48, 0xc1, 0x5c,
	0x9b, 0xff, 0xbf, 0xa1, 0xdf, 0x1a, 0x68, 0x3e, 0x0c, 0x59, 0xa7, 0x45, 0x99, 0xdd, 0x92, 0x1d,
	0x95, 0x9f, 0x78, 0xb2, 0xe6, 0x26, 0xc4, 0xf7, 0x70, 0xb6, 0x92, 0x42, 0x67, 0x85, 0x20, 0x95,
	0xaa, 0xad, 0xd0, 0x45, 0x49, 0xfb, 0x66, 0x7d, 0xeb, 0x9e, 0xee, 0x85, 0xa4, 0xc9, 0x77, 0x6d,
	0x5f, 0xc3, 0x68, 0xbd, 0x5d, 0xd2, 0x86, 0xf4, 0xbe, 0xc2, 0xb3, 0x15, 0xc3, 0x16, 0x77, 0xe2,
	0x1b, 0x08, 0x3f, 0x6f, 0x97, 0x74, 0xa5, 0xe4, 0xc3, 0xae, 0x65, 0x6c, 0x60, 0xcd, 0x27, 0x1c,
	0x27, 0x30, 0xfa, 0x52, 0x91, 0xca, 0x74, 0x21, 0x6e, 0x9b, 0x2b, 0x64, 0xbe, 0x55, 0xff, 0xc5,
	0x18, 0xc1, 0xc9, 0xa5, 0x5a, 0xdd, 0x15, 0x9a, 0x56, 0x7a, 0xab, 0x88, 0x81, 0xd5, 0xfe, 0x62,
	0xcb, 0xbe, 0xfd, 0x79, 0xde, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x51, 0x7c, 0x50, 0x68, 0x4a,
	0x03, 0x00, 0x00,
}
