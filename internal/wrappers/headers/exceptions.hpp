#ifndef _EXCEPTIONS_HPP_
#define _EXCEPTIONS_HPP_

#include <exception>
#include <stdexcept>

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
  UnauthorizedException \
)

#define PARTIAL_EXCEPTION 206
#define BAD_REQUEST_EXCEPTION 400
#define UNAUTHORIZED_EXCEPTION 401
#define FORBIDDEN_EXCEPTION 403
#define NOT_FOUND_EXCEPTION 404
#define PRECONDITION_EXCEPTION 412
#define SIZE_LIMIT_EXCEPTION 413
#define INTERNAL_EXCEPTION 500
#define SERVICE_UNAVAILABLE_EXCEPTION 503
#define GATEWAY_TIMEOUT_EXCEPTION 504

namespace kuzzleio {

  struct KuzzleException : std::runtime_error {
    int status;

    KuzzleException(int status=500, const std::string& message="Internal Exception");
    KuzzleException(const KuzzleException& ke) : status(ke.status), std::runtime_error(ke.getMessage()) {};

    virtual ~KuzzleException() throw() {};
    std::string getMessage() const;
  };

  struct BadRequestException : KuzzleException {
    BadRequestException(const std::string& message="Bad Request Exception")
      : KuzzleException(400, message) {};
    BadRequestException(const BadRequestException& bre) : KuzzleException(bre.status, bre.getMessage()) {}
  };
  struct ForbiddenException: KuzzleException {
    ForbiddenException(const std::string& message="Forbidden Exception")
      : KuzzleException(403, message) {};
    ForbiddenException(const ForbiddenException& fe) : KuzzleException(fe.status, fe.getMessage()) {}
  };
  struct GatewayTimeoutException: KuzzleException {
    GatewayTimeoutException(const std::string& message="Gateway Timeout Exception")
      : KuzzleException(504, message) {};
    GatewayTimeoutException(const GatewayTimeoutException& gte) : KuzzleException(gte.status, gte.getMessage()) {}
  };
  struct InternalException: KuzzleException {
    InternalException(const std::string& message="Internal Exception")
      : KuzzleException(500, message) {};
    InternalException(const InternalException& ie) : KuzzleException(ie.status, ie.getMessage()) {}
  };
  struct NotFoundException: KuzzleException {
    NotFoundException(const std::string& message="Not Found Exception")
      : KuzzleException(404, message) {};
    NotFoundException(const NotFoundException& nfe) : KuzzleException(nfe.status, nfe.getMessage()) {}
  };
  struct PartialException: KuzzleException {
    PartialException(const std::string& message="Partial Exception")
      : KuzzleException(206, message) {};
    PartialException(const PartialException& pe) : KuzzleException(pe.status, pe.getMessage()) {}
  };
  struct PreconditionException: KuzzleException {
    PreconditionException(const std::string& message="Precondition Exception")
      : KuzzleException(412, message) {};
    PreconditionException(const PreconditionException& pe) : KuzzleException(pe.status, pe.getMessage()) {}
  };
  struct ServiceUnavailableException: KuzzleException {
    ServiceUnavailableException(const std::string& message="Service Unavailable Exception")
      : KuzzleException(503, message) {};
    ServiceUnavailableException(const ServiceUnavailableException& sue) : KuzzleException(sue.status, sue.getMessage()) {}
  };
  struct SizeLimitException: KuzzleException {
    SizeLimitException(const std::string& message="Size Limit Exception")
      : KuzzleException(413, message) {};
    SizeLimitException(const SizeLimitException& sle) : KuzzleException(sle.status, sle.getMessage()) {}
  };
  struct UnauthorizedException : KuzzleException {
    UnauthorizedException(const std::string& message="Unauthorized Exception")
     : KuzzleException(401, message) {}
    UnauthorizedException(const UnauthorizedException& ue) : KuzzleException(ue.status, ue.getMessage()) {}
  };

  template <class T>
  void throwExceptionFromStatus(T *result) {
    const std::string error = std::string(result->error);
    delete(result->error);
    if (result->stack) {
      free(result->stack);
    }

    switch(result->status) {
      case PARTIAL_EXCEPTION:
        delete(result);
        throw PartialException(error);
      break;
      case BAD_REQUEST_EXCEPTION:
        delete(result);
        throw BadRequestException(error);
      break;
      case UNAUTHORIZED_EXCEPTION:
        delete(result);
        throw UnauthorizedException(error);
      break;
      case FORBIDDEN_EXCEPTION:
        delete(result);
        throw ForbiddenException(error);      
      break;
      case NOT_FOUND_EXCEPTION:
        delete(result);
        throw NotFoundException(error);  
      break;
      case PRECONDITION_EXCEPTION:
        delete(result);
        throw PreconditionException(error);  
      break;
      case SIZE_LIMIT_EXCEPTION:
        delete(result);
        throw SizeLimitException(error);        
      break;
      case INTERNAL_EXCEPTION:
        delete(result);
        throw InternalException(error); 
      break;
      case SERVICE_UNAVAILABLE_EXCEPTION:
        delete(result);
        throw ServiceUnavailableException(error);         
      break;
      case GATEWAY_TIMEOUT_EXCEPTION:
        delete(result);
        throw GatewayTimeoutException(error);         
      break;
      default:
        delete(result);
        throw KuzzleException(500, error);
    }
  }
}

#endif