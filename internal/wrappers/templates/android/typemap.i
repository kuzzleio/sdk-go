%include <typemaps.i> 

%typemap(jni) std::string* "jobject"
%typemap(jtype) std::string* "java.lang.String"
%typemap(jstype) std::string* "java.lang.String"
%typemap(javain) std::string* "$javainput"
%typemap(in) std::string * (std::string strTemp) {
  jobject oInput = $input;
  if (oInput != NULL) {
    jstring sInput = static_cast<jstring>( oInput );

    const char * $1_pstr = (const char *)jenv->GetStringUTFChars(sInput, 0);
    if (!$1_pstr) return $null;
    strTemp.assign( $1_pstr );
    jenv->ReleaseStringUTFChars( sInput, $1_pstr);
    $1 = &strTemp;
  } else {
    $1 = NULL;
  }
}
%apply std::string * { std::string* };


%typemap(jni) Event& "jobject"
%typemap(jtype) Event& "Event"
%typemap(jstype) Event& "Event"
%typemap(javain) Event& "$javainput"
%typemap(in) Event& (Event tmp) {
  jmethodID swigValueMethod = jenv->GetMethodID(jenv->GetObjectClass($input), "swigValue", "()I");
  const jclass clazz = JCALL1(FindClass, jenv, "io/kuzzle/sdk/Event");
  jint swigValue = jenv->CallIntMethod($input, swigValueMethod, 0);

  Event e = (Event)swigValue;
  $1 = &e;
}
%apply Event& event { Event& event };
