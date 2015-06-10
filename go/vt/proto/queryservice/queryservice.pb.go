// Code generated by protoc-gen-go.
// source: queryservice.proto
// DO NOT EDIT!

/*
Package queryservice is a generated protocol buffer package.

It is generated from these files:
	queryservice.proto

It has these top-level messages:
*/
package queryservice

import proto "github.com/golang/protobuf/proto"
import query "github.com/youtube/vitess/go/vt/proto/query"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

func init() {
}

// Client API for SqlQuery service

type SqlQueryClient interface {
	// Execute executes the specified SQL query (might be in a
	// transaction context, if Query.transaction_id is set).
	Execute(ctx context.Context, in *query.ExecuteRequest, opts ...grpc.CallOption) (*query.ExecuteResponse, error)
	// ExecuteBatch executes a list of queries, and returns the result
	// for each query.
	ExecuteBatch(ctx context.Context, in *query.ExecuteBatchRequest, opts ...grpc.CallOption) (*query.ExecuteBatchResponse, error)
	// StreamExecute executes a streaming query. Use this method if the
	// query returns a large number of rows. The first QueryResult will
	// contain the Fields, subsequent QueryResult messages will contain
	// the rows.
	StreamExecute(ctx context.Context, in *query.StreamExecuteRequest, opts ...grpc.CallOption) (SqlQuery_StreamExecuteClient, error)
	// Begin a transaction.
	Begin(ctx context.Context, in *query.BeginRequest, opts ...grpc.CallOption) (*query.BeginResponse, error)
	// Commit a transaction.
	Commit(ctx context.Context, in *query.CommitRequest, opts ...grpc.CallOption) (*query.CommitResponse, error)
	// Rollback a transaction.
	Rollback(ctx context.Context, in *query.RollbackRequest, opts ...grpc.CallOption) (*query.RollbackResponse, error)
	// SplitQuery is the API to facilitate MapReduce-type iterations
	// over large data sets (like full table dumps).
	SplitQuery(ctx context.Context, in *query.SplitQueryRequest, opts ...grpc.CallOption) (*query.SplitQueryResponse, error)
}

type sqlQueryClient struct {
	cc *grpc.ClientConn
}

func NewSqlQueryClient(cc *grpc.ClientConn) SqlQueryClient {
	return &sqlQueryClient{cc}
}

func (c *sqlQueryClient) Execute(ctx context.Context, in *query.ExecuteRequest, opts ...grpc.CallOption) (*query.ExecuteResponse, error) {
	out := new(query.ExecuteResponse)
	err := grpc.Invoke(ctx, "/queryservice.SqlQuery/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlQueryClient) ExecuteBatch(ctx context.Context, in *query.ExecuteBatchRequest, opts ...grpc.CallOption) (*query.ExecuteBatchResponse, error) {
	out := new(query.ExecuteBatchResponse)
	err := grpc.Invoke(ctx, "/queryservice.SqlQuery/ExecuteBatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlQueryClient) StreamExecute(ctx context.Context, in *query.StreamExecuteRequest, opts ...grpc.CallOption) (SqlQuery_StreamExecuteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SqlQuery_serviceDesc.Streams[0], c.cc, "/queryservice.SqlQuery/StreamExecute", opts...)
	if err != nil {
		return nil, err
	}
	x := &sqlQueryStreamExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SqlQuery_StreamExecuteClient interface {
	Recv() (*query.StreamExecuteResponse, error)
	grpc.ClientStream
}

type sqlQueryStreamExecuteClient struct {
	grpc.ClientStream
}

func (x *sqlQueryStreamExecuteClient) Recv() (*query.StreamExecuteResponse, error) {
	m := new(query.StreamExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sqlQueryClient) Begin(ctx context.Context, in *query.BeginRequest, opts ...grpc.CallOption) (*query.BeginResponse, error) {
	out := new(query.BeginResponse)
	err := grpc.Invoke(ctx, "/queryservice.SqlQuery/Begin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlQueryClient) Commit(ctx context.Context, in *query.CommitRequest, opts ...grpc.CallOption) (*query.CommitResponse, error) {
	out := new(query.CommitResponse)
	err := grpc.Invoke(ctx, "/queryservice.SqlQuery/Commit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlQueryClient) Rollback(ctx context.Context, in *query.RollbackRequest, opts ...grpc.CallOption) (*query.RollbackResponse, error) {
	out := new(query.RollbackResponse)
	err := grpc.Invoke(ctx, "/queryservice.SqlQuery/Rollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlQueryClient) SplitQuery(ctx context.Context, in *query.SplitQueryRequest, opts ...grpc.CallOption) (*query.SplitQueryResponse, error) {
	out := new(query.SplitQueryResponse)
	err := grpc.Invoke(ctx, "/queryservice.SqlQuery/SplitQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SqlQuery service

type SqlQueryServer interface {
	// Execute executes the specified SQL query (might be in a
	// transaction context, if Query.transaction_id is set).
	Execute(context.Context, *query.ExecuteRequest) (*query.ExecuteResponse, error)
	// ExecuteBatch executes a list of queries, and returns the result
	// for each query.
	ExecuteBatch(context.Context, *query.ExecuteBatchRequest) (*query.ExecuteBatchResponse, error)
	// StreamExecute executes a streaming query. Use this method if the
	// query returns a large number of rows. The first QueryResult will
	// contain the Fields, subsequent QueryResult messages will contain
	// the rows.
	StreamExecute(*query.StreamExecuteRequest, SqlQuery_StreamExecuteServer) error
	// Begin a transaction.
	Begin(context.Context, *query.BeginRequest) (*query.BeginResponse, error)
	// Commit a transaction.
	Commit(context.Context, *query.CommitRequest) (*query.CommitResponse, error)
	// Rollback a transaction.
	Rollback(context.Context, *query.RollbackRequest) (*query.RollbackResponse, error)
	// SplitQuery is the API to facilitate MapReduce-type iterations
	// over large data sets (like full table dumps).
	SplitQuery(context.Context, *query.SplitQueryRequest) (*query.SplitQueryResponse, error)
}

func RegisterSqlQueryServer(s *grpc.Server, srv SqlQueryServer) {
	s.RegisterService(&_SqlQuery_serviceDesc, srv)
}

func _SqlQuery_Execute_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(query.ExecuteRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SqlQueryServer).Execute(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SqlQuery_ExecuteBatch_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(query.ExecuteBatchRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SqlQueryServer).ExecuteBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SqlQuery_StreamExecute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.StreamExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SqlQueryServer).StreamExecute(m, &sqlQueryStreamExecuteServer{stream})
}

type SqlQuery_StreamExecuteServer interface {
	Send(*query.StreamExecuteResponse) error
	grpc.ServerStream
}

type sqlQueryStreamExecuteServer struct {
	grpc.ServerStream
}

func (x *sqlQueryStreamExecuteServer) Send(m *query.StreamExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SqlQuery_Begin_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(query.BeginRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SqlQueryServer).Begin(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SqlQuery_Commit_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(query.CommitRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SqlQueryServer).Commit(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SqlQuery_Rollback_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(query.RollbackRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SqlQueryServer).Rollback(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SqlQuery_SplitQuery_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(query.SplitQueryRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(SqlQueryServer).SplitQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _SqlQuery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queryservice.SqlQuery",
	HandlerType: (*SqlQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _SqlQuery_Execute_Handler,
		},
		{
			MethodName: "ExecuteBatch",
			Handler:    _SqlQuery_ExecuteBatch_Handler,
		},
		{
			MethodName: "Begin",
			Handler:    _SqlQuery_Begin_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _SqlQuery_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _SqlQuery_Rollback_Handler,
		},
		{
			MethodName: "SplitQuery",
			Handler:    _SqlQuery_SplitQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamExecute",
			Handler:       _SqlQuery_StreamExecute_Handler,
			ServerStreams: true,
		},
	},
}
