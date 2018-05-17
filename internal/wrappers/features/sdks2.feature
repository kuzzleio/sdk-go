# SDK's integration tests
#========================

Feature: Create document with ids

  Scenario: Do not allow creating a document with an _id that already exist in the same collection
    Given Kuzzle Server is running
    And there is an index 'test-index'
    And it has a collection 'test-collection'
    And the collection has a document with id 'my-document-id'
    When I try to create a new document with id 'my-document-id'
    Then I get an error with message 'document alread exists'

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
    And I update my user custom data with the pair <fieldname> : <fieldvalue>
    When I get my user info
    Then the response '_source' field contains the pair <fieldname>: <fieldname>
    And is a <type>
    Examples:
      | fieldname | fieldvalue      | type   |
      | my_data1  | "mystringvalue" | string |
      | my_data2  | 1234            | number |
      | my_data2  | -1234           | number |
      | my_data2  | 1.234           | number |
      | my_data2  | -1.234          | number |
      | my_data1  | true            | bool   |
      | my_data1  | false           | bool   |


  Scenario: Login out shall revoke the JWT
    Given Kuzzle Server is running
    And there is an user with id 'my-user-id'
    And the user has 'local' credentials with name 'my-user-name' and password 'my-user-pwd'
    And I login using 'local' authentication, with <user-name> and password <user-pwd> as credentials
    And the retrieved JWT is valid
    When I logout
    Then the JWT is no more valid



# Feature: Realtime subscribtion
#
# Scenario: Receive notifications when a document is created
# Given I subscribe to "collection"
# When I create a document in "collection"
# Then I receive a notification
#
# Scenario: Receive notifications when a document is deleted
# Scenario: Receive notifications when a document is updated
# Scenario: Receive notifications when a document is published
# When I create a document in "collection"
# Then I receive a notification

# Scenario: Stop recieving notification when I unsubscribe
# Given I unsubscribe
# When I create a document in "collection"
# Then I do not receive a notification

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
