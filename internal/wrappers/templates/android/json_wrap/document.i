%javamethodmodifiers kuzzleio::Document::count(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::count(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::createOrReplace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::createOrReplace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::deleteByQuery(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::deleteByQuery(const std::string& index, const std::string& collection, const std::string& body) "private";
%rename (_get) kuzzleio::Document::get(const std::string& index, const std::string& collection, const std::string& id, query_options *options);
%rename (_get) kuzzleio::Document::get(const std::string& index, const std::string& collection, const std::string& id);
%javamethodmodifiers kuzzleio::Document::get(const std::string& index, const std::string& collection, const std::string& id, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::get(const std::string& index, const std::string& collection, const std::string& id) "private";
%javamethodmodifiers kuzzleio::Document::replace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::replace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::update(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::update(const std::string& index, const std::string& collection, const std::string& id, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::validate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::validate(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::search(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::search(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::mCreate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::mCreate(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::mCreateOrReplace(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::mCreateOrReplace(const std::string& index, const std::string& collection, const std::string& body) "private";
%rename (_mGet) kuzzleio::Document::mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash, query_options *options);
%rename (_mGet) kuzzleio::Document::mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash);
%javamethodmodifiers kuzzleio::Document::mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash) "private";
%javamethodmodifiers kuzzleio::Document::mReplace(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::mReplace(const std::string& index, const std::string& collection, const std::string& body) "private";
%javamethodmodifiers kuzzleio::Document::mUpdate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) "private";
%javamethodmodifiers kuzzleio::Document::mUpdate(const std::string& index, const std::string& collection, const std::string& body) "private";

%typemap(javacode) kuzzleio::Document %{

  public int count(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return count(index, collection, body.toString(), options);
  }

  public int count(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return count(index, collection, body, null);
  }

  public org.json.JSONObject create(String index, String collection, String id, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = create(index, collection, id, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject create(String index, String collection, String id, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return create(index, collection, id, body, null);
  }

  public org.json.JSONObject createOrReplace(String index, String collection, String id, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = createOrReplace(index, collection, id, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject createOrReplace(String index, String collection, String id, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return createOrReplace(index, collection, id, body, null);
  }

  public StringVector deleteByQuery(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return deleteByQuery(index, collection, body.toString(), options);
  }

  public StringVector deleteByQuery(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return deleteByQuery(index, collection, body, null);
  }

  public org.json.JSONObject get(String index, String collection, String id, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _get(index, collection, id, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject get(String index, String collection, String id) throws org.json.JSONException, KuzzleException {
    return get(index, collection, id, null);
  }

  public org.json.JSONObject replace(String index, String collection, String id, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = replace(index, collection, id, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject replace(String index, String collection, String id, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return replace(index, collection, id, body, null);
  }

  public org.json.JSONObject update(String index, String collection, String id, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = update(index, collection, id, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject update(String index, String collection, String id, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return update(index, collection, id, body, null);
  }

  public boolean validate(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return validate(index, collection, body.toString(), options);
  }

  public boolean validate(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return validate(index, collection, body, null);
  }

  public SearchResult search(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    return search(index, collection, body.toString(), options);
  }

  public SearchResult search(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return search(index, collection, body, null);
  }

  public org.json.JSONObject mCreate(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = mCreate(index, collection, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject mCreate(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return mCreate(index, collection, body, null);
  }

  public org.json.JSONObject mCreateOrReplace(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = mCreateOrReplace(index, collection, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject mCreateOrReplace(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return mCreateOrReplace(index, collection, body, null);
  }

  public org.json.JSONObject mGet(String index, String collection, StringVector ids, boolean includeTrash, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = _mGet(index, collection, ids, includeTrash, options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject mGet(String index, String collection, StringVector ids, boolean includeTrash) throws org.json.JSONException, KuzzleException {
    return mGet(index, collection, ids, includeTrash, null);
  }

  public org.json.JSONObject mReplace(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = mReplace(index, collection, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject mReplace(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return mReplace(index, collection, body, null);
  }

  public org.json.JSONObject mUpdate(String index, String collection, org.json.JSONObject body, QueryOptions options) throws org.json.JSONException, KuzzleException {
    String res = mUpdate(index, collection, body.toString(), options);

    return new org.json.JSONObject(res);
  }

  public org.json.JSONObject mUpdate(String index, String collection, org.json.JSONObject body) throws org.json.JSONException, KuzzleException {
    return mUpdate(index, collection, body, null);
  }

%}
