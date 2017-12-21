#ifndef _KUZZLE_HPP_
#define _KUZZLE_HPP_

#include <exception>
#include <stdexcept>
#include <string>

#include <vector>

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

    KuzzleException(int status=500, const std::string& message="Internal Exception");
    KuzzleException(KuzzleException& ke) : status(ke.status), std::runtime_error(ke.getMessage()) {};

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
  void throwExceptionFromStatus(T result) {
    if (result.status == 206)
        throw PartialException(result.error);
    else if (result.status == 400)
        throw BadRequestException(result.error);
    else if (result.status == 401) {
        throw UnauthorizedException(result.error);
    } else if (result.status == 403)
        throw ForbiddenException(result.error);
    else if (result.status == 404)
        throw NotFoundException(result.error);
    else if (result.status == 412)
        throw PreconditionException(result.error);
    else if (result.status == 413)
        throw SizeLimitException(result.error);
    else if (result.status == 500)
        throw InternalException(result.error);
    else if (result.status == 504)
        throw GatewayTimeoutException(result.error);
    else if (result.status == 503)
        throw ServiceUnavailableException(result.error);
  }

  class Kuzzle {
    kuzzle *_kuzzle;

    public:
      Kuzzle(const std::string& host, options *options=NULL);
      virtual ~Kuzzle();

      token_validity* checkToken(const std::string& token);
      char* connect();
      bool createIndex(const std::string& index, query_options* options=NULL) Kuz_Throw_KuzzleException;
      json_object* createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;

      bool deleteMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      json_object* getMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      json_object* updateMyCredentials(const std::string& strategy, json_object* credentials, query_options *options=NULL) Kuz_Throw_KuzzleException;
      bool validateMyCredentials(const std::string& strategy, json_object* credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string login(const std::string& strategy, json_object* credentials, int expiresIn) Kuz_Throw_KuzzleException;
      std::string login(const std::string& strategy, json_object* credentials) Kuz_Throw_KuzzleException;
      statistics* getAllStatistics(query_options* options=NULL) Kuz_Throw_KuzzleException;
      statistics* getStatistics(time_t start, time_t end, query_options* options=NULL) Kuz_Throw_KuzzleException;
      bool getAutoRefresh(const std::string& index, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getJwt();
      json_object* getMyRights(query_options* options=NULL) Kuz_Throw_KuzzleException;
      json_object* getServerInfo(query_options* options=NULL) Kuz_Throw_KuzzleException;
      collection_entry* listCollections(const std::string& index, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::vector<std::string> listIndexes(query_options* options=NULL) Kuz_Throw_KuzzleException;
      void disconnect();
      void logout();
      kuzzle_response* query(kuzzle_request* query, query_options* options=NULL) Kuz_Throw_KuzzleException;
      shards* refreshIndex(const std::string& index, query_options* options=NULL) Kuz_Throw_KuzzleException;
      long long now(query_options* options=NULL) Kuz_Throw_KuzzleException;
  };
}

#endif