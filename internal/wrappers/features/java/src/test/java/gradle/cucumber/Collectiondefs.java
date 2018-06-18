package gradle.cucumber;

import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import gherkin.deps.com.google.gson.Gson;
import gherkin.deps.com.google.gson.JsonArray;
import gherkin.deps.com.google.gson.JsonObject;
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

    @Then("^the collection should exists$")
    public void the_collection_should_exists() throws Exception {
        Assert.assertTrue(k.getCollection().exists(world.index, world.collection));
    }

    @When("^I check if the collection exists$")
    public void i_check_if_the_collection_exists() throws Exception {
        exists = k.getCollection().exists(world.index, world.collection);
    }

    @Then("^it should exists$")
    public void it_should_exists() throws Exception {
        Assert.assertTrue(exists);
    }

    @When("^I list the collections$")
    public void i_list_the_collections() throws Exception {
        listCollections = k.getCollection().list(world.index);
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

    @When("^I truncate the collection$")
    public void i_truncate_the_collection() throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        k.getCollection().truncate(world.index, world.collection, o);
    }

    @Then("^the collection shall be empty$")
    public void it_should_be_empty() throws Exception {
        SearchResult res = k.getDocument().search(world.index, world.collection, "{}");
        Assert.assertEquals("[]", res.getDocuments());
    }

    @When("^I update the mapping$")
    public void i_update_the_mapping() throws Exception {
        k.getCollection().updateMapping(world.index, world.collection, "{" +
                "\"properties\": {" +
                "    \"foo\": {" +
                "      \"type\": \"string\"" +
                "    }" +
                "}" +
                "}");
    }

    @Then("^the mapping should be updated$")
    public void the_mapping_should_be_updated() throws Exception {
        String mapping = k.getCollection().getMapping(world.index, world.collection);
        Assert.assertEquals("{\"test-index\":{\"mappings\":{\"test-collection\":{\"properties\":{\"foo\":{\"type\":\"text\",\"fields\":{\"keyword\":{\"type\":\"keyword\",\"ignore_above\":256}}}}}}}}", mapping);
    }

    @When("^I update the specifications$")
    public void i_update_the_specifications() throws Exception {
        k.getCollection().updateSpecifications(world.index, world.collection, "{\""+world.index+"\":{\""+world.collection+"\":{\"strict\":true}}}");
    }

    @Then("^they should be updated$")
    public void they_should_be_updated() throws Exception {
        Assert.assertEquals("{\"validation\":{\"strict\":true},\"index\":\"test-index\",\"collection\":\"test-collection\"}", k.getCollection().getSpecifications(world.index, world.collection));
    }

    @When("^I validate the specifications$")
    public void i_validate_the_specifications() throws Exception {
        this.validateSpecs = k.getCollection().validateSpecifications("{\""+world.index+"\":{\""+world.collection+"\":{\"strict\":true}}}");
    }

    @Then("^they should be validated$")
    public void they_should_be_validated() throws Exception {
        Assert.assertTrue(this.validateSpecs);
    }

    @Given("^has specifications$")
    public void has_specifications() throws Exception {
        k.getCollection().updateSpecifications(world.index, world.collection, "{\""+world.index+"\":{\""+world.collection+"\":{\"strict\":true}}}");
    }

    @When("^I delete the specifications$")
    public void i_delete_the_specifications() throws Exception {
        k.getCollection().deleteSpecifications(world.index, world.collection);
    }

    @Then("^the specifications must not exist$")
    public void the_specifications_must_not_exist() throws Exception {
        boolean notFound = false;
        try {
            k.getCollection().getSpecifications(world.index, world.collection);
        } catch (NotFoundException e) {
            notFound = true;
        }
        Assert.assertTrue(notFound);
    }

}
