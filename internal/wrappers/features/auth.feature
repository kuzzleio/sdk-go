Feature: User management

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
    Then the response '_source' field contains the pair <fieldname>: <fieldvalue>
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

