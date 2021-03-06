// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/common/common.proto

package common

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

//* Represents an empty message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

//* A type which contains attestation data for specific platform.
type AttestationData struct {
	//* Type of attestation to perform.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	//* The attestation data.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{1}
}

func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (m *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(m, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestationData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

//* A type which describes the conditions under which a registration
//entry is matched.
type Selector struct {
	//* A selector type represents the type of attestation used in attesting
	//the entity (Eg: AWS, K8).
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	//* The value to be attested.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{2}
}

func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Selector) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

//* Represents a type with a list of Selector.
type Selectors struct {
	//* A list of Selector.
	Entries              []*Selector `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Selectors) Reset()         { *m = Selectors{} }
func (m *Selectors) String() string { return proto.CompactTextString(m) }
func (*Selectors) ProtoMessage()    {}
func (*Selectors) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{3}
}

func (m *Selectors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selectors.Unmarshal(m, b)
}
func (m *Selectors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selectors.Marshal(b, m, deterministic)
}
func (m *Selectors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selectors.Merge(m, src)
}
func (m *Selectors) XXX_Size() int {
	return xxx_messageInfo_Selectors.Size(m)
}
func (m *Selectors) XXX_DiscardUnknown() {
	xxx_messageInfo_Selectors.DiscardUnknown(m)
}

var xxx_messageInfo_Selectors proto.InternalMessageInfo

func (m *Selectors) GetEntries() []*Selector {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Represents an attested SPIRE agent
type AttestedNode struct {
	// Node SPIFFE ID
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Attestation data type
	AttestationDataType string `protobuf:"bytes,2,opt,name=attestation_data_type,json=attestationDataType,proto3" json:"attestation_data_type,omitempty"`
	// Node certificate serial number
	CertSerialNumber string `protobuf:"bytes,3,opt,name=cert_serial_number,json=certSerialNumber,proto3" json:"cert_serial_number,omitempty"`
	// Node certificate not_after (seconds since unix epoch)
	CertNotAfter int64 `protobuf:"varint,4,opt,name=cert_not_after,json=certNotAfter,proto3" json:"cert_not_after,omitempty"`
	// Node certificate serial number
	NewCertSerialNumber string `protobuf:"bytes,5,opt,name=new_cert_serial_number,json=newCertSerialNumber,proto3" json:"new_cert_serial_number,omitempty"`
	// Node certificate not_after (seconds since unix epoch)
	NewCertNotAfter int64 `protobuf:"varint,6,opt,name=new_cert_not_after,json=newCertNotAfter,proto3" json:"new_cert_not_after,omitempty"`
	// Node selectors
	Selectors            []*Selector `protobuf:"bytes,7,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AttestedNode) Reset()         { *m = AttestedNode{} }
func (m *AttestedNode) String() string { return proto.CompactTextString(m) }
func (*AttestedNode) ProtoMessage()    {}
func (*AttestedNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{4}
}

func (m *AttestedNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestedNode.Unmarshal(m, b)
}
func (m *AttestedNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestedNode.Marshal(b, m, deterministic)
}
func (m *AttestedNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestedNode.Merge(m, src)
}
func (m *AttestedNode) XXX_Size() int {
	return xxx_messageInfo_AttestedNode.Size(m)
}
func (m *AttestedNode) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestedNode.DiscardUnknown(m)
}

var xxx_messageInfo_AttestedNode proto.InternalMessageInfo

func (m *AttestedNode) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *AttestedNode) GetAttestationDataType() string {
	if m != nil {
		return m.AttestationDataType
	}
	return ""
}

func (m *AttestedNode) GetCertSerialNumber() string {
	if m != nil {
		return m.CertSerialNumber
	}
	return ""
}

func (m *AttestedNode) GetCertNotAfter() int64 {
	if m != nil {
		return m.CertNotAfter
	}
	return 0
}

func (m *AttestedNode) GetNewCertSerialNumber() string {
	if m != nil {
		return m.NewCertSerialNumber
	}
	return ""
}

func (m *AttestedNode) GetNewCertNotAfter() int64 {
	if m != nil {
		return m.NewCertNotAfter
	}
	return 0
}

func (m *AttestedNode) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

//* This is a curated record that the Server uses to set up and
//manage the various registered nodes and workloads that are controlled by it.
type RegistrationEntry struct {
	//* A list of selectors.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	//* The SPIFFE ID of an entity that is authorized to attest the validity
	//of a selector
	ParentId string `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	//* The SPIFFE ID is a structured string used to identify a resource or
	//caller. It is defined as a URI comprising a “trust domain” and an
	//associated path.
	SpiffeId string `protobuf:"bytes,3,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	//* Time to live.
	Ttl int32 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	//* A list of federated trust domain SPIFFE IDs.
	FederatesWith []string `protobuf:"bytes,5,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
	//* Entry ID
	EntryId string `protobuf:"bytes,6,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	//* Whether or not the workload is an admin workload. Admin workloads
	//can use their SVID's to authenticate with the Registration API, for
	//example.
	Admin bool `protobuf:"varint,7,opt,name=admin,proto3" json:"admin,omitempty"`
	//* To enable signing CA CSR in upstream spire server
	Downstream bool `protobuf:"varint,8,opt,name=downstream,proto3" json:"downstream,omitempty"`
	//* Expiration of this entry, in seconds from epoch
	EntryExpiry int64 `protobuf:"varint,9,opt,name=entryExpiry,proto3" json:"entryExpiry,omitempty"`
	//* DNS entries
	DnsNames             []string `protobuf:"bytes,10,rep,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationEntry) Reset()         { *m = RegistrationEntry{} }
func (m *RegistrationEntry) String() string { return proto.CompactTextString(m) }
func (*RegistrationEntry) ProtoMessage()    {}
func (*RegistrationEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{5}
}

func (m *RegistrationEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationEntry.Unmarshal(m, b)
}
func (m *RegistrationEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationEntry.Marshal(b, m, deterministic)
}
func (m *RegistrationEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationEntry.Merge(m, src)
}
func (m *RegistrationEntry) XXX_Size() int {
	return xxx_messageInfo_RegistrationEntry.Size(m)
}
func (m *RegistrationEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationEntry proto.InternalMessageInfo

func (m *RegistrationEntry) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *RegistrationEntry) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *RegistrationEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *RegistrationEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *RegistrationEntry) GetFederatesWith() []string {
	if m != nil {
		return m.FederatesWith
	}
	return nil
}

func (m *RegistrationEntry) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

func (m *RegistrationEntry) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *RegistrationEntry) GetDownstream() bool {
	if m != nil {
		return m.Downstream
	}
	return false
}

func (m *RegistrationEntry) GetEntryExpiry() int64 {
	if m != nil {
		return m.EntryExpiry
	}
	return 0
}

func (m *RegistrationEntry) GetDnsNames() []string {
	if m != nil {
		return m.DnsNames
	}
	return nil
}

//* A list of registration entries.
type RegistrationEntries struct {
	//* A list of RegistrationEntry.
	Entries              []*RegistrationEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegistrationEntries) Reset()         { *m = RegistrationEntries{} }
func (m *RegistrationEntries) String() string { return proto.CompactTextString(m) }
func (*RegistrationEntries) ProtoMessage()    {}
func (*RegistrationEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{6}
}

func (m *RegistrationEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationEntries.Unmarshal(m, b)
}
func (m *RegistrationEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationEntries.Marshal(b, m, deterministic)
}
func (m *RegistrationEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationEntries.Merge(m, src)
}
func (m *RegistrationEntries) XXX_Size() int {
	return xxx_messageInfo_RegistrationEntries.Size(m)
}
func (m *RegistrationEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationEntries.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationEntries proto.InternalMessageInfo

func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

//* Certificate represents a ASN.1/DER encoded X509 certificate
type Certificate struct {
	DerBytes             []byte   `protobuf:"bytes,1,opt,name=der_bytes,json=derBytes,proto3" json:"der_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{7}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetDerBytes() []byte {
	if m != nil {
		return m.DerBytes
	}
	return nil
}

//* PublicKey represents a PKIX encoded public key
type PublicKey struct {
	//* PKIX encoded key data
	PkixBytes []byte `protobuf:"bytes,1,opt,name=pkix_bytes,json=pkixBytes,proto3" json:"pkix_bytes,omitempty"`
	//* key identifier
	Kid string `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
	//* not after (seconds since unix epoch, 0 means "never expires")
	NotAfter             int64    `protobuf:"varint,3,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{8}
}

func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetPkixBytes() []byte {
	if m != nil {
		return m.PkixBytes
	}
	return nil
}

func (m *PublicKey) GetKid() string {
	if m != nil {
		return m.Kid
	}
	return ""
}

func (m *PublicKey) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

type Bundle struct {
	//* the SPIFFE ID of the trust domain the bundle belongs to
	TrustDomainId string `protobuf:"bytes,1,opt,name=trust_domain_id,json=trustDomainId,proto3" json:"trust_domain_id,omitempty"`
	//* list of root CA certificates
	RootCas []*Certificate `protobuf:"bytes,2,rep,name=root_cas,json=rootCas,proto3" json:"root_cas,omitempty"`
	//* list of JWT signing keys
	JwtSigningKeys []*PublicKey `protobuf:"bytes,3,rep,name=jwt_signing_keys,json=jwtSigningKeys,proto3" json:"jwt_signing_keys,omitempty"`
	//* refresh hint is a hint, in seconds, on how often a bundle consumer
	// should poll for bundle updates
	RefreshHint          int64    `protobuf:"varint,4,opt,name=refresh_hint,json=refreshHint,proto3" json:"refresh_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{9}
}

func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle.Unmarshal(m, b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
}
func (m *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(m, src)
}
func (m *Bundle) XXX_Size() int {
	return xxx_messageInfo_Bundle.Size(m)
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetTrustDomainId() string {
	if m != nil {
		return m.TrustDomainId
	}
	return ""
}

func (m *Bundle) GetRootCas() []*Certificate {
	if m != nil {
		return m.RootCas
	}
	return nil
}

func (m *Bundle) GetJwtSigningKeys() []*PublicKey {
	if m != nil {
		return m.JwtSigningKeys
	}
	return nil
}

func (m *Bundle) GetRefreshHint() int64 {
	if m != nil {
		return m.RefreshHint
	}
	return 0
}

type BundleMask struct {
	RootCas              bool     `protobuf:"varint,1,opt,name=root_cas,json=rootCas,proto3" json:"root_cas,omitempty"`
	JwtSigningKeys       bool     `protobuf:"varint,2,opt,name=jwt_signing_keys,json=jwtSigningKeys,proto3" json:"jwt_signing_keys,omitempty"`
	RefreshHint          bool     `protobuf:"varint,3,opt,name=refresh_hint,json=refreshHint,proto3" json:"refresh_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BundleMask) Reset()         { *m = BundleMask{} }
func (m *BundleMask) String() string { return proto.CompactTextString(m) }
func (*BundleMask) ProtoMessage()    {}
func (*BundleMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11412a53cc81147, []int{10}
}

func (m *BundleMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleMask.Unmarshal(m, b)
}
func (m *BundleMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleMask.Marshal(b, m, deterministic)
}
func (m *BundleMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleMask.Merge(m, src)
}
func (m *BundleMask) XXX_Size() int {
	return xxx_messageInfo_BundleMask.Size(m)
}
func (m *BundleMask) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleMask.DiscardUnknown(m)
}

var xxx_messageInfo_BundleMask proto.InternalMessageInfo

func (m *BundleMask) GetRootCas() bool {
	if m != nil {
		return m.RootCas
	}
	return false
}

func (m *BundleMask) GetJwtSigningKeys() bool {
	if m != nil {
		return m.JwtSigningKeys
	}
	return false
}

func (m *BundleMask) GetRefreshHint() bool {
	if m != nil {
		return m.RefreshHint
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "spire.common.Empty")
	proto.RegisterType((*AttestationData)(nil), "spire.common.AttestationData")
	proto.RegisterType((*Selector)(nil), "spire.common.Selector")
	proto.RegisterType((*Selectors)(nil), "spire.common.Selectors")
	proto.RegisterType((*AttestedNode)(nil), "spire.common.AttestedNode")
	proto.RegisterType((*RegistrationEntry)(nil), "spire.common.RegistrationEntry")
	proto.RegisterType((*RegistrationEntries)(nil), "spire.common.RegistrationEntries")
	proto.RegisterType((*Certificate)(nil), "spire.common.Certificate")
	proto.RegisterType((*PublicKey)(nil), "spire.common.PublicKey")
	proto.RegisterType((*Bundle)(nil), "spire.common.Bundle")
	proto.RegisterType((*BundleMask)(nil), "spire.common.BundleMask")
}

func init() { proto.RegisterFile("spire/common/common.proto", fileDescriptor_c11412a53cc81147) }

var fileDescriptor_c11412a53cc81147 = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6f, 0xe3, 0x44,
	0x10, 0x96, 0xe3, 0x4b, 0x62, 0x4f, 0x72, 0x6d, 0xd9, 0x83, 0xc3, 0x15, 0x02, 0x82, 0x05, 0x28,
	0x3a, 0x4e, 0x09, 0xba, 0xeb, 0xcb, 0x3d, 0xf0, 0xd0, 0x5e, 0x2b, 0x11, 0x55, 0x44, 0x95, 0x8b,
	0x84, 0xe0, 0xc5, 0xda, 0x64, 0x27, 0xc9, 0x36, 0xf1, 0xda, 0xda, 0x9d, 0x90, 0xfa, 0x6f, 0xa4,
	0x7f, 0x14, 0xda, 0x75, 0x7e, 0x53, 0x01, 0x4f, 0xde, 0xfd, 0x76, 0xe6, 0xdb, 0x6f, 0x66, 0x3e,
	0x2f, 0x9c, 0x9b, 0x42, 0x6a, 0xec, 0x8f, 0xf3, 0x2c, 0xcb, 0xd5, 0xfa, 0xd3, 0x2b, 0x74, 0x4e,
	0x39, 0x6b, 0xbb, 0xa3, 0x5e, 0x85, 0xc5, 0x4d, 0xa8, 0xdf, 0x64, 0x05, 0x95, 0xf1, 0x07, 0x38,
	0xbd, 0x24, 0x42, 0x43, 0x9c, 0x64, 0xae, 0xae, 0x39, 0x71, 0xc6, 0xe0, 0x05, 0x95, 0x05, 0x46,
	0x5e, 0xc7, 0xeb, 0x86, 0x89, 0x5b, 0x5b, 0x4c, 0x70, 0xe2, 0x51, 0xad, 0xe3, 0x75, 0xdb, 0x89,
	0x5b, 0xc7, 0x17, 0x10, 0xdc, 0xe3, 0x02, 0xc7, 0x94, 0xeb, 0x67, 0x73, 0x3e, 0x85, 0xfa, 0x9f,
	0x7c, 0xb1, 0x44, 0x97, 0x14, 0x26, 0xd5, 0x26, 0xfe, 0x09, 0xc2, 0x4d, 0x96, 0x61, 0x3f, 0x42,
	0x13, 0x15, 0x69, 0x89, 0x26, 0xf2, 0x3a, 0x7e, 0xb7, 0xf5, 0xee, 0x75, 0x6f, 0x5f, 0x66, 0x6f,
	0x13, 0x99, 0x6c, 0xc2, 0xe2, 0xbf, 0x6a, 0xd0, 0xae, 0x04, 0xa3, 0x18, 0xe6, 0x02, 0xd9, 0x17,
	0x10, 0x9a, 0x42, 0x4e, 0x26, 0x98, 0x4a, 0xb1, 0xbe, 0x3e, 0xa8, 0x80, 0x81, 0x60, 0xef, 0xe0,
	0x33, 0xbe, 0xab, 0x2e, 0xb5, 0xb2, 0x53, 0xa7, 0xb3, 0x92, 0xf4, 0x8a, 0x1f, 0x96, 0xfe, 0xab,
	0x95, 0xfd, 0x16, 0xd8, 0x18, 0x35, 0xa5, 0x06, 0xb5, 0xe4, 0x8b, 0x54, 0x2d, 0xb3, 0x11, 0xea,
	0xc8, 0x77, 0x09, 0x67, 0xf6, 0xe4, 0xde, 0x1d, 0x0c, 0x1d, 0xce, 0xbe, 0x85, 0x13, 0x17, 0xad,
	0x72, 0x4a, 0xf9, 0x84, 0x50, 0x47, 0x2f, 0x3a, 0x5e, 0xd7, 0x4f, 0xda, 0x16, 0x1d, 0xe6, 0x74,
	0x69, 0x31, 0xf6, 0x1e, 0x5e, 0x2b, 0x5c, 0xa5, 0xcf, 0xf0, 0xd6, 0x2b, 0x21, 0x0a, 0x57, 0x1f,
	0x8f, 0xa9, 0x7f, 0x00, 0xb6, 0x4d, 0xda, 0xd1, 0x37, 0x1c, 0xfd, 0xe9, 0x3a, 0x61, 0x7b, 0xc3,
	0x05, 0x84, 0x66, 0xd3, 0xd6, 0xa8, 0xf9, 0xaf, 0xbd, 0xdc, 0x05, 0xc6, 0x4f, 0x35, 0xf8, 0x24,
	0xc1, 0xa9, 0x34, 0xa4, 0x5d, 0x13, 0x6e, 0x14, 0xe9, 0xf2, 0x90, 0xcb, 0xfb, 0x9f, 0x5c, 0x76,
	0x10, 0x05, 0xd7, 0xa8, 0xc8, 0x0e, 0xa2, 0xea, 0x6f, 0x50, 0x01, 0x03, 0x71, 0x38, 0x25, 0xff,
	0x68, 0x4a, 0x67, 0xe0, 0x13, 0x2d, 0x5c, 0xe3, 0xea, 0x89, 0x5d, 0xb2, 0xef, 0xe0, 0x64, 0x82,
	0x02, 0x35, 0x27, 0x34, 0xe9, 0x4a, 0xd2, 0x2c, 0xaa, 0x77, 0xfc, 0x6e, 0x98, 0xbc, 0xdc, 0xa2,
	0xbf, 0x49, 0x9a, 0xb1, 0x73, 0x08, 0xac, 0x2f, 0x4a, 0x4b, 0xda, 0x70, 0xa4, 0xce, 0x27, 0xe5,
	0x40, 0x58, 0xf3, 0x71, 0x91, 0x49, 0x15, 0x35, 0x3b, 0x5e, 0x37, 0x48, 0xaa, 0x0d, 0xfb, 0x0a,
	0x40, 0xe4, 0x2b, 0x65, 0x48, 0x23, 0xcf, 0xa2, 0xc0, 0x1d, 0xed, 0x21, 0xac, 0x03, 0x2d, 0x47,
	0x70, 0xf3, 0x58, 0x48, 0x5d, 0x46, 0xa1, 0xeb, 0xf5, 0x3e, 0x64, 0x0b, 0x11, 0xca, 0xa4, 0x8a,
	0x67, 0x68, 0x22, 0x70, 0xa2, 0x02, 0xa1, 0xcc, 0xd0, 0xee, 0xe3, 0x3b, 0x78, 0x75, 0xdc, 0x4d,
	0x89, 0x86, 0x7d, 0x38, 0x76, 0xf9, 0xd7, 0x87, 0xdd, 0xfc, 0xc7, 0x04, 0x76, 0x76, 0x7f, 0x03,
	0x2d, 0x3b, 0x66, 0x39, 0x91, 0x63, 0x4e, 0xce, 0xec, 0x02, 0x75, 0x3a, 0x2a, 0xc9, 0x71, 0xd9,
	0x7f, 0x31, 0x10, 0xa8, 0xaf, 0xec, 0x3e, 0xfe, 0x1d, 0xc2, 0xbb, 0xe5, 0x68, 0x21, 0xc7, 0xb7,
	0x58, 0xb2, 0x2f, 0x01, 0x8a, 0xb9, 0x7c, 0x3c, 0x08, 0x0d, 0x2d, 0xe2, 0x62, 0x6d, 0xcb, 0xe7,
	0xdb, 0x31, 0xd9, 0xa5, 0xa5, 0xde, 0x99, 0xcc, 0x77, 0x85, 0x07, 0x6a, 0xed, 0xae, 0xf8, 0xc9,
	0x83, 0xc6, 0xd5, 0x52, 0x89, 0x05, 0xb2, 0xef, 0xe1, 0x94, 0xf4, 0xd2, 0x50, 0x2a, 0xf2, 0x8c,
	0x4b, 0xb5, 0xfb, 0xeb, 0x5e, 0x3a, 0xf8, 0xda, 0xa1, 0x03, 0xc1, 0x2e, 0x20, 0xd0, 0x79, 0x4e,
	0xe9, 0x98, 0x9b, 0xa8, 0xe6, 0xaa, 0x3e, 0x3f, 0xac, 0x7a, 0xaf, 0xae, 0xa4, 0x69, 0x43, 0x3f,
	0x72, 0xc3, 0x2e, 0xe1, 0xec, 0x61, 0x45, 0xa9, 0x91, 0x53, 0x25, 0xd5, 0x34, 0x9d, 0x63, 0x69,
	0x22, 0xdf, 0x65, 0x7f, 0x7e, 0x98, 0xbd, 0xad, 0x34, 0x39, 0x79, 0x58, 0xd1, 0x7d, 0x15, 0x7f,
	0x8b, 0xa5, 0x61, 0xdf, 0x40, 0x5b, 0xe3, 0x44, 0xa3, 0x99, 0xa5, 0x33, 0xa9, 0x68, 0xfd, 0x3f,
	0xb6, 0xd6, 0xd8, 0xcf, 0x52, 0x51, 0x4c, 0x00, 0x55, 0x35, 0xbf, 0x70, 0x33, 0xb7, 0x2e, 0xda,
	0x2a, 0xf5, 0x9c, 0x25, 0xb6, 0x72, 0xba, 0xcf, 0xc8, 0xa9, 0xb9, 0x90, 0xff, 0xba, 0xd5, 0x77,
	0x51, 0xfb, 0xb7, 0x5e, 0xbd, 0xfd, 0xe3, 0xcd, 0x54, 0xd2, 0x6c, 0x39, 0xb2, 0x35, 0xf4, 0x2b,
	0xf7, 0xf7, 0xab, 0x07, 0xdb, 0x3d, 0xd1, 0xfd, 0xfd, 0xc7, 0x7b, 0xd4, 0x70, 0xd8, 0xfb, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x97, 0x6b, 0x15, 0x4a, 0xd3, 0x05, 0x00, 0x00,
}
