%rename(TokenValidity) token_validity;
%rename(AckResponse) ack_response;
%rename(queueTTL) queue_ttl;
%rename(Options) options;
%rename(Kuzzle) kuzzle;
%rename(JsonObject) json_object;
%rename(JsonResult) json_result;
%rename(LoginResult) login_result;
%rename(BoolResult) bool_result;
%rename(Statistics) statistics;
%rename(AllStatisticsResult) all_statistics_result;
%rename(StatisticsResult) statistics_result;

%include "typemap.i"
%include "../../kcore.i"
//
//if (strcmp(JSON_C_VERSION, "0.12.99")) {
//    printf("You version of json-c is not equal to 0.12.99, please ensure to have the right version\n");
//    exit(1);
//}


%pragma(java) jniclasscode=%{
  static {
    try {
        System.loadLibrary("kuzzle");
    } catch (UnsatisfiedLinkError e) {
      System.err.println("Native code library failed to load. \n" + e);
      System.exit(1);
    }
  }
%}

%extend options {
    options() {
        options *o = kuzzle_wrapper_new_options();
        return o;
    }

    ~options() {
        free($self);
    }
}

struct json_object { };

%ignore json_object::ptr;
%extend json_object {
    json_object() {
        json_object *j = json_object_new_object();
        return j;
    }

    ~json_object() {
        free($self);
    }

    json_object* put(char* key, char* content) {
        kuzzle_wrapper_json_put($self, key, content, 0);
        return $self;
    }

    json_object* put(char* key, int content) {
        kuzzle_wrapper_json_put($self, key, &content, 1);
        return $self;
    }

    json_object* put(char* key, double content) {
        kuzzle_wrapper_json_put($self, key, &content, 2);
        return $self;
    }

    json_object* put(char* key, bool content) {
        kuzzle_wrapper_json_put($self, key, &content, 3);
        return $self;
    }

    json_object* put(char* key, json_object* content) {
        kuzzle_wrapper_json_put($self, key, content, 4);
        return $self;
    }

    char* getString(char* key) {
        return kuzzle_wrapper_json_get_string($self, key);
    }

    int getInt(char* key) {
        return kuzzle_wrapper_json_get_int($self, key);
    }

    double getDouble(char* key) {
        return kuzzle_wrapper_json_get_double($self, key);
    }

    bool getBoolean(char* key) {
        return kuzzle_wrapper_json_get_bool($self, key);
    }

    json_object* getJsonObject(char* key) {
        return kuzzle_wrapper_json_get_json_object($self, key);
    }
}

%typemap(javaimports) kuzzle "
/* The type Kuzzle. */"

%extend kuzzle {
    // ctors && dtor
    kuzzle(char* host, options *opts) {
        kuzzle *k = malloc(sizeof(kuzzle));
        kuzzle_wrapper_new_kuzzle(k, host, "websocket", opts);
        return k;
    }
    kuzzle(char* host) {
        kuzzle *k;
        k = malloc(sizeof(kuzzle));
        kuzzle_wrapper_new_kuzzle(k, host, "websocket", NULL);
        return k;
    }
    ~kuzzle() {
        unregisterKuzzle($self);
        free($self);
    }

    // checkToken
    token_validity* checkToken(char* token) {
        return kuzzle_wrapper_check_token($self, token);
    }

    // connect
    char* connect() {
        return kuzzle_wrapper_connect($self);
    }

    // createIndex
    bool_result* createIndex(char* index, query_options* options) {
        return kuzzle_wrapper_create_index($self, index, options);
    }
    bool_result* createIndex(char* index) {
        return kuzzle_wrapper_create_index($self, index, NULL);
    }

    // createMyCredentials
    json_result* createMyCredentials(char* strategy, json_object* credentials, query_options* options) {
        return kuzzle_wrapper_create_my_credentials($self, strategy, credentials, options);
    }
    json_result* createMyCredentials(char* strategy, json_object* credentials) {
        return kuzzle_wrapper_create_my_credentials($self, strategy, credentials, NULL);
    }

    // deleteMyCredentials
    bool_result* deleteMyCredentials(char* strategy, query_options *options) {
        return kuzzle_wrapper_delete_my_credentials($self, strategy, options);
    }
    bool_result* deleteMyCredentials(char* strategy) {
        return kuzzle_wrapper_delete_my_credentials($self, strategy, NULL);
    }

    // getMyCredentials
    json_result* getMyCredentials(char *strategy, query_options *options) {
        return kuzzle_wrapper_get_my_credentials($self, strategy, options);
    }
    json_result* getMyCredentials(char *strategy) {
        return kuzzle_wrapper_get_my_credentials($self, strategy, NULL);
    }

    // updateMyCredentials
    json_result* updateMyCredentials(char *strategy, json_object* credentials, query_options *options) {
        return kuzzle_wrapper_update_my_credentials($self, strategy, credentials, options);
    }
    json_result* updateMyCredentials(char *strategy, json_object* credentials) {
        return kuzzle_wrapper_update_my_credentials($self, strategy, credentials, NULL);
    }

    // validateMyCredentials
    bool_result* validateMyCredentials(char *strategy, json_object* credentials, query_options* options) {
        return kuzzle_wrapper_validate_my_credentials($self, strategy, credentials, options);
    }
    bool_result* validateMyCredentials(char *strategy, json_object* credentials) {
        return kuzzle_wrapper_validate_my_credentials($self, strategy, credentials, NULL);
    }

    // login
    string_result* login(char* strategy, json_object* credentials, int expires_in) {
        return kuzzle_wrapper_login($self, strategy, credentials, &expires_in);
    }
    string_result* login(char* strategy, json_object* credentials) {
        return kuzzle_wrapper_login($self, strategy, credentials, NULL);
    }

    // getAllStatistics
    all_statistics_result* getAllStatistics(query_options* options) {
        return kuzzle_wrapper_get_all_statistics($self, options);
    }
    all_statistics_result* getAllStatistics() {
        return kuzzle_wrapper_get_all_statistics($self, NULL);
    }

    // getStatistics
    statistics_result* getStatistics(unsigned long time, query_options* options) {
        return kuzzle_wrapper_get_statistics($self, time, options);
    }
    statistics_result* getStatistics(unsigned long time) {
        return kuzzle_wrapper_get_statistics($self, time, NULL);
    }

    // getAutoRefresh
    bool_result* getAutoRefresh(char* index, query_options* options) {
        return kuzzle_wrapper_get_auto_refresh($self, index, options);
    }
    bool_result* getAutoRefresh(char* index) {
        return kuzzle_wrapper_get_auto_refresh($self, index, NULL);
    }

    // getJwt
    char* getJwt() {
        return kuzzle_wrapper_get_jwt($self);
    }

    // getMyRights
    json_result* getMyRights(query_options* options) {
        return kuzzle_wrapper_get_my_rights($self, options);
    }
    json_result* getMyRights() {
        return kuzzle_wrapper_get_my_rights($self, NULL);
    }

    // getServerInfo
    json_result* getServerInfo(query_options* options) {
        return kuzzle_wrapper_get_server_info($self, options);
    }
    json_result* getServerInfo() {
        return kuzzle_wrapper_get_server_info($self, NULL);
    }
}
