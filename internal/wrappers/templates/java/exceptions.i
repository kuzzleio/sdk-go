%ignore kuzzleio::KuzzleException::stack;

%define TO_JAVA_EXCEPTION(CPPTYPE, JTYPE, JNITYPE)
%typemap(throws,throws=JTYPE) CPPTYPE {
  (void)$1;
  jclass excpcls = jenv->FindClass(JNITYPE);
  if (excpcls) {
    jenv->ThrowNew(excpcls, $1.what());
   }
  return $null;
}
%enddef

%typemap(javabase) kuzzleio::KuzzleException "java.lang.RuntimeException";

TO_JAVA_EXCEPTION(kuzzleio::BadRequestException, "io.kuzzle.sdk.BadRequestException", "io/kuzzle/sdk/BadRequestException");
TO_JAVA_EXCEPTION(kuzzleio::ForbiddenException, "io.kuzzle.sdk.ForbiddenException", "io/kuzzle/sdk/ForbiddenException");
TO_JAVA_EXCEPTION(kuzzleio::GatewayTimeoutException, "io.kuzzle.sdk.GatewayTimeoutException", "io/kuzzle/sdk/GatewayTimeoutException");
TO_JAVA_EXCEPTION(kuzzleio::InternalException, "io.kuzzle.sdk.InternalException", "io/kuzzle/sdk/InternalException");
TO_JAVA_EXCEPTION(kuzzleio::NotFoundException, "io.kuzzle.sdk.NotFoundException", "io/kuzzle/sdk/NotFoundException");
TO_JAVA_EXCEPTION(kuzzleio::PartialException, "io.kuzzle.sdk.PartialException", "io/kuzzle/sdk/PartialException");
TO_JAVA_EXCEPTION(kuzzleio::PreconditionException, "io.kuzzle.sdk.PreconditionException", "io/kuzzle/sdk/PreconditionException");
TO_JAVA_EXCEPTION(kuzzleio::ServiceUnavailableException, "io.kuzzle.sdk.ServiceUnavailableException", "io/kuzzle/sdk/ServiceUnavailableException");
TO_JAVA_EXCEPTION(kuzzleio::SizeLimiException, "io.kuzzle.sdk.SizeLimiException", "io/kuzzle/sdk/SizeLimiException");
TO_JAVA_EXCEPTION(kuzzleio::UnauthorizedException, "io.kuzzle.sdk.UnauthorizedException", "io/kuzzle/sdk/UnauthorizedException");
