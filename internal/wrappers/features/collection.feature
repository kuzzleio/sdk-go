Feature: Collection management

  Scenario: Create a collection
    Given Kuzzle Server is running
    And there is an index 'test-index'
    When I create a collection 'collection-test-collection'
    Then the collection should exists

  Scenario: Check if a collection exists
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    When I check if the collection exists
    Then it should exists

  Scenario: List existing collections
    Given Kuzzle Server is running
    And there is an index 'list-test-index'
    And it has a collection 'test-collection1'
    And it has a collection 'test-collection2'
    When I list the collections
    Then the result contains 2 hits
    And the content should not be null

  Scenario: Truncate a collection
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'my-document-id'
    And the collection has a document with id 'my-document-id2'
    When I truncate the collection
    Then the collection shall be empty

  Scenario: Create a collection with a custom mapping
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    When I update the mapping
    Then the mapping should be updated

  Scenario: Update specifications
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    When I update the specifications
    Then they should be updated

  Scenario: Validate specifications
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    When I validate the specifications
    Then they should be validated

  Scenario: Delete specifications
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And has specifications
    When I delete the specifications
    Then the specifications must not exist