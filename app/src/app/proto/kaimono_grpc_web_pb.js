/**
 * @fileoverview gRPC-Web generated client stub for kaimono
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.kaimono = require('./kaimono_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.kaimono.ApiServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.kaimono.ApiServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.kaimono.HandshakeRequest,
 *   !proto.kaimono.HandshakeResponse>}
 */
const methodDescriptor_ApiService_Handshake = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Handshake',
  grpc.web.MethodType.UNARY,
  proto.kaimono.HandshakeRequest,
  proto.kaimono.HandshakeResponse,
  /**
   * @param {!proto.kaimono.HandshakeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.HandshakeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.HandshakeRequest,
 *   !proto.kaimono.HandshakeResponse>}
 */
const methodInfo_ApiService_Handshake = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.HandshakeResponse,
  /**
   * @param {!proto.kaimono.HandshakeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.HandshakeResponse.deserializeBinary
);


/**
 * @param {!proto.kaimono.HandshakeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.HandshakeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.HandshakeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.handshake =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Handshake',
      request,
      metadata || {},
      methodDescriptor_ApiService_Handshake,
      callback);
};


/**
 * @param {!proto.kaimono.HandshakeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.HandshakeResponse>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.handshake =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Handshake',
      request,
      metadata || {},
      methodDescriptor_ApiService_Handshake);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.LoginRequest,
 *   !proto.kaimono.AuthResponse>}
 */
const methodDescriptor_ApiService_Login = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Login',
  grpc.web.MethodType.UNARY,
  proto.kaimono.LoginRequest,
  proto.kaimono.AuthResponse,
  /**
   * @param {!proto.kaimono.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.AuthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.LoginRequest,
 *   !proto.kaimono.AuthResponse>}
 */
const methodInfo_ApiService_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.AuthResponse,
  /**
   * @param {!proto.kaimono.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.kaimono.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Login',
      request,
      metadata || {},
      methodDescriptor_ApiService_Login,
      callback);
};


/**
 * @param {!proto.kaimono.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.AuthResponse>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Login',
      request,
      metadata || {},
      methodDescriptor_ApiService_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.RegistrationRequest,
 *   !proto.kaimono.AuthResponse>}
 */
const methodDescriptor_ApiService_Register = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Register',
  grpc.web.MethodType.UNARY,
  proto.kaimono.RegistrationRequest,
  proto.kaimono.AuthResponse,
  /**
   * @param {!proto.kaimono.RegistrationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.AuthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.RegistrationRequest,
 *   !proto.kaimono.AuthResponse>}
 */
const methodInfo_ApiService_Register = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.AuthResponse,
  /**
   * @param {!proto.kaimono.RegistrationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.kaimono.RegistrationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Register',
      request,
      metadata || {},
      methodDescriptor_ApiService_Register,
      callback);
};


/**
 * @param {!proto.kaimono.RegistrationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.AuthResponse>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Register',
      request,
      metadata || {},
      methodDescriptor_ApiService_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodDescriptor_ApiService_Select = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Select',
  grpc.web.MethodType.UNARY,
  proto.kaimono.RequestBody,
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodInfo_ApiService_Select = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResponseBody)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResponseBody>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.select =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Select',
      request,
      metadata || {},
      methodDescriptor_ApiService_Select,
      callback);
};


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResponseBody>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.select =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Select',
      request,
      metadata || {},
      methodDescriptor_ApiService_Select);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodDescriptor_ApiService_Insert = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Insert',
  grpc.web.MethodType.UNARY,
  proto.kaimono.RequestBody,
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodInfo_ApiService_Insert = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResponseBody)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResponseBody>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.insert =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Insert',
      request,
      metadata || {},
      methodDescriptor_ApiService_Insert,
      callback);
};


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResponseBody>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.insert =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Insert',
      request,
      metadata || {},
      methodDescriptor_ApiService_Insert);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodDescriptor_ApiService_Update = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Update',
  grpc.web.MethodType.UNARY,
  proto.kaimono.RequestBody,
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodInfo_ApiService_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResponseBody)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResponseBody>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Update',
      request,
      metadata || {},
      methodDescriptor_ApiService_Update,
      callback);
};


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResponseBody>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Update',
      request,
      metadata || {},
      methodDescriptor_ApiService_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodDescriptor_ApiService_Delete = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Delete',
  grpc.web.MethodType.UNARY,
  proto.kaimono.RequestBody,
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodInfo_ApiService_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResponseBody)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResponseBody>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Delete',
      request,
      metadata || {},
      methodDescriptor_ApiService_Delete,
      callback);
};


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResponseBody>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Delete',
      request,
      metadata || {},
      methodDescriptor_ApiService_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodDescriptor_ApiService_Create = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/Create',
  grpc.web.MethodType.UNARY,
  proto.kaimono.RequestBody,
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.RequestBody,
 *   !proto.kaimono.ResponseBody>}
 */
const methodInfo_ApiService_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResponseBody,
  /**
   * @param {!proto.kaimono.RequestBody} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResponseBody.deserializeBinary
);


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResponseBody)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResponseBody>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/Create',
      request,
      metadata || {},
      methodDescriptor_ApiService_Create,
      callback);
};


/**
 * @param {!proto.kaimono.RequestBody} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResponseBody>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/Create',
      request,
      metadata || {},
      methodDescriptor_ApiService_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.ResourceRequest,
 *   !proto.kaimono.ResourceOverviewResponse>}
 */
const methodDescriptor_ApiService_PublicResources = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/PublicResources',
  grpc.web.MethodType.UNARY,
  proto.kaimono.ResourceRequest,
  proto.kaimono.ResourceOverviewResponse,
  /**
   * @param {!proto.kaimono.ResourceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResourceOverviewResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.ResourceRequest,
 *   !proto.kaimono.ResourceOverviewResponse>}
 */
const methodInfo_ApiService_PublicResources = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResourceOverviewResponse,
  /**
   * @param {!proto.kaimono.ResourceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResourceOverviewResponse.deserializeBinary
);


/**
 * @param {!proto.kaimono.ResourceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResourceOverviewResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResourceOverviewResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.publicResources =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/PublicResources',
      request,
      metadata || {},
      methodDescriptor_ApiService_PublicResources,
      callback);
};


/**
 * @param {!proto.kaimono.ResourceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResourceOverviewResponse>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.publicResources =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/PublicResources',
      request,
      metadata || {},
      methodDescriptor_ApiService_PublicResources);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kaimono.ResourceRequest,
 *   !proto.kaimono.ResourceOverviewResponse>}
 */
const methodDescriptor_ApiService_PrivateResources = new grpc.web.MethodDescriptor(
  '/kaimono.ApiService/PrivateResources',
  grpc.web.MethodType.UNARY,
  proto.kaimono.ResourceRequest,
  proto.kaimono.ResourceOverviewResponse,
  /**
   * @param {!proto.kaimono.ResourceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResourceOverviewResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.kaimono.ResourceRequest,
 *   !proto.kaimono.ResourceOverviewResponse>}
 */
const methodInfo_ApiService_PrivateResources = new grpc.web.AbstractClientBase.MethodInfo(
  proto.kaimono.ResourceOverviewResponse,
  /**
   * @param {!proto.kaimono.ResourceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kaimono.ResourceOverviewResponse.deserializeBinary
);


/**
 * @param {!proto.kaimono.ResourceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.kaimono.ResourceOverviewResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kaimono.ResourceOverviewResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kaimono.ApiServiceClient.prototype.privateResources =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kaimono.ApiService/PrivateResources',
      request,
      metadata || {},
      methodDescriptor_ApiService_PrivateResources,
      callback);
};


/**
 * @param {!proto.kaimono.ResourceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kaimono.ResourceOverviewResponse>}
 *     A native promise that resolves to the response
 */
proto.kaimono.ApiServicePromiseClient.prototype.privateResources =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kaimono.ApiService/PrivateResources',
      request,
      metadata || {},
      methodDescriptor_ApiService_PrivateResources);
};


module.exports = proto.kaimono;

