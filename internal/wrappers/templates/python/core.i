%ignore kuzzle;
%rename(QueryOptions) query_options;

%{
#define SWIG_FILE_WITH_INIT
%}
%include "exception.i"
%include "../../kcore.i"

%inline %{
  typedef struct {
    kuzzle* _kuzzle;
  } Kuzzle;
%}

%extend Kuzzle {
    Kuzzle(char* host, options *opts) {
        kuzzle *k = malloc(sizeof(kuzzle));
        kuzzle_new_kuzzle(k, host, "websocket", opts);

        Kuzzle *K = calloc(1, sizeof(Kuzzle));
        K->_kuzzle = k;
        return K;
    }
    Kuzzle(char* host) {
        kuzzle *k;
        k = malloc(sizeof(kuzzle));
        kuzzle_new_kuzzle(k, host, "websocket", NULL);

        Kuzzle *K = calloc(1, sizeof(Kuzzle));
        K->_kuzzle = k;
        return K;
    }
    ~Kuzzle() {
        unregisterKuzzle($self->_kuzzle);
        free($self->_kuzzle);
        free($self);
    }

    long long now(query_options* options=NULL) {
      int_result *result = kuzzle_now($self->_kuzzle, options);
      if (result->error != NULL) {
        PyErr_SetString(PyExc_ValueError, result->error);
        return -1;
      }

      return result->result;
    }
}
