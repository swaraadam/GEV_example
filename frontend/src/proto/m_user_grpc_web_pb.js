/**
 * @fileoverview gRPC-Web generated client stub for m_user
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.m_user = require('./m_user_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.m_user.UserServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.m_user.UserServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.m_user.RegisterRequest,
 *   !proto.m_user.RegisterResponse>}
 */
const methodDescriptor_UserService_Register = new grpc.web.MethodDescriptor(
  '/m_user.UserService/Register',
  grpc.web.MethodType.UNARY,
  proto.m_user.RegisterRequest,
  proto.m_user.RegisterResponse,
  /**
   * @param {!proto.m_user.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.m_user.RegisterResponse.deserializeBinary
);


/**
 * @param {!proto.m_user.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.m_user.RegisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.m_user.RegisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.m_user.UserServiceClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/m_user.UserService/Register',
      request,
      metadata || {},
      methodDescriptor_UserService_Register,
      callback);
};


/**
 * @param {!proto.m_user.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.m_user.RegisterResponse>}
 *     Promise that resolves to the response
 */
proto.m_user.UserServicePromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/m_user.UserService/Register',
      request,
      metadata || {},
      methodDescriptor_UserService_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.m_user.LoginRequest,
 *   !proto.m_user.LoginResponse>}
 */
const methodDescriptor_UserService_Login = new grpc.web.MethodDescriptor(
  '/m_user.UserService/Login',
  grpc.web.MethodType.UNARY,
  proto.m_user.LoginRequest,
  proto.m_user.LoginResponse,
  /**
   * @param {!proto.m_user.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.m_user.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.m_user.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.m_user.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.m_user.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.m_user.UserServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/m_user.UserService/Login',
      request,
      metadata || {},
      methodDescriptor_UserService_Login,
      callback);
};


/**
 * @param {!proto.m_user.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.m_user.LoginResponse>}
 *     Promise that resolves to the response
 */
proto.m_user.UserServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/m_user.UserService/Login',
      request,
      metadata || {},
      methodDescriptor_UserService_Login);
};


module.exports = proto.m_user;

