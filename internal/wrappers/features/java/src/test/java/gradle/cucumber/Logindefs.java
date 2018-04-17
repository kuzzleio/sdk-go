package gradle.cucumber;

import cucumber.api.java.Before;
import cucumber.api.java.After;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;

public class Logindefs {
    private Kuzzle k;

    @Before
    public void before() {
        k = new Kuzzle("localhost", null);

        // @todo do this with fixture
        try {
            k.getIndex().create("index");
            k.getCollection().create("index", "collection");
        } catch(Exception e) {}
    }

    @After
    public void after() {
        //@todo do this with fixture
        KuzzleRequest request = new KuzzleRequest();
        request.setController("collection");
        request.setAction("truncate");
        request.setIndex("index");
        request.setCollection("collection");
        k.query(request);

        request.setController("security");
        request.setAction("deleteUser");
        request.setId("useradmin-id");
        k.query(request);

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
        k.getDocument().create("index", "collection", id, "{\"foo\":\"bar\"}");
    }

    /*

    @Then("^I check if the document with id \"([^\"]*)\"( does not)? exists$")
    public void GetDocument(String id, String does_not_exists) {
        boolean raised = false;
        Collection c = new Collection(k, "collection", "index");
        try {
            c.fetchDocument(id);
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
        JsonObject query = new JsonObject()
                .put("username", user)
                .put("password", password);
        String jwt = k.login("local", query, 1000000);
        Assert.assertNotNull("Jwt should not be null after a right login", jwt);
    }

    @Then("^I check the JWT is( not)? null$")
    public void CheckJwt(String not) {
        if (not != null) {
            Assert.assertNotNull(k.getJwt());
        } else {
            Assert.assertNull(k.getJwt());
        }
    }

    @When("^I update my credentials with username \"([^\"]*)\" and \"([^\"]*)\" = \"([^\"]*)\"$")
    public void UpdateMyCredentials(String username, String key, String value) {
        JsonObject credentials = new JsonObject()
                .put("username", username)
                .put(key, value);
        try {
            k.updateMyCredentials("local", credentials);
        } catch(Exception e) {}
    }

    @Then("^I check my new credentials are( not)? valid with username \"([^\"]*)\", password \"([^\"]*)\" and \"([^\"]*)\" = \"([^\"]*)\"$")
    public void CheckMyCredentials(String not, String username, String password, String key, String value) {
        JsonObject credentials = new JsonObject()
                .put("username", username)
                .put("password", password)
                .put(key, value);
        if (not == null) {
            Assert.assertTrue(k.validateMyCredentials("local", credentials));
        } else {
            try {
                Assert.assertFalse(k.validateMyCredentials("local", credentials));
            } catch(UnauthorizedException e) {}
        }
    }

    @Given("^I logout$")
    public void Logout() {
        k.logout();
    }
    */
}
