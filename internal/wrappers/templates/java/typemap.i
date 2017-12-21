%rename (_now) kuzzleio::Kuzzle::now(query_options*);
%rename (_now) kuzzleio::Kuzzle::now();

%javamethodmodifiers kuzzleio::Kuzzle::now(query_options* options) "private";
%javamethodmodifiers kuzzleio::Kuzzle::now() "private";
%typemap(javacode) kuzzleio::Kuzzle %{
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
