// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Pb

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

// ClienteSvClient is the client API for ClienteSv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClienteSvClient interface {
	DimeHola(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Mensaje, error)
	MandarJugadores(ctx context.Context, in *Mensajito2, opts ...grpc.CallOption) (*Mensajito2, error)
}

type clienteSvClient struct {
	cc grpc.ClientConnInterface
}

func NewClienteSvClient(cc grpc.ClientConnInterface) ClienteSvClient {
	return &clienteSvClient{cc}
}

func (c *clienteSvClient) DimeHola(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Mensaje, error) {
	out := new(Mensaje)
	err := c.cc.Invoke(ctx, "/pb.ClienteSv/DimeHola", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clienteSvClient) MandarJugadores(ctx context.Context, in *Mensajito2, opts ...grpc.CallOption) (*Mensajito2, error) {
	out := new(Mensajito2)
	err := c.cc.Invoke(ctx, "/pb.ClienteSv/MandarJugadores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClienteSvServer is the server API for ClienteSv service.
// All implementations must embed UnimplementedClienteSvServer
// for forward compatibility
type ClienteSvServer interface {
	DimeHola(context.Context, *Mensaje) (*Mensaje, error)
	MandarJugadores(context.Context, *Mensajito2) (*Mensajito2, error)
	mustEmbedUnimplementedClienteSvServer()
}

// UnimplementedClienteSvServer must be embedded to have forward compatible implementations.
type UnimplementedClienteSvServer struct {
}

func (UnimplementedClienteSvServer) DimeHola(context.Context, *Mensaje) (*Mensaje, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DimeHola not implemented")
}
func (UnimplementedClienteSvServer) MandarJugadores(context.Context, *Mensajito2) (*Mensajito2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarJugadores not implemented")
}
func (UnimplementedClienteSvServer) mustEmbedUnimplementedClienteSvServer() {}

// UnsafeClienteSvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClienteSvServer will
// result in compilation errors.
type UnsafeClienteSvServer interface {
	mustEmbedUnimplementedClienteSvServer()
}

func RegisterClienteSvServer(s grpc.ServiceRegistrar, srv ClienteSvServer) {
	s.RegisterService(&ClienteSv_ServiceDesc, srv)
}

func _ClienteSv_DimeHola_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClienteSvServer).DimeHola(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClienteSv/DimeHola",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClienteSvServer).DimeHola(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClienteSv_MandarJugadores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensajito2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClienteSvServer).MandarJugadores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClienteSv/MandarJugadores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClienteSvServer).MandarJugadores(ctx, req.(*Mensajito2))
	}
	return interceptor(ctx, in, info, handler)
}

// ClienteSv_ServiceDesc is the grpc.ServiceDesc for ClienteSv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClienteSv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ClienteSv",
	HandlerType: (*ClienteSvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DimeHola",
			Handler:    _ClienteSv_DimeHola_Handler,
		},
		{
			MethodName: "MandarJugadores",
			Handler:    _ClienteSv_MandarJugadores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}
