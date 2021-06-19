// package: main
// file: proto/service.proto

var proto_service_pb = require("../proto/service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var MyGRPC = (function () {
  function MyGRPC() {}
  MyGRPC.serviceName = "main.MyGRPC";
  return MyGRPC;
}());

MyGRPC.newMLResult = {
  methodName: "newMLResult",
  service: MyGRPC,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_pb.ReqMLResult,
  responseType: proto_service_pb.ResEmpty
};

MyGRPC.pullMLResult = {
  methodName: "pullMLResult",
  service: MyGRPC,
  requestStream: false,
  responseStream: true,
  requestType: proto_service_pb.ReqEmpty,
  responseType: proto_service_pb.ResMLResult
};

MyGRPC.confirmContainerID = {
  methodName: "confirmContainerID",
  service: MyGRPC,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_pb.ReqConfirmContainerID,
  responseType: proto_service_pb.ResEmpty
};

exports.MyGRPC = MyGRPC;

function MyGRPCClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MyGRPCClient.prototype.newMLResult = function newMLResult(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MyGRPC.newMLResult, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MyGRPCClient.prototype.pullMLResult = function pullMLResult(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(MyGRPC.pullMLResult, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

MyGRPCClient.prototype.confirmContainerID = function confirmContainerID(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MyGRPC.confirmContainerID, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.MyGRPCClient = MyGRPCClient;

