Feature: Create document with ids

  Scenario: Do not allow creating a document with an _id that already exist in the same collection
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'my-document-id'
    When I try to create a new document with id 'my-document-id'
    Then I get an error with message 'Document already exists'

  Scenario: Allow creating a document with an _id when the _id isn't used yet
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection doesn't have a document with id 'my-document-id'
    When I try to create a new document with id 'my-document-id'
    Then the document is successfully created
