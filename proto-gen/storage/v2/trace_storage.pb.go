// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace_storage.proto

package storage

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	v11 "go.opentelemetry.io/proto/otlp/common/v1"
	v1 "go.opentelemetry.io/proto/otlp/trace/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GetTraceParams represents the query for a single trace from the storage backend.
type GetTraceParams struct {
	// trace_id is a 16 byte array containing the unique identifier for the trace to query.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// start_time is the start of the time interval to search for the trace_id.
	//
	// This field is optional.
	StartTime *time.Time `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time,omitempty"`
	// end_time is the end of the time interval to search for the trace_id.
	//
	// This field is optional.
	EndTime              *time.Time `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetTraceParams) Reset()         { *m = GetTraceParams{} }
func (m *GetTraceParams) String() string { return proto.CompactTextString(m) }
func (*GetTraceParams) ProtoMessage()    {}
func (*GetTraceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{0}
}
func (m *GetTraceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTraceParams.Unmarshal(m, b)
}
func (m *GetTraceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTraceParams.Marshal(b, m, deterministic)
}
func (m *GetTraceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTraceParams.Merge(m, src)
}
func (m *GetTraceParams) XXX_Size() int {
	return xxx_messageInfo_GetTraceParams.Size(m)
}
func (m *GetTraceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTraceParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetTraceParams proto.InternalMessageInfo

func (m *GetTraceParams) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *GetTraceParams) GetStartTime() *time.Time {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *GetTraceParams) GetEndTime() *time.Time {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// GetTracesRequest represents a request to retrieve multiple traces.
type GetTracesRequest struct {
	Query                []*GetTraceParams `protobuf:"bytes,1,rep,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetTracesRequest) Reset()         { *m = GetTracesRequest{} }
func (m *GetTracesRequest) String() string { return proto.CompactTextString(m) }
func (*GetTracesRequest) ProtoMessage()    {}
func (*GetTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{1}
}
func (m *GetTracesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTracesRequest.Unmarshal(m, b)
}
func (m *GetTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTracesRequest.Marshal(b, m, deterministic)
}
func (m *GetTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTracesRequest.Merge(m, src)
}
func (m *GetTracesRequest) XXX_Size() int {
	return xxx_messageInfo_GetTracesRequest.Size(m)
}
func (m *GetTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTracesRequest proto.InternalMessageInfo

func (m *GetTracesRequest) GetQuery() []*GetTraceParams {
	if m != nil {
		return m.Query
	}
	return nil
}

// GetServicesRequest represents a request to get service names.
type GetServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesRequest) Reset()         { *m = GetServicesRequest{} }
func (m *GetServicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetServicesRequest) ProtoMessage()    {}
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{2}
}
func (m *GetServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesRequest.Unmarshal(m, b)
}
func (m *GetServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesRequest.Marshal(b, m, deterministic)
}
func (m *GetServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesRequest.Merge(m, src)
}
func (m *GetServicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetServicesRequest.Size(m)
}
func (m *GetServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesRequest proto.InternalMessageInfo

// GetServicesResponse represents the response for GetServicesRequest.
type GetServicesResponse struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesResponse) Reset()         { *m = GetServicesResponse{} }
func (m *GetServicesResponse) String() string { return proto.CompactTextString(m) }
func (*GetServicesResponse) ProtoMessage()    {}
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{3}
}
func (m *GetServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesResponse.Unmarshal(m, b)
}
func (m *GetServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesResponse.Marshal(b, m, deterministic)
}
func (m *GetServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesResponse.Merge(m, src)
}
func (m *GetServicesResponse) XXX_Size() int {
	return xxx_messageInfo_GetServicesResponse.Size(m)
}
func (m *GetServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesResponse proto.InternalMessageInfo

func (m *GetServicesResponse) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

// GetOperationsRequest represents a request to get operation names.
type GetOperationsRequest struct {
	// service is the name of the service for which to get operation names.
	//
	// This field is required.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// span_kind is the type of span which is used to distinguish between
	// spans generated in a particular context.
	//
	// This field is optional.
	SpanKind             v1.Span_SpanKind `protobuf:"varint,2,opt,name=span_kind,json=spanKind,proto3,enum=opentelemetry.proto.trace.v1.Span_SpanKind" json:"span_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetOperationsRequest) Reset()         { *m = GetOperationsRequest{} }
func (m *GetOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationsRequest) ProtoMessage()    {}
func (*GetOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{4}
}
func (m *GetOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationsRequest.Unmarshal(m, b)
}
func (m *GetOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationsRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationsRequest.Merge(m, src)
}
func (m *GetOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationsRequest.Size(m)
}
func (m *GetOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationsRequest proto.InternalMessageInfo

func (m *GetOperationsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GetOperationsRequest) GetSpanKind() v1.Span_SpanKind {
	if m != nil {
		return m.SpanKind
	}
	return v1.Span_SPAN_KIND_UNSPECIFIED
}

// Operation contains information about an operation for a given service.
type Operation struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SpanKind             v1.Span_SpanKind `protobuf:"varint,2,opt,name=span_kind,json=spanKind,proto3,enum=opentelemetry.proto.trace.v1.Span_SpanKind" json:"span_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{5}
}
func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetSpanKind() v1.Span_SpanKind {
	if m != nil {
		return m.SpanKind
	}
	return v1.Span_SPAN_KIND_UNSPECIFIED
}

// GetOperationsResponse represents the response for GetOperationsRequest.
type GetOperationsResponse struct {
	Operations           []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetOperationsResponse) Reset()         { *m = GetOperationsResponse{} }
func (m *GetOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetOperationsResponse) ProtoMessage()    {}
func (*GetOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{6}
}
func (m *GetOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationsResponse.Unmarshal(m, b)
}
func (m *GetOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationsResponse.Marshal(b, m, deterministic)
}
func (m *GetOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationsResponse.Merge(m, src)
}
func (m *GetOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetOperationsResponse.Size(m)
}
func (m *GetOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationsResponse proto.InternalMessageInfo

func (m *GetOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// TraceQueryParameters contains query paramters to find traces. For a detailed
// definition of each field in this message, refer to `TraceQueryParameters` in `jaeger.api_v3`
// (https://github.com/jaegertracing/jaeger-idl/blob/main/proto/api_v3/query_service.proto).
type TraceQueryParameters struct {
	ServiceName          string          `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	OperationName        string          `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	Attributes           []*v11.KeyValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	StartTimeMin         time.Time       `protobuf:"bytes,4,opt,name=start_time_min,json=startTimeMin,proto3,stdtime" json:"start_time_min"`
	StartTimeMax         time.Time       `protobuf:"bytes,5,opt,name=start_time_max,json=startTimeMax,proto3,stdtime" json:"start_time_max"`
	DurationMin          time.Duration   `protobuf:"bytes,6,opt,name=duration_min,json=durationMin,proto3,stdduration" json:"duration_min"`
	DurationMax          time.Duration   `protobuf:"bytes,7,opt,name=duration_max,json=durationMax,proto3,stdduration" json:"duration_max"`
	SearchDepth          int32           `protobuf:"varint,8,opt,name=search_depth,json=searchDepth,proto3" json:"search_depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TraceQueryParameters) Reset()         { *m = TraceQueryParameters{} }
func (m *TraceQueryParameters) String() string { return proto.CompactTextString(m) }
func (*TraceQueryParameters) ProtoMessage()    {}
func (*TraceQueryParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{7}
}
func (m *TraceQueryParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceQueryParameters.Unmarshal(m, b)
}
func (m *TraceQueryParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceQueryParameters.Marshal(b, m, deterministic)
}
func (m *TraceQueryParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceQueryParameters.Merge(m, src)
}
func (m *TraceQueryParameters) XXX_Size() int {
	return xxx_messageInfo_TraceQueryParameters.Size(m)
}
func (m *TraceQueryParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceQueryParameters.DiscardUnknown(m)
}

var xxx_messageInfo_TraceQueryParameters proto.InternalMessageInfo

func (m *TraceQueryParameters) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *TraceQueryParameters) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *TraceQueryParameters) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *TraceQueryParameters) GetStartTimeMin() time.Time {
	if m != nil {
		return m.StartTimeMin
	}
	return time.Time{}
}

func (m *TraceQueryParameters) GetStartTimeMax() time.Time {
	if m != nil {
		return m.StartTimeMax
	}
	return time.Time{}
}

func (m *TraceQueryParameters) GetDurationMin() time.Duration {
	if m != nil {
		return m.DurationMin
	}
	return 0
}

func (m *TraceQueryParameters) GetDurationMax() time.Duration {
	if m != nil {
		return m.DurationMax
	}
	return 0
}

func (m *TraceQueryParameters) GetSearchDepth() int32 {
	if m != nil {
		return m.SearchDepth
	}
	return 0
}

// FindTracesRequest represents a request to find traces.
// It can be used to retrieve the traces (FindTraces) or simply
// the trace IDs (FindTraceIDs).
type FindTracesRequest struct {
	Query                *TraceQueryParameters `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FindTracesRequest) Reset()         { *m = FindTracesRequest{} }
func (m *FindTracesRequest) String() string { return proto.CompactTextString(m) }
func (*FindTracesRequest) ProtoMessage()    {}
func (*FindTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{8}
}
func (m *FindTracesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTracesRequest.Unmarshal(m, b)
}
func (m *FindTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTracesRequest.Marshal(b, m, deterministic)
}
func (m *FindTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTracesRequest.Merge(m, src)
}
func (m *FindTracesRequest) XXX_Size() int {
	return xxx_messageInfo_FindTracesRequest.Size(m)
}
func (m *FindTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindTracesRequest proto.InternalMessageInfo

func (m *FindTracesRequest) GetQuery() *TraceQueryParameters {
	if m != nil {
		return m.Query
	}
	return nil
}

// FindTraceIDsResponse represents the response for FindTracesRequest.
type FindTraceIDsResponse struct {
	TraceIds             [][]byte `protobuf:"bytes,1,rep,name=trace_ids,json=traceIds,proto3" json:"trace_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTraceIDsResponse) Reset()         { *m = FindTraceIDsResponse{} }
func (m *FindTraceIDsResponse) String() string { return proto.CompactTextString(m) }
func (*FindTraceIDsResponse) ProtoMessage()    {}
func (*FindTraceIDsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441c0fd9397413c, []int{9}
}
func (m *FindTraceIDsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTraceIDsResponse.Unmarshal(m, b)
}
func (m *FindTraceIDsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTraceIDsResponse.Marshal(b, m, deterministic)
}
func (m *FindTraceIDsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTraceIDsResponse.Merge(m, src)
}
func (m *FindTraceIDsResponse) XXX_Size() int {
	return xxx_messageInfo_FindTraceIDsResponse.Size(m)
}
func (m *FindTraceIDsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTraceIDsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindTraceIDsResponse proto.InternalMessageInfo

func (m *FindTraceIDsResponse) GetTraceIds() [][]byte {
	if m != nil {
		return m.TraceIds
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTraceParams)(nil), "jaeger.storage.v2.GetTraceParams")
	proto.RegisterType((*GetTracesRequest)(nil), "jaeger.storage.v2.GetTracesRequest")
	proto.RegisterType((*GetServicesRequest)(nil), "jaeger.storage.v2.GetServicesRequest")
	proto.RegisterType((*GetServicesResponse)(nil), "jaeger.storage.v2.GetServicesResponse")
	proto.RegisterType((*GetOperationsRequest)(nil), "jaeger.storage.v2.GetOperationsRequest")
	proto.RegisterType((*Operation)(nil), "jaeger.storage.v2.Operation")
	proto.RegisterType((*GetOperationsResponse)(nil), "jaeger.storage.v2.GetOperationsResponse")
	proto.RegisterType((*TraceQueryParameters)(nil), "jaeger.storage.v2.TraceQueryParameters")
	proto.RegisterType((*FindTracesRequest)(nil), "jaeger.storage.v2.FindTracesRequest")
	proto.RegisterType((*FindTraceIDsResponse)(nil), "jaeger.storage.v2.FindTraceIDsResponse")
}

func init() { proto.RegisterFile("trace_storage.proto", fileDescriptor_3441c0fd9397413c) }

var fileDescriptor_3441c0fd9397413c = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x8e, 0x9b, 0xb6, 0x89, 0x27, 0x69, 0x74, 0xba, 0xcd, 0x91, 0xdc, 0x9c, 0x23, 0x9a, 0x1a,
	0x4a, 0x23, 0x90, 0x1c, 0x92, 0x5e, 0x70, 0x01, 0x08, 0xa9, 0x8a, 0x1a, 0x4a, 0xc5, 0x9f, 0x5b,
	0xb8, 0xe0, 0xa2, 0x61, 0x5b, 0x0f, 0xa9, 0xa1, 0xfe, 0xa9, 0x77, 0x13, 0xa5, 0x3c, 0x05, 0x97,
	0xbc, 0x02, 0x6f, 0xc2, 0x53, 0xc0, 0x4b, 0x20, 0xae, 0x91, 0x77, 0xd7, 0xce, 0xaf, 0xda, 0x52,
	0x71, 0x63, 0xed, 0x8e, 0xbf, 0x99, 0xf9, 0x76, 0xbe, 0x99, 0x81, 0x15, 0x1e, 0xd1, 0x63, 0xec,
	0x30, 0x1e, 0x44, 0xb4, 0x8b, 0x56, 0x18, 0x05, 0x3c, 0x20, 0xcb, 0x1f, 0x28, 0x76, 0x31, 0xb2,
	0x12, 0x6b, 0xbf, 0x59, 0x29, 0x77, 0x83, 0x6e, 0x20, 0xfe, 0xd6, 0xe3, 0x93, 0x04, 0x56, 0x6e,
	0x74, 0x83, 0xa0, 0x7b, 0x8a, 0x75, 0x71, 0x3b, 0xea, 0xbd, 0xaf, 0x3b, 0xbd, 0x88, 0x72, 0x37,
	0xf0, 0xd5, 0xff, 0xb5, 0xc9, 0xff, 0xdc, 0xf5, 0x90, 0x71, 0xea, 0x85, 0x0a, 0x70, 0x27, 0x08,
	0xd1, 0xe7, 0x78, 0x8a, 0x1e, 0xf2, 0xe8, 0x5c, 0xe2, 0xea, 0xc7, 0x81, 0xe7, 0x05, 0x7e, 0xbd,
	0xdf, 0x50, 0x27, 0x85, 0xad, 0xcd, 0xc2, 0x0a, 0xfa, 0x31, 0x54, 0x1c, 0x24, 0xd2, 0xfc, 0xaa,
	0x41, 0xa9, 0x8d, 0xfc, 0x20, 0x36, 0xbd, 0xa4, 0x11, 0xf5, 0x18, 0x59, 0x85, 0xbc, 0x7c, 0xa9,
	0xeb, 0x18, 0x5a, 0x55, 0xab, 0x15, 0xed, 0x9c, 0xb8, 0xef, 0x3a, 0xe4, 0x31, 0x00, 0xe3, 0x34,
	0xe2, 0x9d, 0x98, 0x9c, 0x31, 0x57, 0xd5, 0x6a, 0x85, 0x66, 0xc5, 0x92, 0xcc, 0xad, 0x84, 0xb9,
	0x75, 0x90, 0x30, 0xdf, 0x9e, 0xff, 0xfc, 0x63, 0x4d, 0xb3, 0x75, 0xe1, 0x13, 0x5b, 0xc9, 0x03,
	0xc8, 0xa3, 0xef, 0x48, 0xf7, 0xec, 0x15, 0xdd, 0x73, 0xe8, 0x3b, 0xb1, 0xcd, 0xdc, 0x83, 0x7f,
	0x12, 0xaa, 0xcc, 0xc6, 0xb3, 0x1e, 0x32, 0x4e, 0xee, 0xc3, 0xc2, 0x59, 0x0f, 0xa3, 0x73, 0x43,
	0xab, 0x66, 0x6b, 0x85, 0xe6, 0xba, 0x35, 0xa5, 0x87, 0x35, 0xfe, 0x3c, 0x5b, 0xe2, 0xcd, 0x32,
	0x90, 0x36, 0xf2, 0x7d, 0x8c, 0xfa, 0xee, 0x30, 0x9c, 0xd9, 0x80, 0x95, 0x31, 0x2b, 0x0b, 0x03,
	0x9f, 0x21, 0xa9, 0x40, 0x9e, 0x29, 0x9b, 0x48, 0xa4, 0xdb, 0xe9, 0xdd, 0xfc, 0x04, 0xe5, 0x36,
	0xf2, 0x17, 0x21, 0x4a, 0x39, 0x53, 0x66, 0x06, 0xe4, 0x14, 0x46, 0x54, 0x51, 0xb7, 0x93, 0x2b,
	0x79, 0x02, 0x3a, 0x0b, 0xa9, 0xdf, 0xf9, 0xe8, 0xfa, 0x8e, 0x28, 0x62, 0xa9, 0x79, 0xd7, 0x1a,
	0x53, 0x4c, 0x16, 0xc3, 0x92, 0x42, 0xf5, 0x1b, 0xd6, 0x7e, 0x48, 0x7d, 0xf1, 0xd9, 0x73, 0x7d,
	0xc7, 0xce, 0x33, 0x75, 0x32, 0x5d, 0xd0, 0xd3, 0xc4, 0x84, 0xc0, 0xbc, 0x4f, 0xbd, 0x24, 0x9b,
	0x38, 0xff, 0xc5, 0x54, 0xaf, 0xe1, 0xdf, 0x89, 0x67, 0xaa, 0xda, 0x3c, 0x04, 0x08, 0x52, 0xab,
	0x92, 0xe1, 0xff, 0x19, 0x32, 0xa4, 0xae, 0xf6, 0x08, 0xde, 0xfc, 0x99, 0x85, 0xb2, 0x50, 0xe7,
	0x55, 0xac, 0x8a, 0x90, 0x08, 0x39, 0x46, 0x8c, 0xac, 0x43, 0x51, 0xd5, 0xab, 0x33, 0xf2, 0xaa,
	0x82, 0xb2, 0x3d, 0x8f, 0x1f, 0xb7, 0x01, 0xa5, 0x34, 0x92, 0x04, 0xcd, 0x09, 0xd0, 0x52, 0x6a,
	0x15, 0xb0, 0x36, 0x00, 0xe5, 0x3c, 0x72, 0x8f, 0x7a, 0x1c, 0x99, 0x91, 0x15, 0x04, 0x37, 0x67,
	0x16, 0x41, 0xcd, 0x50, 0xbf, 0x61, 0xed, 0xe1, 0xf9, 0x1b, 0x7a, 0xda, 0x43, 0x7b, 0xc4, 0x95,
	0x3c, 0x85, 0xd2, 0xb0, 0xfb, 0x3b, 0x9e, 0xeb, 0x1b, 0xf3, 0x97, 0xb6, 0x70, 0xfe, 0xdb, 0xf7,
	0xb5, 0x8c, 0x68, 0xe3, 0x62, 0x3a, 0x05, 0xcf, 0x5c, 0x7f, 0x32, 0x16, 0x1d, 0x18, 0x0b, 0xd7,
	0x8b, 0x45, 0x07, 0x64, 0x07, 0x8a, 0xc9, 0x32, 0x11, 0xac, 0x16, 0x45, 0xa4, 0xd5, 0xa9, 0x48,
	0x2d, 0x05, 0x92, 0x81, 0xbe, 0xc4, 0x81, 0x0a, 0x89, 0x63, 0xcc, 0x69, 0x2c, 0x0e, 0x1d, 0x18,
	0xb9, 0xeb, 0xc4, 0xa1, 0x03, 0x29, 0x1d, 0x8d, 0x8e, 0x4f, 0x3a, 0x0e, 0x86, 0xfc, 0xc4, 0xc8,
	0x57, 0xb5, 0xda, 0x42, 0x2c, 0x5d, 0x6c, 0x6b, 0xc5, 0x26, 0xd3, 0x86, 0xe5, 0x1d, 0xd7, 0x77,
	0xc6, 0x67, 0xf9, 0xd1, 0x70, 0x96, 0x35, 0xa1, 0xd1, 0x74, 0x13, 0xcd, 0x6a, 0x95, 0x64, 0xa2,
	0xb7, 0xa0, 0x9c, 0xc6, 0xdc, 0x6d, 0x0d, 0x1b, 0xf4, 0x3f, 0xd0, 0x93, 0x7d, 0x26, 0xfb, 0xb3,
	0x68, 0xe7, 0xd5, 0x42, 0x63, 0xcd, 0x5f, 0x59, 0x28, 0x08, 0x0f, 0x1b, 0xa9, 0x83, 0x11, 0x39,
	0x04, 0x3d, 0xdd, 0x31, 0xe4, 0xe6, 0x05, 0xdb, 0x24, 0x61, 0x5d, 0xa9, 0x5d, 0x3c, 0x4f, 0x12,
	0xdc, 0xa2, 0x9c, 0x9a, 0x99, 0x7b, 0x1a, 0x39, 0x84, 0xc2, 0xc8, 0x82, 0x21, 0x1b, 0xb3, 0x33,
	0x4c, 0xac, 0xa5, 0xca, 0xed, 0xcb, 0x60, 0xf2, 0xa9, 0x66, 0x86, 0x38, 0xb0, 0x34, 0x36, 0xa6,
	0x64, 0x73, 0xb6, 0xeb, 0xd4, 0xbe, 0xaa, 0xd4, 0x2e, 0x07, 0xa6, 0x59, 0xde, 0x01, 0x0c, 0xe5,
	0x23, 0xb7, 0x66, 0x78, 0x4e, 0xa9, 0xfb, 0x87, 0x75, 0xea, 0x40, 0x71, 0x54, 0xcc, 0x2b, 0xe6,
	0xd8, 0xbc, 0x08, 0x35, 0xd2, 0x13, 0x66, 0x66, 0x5b, 0x7f, 0x9b, 0x53, 0xa0, 0xa3, 0x45, 0x41,
	0x65, 0xeb, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xf5, 0x16, 0xea, 0xe1, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceReaderClient is the client API for TraceReader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceReaderClient interface {
	// GetTraces returns a stream that retrieves all traces with given IDs.
	//
	// Chunking requirements:
	// - A single TracesData chunk MUST NOT contain spans from multiple traces.
	// - Large traces MAY be split across multiple, *consecutive* TracesData chunks.
	// - Each returned TracesData object MUST NOT be empty.
	//
	// Edge cases:
	// - If no spans are found for any given trace ID, the ID is ignored.
	// - If none of the trace IDs are found in the storage, an empty response is returned.
	GetTraces(ctx context.Context, in *GetTracesRequest, opts ...grpc.CallOption) (TraceReader_GetTracesClient, error)
	// GetServices returns all service names known to the backend from traces
	// within its retention period.
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	// GetOperations returns all operation names for a given service
	// known to the backend from traces within its retention period.
	GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error)
	// FindTraces returns a stream that retrieves traces matching query parameters.
	//
	// The chunking rules are the same as for GetTraces.
	//
	// If no matching traces are found, an empty stream is returned.
	FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (TraceReader_FindTracesClient, error)
	// FindTraceIDs returns a stream that retrieves IDs of traces matching query parameters.
	//
	// If no matching traces are found, an empty stream is returned.
	//
	// This call behaves identically to FindTraces, except that it returns only the list
	// of matching trace IDs. This is useful in some contexts, such as batch jobs, where a
	// large list of trace IDs may be queried first and then the full traces are loaded
	// in batches.
	FindTraceIDs(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (*FindTraceIDsResponse, error)
}

type traceReaderClient struct {
	cc *grpc.ClientConn
}

func NewTraceReaderClient(cc *grpc.ClientConn) TraceReaderClient {
	return &traceReaderClient{cc}
}

func (c *traceReaderClient) GetTraces(ctx context.Context, in *GetTracesRequest, opts ...grpc.CallOption) (TraceReader_GetTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceReader_serviceDesc.Streams[0], "/jaeger.storage.v2.TraceReader/GetTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceReaderGetTracesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TraceReader_GetTracesClient interface {
	Recv() (*v1.TracesData, error)
	grpc.ClientStream
}

type traceReaderGetTracesClient struct {
	grpc.ClientStream
}

func (x *traceReaderGetTracesClient) Recv() (*v1.TracesData, error) {
	m := new(v1.TracesData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *traceReaderClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v2.TraceReader/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error) {
	out := new(GetOperationsResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v2.TraceReader/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (TraceReader_FindTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceReader_serviceDesc.Streams[1], "/jaeger.storage.v2.TraceReader/FindTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceReaderFindTracesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TraceReader_FindTracesClient interface {
	Recv() (*v1.TracesData, error)
	grpc.ClientStream
}

type traceReaderFindTracesClient struct {
	grpc.ClientStream
}

func (x *traceReaderFindTracesClient) Recv() (*v1.TracesData, error) {
	m := new(v1.TracesData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *traceReaderClient) FindTraceIDs(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (*FindTraceIDsResponse, error) {
	out := new(FindTraceIDsResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v2.TraceReader/FindTraceIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceReaderServer is the server API for TraceReader service.
type TraceReaderServer interface {
	// GetTraces returns a stream that retrieves all traces with given IDs.
	//
	// Chunking requirements:
	// - A single TracesData chunk MUST NOT contain spans from multiple traces.
	// - Large traces MAY be split across multiple, *consecutive* TracesData chunks.
	// - Each returned TracesData object MUST NOT be empty.
	//
	// Edge cases:
	// - If no spans are found for any given trace ID, the ID is ignored.
	// - If none of the trace IDs are found in the storage, an empty response is returned.
	GetTraces(*GetTracesRequest, TraceReader_GetTracesServer) error
	// GetServices returns all service names known to the backend from traces
	// within its retention period.
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	// GetOperations returns all operation names for a given service
	// known to the backend from traces within its retention period.
	GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error)
	// FindTraces returns a stream that retrieves traces matching query parameters.
	//
	// The chunking rules are the same as for GetTraces.
	//
	// If no matching traces are found, an empty stream is returned.
	FindTraces(*FindTracesRequest, TraceReader_FindTracesServer) error
	// FindTraceIDs returns a stream that retrieves IDs of traces matching query parameters.
	//
	// If no matching traces are found, an empty stream is returned.
	//
	// This call behaves identically to FindTraces, except that it returns only the list
	// of matching trace IDs. This is useful in some contexts, such as batch jobs, where a
	// large list of trace IDs may be queried first and then the full traces are loaded
	// in batches.
	FindTraceIDs(context.Context, *FindTracesRequest) (*FindTraceIDsResponse, error)
}

// UnimplementedTraceReaderServer can be embedded to have forward compatible implementations.
type UnimplementedTraceReaderServer struct {
}

func (*UnimplementedTraceReaderServer) GetTraces(req *GetTracesRequest, srv TraceReader_GetTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTraces not implemented")
}
func (*UnimplementedTraceReaderServer) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (*UnimplementedTraceReaderServer) GetOperations(ctx context.Context, req *GetOperationsRequest) (*GetOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (*UnimplementedTraceReaderServer) FindTraces(req *FindTracesRequest, srv TraceReader_FindTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method FindTraces not implemented")
}
func (*UnimplementedTraceReaderServer) FindTraceIDs(ctx context.Context, req *FindTracesRequest) (*FindTraceIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTraceIDs not implemented")
}

func RegisterTraceReaderServer(s *grpc.Server, srv TraceReaderServer) {
	s.RegisterService(&_TraceReader_serviceDesc, srv)
}

func _TraceReader_GetTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTracesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TraceReaderServer).GetTraces(m, &traceReaderGetTracesServer{stream})
}

type TraceReader_GetTracesServer interface {
	Send(*v1.TracesData) error
	grpc.ServerStream
}

type traceReaderGetTracesServer struct {
	grpc.ServerStream
}

func (x *traceReaderGetTracesServer) Send(m *v1.TracesData) error {
	return x.ServerStream.SendMsg(m)
}

func _TraceReader_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v2.TraceReader/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v2.TraceReader/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetOperations(ctx, req.(*GetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_FindTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindTracesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TraceReaderServer).FindTraces(m, &traceReaderFindTracesServer{stream})
}

type TraceReader_FindTracesServer interface {
	Send(*v1.TracesData) error
	grpc.ServerStream
}

type traceReaderFindTracesServer struct {
	grpc.ServerStream
}

func (x *traceReaderFindTracesServer) Send(m *v1.TracesData) error {
	return x.ServerStream.SendMsg(m)
}

func _TraceReader_FindTraceIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).FindTraceIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v2.TraceReader/FindTraceIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).FindTraceIDs(ctx, req.(*FindTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceReader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v2.TraceReader",
	HandlerType: (*TraceReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _TraceReader_GetServices_Handler,
		},
		{
			MethodName: "GetOperations",
			Handler:    _TraceReader_GetOperations_Handler,
		},
		{
			MethodName: "FindTraceIDs",
			Handler:    _TraceReader_FindTraceIDs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTraces",
			Handler:       _TraceReader_GetTraces_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindTraces",
			Handler:       _TraceReader_FindTraces_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "trace_storage.proto",
}
