# SDK's integration tests
#========================

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

  #Feature: User management

  Scenario Outline: Get a valid JWT when login
    Given Kuzzle Server is running
    And there is an user with id 'my-user-id'
    And the user has 'local' credentials with name 'my-user-name' and password 'my-user-pwd'
    When I log in as <user-name>:<user-pwd>
    Then the JWT is <jwt_validity>
    Examples:
      | user-name        | user-pwd        | jwt_validity |
      | 'my-user-name'   | 'my-user-pwd'   | valid        |
      | 'my-user-name-w' | 'my-user-pwd'   | invalid      |
      | 'my-user-name'   | 'my-user-pwd-w' | invalid      |
      | 'my-user-name-w' | 'my-user-pwd-w' | invalid      |

  Scenario Outline: Set user custom data (updateSelf)
    Given Kuzzle Server is running
    And there is an user with id 'my-user-id'
    And the user has 'local' credentials with name 'my-user-name' and password 'my-user-pwd'
    And I log in as 'my-user-name':'my-user-pwd'
    And I update my user custom data with the pair <field-name>:<field-value>
    When I get my user info
    Then the response '_source' field contains the pair <field-name>:<field-value>
    Examples:
      | field-name | field-value |
      | 'my_data' | 'mystringvalue' |
      
  Scenario: Login out shall revoke the JWT
    Given Kuzzle Server is running
    And there is an user with id 'my-user-id'
    And the user has 'local' credentials with name 'my-user-name' and password 'my-user-pwd'
    And I log in as 'my-user-name':'my-user-pwd'
    Then the JWT is valid
    When I logout
    Then the JWT is invalid

  #Feature: Realtime subscription

  Scenario: Receive notifications when a document is created
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And I subscribe to 'test-collection'
    When I create a document in "test-collection"
    Then I receive a notification

  Scenario: Receive notifications when a document is updated
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'my-document-id'
    And I subscribe to 'test-collection' with filter '{"equals": {"foo": "barz"}}'
    When I update the document with id 'my-document-id' and content 'foo' = 'barz'
    Then I receive a notification

  Scenario: Receive notifications when a document is deleted
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'my-document-id'
    And I subscribe to 'test-collection'
    When I delete the document with id 'my-document-id'
    Then I receive a notification

  Scenario: Receive notifications when a document is published
    Given Kuzzle Server is running
    And I subscribe to 'test-collection'
    When I publish a document
    Then I receive a notification

  Scenario: Stop receiving notification when I unsubscribe
    Given Kuzzle Server is running
    And I subscribe to 'test-collection'
    And the collection has a document with id 'my-document-id'
    And I received a notification
    And I unsubscribe from 'test-collection'
    When I publish a document
    Then I do not receive a notification

# Feature: Geofencing subscriptions

# Scenario Outline: Subscribe to a geofence filter
# Given I subscribe to "geofence" collection
# When I create a document in "geofence" and inside my geofence
# Then I receive a "in" notification
#
# When I update this document to be outside my geofence
# Then I receive a "out" notification
# Examples:
# | lon     | lat    | in  |
# | 43.6108 | 3.8767 | yes |
