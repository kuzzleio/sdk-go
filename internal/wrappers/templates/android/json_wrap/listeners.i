%define TYPEMAP_DIRECTOR_INPUT(JNI, CppType, JavaType, JNITypeFrom, JNITypeTo)
  %typemap(jni) CppType JNI
  %typemap(jtype) CppType JavaType
  %typemap(jstype) CppType JavaType
  %typemap(javain) CppType "$javainput"
  %typemap(in) CppType {}
  %typemap(javadirectorin) CppType "$jniinput";
  %typemap(directorin, descriptor="L"JNITypeTo";", noblock=1) CppType {
    if ($input) {
      const jclass clazz = JCALL1(FindClass, jenv, JNITypeTo);
      jmethodID methodID = jenv->GetMethodID(clazz, "<init>", "(L"JNITypeFrom";)V");
      jobject a = jenv->NewObject(clazz, methodID, $input);
      if (!a) return $null;
      $input = a;
    } else {
      $input = NULL;
    }
    Swig::LocalRefGuard $1_refguard(jenv, $input);
  }
  %apply CppType { CppType };
%enddef


/********************************************/
/*               EventListener              */
/********************************************/

TYPEMAP_DIRECTOR_INPUT("jobject", const std::string& jsonResponse, "org.json.JSONObject", "java/lang/String", "org/json/JSONObject")


/********************************************/
/*               NotificationListener       */
/********************************************/

%ignore kuzzleio::Realtime::getListener(const std::string& roomId);
%typemap(jni) char* kuzzleio::notification_content::content "jstring"
%typemap(jstype) char* kuzzleio::notification_content::content "org.json.JSONObject"

%typemap(javaout) char* kuzzleio::notification_content::content {
    org.json.JSONObject res = null;
    try {
      res = new org.json.JSONObject($jnicall);
    } catch (org.json.JSONException e) {
      throw new RuntimeException(e);
    }

    return res;
  }