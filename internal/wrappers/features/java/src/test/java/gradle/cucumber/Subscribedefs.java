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
    private String roomId;
    private NotificationContent content = null;
    private World world;

    @Before
    public void before() {
        k = KuzzleSingleton.getInstance();
    }

    @After
    public void after() {
        if (roomId != null) {
            k.getRealtime().unsubscribe(roomId);
        }
    }

    @Given("^I subscribe to \'(.*)\'$")
    public void i_subscribe_to(String collection) throws Exception {
        roomId = k.getRealtime().subscribe(world.index, collection, "{}", new NotificationListener() {
            @Override
            public void onMessage(NotificationResult res) {
                content = res.getResult();
            }
        });
    }

    @Given("^I subscribe to \'(.*)\' with \'(.*)\' as filter$")
    public void i_subscribe_with_filter(String collection, String filter) throws Exception {
        content = null;
        roomId = k.getRealtime().subscribe(world.index, collection, filter, new NotificationListener() {
            @Override
            public void onMessage(NotificationResult res) {
                content = res.getResult();
            }
        });
    }

    @When("^I create a document in \'([^\"]*)\'$")
    public void i_create_a_document_in(String collection) throws Exception {
        k.getDocument().create(world.index, collection, "", "{}");
        Thread.sleep(1000);
    }

    @Then("^I( do not)? receive a notification$")
    public void i_receive_a_notification(String dont) throws Exception {
        Thread.sleep(1000);
        if (dont == null) {
            Assert.assertNotNull(content);
            Assert.assertNotNull(content.getContent());
            content = null;
            return;
        }
        Assert.assertNull(content);
    }

    @Given("^I received a notification$")
    public void i_received_a_notification() throws Exception {
        Thread.sleep(1000);
        Assert.assertNotNull(content);
        Assert.assertNotNull(content.getContent());
        content = null;
    }

    @Given("^I unsubscribe")
    public void i_unsubscribe() throws Exception {
        k.getRealtime().unsubscribe(roomId);
    }

    @When("^I delete the document with id \'([^\"]*)\'$")
    public void i_delete_the_document_with_id(String id) throws Exception {
        k.getDocument().delete(world.index, world.collection, id);
        Thread.sleep(1000);
    }

    @When("^I publish a document$")
    public void i_publish_a_document() throws Exception {
        k.getRealtime().publish(world.index, world.collection, "{}");
    }

}
