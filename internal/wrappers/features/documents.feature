Feature: Document management

  Scenario: Do not allow creating a document with an _id that already exist in the same collection
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'my-document-id'
    When I create a document with id 'my-document-id'
    Then I get an error with message 'Document already exists'

  Scenario: Create a document with create
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection doesn't have a document with id 'my-document-id'
    When I create a document with id 'my-document-id'
    Then the document is successfully created

  Scenario: Create a document with createOrReplace if it does not exists
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection doesn't have a document with id 'my-document-id'
    When I createOrReplace a document with id 'my-document-id'
    Then the document is successfully created

  Scenario: Replace a document if it exists
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    When I createOrReplace a document with id 'my-document-id'
    Then the document is successfully replaced

  Scenario: Replace a document
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'replace-my-document-id'
    When I replace a document with id 'replace-my-document-id'
    Then the document is successfully replaced

  Scenario: Delete a document
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'delete-my-document-id'
    When I delete the document with id 'delete-my-document-id'
    Then the document is successfully deleted

  Scenario: Update a document
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'update-my-document-id'
    When I update a document with id 'update-my-document-id'
    Then the document is successfully updated

  Scenario: Search a document by id and find it
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'search-my-document-id'
    When I search a document with id 'search-my-document-id'
    Then the document is successfully found

  Scenario: Search a document by id and don't find it
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'search-my-document-id'
    When I search a document with id 'fake-id'
    Then the document is not found