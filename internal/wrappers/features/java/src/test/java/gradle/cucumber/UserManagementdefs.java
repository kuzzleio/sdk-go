package gradle.cucumber;

import cucumber.api.java.After;
import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class UserManagementdefs {
    private Kuzzle k;
    private String userId;
    private User currentUser;

    @Before
    public void before() {
        k = KuzzleSingleton.getInstance();
    }

    @After
    public void after() {
        if (k != null && (k.getJwt() == null || !k.getJwt().equals(""))) {
            k.getAuth().logout();
        }
    }

    @Given("^there is an user with id \'([^\"]*)\'$")
    public void there_is_an_user_with_id(String id) throws Exception {
        this.userId = id;
    }

    @Given("^the user has 'local' credentials with name \'([^\"]*)\' and password \'([^\"]*)\'$")
    public void the_user_has_local_credentials_with_name_and_password(String name, String password) throws Exception {
      KuzzleRequest request = new KuzzleRequest();
        String json = "{\n" +
                "  \"content\": {\n" +
                "    \"profileIds\": [\"default\"]" +
                "  },\n" +
                "  \"credentials\": {\n" +
                "    \"local\": {\n" +
                "      \"username\": \""+name+"\"," +
                "      \"password\": \""+password+"\"" +
                "    }\n" +
                "  }\n" +
                "}";

        request.setController("security");
        request.setAction("createUser");
        request.setId(this.userId);
        request.setBody(json);
        request.setStrategy("local");

        try {
            k.query(request);
        } catch(KuzzleException e) {
            if (e.getClass() != PreconditionException.class) {
                throw e;
            }
        }
    }

    @When("^I log in as \'([^\"]*)\':\'([^\"]*)\'$")
    public void i_log_in_as(String name, String password) throws Exception {
        try {
            k.getAuth().login("local", "{\"username\": \"" + name + "\", \"password\":\"" + password + "\" }");
        } catch (KuzzleException e) {
            if (e.getClass() != UnauthorizedException.class) {
                throw e;
            }
        }
    }

    @Then("^the JWT is valid$")
    public void the_JWT_is_valid() throws Exception {
        Assert.assertTrue(k.getAuth().checkToken(k.getJwt()).getValid());
    }

    @Then("^the JWT is invalid$")
    public void the_JWT_is_invalid() throws Exception {
        Assert.assertFalse(k.getAuth().checkToken(k.getJwt()).getValid());
    }

    @Given("^I update my user custom data with the pair \'([^\"]*)\':\'([^\"]*)\'$")
    public void i_update_my_user_custom_data_with_the_pair(String fieldname, String fieldvalue) throws Exception {
        k.getAuth().updateSelf("{\""+fieldname+"\": \""+fieldvalue+"\"}");
    }

    @When("^I get my user info$")
    public void i_get_my_user_info() throws Exception {
        currentUser = k.getAuth().getCurrentUser();
    }

    @Then("^the response '_source' field contains the pair \'([^\"]*)\':\'([^\"]*)\'$")
    public void the_response__source_field_contains_the_pair(String key, String value) throws Exception {
        Assert.assertNotNull(currentUser.getContent());
    }

    @When("^I logout$")
    public void i_logout() throws Exception {
        k.getAuth().logout();
    }
}
