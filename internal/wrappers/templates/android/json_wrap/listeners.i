/********************************************/
/*               EventListener              */
/********************************************/

%typemap(javaclassmodifiers) kuzzleio::EventListener "public abstract class"
%javamethodmodifiers kuzzleio::EventListener::trigger "public abstract"
%typemap(javaout) void kuzzleio::EventListener::trigger ";"

%typemap(jni) const std::string& jsonResponse "jobject"
%typemap(jtype) const std::string& jsonResponse "org.json.JSONObject"
%typemap(jstype) const std::string& jsonResponse "org.json.JSONObject"
%typemap(javain) const std::string& jsonResponse "$javainput"
%typemap(in) const std::string& jsonResponse {}
%typemap(javadirectorin) const std::string& jsonResponse "$jniinput";
%typemap(directorin, descriptor="Lorg/json/JSONObject;", noblock=1) const std::string& jsonResponse {
  const jclass clazz = JCALL1(FindClass, jenv, "org/json/JSONObject");
  
  if ($input) {
    jmethodID methodID = jenv->GetMethodID(clazz, "<init>", "(Ljava/lang/String;)V");
    jobject a = jenv->NewObject(clazz, methodID, $input);
    if (!a) return $null;
    $input = a;
  } else {
    $input = NULL;
  }
  Swig::LocalRefGuard $1_refguard(jenv, $input);
}
%apply const std::string& jsonResponse { const std::string& jsonResponse };


/********************************************/
/*               SubscribeListener          */
/********************************************/

%typemap(javaclassmodifiers) kuzzleio::SubscribeListener "public abstract class"
%javamethodmodifiers kuzzleio::SubscribeListener::onSubscribe "public abstract"
%typemap(javaout) void kuzzleio::SubscribeListener::onSubscribe ";"
