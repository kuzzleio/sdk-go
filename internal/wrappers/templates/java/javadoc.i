%javamethodmodifiers kuzzle::kuzzle(char*) "
  /**
   * Constructor
   *
   * @param host - Target Kuzzle host name or IP address
   */
  public";

%javamethodmodifiers kuzzle::kuzzle(char*, options*) "
  /**
   * Constructor
   *
   * @param host - Target Kuzzle host name or IP address
   * @param options - Request options
   */
  public";

%javamethodmodifiers kuzzle::checkToken(char*) "
  /**
   * Check an authentication token validity
   *
   * @param token - Token to check (JWT)
   * @return a TokenValidity object
   */
  public";

%javamethodmodifiers kuzzle::connect() "
  /**
   * Connects to a Kuzzle instance using the provided host and port.
   *
   * @return a string which represent an error or null
   */
  public";

%javamethodmodifiers kuzzle::createIndex(char*, query_options*) "
  /**
   * Create a new data index
   *
   * @param index - index name to create
   * @param options - Request options
   * @return a BoolResult object
   */
  public";

%javamethodmodifiers kuzzle::createIndex(char*) "
  /**
   * {@link #createIndex(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::createMyCredentials(char*, json_object*, query_options*) "
  /**
   * Create credentials of the specified strategy for the current user.
   *
   * @param strategy - impacted strategy name
   * @param credentials - credentials to create
   * @param options - Request options
   * @return a JsonResult object
   */
  public";

%javamethodmodifiers kuzzle::createMyCredentials(char*, json_object*) "
  /**
   * {@link #createIndex(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::deleteMyCredentials(char*, query_options*) "
  /**
   * Delete credentials of the specified strategy for the current user.
   *
   * @param strategy- Name of the strategy to remove
   * @param options - Request options
   * @return a BoolResult object
   */
  public";

%javamethodmodifiers kuzzle::deleteMyCredentials(char*) "
  /**
   * {@link #deleteMyCredentials(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::getMyCredentials(char *strategy, query_options *options) "
  /**
   * Get credential information of the specified strategy for the current user.
   *
   * @param strategy - Strategy name to get
   * @param options - Request options
   * @return a JsonResult
   */
  public";

%javamethodmodifiers kuzzle::getMyCredentials(char *strategy) "
  /**
   * {@link #getMyCredentials(String, QueryOptions, ResponseListener)}
   */
  public";

%javamethodmodifiers kuzzle::updateMyCredentials(char *strategy, json_object* credentials, query_options *options) "
  /**
   * Update credentials of the specified strategy for the current user.
   *
   * @param strategy - Strategy name to update
   * @param credentials - Updated credentials content
   * @param options - Request options
   * @return a JsonResult
   */
  public";

%javamethodmodifiers kuzzle::updateMyCredentials(char *strategy, json_object* credentials) "
  /**
   * {@link #updateMyCredentials(String, JSONObject, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::validateMyCredentials(char *strategy, json_object* credentials, query_options* options) "
  /**
   * Validate credentials of the specified strategy for the current user.
   *
   * @param strategy - Strategy name to validate
   * @param credentials - Credentials content
   * @param options - Request options
   * @return a Bool result
   */
  public";

%javamethodmodifiers kuzzle::validateMyCredentials(char *strategy, json_object* credentials) "
  /**
   * {@link #validateMyCredentials(String, JSONObject, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::logout() "
  /**
   * Logout method
   *
   * @param listener - Response callback listener
   */
  public";

%javamethodmodifiers kuzzle::login(char*, json_object*, int) "
  /**
   * Log-in Strategy name to use for the authentication
   *
   * @param strategy - Strategy name to use for the authentication
   * @param credentials - Login credentials
   * @param expiresIn - Token expiration delay
   * @return StringResult
   */
  public";

%javamethodmodifiers kuzzle::login(char*, json_object*) "
  /**
   * Log-in Strategy name to use for the authentication
   *
   * @param strategy - Strategy name to use for the authentication
   * @param credentials - Login credentials
   */
  public";

%javamethodmodifiers kuzzle::login(char*) "
  /**
   * Log-in Strategy name to use for the authentication
   *
   * @param strategy - Strategy name to use for the authentication
   */
  public";

%javamethodmodifiers kuzzle::login(char*) "
  /**
   * Get all Kuzzle usage statistics frames
   *
   * @param options - Request options
   * @param listener - Response callback listener
   */
  public";

%javamethodmodifiers kuzzle::getAllStatistics(query_options*) "
  /**
   * Get all Kuzzle usage statistics frames
   *
   * @param options - Request options
   * @return a AllStatisticsResult
   */
  public";

%javamethodmodifiers kuzzle::getAllStatistics() "
  /**
   * {@link #getAllStatistics(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::getStatistics(unsigned long, query_options*) "
  /**
   * Get Kuzzle usage statistics
   *
   * @param options - Request options
   * @return a StatisticsResult
   */
  public";

%javamethodmodifiers kuzzle::getStatistics(unsigned long) "
  /**
   * {@link #getStatistics(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::getAutoRefresh(char*, query_options*) "
  /**
   * Gets the autoRefresh value for the provided data index name
   *
   * @param index - Data index name
   * @param options - Request options
   */
  public";

%javamethodmodifiers kuzzle::getAutoRefresh(char*) "
  /**
   * {@link #getAutoRefresh(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::getAutoRefresh() "
  /**
   * {@link #getAutoRefresh(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::getJwt() "
  /**
   * Authentication token getter
   *
   * @return a string which is the jwt
   */
  public";

%javamethodmodifiers kuzzle::getMyRights(query_options*) "
  /**
   * Gets the rights array for the currently logged user.
   *
   * @param options - Request options
   * @return a JsonResult
   */
  public";

%javamethodmodifiers kuzzle::getMyRights() "
  /**
   * {@link #getMyRights(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::getServerInfo(query_options*) "
  /**
   * Gets server info.
   *
   * @param options - Request options
   * #return a JsonResult
   */
  public";

%javamethodmodifiers kuzzle::getServerInfo() "
  /**
   * {@link #getServerInfo(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::listCollections(char *, query_options*) "
  /**
   * List data collections
   *
   * @param index - Parent data index name
   * @param options - Request options
   * @return a CollectionListResult
   */
  public";

%javamethodmodifiers kuzzle::listCollections(char *) "
  /**
   * {@link #listCollections(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::listCollections() "
  /**
   * {@link #listCollections(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::listIndexes(query_options*) "
  /**
   * List data indexes
   *
   * @param options - Request options
   */
  public";

%javamethodmodifiers kuzzle::listIndexes() "
  /**
   * {@link #listIndexes(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzle::disconnect() "
  /**
   * Disconnect from Kuzzle and invalidate this instance.
   * Does not fire a disconnected event.
   */
  public";

%javamethodmodifiers kuzzle::logout() "
  /**
   * Logout method
   */
  public";