package gradle.cucumber;

import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class Subscribedefs {
    private Kuzzle k = new Kuzzle((System.getenv().get("KUZZLE_HOST") != null) ? (System.getenv().get("KUZZLE_HOST")) : "localhost");
    private String roomId;
    private NotificationContent content = null;

    @Given("^I subscribe to \"([^\"]*)\"$")
    public void i_subscribe_to(String collection) throws Exception {
        k.getAuth().login("local", "{\"username\": \""+"useradmin"+"\", \"password\":\""+"testpwd"+"\" }", 40000);

        roomId = k.getRealtime().subscribe("index", collection, "{}", new NotificationListener() {
            @Override
            public void onMessage(NotificationResult res) {
                content = res.getResult();
            }
        });
    }

    @When("^I create a document in \"([^\"]*)\"$")
    public void i_create_a_document_in(String collection) throws Exception {
        k.getDocument().create("index", collection, "", "{}");
        Thread.sleep(1000);
    }

    @Then("^I( do not)? receive a notification$")
    public void i_receive_a_notification(String dont) throws Exception {
        if (dont == null) {
            Assert.assertNotNull(content);
            Assert.assertNotNull(content.getContent());
            content = null;
            return;
        }
        Assert.assertNull(content);
    }

    @Given("^I unsubscribe")
    public void i_unsubscribe() throws Exception {
        k.getRealtime().unsubscribe(roomId);
    }

}
