%typemap(javaimports) kuzzleio::Collection "
/* The type Collection. */"

%javamethodmodifiers kuzzleio::Collection::Collection(Kuzzle *kuzzle, const std::string& collection, const std::string& index) "
  /**
   * Constructor
   *
   * @param kuzzle  Kuzzle instance
   * @param collection  Data collection name
   * @param index  Parent data index name
   */
  public";

%javamethodmodifiers kuzzleio::Collection::count(search_filters* filters, query_options* options) "
  /**
   * Returns the number of documents matching the provided set of filters.
   * There is a small delay between documents creation and their existence in our search layer,
   * usually a couple of seconds.
   * That means that a document that was just been created won’t be returned by this function
   *
   * @param filters  Search filters
   * @param options  Request options
   * @returns the number of documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::count(search_filters* filters) "
  /**
   * {@link #count(SearchFilters filters, QueryOptions options)}
   */
  public";


%javamethodmodifiers kuzzleio::Collection::deleteDocument(const std::string& id, query_options* options) "
  /**
   * Delete a single document
   *
   * @param id Document unique identifier
   * @param options  Request options
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Collection::deleteDocument(const std::string& id) "
  /**
   * {@link #deleteDocument(String id, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::fetchDocument(const std::string& id, query_options* options) "
  /**
   * Fetch a document from Kuzzle
   *
   * @param id  Document unique identifier
   * @param options  Request options
   */
  public";

%javamethodmodifiers kuzzleio::Collection::fetchDocument(const std::string& id) "
  /**
   * {@link #fetchDocument(String id, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mCreateDocument(std::vector<Document*>& documents, query_options* options) "
  /**
   * Create multiple documents
   *
   * @param documents  List of Document objects to create
   * @param options  Request options
   * @return a list of all document created 
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mCreateDocument(std::vector<Document*>& documents) "
  /**
   * {@link #mCreateDocument(DocumentVector documents, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mCreateOrReplaceDocument(std::vector<Document*>& documents, query_options* options) "
  /**
   * Create or replace multiple documents
   *
   * @param documents  Array of Document objects to create or replace
   * @param options  Request options
   * @return a list of all created or updated documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mCreateOrReplaceDocument(std::vector<Document*>& documents) "
  /**
   * {@link #mCreateOrReplaceDocument(DocumentVector documents, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mDeleteDocument(std::vector<std::string>& ids, query_options* options) "
  /**
   * Delete multiple documents using their unique IDs
   *
   * @param ids  Array of document IDs to delete
   * @param options  Request options
   * @return a list of all deleted ids's documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mDeleteDocument(std::vector<std::string>& ids) "
  /**
   * {@link #mDeleteDocument(StringVector ids, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mGetDocument(std::vector<std::string>& ids, query_options* options) "
  /**
   * Fetch multiple documents
   *
   * @param ids  Array of document IDs to get
   * @param options  Request options
   * @return a list of Documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mGetDocument(std::vector<std::string>& ids) "
  /**
   * {@link #mGetDocument(StringVector ids, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mReplaceDocument(std::vector<Document*>& documents, query_options* options) "
  /**
   * Replace multiple documents
   *
   * @param documents  Array of Document objects to replace
   * @param options  Request options
   * @return a list of all updated documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mReplaceDocument(std::vector<Document*>& documents) "
  /**
   * {@link #mReplaceDocument(DocumentVector documents, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mUpdateDocument(std::vector<Document*>& documents, query_options* options) "
  /**
   * Update multiple documents
   *
   * @param documents  Array of Document objects to replace
   * @param options  Request options
   * @return a list of all updated documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mUpdateDocument(std::vector<Document*>& documents) "
  /**
   * {@link #mUpdateDocument(DocumentVector documents, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::publishMessage(json_object* content, query_options* options) "
  /**
   * Publish a real-time message
   *
   * @param content  Content to publish
   * @param options  Request options
   * @return boolean
   */
  public";

%javamethodmodifiers kuzzleio::Collection::publishMessage(json_object* content) "
  /**
   * {@link #publishMessage(JsonObject content, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::replaceDocument(const std::string& id, Document* document, query_options* options) "
  /**
   * Replace an existing document with a new one.
   *
   * @param id  Document unique identifier
   * @param document  New document
   * @param options  Request options
   * @return the document
   */
  public";

%javamethodmodifiers kuzzleio::Collection::replaceDocument(const std::string& id, Document* document) "
  /**
   * {@link #replaceDocument(String id, Document document, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::scroll(const std::string& id, query_options* options) "
  /**
   * Gets the next page of results from a previous search or scroll request
   * 
   * @param id  Scroll unique identifier
   * @param options  Request options
   * @returns a SearchResult
   */
  public";

%javamethodmodifiers kuzzleio::Collection::scroll(const std::string& id) "
  /**
   * {@link #scroll(String id, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::search(search_filters* filters, query_options* options) "
  /**
   * Executes a search on the data collection.
   * /!\ There is a small delay between documents creation and their existence in our search layer,
   * usually a couple of seconds.
   * That means that a document that was just been created won’t be returned by this function.
   *
   * @param filters  Search filters to apply
   * @param options  Request options
   * @returns a SearchResult
   */
  public";

%javamethodmodifiers kuzzleio::Collection::search(search_filters* filters) "
  /**
   * {@link #search(SearchFilters filters, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::subscribe(search_filters* filters, NotificationListener *listener, room_options* options) "
  /**
   * Subscribes to this data collection with a set of Kuzzle DSL filters.
   *
   * @param filters  Subscription filters
   * @param options  Request options
   * @param listener  Response callback listener
   * @return an object with a onDone() callback triggered when the subscription is active
   */
  public";

%javamethodmodifiers kuzzleio::Collection::subscribe(search_filters* filters, NotificationListener *listener) "
  /**
   * {@link #subscribe(SearchFilters filters, NotificationListener listeners)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::updateDocument(const std::string& id, Document *document, query_options* options) "
  /**
   * Update parts of a document
   *
   * @param id  Document unique identifier
   * @param content  Document content to update
   * @param options  Request options
   * @return the document
   */
  public";

%javamethodmodifiers kuzzleio::Collection::updateDocument(const std::string& id, Document *document) "
  /**
   * {@link #updateDocument(String id, Document document)}
   */
  public";
