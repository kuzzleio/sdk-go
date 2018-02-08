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

%javamethodmodifiers kuzzleio::Collection::createDocument(Document* document, const std::string& id, query_options* options) "
  /**
   * Create a new document in kuzzle
   *
   * @param document the document
   * @param id the document id
   * @param options  Request options
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Collection::createDocument(Document* document, const std::string& id) "
  /**
   * {@link #createDocument(Document document, String id, QueryOptions options)}
   */
  public";

%javamethodmodifiers kuzzleio::Collection::createDocument(Document* document) "
  /**
   * {@link #createDocument(Document document, String id, QueryOptions options)}
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
   * @return a lit of Documents
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
   * @return a list of all created or updated documents
   */
  public";

%javamethodmodifiers kuzzleio::Collection::mReplaceDocument(std::vector<Document*>& documents) "
  /**
   * {@link #mReplaceDocument(DocumentVector documents, QueryOptions options)}
   */
  public";