// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tarorpc

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

// TaroClient is the client API for Taro service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaroClient interface {
	// tarocli: `assets mint`
	//MintAsset will attempts to mint the set of assets (async by default to
	//ensure proper batching) specified in the request.
	MintAsset(ctx context.Context, in *MintAssetRequest, opts ...grpc.CallOption) (*MintAssetResponse, error)
	// tarocli: `assets list`
	//ListAssets lists the set of assets owned by the target daemon.
	ListAssets(ctx context.Context, in *ListAssetRequest, opts ...grpc.CallOption) (*ListAssetResponse, error)
	// tarocli: `assets utxos`
	//ListUtxos lists the UTXOs managed by the target daemon, and the assets they
	//hold.
	ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error)
	// tarocli: `assets groups`
	//ListGroups lists the asset groups known to the target daemon, and the assets
	//held in each group.
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// tarocli: `assets balance`
	//ListBalances lists asset balances
	ListBalances(ctx context.Context, in *ListBalancesRequest, opts ...grpc.CallOption) (*ListBalancesResponse, error)
	// tarocli: `assets transfers`
	//ListTransfers lists outbound asset transfers tracked by the target daemon.
	ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error)
	// tarocli: `stop`
	//StopDaemon will send a shutdown request to the interrupt handler, triggering
	//a graceful shutdown of the daemon.
	StopDaemon(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	// tarocli: `debuglevel`
	//DebugLevel allows a caller to programmatically set the logging verbosity of
	//tarod. The logging can be targeted according to a coarse daemon-wide logging
	//level, or in a granular fashion to specify the logging for a target
	//sub-system.
	DebugLevel(ctx context.Context, in *DebugLevelRequest, opts ...grpc.CallOption) (*DebugLevelResponse, error)
	// tarocli: `addrs query`
	//QueryTaroAddrs queries the set of Taro addresses stored in the database.
	QueryAddrs(ctx context.Context, in *QueryAddrRequest, opts ...grpc.CallOption) (*QueryAddrResponse, error)
	// tarocli: `addrs new`
	//NewAddr makes a new address from the set of request params.
	NewAddr(ctx context.Context, in *NewAddrRequest, opts ...grpc.CallOption) (*Addr, error)
	// tarocli: `addrs decode`
	//DecodeAddr decode a Taro address into a partial asset message that
	//represents the asset it wants to receive.
	DecodeAddr(ctx context.Context, in *DecodeAddrRequest, opts ...grpc.CallOption) (*Addr, error)
	// tarocli: `addrs receives`
	//List all receives for incoming asset transfers for addresses that were
	//created previously.
	AddrReceives(ctx context.Context, in *AddrReceivesRequest, opts ...grpc.CallOption) (*AddrReceivesResponse, error)
	// tarocli: `proofs verify`
	//VerifyProof attempts to verify a given proof file that claims to be anchored
	//at the specified genesis point.
	VerifyProof(ctx context.Context, in *ProofFile, opts ...grpc.CallOption) (*ProofVerifyResponse, error)
	// tarocli: `proofs export`
	//ExportProof exports the latest raw proof file anchored at the specified
	//script_key.
	ExportProof(ctx context.Context, in *ExportProofRequest, opts ...grpc.CallOption) (*ProofFile, error)
	// tarocli: `proofs import`
	//ImportProof attempts to import a proof file into the daemon. If successful,
	//a new asset will be inserted on disk, spendable using the specified target
	//script key, and internal key.
	ImportProof(ctx context.Context, in *ImportProofRequest, opts ...grpc.CallOption) (*ImportProofResponse, error)
	// tarocli: `assets send`
	//SendAsset uses a passed taro address to attempt to complete an asset send.
	//The method returns information w.r.t the on chain send, as well as the
	//proof file information the receiver needs to fully receive the asset.
	SendAsset(ctx context.Context, in *SendAssetRequest, opts ...grpc.CallOption) (*SendAssetResponse, error)
	//
	//SubscribeSendAssetEventNtfns registers a subscription to the event
	//notification stream which relates to the asset sending process.
	SubscribeSendAssetEventNtfns(ctx context.Context, in *SubscribeSendAssetEventNtfnsRequest, opts ...grpc.CallOption) (Taro_SubscribeSendAssetEventNtfnsClient, error)
	//
	//FetchAssetMeta allows a caller to fetch the reveal meta data for an asset
	//either by the asset ID for that asset, or a meta hash.
	FetchAssetMeta(ctx context.Context, in *FetchAssetMetaRequest, opts ...grpc.CallOption) (*AssetMeta, error)
}

type taroClient struct {
	cc grpc.ClientConnInterface
}

func NewTaroClient(cc grpc.ClientConnInterface) TaroClient {
	return &taroClient{cc}
}

func (c *taroClient) MintAsset(ctx context.Context, in *MintAssetRequest, opts ...grpc.CallOption) (*MintAssetResponse, error) {
	out := new(MintAssetResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/MintAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ListAssets(ctx context.Context, in *ListAssetRequest, opts ...grpc.CallOption) (*ListAssetResponse, error) {
	out := new(ListAssetResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ListAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error) {
	out := new(ListUtxosResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ListUtxos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ListBalances(ctx context.Context, in *ListBalancesRequest, opts ...grpc.CallOption) (*ListBalancesResponse, error) {
	out := new(ListBalancesResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ListBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error) {
	out := new(ListTransfersResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ListTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) StopDaemon(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/StopDaemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) DebugLevel(ctx context.Context, in *DebugLevelRequest, opts ...grpc.CallOption) (*DebugLevelResponse, error) {
	out := new(DebugLevelResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/DebugLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) QueryAddrs(ctx context.Context, in *QueryAddrRequest, opts ...grpc.CallOption) (*QueryAddrResponse, error) {
	out := new(QueryAddrResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/QueryAddrs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) NewAddr(ctx context.Context, in *NewAddrRequest, opts ...grpc.CallOption) (*Addr, error) {
	out := new(Addr)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/NewAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) DecodeAddr(ctx context.Context, in *DecodeAddrRequest, opts ...grpc.CallOption) (*Addr, error) {
	out := new(Addr)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/DecodeAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) AddrReceives(ctx context.Context, in *AddrReceivesRequest, opts ...grpc.CallOption) (*AddrReceivesResponse, error) {
	out := new(AddrReceivesResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/AddrReceives", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) VerifyProof(ctx context.Context, in *ProofFile, opts ...grpc.CallOption) (*ProofVerifyResponse, error) {
	out := new(ProofVerifyResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/VerifyProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ExportProof(ctx context.Context, in *ExportProofRequest, opts ...grpc.CallOption) (*ProofFile, error) {
	out := new(ProofFile)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ExportProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) ImportProof(ctx context.Context, in *ImportProofRequest, opts ...grpc.CallOption) (*ImportProofResponse, error) {
	out := new(ImportProofResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/ImportProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) SendAsset(ctx context.Context, in *SendAssetRequest, opts ...grpc.CallOption) (*SendAssetResponse, error) {
	out := new(SendAssetResponse)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/SendAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taroClient) SubscribeSendAssetEventNtfns(ctx context.Context, in *SubscribeSendAssetEventNtfnsRequest, opts ...grpc.CallOption) (Taro_SubscribeSendAssetEventNtfnsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Taro_ServiceDesc.Streams[0], "/tarorpc.Taro/SubscribeSendAssetEventNtfns", opts...)
	if err != nil {
		return nil, err
	}
	x := &taroSubscribeSendAssetEventNtfnsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Taro_SubscribeSendAssetEventNtfnsClient interface {
	Recv() (*SendAssetEvent, error)
	grpc.ClientStream
}

type taroSubscribeSendAssetEventNtfnsClient struct {
	grpc.ClientStream
}

func (x *taroSubscribeSendAssetEventNtfnsClient) Recv() (*SendAssetEvent, error) {
	m := new(SendAssetEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taroClient) FetchAssetMeta(ctx context.Context, in *FetchAssetMetaRequest, opts ...grpc.CallOption) (*AssetMeta, error) {
	out := new(AssetMeta)
	err := c.cc.Invoke(ctx, "/tarorpc.Taro/FetchAssetMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaroServer is the server API for Taro service.
// All implementations must embed UnimplementedTaroServer
// for forward compatibility
type TaroServer interface {
	// tarocli: `assets mint`
	//MintAsset will attempts to mint the set of assets (async by default to
	//ensure proper batching) specified in the request.
	MintAsset(context.Context, *MintAssetRequest) (*MintAssetResponse, error)
	// tarocli: `assets list`
	//ListAssets lists the set of assets owned by the target daemon.
	ListAssets(context.Context, *ListAssetRequest) (*ListAssetResponse, error)
	// tarocli: `assets utxos`
	//ListUtxos lists the UTXOs managed by the target daemon, and the assets they
	//hold.
	ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error)
	// tarocli: `assets groups`
	//ListGroups lists the asset groups known to the target daemon, and the assets
	//held in each group.
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	// tarocli: `assets balance`
	//ListBalances lists asset balances
	ListBalances(context.Context, *ListBalancesRequest) (*ListBalancesResponse, error)
	// tarocli: `assets transfers`
	//ListTransfers lists outbound asset transfers tracked by the target daemon.
	ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error)
	// tarocli: `stop`
	//StopDaemon will send a shutdown request to the interrupt handler, triggering
	//a graceful shutdown of the daemon.
	StopDaemon(context.Context, *StopRequest) (*StopResponse, error)
	// tarocli: `debuglevel`
	//DebugLevel allows a caller to programmatically set the logging verbosity of
	//tarod. The logging can be targeted according to a coarse daemon-wide logging
	//level, or in a granular fashion to specify the logging for a target
	//sub-system.
	DebugLevel(context.Context, *DebugLevelRequest) (*DebugLevelResponse, error)
	// tarocli: `addrs query`
	//QueryTaroAddrs queries the set of Taro addresses stored in the database.
	QueryAddrs(context.Context, *QueryAddrRequest) (*QueryAddrResponse, error)
	// tarocli: `addrs new`
	//NewAddr makes a new address from the set of request params.
	NewAddr(context.Context, *NewAddrRequest) (*Addr, error)
	// tarocli: `addrs decode`
	//DecodeAddr decode a Taro address into a partial asset message that
	//represents the asset it wants to receive.
	DecodeAddr(context.Context, *DecodeAddrRequest) (*Addr, error)
	// tarocli: `addrs receives`
	//List all receives for incoming asset transfers for addresses that were
	//created previously.
	AddrReceives(context.Context, *AddrReceivesRequest) (*AddrReceivesResponse, error)
	// tarocli: `proofs verify`
	//VerifyProof attempts to verify a given proof file that claims to be anchored
	//at the specified genesis point.
	VerifyProof(context.Context, *ProofFile) (*ProofVerifyResponse, error)
	// tarocli: `proofs export`
	//ExportProof exports the latest raw proof file anchored at the specified
	//script_key.
	ExportProof(context.Context, *ExportProofRequest) (*ProofFile, error)
	// tarocli: `proofs import`
	//ImportProof attempts to import a proof file into the daemon. If successful,
	//a new asset will be inserted on disk, spendable using the specified target
	//script key, and internal key.
	ImportProof(context.Context, *ImportProofRequest) (*ImportProofResponse, error)
	// tarocli: `assets send`
	//SendAsset uses a passed taro address to attempt to complete an asset send.
	//The method returns information w.r.t the on chain send, as well as the
	//proof file information the receiver needs to fully receive the asset.
	SendAsset(context.Context, *SendAssetRequest) (*SendAssetResponse, error)
	//
	//SubscribeSendAssetEventNtfns registers a subscription to the event
	//notification stream which relates to the asset sending process.
	SubscribeSendAssetEventNtfns(*SubscribeSendAssetEventNtfnsRequest, Taro_SubscribeSendAssetEventNtfnsServer) error
	//
	//FetchAssetMeta allows a caller to fetch the reveal meta data for an asset
	//either by the asset ID for that asset, or a meta hash.
	FetchAssetMeta(context.Context, *FetchAssetMetaRequest) (*AssetMeta, error)
	mustEmbedUnimplementedTaroServer()
}

// UnimplementedTaroServer must be embedded to have forward compatible implementations.
type UnimplementedTaroServer struct {
}

func (UnimplementedTaroServer) MintAsset(context.Context, *MintAssetRequest) (*MintAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintAsset not implemented")
}
func (UnimplementedTaroServer) ListAssets(context.Context, *ListAssetRequest) (*ListAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssets not implemented")
}
func (UnimplementedTaroServer) ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUtxos not implemented")
}
func (UnimplementedTaroServer) ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedTaroServer) ListBalances(context.Context, *ListBalancesRequest) (*ListBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalances not implemented")
}
func (UnimplementedTaroServer) ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransfers not implemented")
}
func (UnimplementedTaroServer) StopDaemon(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopDaemon not implemented")
}
func (UnimplementedTaroServer) DebugLevel(context.Context, *DebugLevelRequest) (*DebugLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebugLevel not implemented")
}
func (UnimplementedTaroServer) QueryAddrs(context.Context, *QueryAddrRequest) (*QueryAddrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAddrs not implemented")
}
func (UnimplementedTaroServer) NewAddr(context.Context, *NewAddrRequest) (*Addr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddr not implemented")
}
func (UnimplementedTaroServer) DecodeAddr(context.Context, *DecodeAddrRequest) (*Addr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeAddr not implemented")
}
func (UnimplementedTaroServer) AddrReceives(context.Context, *AddrReceivesRequest) (*AddrReceivesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddrReceives not implemented")
}
func (UnimplementedTaroServer) VerifyProof(context.Context, *ProofFile) (*ProofVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyProof not implemented")
}
func (UnimplementedTaroServer) ExportProof(context.Context, *ExportProofRequest) (*ProofFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportProof not implemented")
}
func (UnimplementedTaroServer) ImportProof(context.Context, *ImportProofRequest) (*ImportProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProof not implemented")
}
func (UnimplementedTaroServer) SendAsset(context.Context, *SendAssetRequest) (*SendAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAsset not implemented")
}
func (UnimplementedTaroServer) SubscribeSendAssetEventNtfns(*SubscribeSendAssetEventNtfnsRequest, Taro_SubscribeSendAssetEventNtfnsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeSendAssetEventNtfns not implemented")
}
func (UnimplementedTaroServer) FetchAssetMeta(context.Context, *FetchAssetMetaRequest) (*AssetMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAssetMeta not implemented")
}
func (UnimplementedTaroServer) mustEmbedUnimplementedTaroServer() {}

// UnsafeTaroServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaroServer will
// result in compilation errors.
type UnsafeTaroServer interface {
	mustEmbedUnimplementedTaroServer()
}

func RegisterTaroServer(s grpc.ServiceRegistrar, srv TaroServer) {
	s.RegisterService(&Taro_ServiceDesc, srv)
}

func _Taro_MintAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).MintAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/MintAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).MintAsset(ctx, req.(*MintAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ListAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ListAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ListAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ListAssets(ctx, req.(*ListAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ListUtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUtxosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ListUtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ListUtxos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ListUtxos(ctx, req.(*ListUtxosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ListBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ListBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ListBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ListBalances(ctx, req.(*ListBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ListTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ListTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ListTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ListTransfers(ctx, req.(*ListTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_StopDaemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).StopDaemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/StopDaemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).StopDaemon(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_DebugLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).DebugLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/DebugLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).DebugLevel(ctx, req.(*DebugLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_QueryAddrs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).QueryAddrs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/QueryAddrs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).QueryAddrs(ctx, req.(*QueryAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_NewAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).NewAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/NewAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).NewAddr(ctx, req.(*NewAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_DecodeAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).DecodeAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/DecodeAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).DecodeAddr(ctx, req.(*DecodeAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_AddrReceives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddrReceivesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).AddrReceives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/AddrReceives",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).AddrReceives(ctx, req.(*AddrReceivesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_VerifyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProofFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).VerifyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/VerifyProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).VerifyProof(ctx, req.(*ProofFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ExportProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ExportProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ExportProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ExportProof(ctx, req.(*ExportProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_ImportProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).ImportProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/ImportProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).ImportProof(ctx, req.(*ImportProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_SendAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).SendAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/SendAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).SendAsset(ctx, req.(*SendAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taro_SubscribeSendAssetEventNtfns_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeSendAssetEventNtfnsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaroServer).SubscribeSendAssetEventNtfns(m, &taroSubscribeSendAssetEventNtfnsServer{stream})
}

type Taro_SubscribeSendAssetEventNtfnsServer interface {
	Send(*SendAssetEvent) error
	grpc.ServerStream
}

type taroSubscribeSendAssetEventNtfnsServer struct {
	grpc.ServerStream
}

func (x *taroSubscribeSendAssetEventNtfnsServer) Send(m *SendAssetEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Taro_FetchAssetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAssetMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaroServer).FetchAssetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarorpc.Taro/FetchAssetMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaroServer).FetchAssetMeta(ctx, req.(*FetchAssetMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Taro_ServiceDesc is the grpc.ServiceDesc for Taro service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Taro_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tarorpc.Taro",
	HandlerType: (*TaroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintAsset",
			Handler:    _Taro_MintAsset_Handler,
		},
		{
			MethodName: "ListAssets",
			Handler:    _Taro_ListAssets_Handler,
		},
		{
			MethodName: "ListUtxos",
			Handler:    _Taro_ListUtxos_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _Taro_ListGroups_Handler,
		},
		{
			MethodName: "ListBalances",
			Handler:    _Taro_ListBalances_Handler,
		},
		{
			MethodName: "ListTransfers",
			Handler:    _Taro_ListTransfers_Handler,
		},
		{
			MethodName: "StopDaemon",
			Handler:    _Taro_StopDaemon_Handler,
		},
		{
			MethodName: "DebugLevel",
			Handler:    _Taro_DebugLevel_Handler,
		},
		{
			MethodName: "QueryAddrs",
			Handler:    _Taro_QueryAddrs_Handler,
		},
		{
			MethodName: "NewAddr",
			Handler:    _Taro_NewAddr_Handler,
		},
		{
			MethodName: "DecodeAddr",
			Handler:    _Taro_DecodeAddr_Handler,
		},
		{
			MethodName: "AddrReceives",
			Handler:    _Taro_AddrReceives_Handler,
		},
		{
			MethodName: "VerifyProof",
			Handler:    _Taro_VerifyProof_Handler,
		},
		{
			MethodName: "ExportProof",
			Handler:    _Taro_ExportProof_Handler,
		},
		{
			MethodName: "ImportProof",
			Handler:    _Taro_ImportProof_Handler,
		},
		{
			MethodName: "SendAsset",
			Handler:    _Taro_SendAsset_Handler,
		},
		{
			MethodName: "FetchAssetMeta",
			Handler:    _Taro_FetchAssetMeta_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeSendAssetEventNtfns",
			Handler:       _Taro_SubscribeSendAssetEventNtfns_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "taro.proto",
}
