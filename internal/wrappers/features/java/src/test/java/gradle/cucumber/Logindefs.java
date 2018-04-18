package gradle.cucumber;

import cucumber.api.java.Before;
import cucumber.api.java.After;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class Logindefs {
    private Kuzzle k;

    @Before
    public void before() {
        k = new Kuzzle("localhost", null);
    }

    @After
    public void after() {
        k.disconnect();
    }

    @Given("^I create a user \"([^\"]*)\" with password \"([^\"]*)\" with id \"([^\"]*)\"$")
    public void CreateUser(String user, String password, String id) {
        String json = "{\n" +
                "  \"content\": {\n" +
                "    \"profileIds\": [\"default\"]" +
                "  },\n" +
                "  \"credentials\": {\n" +
                "    \"local\": {\n" +
                "      \"username\": \""+user+"\"," +
                "      \"password\": \""+password+"\"" +
                "    }\n" +
                "  }\n" +
                "}";

        KuzzleRequest request = new KuzzleRequest();
        request.setController("security");
        request.setAction("createUser");
        request.setBody(json);
        request.setId(id);

        k.query(request);
    }

    @When("^I try to create a document with id \"([^\"]*)\"$")
    public void CreateDocument(String id) {
        try {
            k.getDocument().create("index", "collection", id, "{\"foo\":\"bar\"}");
        } catch(Exception e) {}
    }

    @Then("^I check if the document with id \"([^\"]*)\"( does not)? exists$")
    public void GetDocument(String id, String does_not_exists) {
        boolean raised = false;
        try {
            k.getDocument().get("index", "collection", id);
        } catch (Exception e) {
            if (does_not_exists == null) {
                Assert.fail("Exception raised: " + e.toString());
            } else {
                raised = true;
            }
        }
        if (does_not_exists != null) {
            Assert.assertTrue("The document has been created and should not have been.", raised);
        }
    }

    @When("^I log in as \"([^\"]*)\":\"([^\"]*)\"$")
    public void Login(String user, String password) {
        String jwt = k.getAuth().login("local", "{\"username\": \""+user+"\", \"password\":\""+password+"\" }");
        Assert.assertNotNull("Jwt should not be null after a right login", jwt);
    }

    @Then("^I check the JWT is( not)? null$")
    public void CheckJwt(String not) {
        if (not != null) {
            Assert.assertNotEquals("", k.getJwt());
        } else {
            Assert.assertEquals("", k.getJwt());
        }
    }


    @When("^I update my credentials with username \"([^\"]*)\" and \"([^\"]*)\" = \"([^\"]*)\"$")
    public void UpdateMyCredentials(String username, String key, String value) {
        try {
            k.getAuth().updateMyCredentials("local", "{\"username\": \""+username+"\", \""+key+"\":\""+value+"\"}");
        } catch(Exception e) {}
    }

    @Then("^I check my new credentials are( not)? valid with username \"([^\"]*)\", password \"([^\"]*)\" and \"([^\"]*)\" = \"([^\"]*)\"$")
    public void CheckMyCredentials(String not, String username, String password, String key, String value) {
        String credentials = "{\"username\": \""+username+"\", \"password\":\""+password+"\", \""+key+"\":\""+value+"\"}";
        if (not == null) {
            Assert.assertTrue(k.getAuth().validateMyCredentials("local", credentials));
        } else {
            try {
                Assert.assertFalse(k.getAuth().validateMyCredentials("local", credentials));
            } catch(Exception e) {}
        }
    }

    @Given("^I logout$")
    public void Logout() {
        k.getAuth().logout();
    }
}
