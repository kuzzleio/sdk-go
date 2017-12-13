#ifndef _KUZZLE_HPP_
#define _KUZZLE_HPP_

#include <exception>
#include <stdexcept>
#include <string>

extern "C" {
  #define _Complex
  #include <stdio.h>

  #include "kuzzle.h"
  #include "kuzzlesdk.h"
  #include "swig.h"
}

#define Kuz_Throw_KuzzleException throw(\
  BadRequestException, \
  ForbiddenException, \
  GatewayTimeoutException, \
  InternalException, \
  NotFoundException, \
  PartialException, \
  PreconditionException, \
  ServiceUnavailableException, \
  SizeLimitException, \
  UnauthorizedException, \
  KuzzleException \
)

namespace kuzzleio {

  // Exceptions
  struct KuzzleException : std::runtime_error {
    int status;
    std::string stack;

    KuzzleException(int status=500, const std::string& message="Kuzzle Exception", const std::string& stack="");
    virtual ~KuzzleException() throw() {};
    std::string getMessage();
  };

  struct BadRequestException : KuzzleException {
    BadRequestException(const std::string& message="Bad Request Exception", const std::string& stack="")
      : KuzzleException(400, message, stack) {};
  };
  struct ForbiddenException: KuzzleException {
    ForbiddenException(const std::string& message="Forbidden Exception", const std::string& stack="")
      : KuzzleException(403, message, stack) {};
  };
  struct GatewayTimeoutException: KuzzleException {
    GatewayTimeoutException(const std::string& message="Gateway Timeout Exception", const std::string& stack="")
      : KuzzleException(504, message, stack) {};
  };
  struct InternalException: KuzzleException {
    InternalException(const std::string& message="Internal Exception", const std::string& stack="")
      : KuzzleException(500, message, stack) {};
  };
  struct NotFoundException: KuzzleException {
    NotFoundException(const std::string& message="Not Found Exception", const std::string& stack="")
      : KuzzleException(404, message, stack) {};
  };
  struct PartialException: KuzzleException {
    PartialException(const std::string& message="Partial Exception", const std::string& stack="")
      : KuzzleException(206, message, stack) {};
  };
  struct PreconditionException: KuzzleException {
    PreconditionException(const std::string& message="Precondition Exception", const std::string& stack="")
      : KuzzleException(412, message, stack) {};
  };
  struct ServiceUnavailableException: KuzzleException {
    ServiceUnavailableException(const std::string& message="Service Unavailable Exception", const std::string& stack="")
      : KuzzleException(503, message, stack) {};
  };
  struct SizeLimitException: KuzzleException {
    SizeLimitException(const std::string& message="Size Limit Exception", const std::string& stack="")
      : KuzzleException(413, message, stack) {};
  };
  struct UnauthorizedException: KuzzleException {
    UnauthorizedException(const std::string& message="Unauthorized Exception", const std::string& stack="")
      : KuzzleException(401, message, stack) {};
  };

  template <class T>
  void throwExceptionFromStatus(T result) Kuz_Throw_KuzzleException {
    printf("-- %s\n", result.stack);
    if (result.status == 206)
        throw PartialException(result.error, result.stack);
    else if (result.status == 400)
        throw BadRequestException(result.error, result.stack);
    else if (result.status == 401)
        throw UnauthorizedException(result.error, result.stack);
    else if (result.status == 403)
        throw ForbiddenException(result.error, result.stack);
    else if (result.status == 404)
        throw NotFoundException(result.error, result.stack);
    else if (result.status == 412)
        throw PreconditionException(result.error, result.stack);
    else if (result.status == 413)
        throw SizeLimitException(result.error, result.stack);
    else if (result.status == 500)
        throw InternalException(result.error, result.stack);
    else if (result.status == 504)
        throw GatewayTimeoutException(result.error, result.stack);
    else if (result.status == 503)
        throw ServiceUnavailableException(result.error, result.stack);
  }

  class Kuzzle {
    kuzzle *_kuzzle;

    public:
      Kuzzle(const std::string& host, options *options=NULL);
      virtual ~Kuzzle();

      token_validity* checkToken(const std::string&);
      char* connect();
      bool_result* createIndex(const std::string&, query_options* options=NULL);
      json_result* createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options=NULL);

      bool deleteMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
  };
}

#endif