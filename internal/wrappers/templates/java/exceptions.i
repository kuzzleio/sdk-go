%define CATCH(cppException, JNIException)
  catch (cppException &e) {
    jclass clazz = jenv->FindClass(JNIException);
    jenv->ThrowNew(clazz, e.what());
    return $null;
  }
%enddef

%javaexception("io.kuzzle.sdk.BadRequestException, io.kuzzle.sdk.ForbiddenException, 
                io.kuzzle.sdk.GatewayTimeoutException, io.kuzzle.sdk.InternalException, 
                io.kuzzle.sdk.NotFoundException, io.kuzzle.sdk.PartialException,
                io.kuzzle.sdk.PreconditionException, io.kuzzle.sdk.ServiceUnavailableException,
                io.kuzzle.sdk.SizeLimitException, io.kuzzle.sdk.UnauthorizedException,
                io.kuzzle.sdk.KuzzleException") {
    try {
        $action
    } CATCH (kuzzleio::BadRequestException, "io/kuzzle/sdk/BadRequestException")
    CATCH (kuzzleio::ForbiddenException, "io/kuzzle/sdk/ForbiddenException")
    CATCH (kuzzleio::GatewayTimeoutException, "io/kuzzle/sdk/GatewayTimeoutException")
    CATCH (kuzzleio::InternalException, "io/kuzzle/sdk/InternalException")
    CATCH (kuzzleio::NotFoundException, "io/kuzzle/sdk/NotFoundException")
    CATCH (kuzzleio::PartialException, "io/kuzzle/sdk/PartialException")
    CATCH (kuzzleio::PreconditionException, "io/kuzzle/sdk/PreconditionException")
    CATCH (kuzzleio::ServiceUnavailableException, "io/kuzzle/sdk/ServiceUnavailableException")
    CATCH (kuzzleio::SizeLimitException, "io/kuzzle/sdk/SizeLimitException")
    CATCH (kuzzleio::UnauthorizedException, "io/kuzzle/sdk/UnauthorizedException")
    CATCH (kuzzleio::KuzzleException, "io/kuzzle/sdk/KuzzleException")   
}


%typemap(javabase) kuzzleio::KuzzleException "java.lang.RuntimeException";
