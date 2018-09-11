Feature: Index controller

  Scenario: Create should create an index
    Given Kuzzle Server is running
    And there is no index called 'test-index'
    When I create an index called 'test-index'
    Then the index should exist

  Scenario: Create should return an error when the index already exist
    Given Kuzzle Server is running
    And there is an index 'test-index'
    When I create an index called 'test-index'
    Then I get an error

  Scenario: Delete multiple indexes
    Given Kuzzle Server is running
    And there is the indexes 'test-index1' and 'test-index2'
    When I delete the indexes 'test-index1' and 'test-index2'
    Then indexes 'test-index1' and 'test-index2' don't exist

  Scenario: List indexes
    Given Kuzzle Server is running
    And there is the indexes 'test-index1' and 'test-index2'
    When I list indexes
    Then I get 'test-index1' and 'test-index2'
