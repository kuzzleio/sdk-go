%typemap(javaimports) kuzzleio::Document "
/* The type Document. */"

%javamethodmodifiers kuzzleio::Document::Document(Collection *collection, const std::string& id, json_object* content) "
  /**
   * Kuzzle handles documents either as real-time messages or as stored documents.
   * Document is the object representation of one of these documents.
   *
   * @param collection           - An instantiated Collection object
   * @param id                   - Unique document identifier
   * @param content              - The content of the document
   */
  public";

%javamethodmodifiers kuzzleio::Document::Document(Collection *collection, const std::string& id) "
  /**
   * Kuzzle handles documents either as real-time messages or as stored documents.
   * Document is the object representation of one of these documents.
   *
   * @param collection           - An instantiated Collection object
   * @param id                   - Unique document identifier
   */
  public";

%javamethodmodifiers kuzzleio::Document::Document(Collection *collection) "
  /**
   * Kuzzle handles documents either as real-time messages or as stored documents.
   * Document is the object representation of one of these documents.
   *
   * @param collection - An instantiated Collection object
   */
  public";


%javamethodmodifiers kuzzleio::Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body="", query_options *options) "
  /**
   * Create a new document in kuzzle
   *
   * @param index the index where to create the document
   * @param collection the collection where to create the document
   * @param id the document id
   * @param body the content of the document
   * @param options  Request options
   * @return document id
   */
  public";

%javamethodmodifiers kuzzleio::Collection::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body="") "
  /**
   * {@link #create(String index, String collection, String id, String body, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::delete_(query_options *) "
  /**
   * Delete this document from Kuzzle
   *
   * @param options - Request options
   * @return string - id of the deleted document
   */
  public";

%javamethodmodifiers kuzzleio::Document::delete_() "
  /**
   * {@link #delete_(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::exists(query_options *) "
  /**
   * Ask Kuzzle if this document exists
   * 
   * @param options - Request options
   * @return bool
   */
  public";

%javamethodmodifiers kuzzleio::Document::exists() "
  /**
   * {@link #exists(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::publish(query_options *) "
  /**
   * Sends the content of this document as a real-time message.
   *
   * @param options - Request options
   * @return bool
   */
  public";

%javamethodmodifiers kuzzleio::Document::publish() "
  /**
   * {@link #publish(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::refresh(query_options* options) "
  /**
   * Gets a refreshed copy of this document from Kuzzle
   *
   * @param options - Request options
   * @return Document - refreshed document
   */
  public";

%javamethodmodifiers kuzzleio::Document::refresh() "
  /**
   * {@link #refresh(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::save(query_options* options) "
  /**
   * Saves this document into Kuzzle.
   * If this is a new document, this function will create it in Kuzzle and the id property will be made available.
   * Otherwise, this method will replace the latest version of this document in Kuzzle by the current content of this object.
   *
   * @param options - Request options
   * @return Document - saved document
   */
  public";

%javamethodmodifiers kuzzleio::Document::save() "
  /**
   * {@link #save(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::setContent(json_object* content, bool replace) "
  /**
   * Sets this document content
   *
   * @param content - New content for this document
   * @param replace - true: replace the current content, false (default): update/append it
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Document::setContent(json_object* content) "
  /**
   * {@link #setContent(JsonObject, boolean)}
   */
  public";

%javamethodmodifiers kuzzleio::Document::getContent() "
  /**
   * Document content getter
   *
   * @return current document content
   */
  public";

%javamethodmodifiers kuzzleio::Document::subscribe(NotificationListener* listener, room_options* options) "
  /**
   * Subscribe to changes occuring on this document.
   * Throws an error if this document has not yet been created in Kuzzle.
   *
   * @param options - Room object constructor options
   * @param listener - Response callback listener
   * @return a Room
   */
  public";

%javamethodmodifiers kuzzleio::Document::subscribe(NotificationListener* cb) "
  /**
   * {@link #subscribe(NotificationListener listener, room_options options)}
   */
  public";
