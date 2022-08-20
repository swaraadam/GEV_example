package auth

import (
	"context"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/**
 * AuthInterceptor is a server interceptor for authentication and authorization
 */
type AuthInterceptor struct{}

/**
 * wrappedStream wraps around the embedded grpc.ServerStream, and intercepts the RecvMsg and SendMsg method call
 */
type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

/**
 * NewAuthInterceptor returns a new auth interceptor
 */
func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

/**
 * Unary returns a server interceptor function to authenticate and authorize Unary RPC
 */
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("-- Unary Interceptor: ", info.FullMethod)
		stringArr := strings.Split(info.FullMethod, "/")
		field := stringArr[len(stringArr)-1]
		if Endpoints[field].RequireAuth {
			md, _ := metadata.FromIncomingContext(ctx)
			token, err := interceptor.ParseToken(md, info.FullMethod)
			if err != nil {
				return nil, err
			}
			// if !interceptor.isAuthorized(token.Role, Endpoints[field].Roles) {
			// 	return nil, status.Error(codes.InvalidArgument, "User Unauthorized!")
			// }
			header := metadata.Pairs("tokenid", token.ID.String(), "gamerid", token.UserId)
			newContext := metadata.NewIncomingContext(ctx, header)
			return handler(newContext, req)
		}
		return handler(ctx, req)
	}
}

/**
 * Stream returns a server interceptor function to authenticate and authorize Stream RPC
 */
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Println("-- Stream Interceptor: ", info.FullMethod)
		stringArr := strings.Split(info.FullMethod, "/")
		field := stringArr[len(stringArr)-1]
		if Endpoints[field].RequireAuth {
			md, _ := metadata.FromIncomingContext(stream.Context())
			token, err := interceptor.ParseToken(md, info.FullMethod)
			if err != nil {
				return err
			}
			// if !interceptor.isAuthorized(token.Role, Endpoints[field].Roles) {
			// 	return status.Error(codes.InvalidArgument, "User Unauthorized!")
			// }
			header := metadata.Pairs("tokenid", token.ID.String(), "gamerid", token.UserId)
			newContext := metadata.NewIncomingContext(stream.Context(), header)
			return handler(srv, &wrappedStream{
				ServerStream: stream,
				ctx:          newContext,
			})
		}
		return handler(srv, stream)
	}
}

/**
 * ParseToken reads accesstoken from metadata, verifies and returns claims
 */
func (interceptor *AuthInterceptor) ParseToken(md metadata.MD, info string) (*Claim, error) {
	// check if AccessToken is present in metadata
	if accessToken, ok := md["accesstoken"]; !ok {
		return nil, status.Error(codes.InvalidArgument, "User Unauthorized!")
	} else {
		// parse token (expects GamerID within claim)
		jwt, err := NewJWTMaker()
		if err != nil {
			return nil, err
		}
		token, err := jwt.VerifyToken(accessToken[0])
		if err != nil {
			return nil, err
		}
		return token, nil
	}
}

/**
 * Check USER's role to see if they are allowed to access the endpoint
 */
func (interceptor *AuthInterceptor) isAuthorized(role string, roles []string) bool {
	for i := range roles {
		if roles[i] == role {
			return true
		}
	}
	return false
}

/**
 * Custom context for stream interceptor
 */
func (w *wrappedStream) Context() context.Context { return w.ctx }
