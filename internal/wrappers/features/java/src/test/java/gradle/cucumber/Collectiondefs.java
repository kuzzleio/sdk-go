package gradle.cucumber;

import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import gherkin.deps.com.google.gson.Gson;
import gherkin.deps.com.google.gson.JsonArray;
import io.kuzzle.sdk.Kuzzle;
import io.kuzzle.sdk.NotFoundException;
import io.kuzzle.sdk.QueryOptions;
import io.kuzzle.sdk.SearchResult;
import org.junit.Assert;

import javax.management.remote.NotificationResult;

public class Collectiondefs {
    private Kuzzle k;
    private World world;
    private boolean exists = false;
    String listCollections;
    private boolean validateSpecs = false;

    @Before
    public void before() {
        k = KuzzleSingleton.getInstance();
    }

    @When("^I create a collection \'([^\"]*)\'$")
    public void i_create_a_collection_test_collection(String collection) throws Exception {
        k.getCollection().create(world.index, collection);
        world.collection = collection;
    }

    @Then("^the collection \'([^\"]*)\' should exist$")
    public void the_collection_should_exists(String collection) throws Exception {
        Assert.assertTrue(k.getCollection().exists(world.index, collection));
    }

    @When("^I check if the collection \'([^\"]*)\' exists$")
    public void i_check_if_the_collection_exists(String collection) throws Exception {
        exists = k.getCollection().exists(world.index, collection);
    }

    @Then("^the collection should exist$")
    public void it_should_exists() throws Exception {
        Assert.assertTrue(exists);
    }

    @When("^I list the collections of \'([^\"]*)\'$")
    public void i_list_the_collections(String index) throws Exception {
        listCollections = k.getCollection().list(index);
    }

    class Col {
        JsonArray collections;
    }

    @Then("^the result contains (\\d+) hits$")
    public void the_result_contains_hits(int nb) throws Exception {
        Gson gson = new Gson();
        Col c = gson.fromJson(listCollections, Collectiondefs.Col.class);

        Assert.assertEquals(nb, c.collections.size());

    }

    @Then("^the content should not be null$")
    public void i() {
        Assert.assertNotNull(this.listCollections);
        Assert.assertEquals("{\"type\":\"all\",\"collections\":[{\"name\":\"test-collection1\",\"type\":\"stored\"},{\"name\":\"test-collection2\",\"type\":\"stored\"}],\"from\":0,\"size\":10}", this.listCollections);
    }

    @When("^I truncate the collection \'([^\"]*)\'$")
    public void i_truncate_the_collection(String collection) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        k.getCollection().truncate(world.index, collection, o);
    }

    @Then("^the collection \'([^\"]*)\' should be empty$")
    public void it_should_be_empty(String collection) throws Exception {
        SearchResult res = k.getDocument().search(world.index, collection, "{}");
        Assert.assertEquals("[]", res.getDocuments());
    }

    @When("^I update the mapping of collection \'([^\"]*)\'$")
    public void i_update_the_mapping(String collection) throws Exception {
        k.getCollection().updateMapping(world.index, collection, "{" +
                "\"properties\": {" +
                "    \"foo\": {" +
                "      \"type\": \"string\"," +
                "      \"fields\": {\"keyword\":{\"type\":\"keyword\",\"ignore_above\":256}}" +
                "    }" +
                "}" +
                "}");
    }

    @Then("^the mapping of \'([^\"]*)\' should be updated$")
    public void the_mapping_should_be_updated(String collection) throws Exception {
        String mapping = k.getCollection().getMapping(world.index, collection);
        Assert.assertEquals(
            "{\"" + world.index + "\":{\"mappings\":{\"" + collection + "\":{\"properties\":{\"foo\":{\"type\":\"text\",\"fields\":{\"keyword\":{\"type\":\"keyword\",\"ignore_above\":256}}}}}}}}", mapping);
    }

    @When("^I update the specifications of the collection \'([^\"]*)\'$")
    public void i_update_the_specifications(String collection) throws Exception {
        k.getCollection().updateSpecifications(world.index, world.collection, "{\"" + world.index + "\":{\""+ collection +"\":{\"strict\":true}}}");
    }

    @Then("^the specifications of \'([^\"]*)\' should be updated$")
    public void the_specifications_of_collection_should_be_updated(String collection) throws Exception {
        Assert.assertEquals("{\"validation\":{\"strict\":true},\"index\":\"test-index\",\"collection\":\"test-collection\"}", k.getCollection().getSpecifications(world.index, collection));
    }

    @When("^I validate the specifications of \'([^\"]*)\'$")
    public void i_validate_the_specifications(String collection) throws Exception {
        this.validateSpecs = k.getCollection().validateSpecifications("{\""+world.index+"\":{\""+collection+"\":{\"strict\":true}}}");
    }

    @Then("^they should be validated$")
    public void they_should_be_validated() throws Exception {
        Assert.assertTrue(this.validateSpecs);
    }

    @Given("^has specifications$")
    public void has_specifications() throws Exception {
        k.getCollection().updateSpecifications(world.index, world.collection, "{\""+world.index+"\":{\""+world.collection+"\":{\"strict\":true}}}");
    }

    @When("^I delete the specifications of \'([^\"]*)\'$")
    public void i_delete_the_specifications(String collection) throws Exception {
        k.getCollection().deleteSpecifications(world.index, collection);
    }

    @Then("^the specifications of \'([^\"]*)\' must not exist$")
    public void the_specifications_must_not_exist(String collection) throws Exception {
        boolean notFound = false;
        try {
            k.getCollection().getSpecifications(world.index, collection);
        } catch (NotFoundException e) {
            notFound = true;
        }
        Assert.assertTrue(notFound);
    }

    @When("^I create a collection \'([^\"]*)\' with a mapping$")
    public void i_create_a_collection_test_collection_with_mapping(String collection) throws Exception {
        String mapping = "{\"properties\": {\"foo\": {"
            + "\"type\": \"string\", \"fields\": {"
            + "\"keyword\":{\"type\":\"keyword\",\"ignore_above\":256}}"
            + "}}}";
        k.getCollection().create(world.index, collection, mapping);
        world.collection = collection;
    }
}
