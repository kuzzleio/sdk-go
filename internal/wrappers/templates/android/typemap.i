%rename (_now) kuzzleio::Server::now(query_options*);
%rename (_now) kuzzleio::Server::now();

%javamethodmodifiers kuzzleio::Server::now(query_options* options) "private";
%javamethodmodifiers kuzzleio::Server::now() "private";
%typemap(javacode) kuzzleio::Server %{
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
   * @return a DateResult
   */
  public java.util.Date now(QueryOptions options) {
    long res = _now(options);

    return new java.util.Date(res);
  }
%}

%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body) "private";
%typemap(javacode) kuzzleio::Document %{

  public org.json.JSONObject create(String index, String collection, String id, org.json.JSONObject body, QueryOptions options) throws Exception {
    String res = create(index, collection, id, body.toString());

    return new org.json.JSONObject(res);
  }

%}
