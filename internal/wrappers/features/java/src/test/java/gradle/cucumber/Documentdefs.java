package gradle.cucumber;

import cucumber.api.java.After;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import gherkin.deps.com.google.gson.Gson;
import gherkin.deps.com.google.gson.JsonObject;
import io.kuzzle.sdk.*;
import org.junit.Assert;

public class Documentdefs {
    private Kuzzle k;
    private String errorMessage;
    private World world;
    private String documentId;
    private SearchResult documents;

    class Source {
        JsonObject _source;
    }

    @Given("^Kuzzle Server is running$")
    public void kuzzle_Server_is_running() throws Exception {
        k = KuzzleSingleton.getInstance();
    }

    @Given("^there is an index \'([^\"]*)\'$")
    public void there_is_an_index(String index) throws Exception {
        world.index = index;
        if (!k.getIndex().exists(index)) {
            k.getIndex().create(index);
        }
    }

    @Given("^it has a collection \'([^\"]*)\'$")
    public void it_has_a_collection(String collection) throws Exception {
        world.collection = collection;
        if (!k.getCollection().exists(world.index, collection)) {
            k.getCollection().create(world.index, collection);
        }
    }

    @Given("^the collection has a document with id \'([^\"]*)\'$")
    public void the_collection_has_a_document_with_id(String id) throws Exception {
        QueryOptions options = new QueryOptions();
        options.setRefresh("wait_for");

        try {
            k.getDocument().create(world.index, world.collection, id, "{\"foo\":\"bar\"}", options);
            this.documentId = id;
        } catch (BadRequestException e) {
            if (!e.getMessage().equals("Document already exists")) {
                throw e;
            }
        }
    }

    @When("^I create a document with id \'([^\"]*)\'$")
    public void i_try_to_create_a_new_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            k.getDocument().create(world.index, world.collection, id, "{\"foo\": \"bar\"}");
            this.documentId = id;
            this.errorMessage = null;
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^I get an error with message \'([^\"]*)\'$")
    public void i_get_an_error_with_message(String message) throws Exception {
        Assert.assertEquals(message, this.errorMessage);
    }

    @Then("^the document is successfully created$")
    public void the_document_is_successfully_created() throws Exception {
        Assert.assertNull(this.errorMessage);
    }

    @When("^I update the document with id \'([^\"]*)\' and content \'([^\"]*)\' = \'([^\"]*)\'$")
    public void i_update_the_document_with_id_and_content(String id, String key, String value) throws Exception {
        QueryOptions options = new QueryOptions();
        options.setRefresh("wait_for");

        k.getDocument().update(world.index, world.collection, id, "{\""+key+"\":\""+value+"\"}");
    }

    @When("^I createOrReplace a document with id \'([^\"]*)\'$")
    public void i_createOrReplace_a_new_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            k.getDocument().createOrReplace(world.index, world.collection, id, "{\"foo\": \"barz\"}");
            this.documentId = id;
            this.errorMessage = null;
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document is successfully replaced$")
    public void the_document_is_successfully_replaced() throws Exception {
        String doc = k.getDocument().get(world.index, world.collection, this.documentId);
        Assert.assertNull(this.errorMessage);

        Gson gson = new Gson();
        Source s = gson.fromJson(doc, Source.class);
        Assert.assertEquals("{\"foo\":\"barz\"}", s._source.toString());
    }

    @Given("^the collection doesn't have a document with id \'([^\"]*)\'$")
    public void the_collection_doesn_t_have_a_document_with_id(String id) throws Exception {
        QueryOptions options = new QueryOptions();
        options.setRefresh("wait_for");

        k.getDocument().delete_(world.index, world.collection, id, options);
    }

    @When("^I replace a document with id \'([^\"]*)\'$")
    public void i_replace_a_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            k.getDocument().replace(world.index, world.collection, id, "{\"foo\": \"barz\"}");
            this.documentId = id;
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @When("^I delete a document with id \'([^\"]*)\'$")
    public void i_delete_a_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            k.getDocument().delete_(world.index, world.collection, id);
        } catch (KuzzleException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document is successfully deleted")
    public void the_document_is_successfully_deleted() throws Exception {
        Assert.assertNull(this.errorMessage);
    }

    @When("^I update a document with id \'([^\"]*)\'$")
    public void i_update_a_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            k.getDocument().update(world.index, world.collection, id, "{\"foo\": \"barz\"}");
            this.documentId = id;
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document is successfully updated")
    public void the_document_is_successfully_updated() throws Exception {
        String doc = k.getDocument().get(world.index, world.collection, this.documentId);
        Assert.assertNull(this.errorMessage);

        Gson gson = new Gson();
        Source s = gson.fromJson(doc, Source.class);
        Assert.assertEquals("{\"foo\":\"barz\"}", s._source.toString());
    }

    @When("^I search a document with id \'([^\"]*)\'$")
    public void i_search_a_document_with_id(String id) throws Exception {
        this.errorMessage = null;
        try {
            String document = k.getDocument().search(world.index, world.collection, id, "{\"foo\": \"barz\"}");
            this.document = id;
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }
}
