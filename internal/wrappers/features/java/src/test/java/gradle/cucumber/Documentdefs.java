package gradle.cucumber;

import cucumber.api.java.After;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.BadRequestException;
import io.kuzzle.sdk.Kuzzle;
import io.kuzzle.sdk.KuzzleException;
import org.junit.Assert;

public class Documentdefs {
    private Kuzzle k;
    private String index;
    private String collection;
    private String errorMessage;
    private String documentId;

    @After
    public void after() {
        if (documentId != null) {
            k.getDocument().delete_(this.index, this.collection, this.documentId);
        }
    }

    @Given("^Kuzzle Server is running$")
    public void kuzzle_Server_is_running() throws Exception {
        k = new Kuzzle((System.getenv().get("KUZZLE_HOST") != null) ? (System.getenv().get("KUZZLE_HOST")) : "localhost");
    }

    @Given("^there is an index \'([^\"]*)\'$")
    public void there_is_an_index(String index) throws Exception {
        this.index = index;
        if (!k.getIndex().exists(index)) {
            k.getIndex().create(index);
        }
    }

    @Given("^it has a collection \'([^\"]*)\'$")
    public void it_has_a_collection(String collection) throws Exception {
        this.collection = collection;
        if (!k.getCollection().exists(this.index, collection)) {
            k.getCollection().create(this.index, collection);
        }
    }

    @Given("^the collection has a document with id \'([^\"]*)\'$")
    public void the_collection_has_a_document_with_id(String id) throws Exception {
        this.documentId = id;
        k.getDocument().create(index, collection, id, "{\"foo\":\"bar\"}");
    }

    @When("^I try to create a new document with id \'([^\"]*)\'$")
    public void i_try_to_create_a_new_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            this.documentId = id;
            k.getDocument().create(this.index, this.collection, id, "{\"foo\": \"bar\"}");
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^I get an error with message \'([^\"]*)\'$")
    public void i_get_an_error_with_message(String message) throws Exception {
        Assert.assertEquals(message, this.errorMessage);
    }

    @Given("^the collection doesn't have a document with id \'([^\"]*)\'$")
    public void the_collection_doesn_t_have_a_document_with_id(String id) throws Exception {
        Assert.assertFalse(k.getDocument().exists(this.index, this.collection, id));
    }

    @Then("^the document is successfully created$")
    public void the_document_is_successfully_created() throws Exception {
        Assert.assertNull(this.errorMessage);
    }
}
