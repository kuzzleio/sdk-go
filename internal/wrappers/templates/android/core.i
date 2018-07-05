%include "../java/common.i"

%include "../java/exceptions.i"
%include "std_string.i"
%include "typemap.i"
%include "../../kcore.i"

%include "std_vector.i"

%template(StringVector) std::vector<std::string>;

%typemap(out) const StringVector& %{
    return $1;
%}

%pragma(java) jniclasscode=%{
  static {
    try {
      System.loadLibrary("kuzzle-wrapper-android");
    } catch (Exception e) {
      System.err.println("Native code library failed to load. \n");
      e.printStackTrace();
      System.exit(1);
    }
  }
%}

%extend options {
    options() {
        options *o = kuzzle_new_options();
        return o;
    }

    ~options() {
        free($self);
    }
}

%include "kuzzle.cpp"
%include "collection.cpp"
%include "document.cpp"
%include "realtime.cpp"
%include "auth.cpp"
%include "index.cpp"
%include "server.cpp"
