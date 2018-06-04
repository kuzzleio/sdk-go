Feature: Index controller

  Scenario: Create should create an index
    Given Kuzzle Server is running
    When I create an index called 'test-index'
    Then the index should exists

  Scenario: Create should return an error when the index already exists
    Given Given Kuzzle Server is running
    And there is an index 'test-index'
    When I create an index called 'test-index'
    Then I get an error