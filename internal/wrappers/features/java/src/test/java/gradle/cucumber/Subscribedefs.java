package gradle.cucumber;

import cucumber.api.java.After;
import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class Subscribedefs {
    private Kuzzle k;
    private NotificationContent content = null;

    @Before
    public void before() {
        k = new Kuzzle("localhost", null);
    }

    @After
    public void after() {
        k.disconnect();
    }

    @Given("^I subscribe to \"([^\"]*)\"$")
    public void i_subscribe_to(String arg1) throws Exception {
        k.getRealtime().subscribe("index", "collection", "{}", new NotificationListener() {
            @Override
            public void onMessage(NotificationResult res) {
                content = res.getResult();
            }
        });
    }

    @When("^I create a document in \"([^\"]*)\"$")
    public void i_create_a_document_in(String arg1) throws Exception {
        k.getAuth().login("local", "{\"username\": \""+"useradmin"+"\", \"password\":\""+"testpwd"+"\" }");

        k.getDocument().create("index", "collection", "", "{}");
    }

    @Then("^I receive a notification$")
    public void i_receive_a_notification() throws Exception {
        //Assert.assertEquals("foo", content.getContent());
    }
}
