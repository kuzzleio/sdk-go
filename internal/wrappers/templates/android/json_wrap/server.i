%rename (_getAllStats) kuzzleio::Server::getAllStats(query_options *options);
%rename (_getAllStats) kuzzleio::Server::getAllStats();
%javamethodmodifiers kuzzleio::Server::getAllStats(query_options *options) "private";
%javamethodmodifiers kuzzleio::Server::getAllStats() "private";
%rename (_getStats) kuzzleio::Server::getStats(time_t start, time_t end, query_options* options);
%rename (_getStats) kuzzleio::Server::getAllStats(time_t start, time_t end);
%javamethodmodifiers kuzzleio::Server::getStats(time_t start, time_t end, query_options* options) "private";
%javamethodmodifiers kuzzleio::Server::getStats(time_t start, time_t end) "private";
%rename (_getLastStats) kuzzleio::Server::getLastStats(query_options *options);
%rename (_getLastStats) kuzzleio::Server::getLastStats();
%javamethodmodifiers kuzzleio::Server::getLastStats(query_options *options) "private";
%javamethodmodifiers kuzzleio::Server::getLastStats() "private";
%rename (_getConfig) kuzzleio::Server::getConfig(query_options *options);
%rename (_getConfig) kuzzleio::Server::getConfig();
%javamethodmodifiers kuzzleio::Server::getConfig(query_options *options) "private";
%javamethodmodifiers kuzzleio::Server::getConfig() "private";
%rename (_info) kuzzleio::Server::info(query_options *options);
%rename (_info) kuzzleio::Server::info();
%javamethodmodifiers kuzzleio::Server::info(query_options *options) "private";
%javamethodmodifiers kuzzleio::Server::info() "private";
%rename (_now) kuzzleio::Server::now(query_options*);
%rename (_now) kuzzleio::Server::now();
%javamethodmodifiers kuzzleio::Server::now(query_options* options) "private";
%javamethodmodifiers kuzzleio::Server::now() "private";

%typemap(javacode) kuzzleio::Server %{
  public org.json.JSONObject getAllStats(QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _getAllStats(options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getAllStats() throws org.json.JSONException, KuzzleException {
    return getAllStats(null);
  }

  public org.json.JSONObject getStats(java.util.Date start, java.util.Date end, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _getStats(start.getTime(), end.getTime(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getStats(java.util.Date start, java.util.Date end) throws org.json.JSONException, KuzzleException {
    return getStats(start, end, null);
  }

  public org.json.JSONObject getLastStats(QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _getLastStats(options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getLastStats() throws org.json.JSONException, KuzzleException {
    return getLastStats(null);
  }

  public org.json.JSONObject getConfig(QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _getConfig(options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject getConfig() throws org.json.JSONException, KuzzleException {
    return getConfig(null);
  }

  public org.json.JSONObject info(QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _info(options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject info() throws org.json.JSONException, KuzzleException {
    return info(null);
  }

  /**
   * {@link #now(QueryOptions)}
   */
  public java.util.Date now() {
    long res = _now();

    return new java.util.Date(res);
  }

  /**
   * Returns the current Kuzzle UTC timestamp
   *
   * @param options - Request options
   * @return a Date
   */
  public java.util.Date now(QueryOptions options) {
    long res = _now(options);

    return new java.util.Date(res);
  }
%}