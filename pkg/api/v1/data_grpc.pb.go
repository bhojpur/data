// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	// StartLocalEngine starts a Data Processing Engine on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the data/config.yaml
	//   3. all bytes constituting the Engine YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalEngine(ctx context.Context, opts ...grpc.CallOption) (DataService_StartLocalEngineClient, error)
	// StartFromPreviousEngine starts a new Engine based on a previous one.
	// If the previous Engine does not have the can-replay condition set this call will result in an error.
	StartFromPreviousEngine(ctx context.Context, in *StartFromPreviousEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error)
	// StartEngineRequest starts a new Engine based on its specification.
	StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error)
	// Searches for Engine(s) known to this Engine
	ListEngines(ctx context.Context, in *ListEnginesRequest, opts ...grpc.CallOption) (*ListEnginesResponse, error)
	// Subscribe listens to new Engine(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DataService_SubscribeClient, error)
	// GetEngine retrieves details of a single Engine
	GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineResponse, error)
	// Listen listens to Engine updates and log output of a running Engine
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (DataService_ListenClient, error)
	// StopEngine stops a currently running Engine
	StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*StopEngineResponse, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) StartLocalEngine(ctx context.Context, opts ...grpc.CallOption) (DataService_StartLocalEngineClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], "/v1.DataService/StartLocalEngine", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceStartLocalEngineClient{stream}
	return x, nil
}

type DataService_StartLocalEngineClient interface {
	Send(*StartLocalEngineRequest) error
	CloseAndRecv() (*StartEngineResponse, error)
	grpc.ClientStream
}

type dataServiceStartLocalEngineClient struct {
	grpc.ClientStream
}

func (x *dataServiceStartLocalEngineClient) Send(m *StartLocalEngineRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServiceStartLocalEngineClient) CloseAndRecv() (*StartEngineResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartEngineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) StartFromPreviousEngine(ctx context.Context, in *StartFromPreviousEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error) {
	out := new(StartEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.DataService/StartFromPreviousEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error) {
	out := new(StartEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.DataService/StartEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) ListEngines(ctx context.Context, in *ListEnginesRequest, opts ...grpc.CallOption) (*ListEnginesResponse, error) {
	out := new(ListEnginesResponse)
	err := c.cc.Invoke(ctx, "/v1.DataService/ListEngines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DataService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[1], "/v1.DataService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type dataServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *dataServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineResponse, error) {
	out := new(GetEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.DataService/GetEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (DataService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[2], "/v1.DataService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type dataServiceListenClient struct {
	grpc.ClientStream
}

func (x *dataServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*StopEngineResponse, error) {
	out := new(StopEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.DataService/StopEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	// StartLocalEngine starts a Data Processing Engine on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the data/config.yaml
	//   3. all bytes constituting the Engine YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalEngine(DataService_StartLocalEngineServer) error
	// StartFromPreviousEngine starts a new Engine based on a previous one.
	// If the previous Engine does not have the can-replay condition set this call will result in an error.
	StartFromPreviousEngine(context.Context, *StartFromPreviousEngineRequest) (*StartEngineResponse, error)
	// StartEngineRequest starts a new Engine based on its specification.
	StartEngine(context.Context, *StartEngineRequest) (*StartEngineResponse, error)
	// Searches for Engine(s) known to this Engine
	ListEngines(context.Context, *ListEnginesRequest) (*ListEnginesResponse, error)
	// Subscribe listens to new Engine(s) updates
	Subscribe(*SubscribeRequest, DataService_SubscribeServer) error
	// GetEngine retrieves details of a single Engine
	GetEngine(context.Context, *GetEngineRequest) (*GetEngineResponse, error)
	// Listen listens to Engine updates and log output of a running Engine
	Listen(*ListenRequest, DataService_ListenServer) error
	// StopEngine stops a currently running Engine
	StopEngine(context.Context, *StopEngineRequest) (*StopEngineResponse, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) StartLocalEngine(DataService_StartLocalEngineServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalEngine not implemented")
}
func (UnimplementedDataServiceServer) StartFromPreviousEngine(context.Context, *StartFromPreviousEngineRequest) (*StartEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousEngine not implemented")
}
func (UnimplementedDataServiceServer) StartEngine(context.Context, *StartEngineRequest) (*StartEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartEngine not implemented")
}
func (UnimplementedDataServiceServer) ListEngines(context.Context, *ListEnginesRequest) (*ListEnginesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEngines not implemented")
}
func (UnimplementedDataServiceServer) Subscribe(*SubscribeRequest, DataService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedDataServiceServer) GetEngine(context.Context, *GetEngineRequest) (*GetEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEngine not implemented")
}
func (UnimplementedDataServiceServer) Listen(*ListenRequest, DataService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedDataServiceServer) StopEngine(context.Context, *StopEngineRequest) (*StopEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopEngine not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_StartLocalEngine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServiceServer).StartLocalEngine(&dataServiceStartLocalEngineServer{stream})
}

type DataService_StartLocalEngineServer interface {
	SendAndClose(*StartEngineResponse) error
	Recv() (*StartLocalEngineRequest, error)
	grpc.ServerStream
}

type dataServiceStartLocalEngineServer struct {
	grpc.ServerStream
}

func (x *dataServiceStartLocalEngineServer) SendAndClose(m *StartEngineResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServiceStartLocalEngineServer) Recv() (*StartLocalEngineRequest, error) {
	m := new(StartLocalEngineRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataService_StartFromPreviousEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).StartFromPreviousEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataService/StartFromPreviousEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).StartFromPreviousEngine(ctx, req.(*StartFromPreviousEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_StartEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).StartEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataService/StartEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).StartEngine(ctx, req.(*StartEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_ListEngines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnginesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).ListEngines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataService/ListEngines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).ListEngines(ctx, req.(*ListEnginesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).Subscribe(m, &dataServiceSubscribeServer{stream})
}

type DataService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type dataServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *dataServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataService_GetEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataService/GetEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetEngine(ctx, req.(*GetEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).Listen(m, &dataServiceListenServer{stream})
}

type DataService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type dataServiceListenServer struct {
	grpc.ServerStream
}

func (x *dataServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataService_StopEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).StopEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataService/StopEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).StopEngine(ctx, req.(*StopEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousEngine",
			Handler:    _DataService_StartFromPreviousEngine_Handler,
		},
		{
			MethodName: "StartEngine",
			Handler:    _DataService_StartEngine_Handler,
		},
		{
			MethodName: "ListEngines",
			Handler:    _DataService_ListEngines_Handler,
		},
		{
			MethodName: "GetEngine",
			Handler:    _DataService_GetEngine_Handler,
		},
		{
			MethodName: "StopEngine",
			Handler:    _DataService_StopEngine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalEngine",
			Handler:       _DataService_StartLocalEngine_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _DataService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _DataService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data.proto",
}