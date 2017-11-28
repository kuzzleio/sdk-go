#ifndef _KUZZLE_HPP_
#define _KUZZLE_HPP_

#include <exception>
#include <stdexcept>
#include <string>

extern "C" {
  #define _Complex
  #include <Python.h>
  #include <stdio.h>

  #include "kuzzle.h"
  #include "kuzzlesdk.h"
  #include "swig.h"
}

#define Kuz_Throw_KuzzleError throw(\
  BadRequestError, \
  ExternalServiceError, \
  ForbiddenError, \
  GatewayTimeoutError, \
  InternalError, \
  NotFoundError, \
  ParseError, \
  PartialError, \
  PluginImplementationError, \
  PreconditionError, \
  ServiceUnavailableError, \
  SizeLimiError, \
  UnauthorizedError, \
  KuzzleError \
)

namespace kuzzleio {

  // errors
  struct KuzzleError: std::runtime_error {
    int status;
    std::string stack;

    KuzzleError(int status=500, const std::string& error="Kuzzle Error", const std::string& stack="");
    ~KuzzleError() throw() {};
  };
  struct BadRequestError: KuzzleError {
    BadRequestError(const std::string& error="Bad Request Error", const std::string& stack="")
      : KuzzleError(400, error, stack) {};
  };
  struct ExternalServiceError: KuzzleError {
    ExternalServiceError(const std::string& error="External Service Error", const std::string& stack="")
      : KuzzleError(500, error, stack) {};
  };
  struct ForbiddenError: KuzzleError {
    ForbiddenError(const std::string& error="Forbidden Error", const std::string& stack="")
      : KuzzleError(403, error, stack) {};
  };
  struct GatewayTimeoutError: KuzzleError {
    GatewayTimeoutError(const std::string& error="Gateway Timeout Error", const std::string& stack="")
      : KuzzleError(504, error, stack) {};
  };
  struct InternalError: KuzzleError {
    InternalError(const std::string& error="Internal Error", const std::string& stack="")
      : KuzzleError(500, error, stack) {};
  };
  struct NotFoundError: KuzzleError {
    NotFoundError(const std::string& error="Not Found Error", const std::string& stack="")
      : KuzzleError(404, error, stack) {};
  };
  struct ParseError: KuzzleError {
    ParseError(const std::string& error="Parse Error", const std::string& stack="")
      :KuzzleError(400, error, stack) {};
  };
  struct PartialError: KuzzleError {
    PartialError(const std::string& error="Partial Error", const std::string& stack="")
      : KuzzleError(206, error, stack) {};
  };
  struct PluginImplementationError: KuzzleError {
    PluginImplementationError(const std::string& error="Plugin Implementation Error", const std::string& stack="")
      : KuzzleError(500, error, stack) {};
  };
  struct PreconditionError: KuzzleError {
    PreconditionError(const std::string& error="Precondition Error", const std::string& stack="")
      : KuzzleError(412, error, stack) {};
  };
  struct ServiceUnavailableError: KuzzleError {
    ServiceUnavailableError(const std::string& error="Service Unavailable Error", const std::string& stack="")
      : KuzzleError(503, error, stack) {};
  };
  struct SizeLimiError: KuzzleError {
    SizeLimiError(const std::string& error="Size Limit Error", const std::string& stack="")
      : KuzzleError(413, error, stack) {};
  };
  struct UnauthorizedError: KuzzleError {
    UnauthorizedError(const std::string& error="Unauthorized Error", const std::string& stack="")
      : KuzzleError(401, error, stack) {};
  };


  class Kuzzle {
    kuzzle *_kuzzle;

    public:
      Kuzzle(std::string host, options *options=NULL);
      ~Kuzzle();

      long long now(query_options *options=NULL) Kuz_Throw_KuzzleError;
  };


}

#endif
