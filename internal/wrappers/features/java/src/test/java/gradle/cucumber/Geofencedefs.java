package gradle.cucumber;

import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class Geofencedefs {
    private Kuzzle k;
    private NotificationResult content = null;
    private String documentId;

    @Before
    public void before() {
        k = KuzzleSingleton.getInstance();
    }

    @Given("^I subscribe to \"([^\"]*)\" collection$")
    public void i_subscribe_to_collection(String collection) throws Exception {
        String filter = "{" +
                "\"geoBoundingBox\": {" +
                "\"position\": {" +
                "\"topLeft\": {" +
                "\"lat\": 10," +
                "\"lon\": 1" +
                "}," +
                "\"bottomRight\": {" +
                "\"lat\": 1," +
                "\"lon\": 10" +
                "}" +
                "}" +
                "}" +
                "}";

        RoomOptions opts = new RoomOptions();
        opts.setUser("all");
        opts.setState("all");
        opts.setScope("all");
        opts.setSubscribeToSelf(true);

        k.getAuth().login("local", "{\"username\": \""+"useradmin"+"\", \"password\":\""+"testpwd"+"\" }", 40000);

        k.getRealtime().subscribe("index", collection, filter, new NotificationListener() {
            @Override
            public void onMessage(NotificationResult res) {
                content = res;
            }
        }, opts);
    }

    @When("^I create a document in \"([^\"]*)\" and inside my geofence$")
    public void i_create_a_document_in_and_inside_my_geofence(String collection) throws Exception {
        System.err.println("create document");
        documentId = k.getDocument().create("index", collection, "", "{\"position\":{\"lat\": 2, \"lon\": 2}}");
        Thread.sleep(1000);
    }

    @Then("^I receive a \"([^\"]*)\" notification$")
    public void i_receive_a_notification(String scope) throws Exception {
        Assert.assertNotNull(content);
        Assert.assertEquals(scope, content.getScope());
    }

    @When("^I update this document to be outside my geofence$")
    public void i_update_this_document_to_be_outside_my_geofence() throws Exception {
        content = null;
        k.getDocument().update("index", "geofence", documentId, "{\"position\":{\"lat\": 30, \"lon\": 20}}");
        Thread.sleep(1000);
    }
}
