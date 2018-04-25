Feature: SDK's integration tests

  Scenario: Login user
    Given I create a user "useradmin" with password "testpwd" with id "useradmin-id"
    When I try to create a document with id "my-document"
    Then I check if the document with id "my-document" does not exists
    
    When I log in as "useradmin":"testpwd"
    Then I check the JWT is not null
    
    When I update my credentials with username "useradmin" and "foo" = "bar"
    Then I check my new credentials are valid with username "useradmin", password "testpwd" and "foo" = "bar"
    
    When I try to create a document with id "my-document"
    Then I check if the document with id "my-document" exists

    Given I logout
    When I update my credentials with username "useradmin" and "foo" = "barz"
    Then I check my new credentials are not valid with username "useradmin", password "testpwd" and "foo" = "barz"
    Then I check the JWT is null

  Scenario: Subscribe to a collection and receive notifications
    Given I subscribe to "collection"
    When I create a document in "collection"
    Then I receive a notification

    When I create a document in "collection"
    Then I receive a notification

    Given I unsubscribe
    When I create a document in "collection"
    Then I do not receive a notification

  Scenario: Subscribe to a geofence filter
    Given I subscribe to "geofence" collection
    When I create a document in "geofence" and inside my geofence
    Then I receive a "in" notification

    When I update this document to be outside my geofence
    Then I receive a "out" notification