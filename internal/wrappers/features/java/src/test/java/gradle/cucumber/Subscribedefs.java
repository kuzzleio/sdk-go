package gradle.cucumber;

import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class Subscribedefs {
    private Kuzzle k = new Kuzzle("localhost");
    private NotificationContent content = null;

    @Given("^I subscribe to \"([^\"]*)\"$")
    public void i_subscribe_to(String arg1) throws Exception {
        k.getAuth().login("local", "{\"username\": \""+"useradmin"+"\", \"password\":\""+"testpwd"+"\" }", 40000);

        RoomOptions opts = new RoomOptions();
        opts.setSubscribeToSelf(true);

        k.getRealtime().subscribe("index", "collection", "{}", new NotificationListener() {
            @Override
            public void onMessage(NotificationResult res) {
                content = res.getResult();
            }
        }, opts);
    }

    @When("^I create a document in \"([^\"]*)\"$")
    public void i_create_a_document_in(String arg1) throws Exception {
        k.getDocument().create("index", "collection", "", "{}");
        Thread.sleep(1000);
    }

    @Then("^I receive a notification$")
    public void i_receive_a_notification() throws Exception {
        Assert.assertNotNull(content);
        Assert.assertNotNull(content.getContent());
    }
}
