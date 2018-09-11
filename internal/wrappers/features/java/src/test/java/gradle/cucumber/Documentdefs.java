package gradle.cucumber;

import cucumber.api.java.en.And;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import gherkin.deps.com.google.gson.Gson;
import gherkin.deps.com.google.gson.JsonObject;
import io.kuzzle.sdk.*;
import org.junit.Assert;

import java.util.List;

public class Documentdefs {
    private Kuzzle k;
    private String errorMessage;
    private World world;
    private String documentId;
    private SearchResult documents;
    private int nbDocuments;
    private boolean partialException = false;
    private boolean documentExists = false;
    private String jsonDocuments;

    class Source {
        JsonObject _source;
    }

    @Given("^Kuzzle Server is running$")
    public void kuzzle_Server_is_running() throws Exception {
        k = KuzzleSingleton.getInstance();

        k.connect();
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
        Assert.assertEquals("\"barz\"", s._source.get("foo").toString());
    }

    @Given("^the collection doesn't have a document with id \'([^\"]*)\'$")
    public void the_collection_doesn_t_have_a_document_with_id(String id) throws Exception {
        QueryOptions options = new QueryOptions();
        options.setRefresh("wait_for");

        try {
            k.getDocument().delete(world.index, world.collection, id, options);
        } catch (KuzzleException e) {}
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
            k.getDocument().delete(world.index, world.collection, id);
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
        Assert.assertEquals("\"barz\"", s._source.get("foo").toString());
    }

    @When("^I search a document with id \'([^\"]*)\'$")
    public void i_search_a_document_with_id(String id) throws Exception {
        this.errorMessage = null;

        try {
            QueryOptions options = new QueryOptions();
            options.setSize(42);
            this.documents = k.getDocument().search(world.index, world.collection, "{\"query\": {\"bool\": {\"should\":[{\"match\":{\"_id\": \""+id+"\"}}]}}}");
        } catch (BadRequestException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document is successfully found$")
    public void the_document_is_successfully_found() throws Exception {
        Assert.assertNotNull(this.documents);
        Assert.assertNotNull(this.documents.getDocuments());
        Assert.assertNotEquals("[]", this.documents.getDocuments());
    }

    @Then("^the document is not found$")
    public void the_document_is_not_found() throws Exception {
        Assert.assertNotNull(this.documents.getDocuments());
        Assert.assertEquals("[]", this.documents.getDocuments());
    }

    @Then("^I shall receive (\\d+)$")
    public void i_shall_receive(int nbDocuments) throws Exception {
        Assert.assertEquals(nbDocuments, this.nbDocuments);
        this.nbDocuments = 0;
    }

    @Then("^I must have (\\d+) documents in the collection$")
    public void i_must_have_documents_in_the_collection(int nb) {
        Assert.assertEquals(nb, k.getDocument().count(world.index, world.collection, "{}"));
    }

    @When("^I count how many documents there is in the collection$")
    public void i_count_how_many_documents_there_is_in_the_collection() throws Exception {
        nbDocuments = k.getDocument().count(world.index, world.collection, "{}");
    }

    @When("^I delete the documents \\[\'(.*)\', \'(.*)\'\\]$")
    public void i_delete_the_documents(String doc1, String doc2) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        StringVector v = new StringVector();
        v.add(doc1);
        v.add(doc2);
        try {
            k.getDocument().mDelete(world.index, world.collection, v, o);
            this.partialException = false;
        } catch (PartialException e) {
            this.partialException = true;
        }
    }

    @Then("^the collection must be empty$")
    public void the_collection_must_be_empty() {
        Assert.assertEquals(0, k.getDocument().count(world.index, world.collection, "{}"));
    }

    @And("^I get a partial error$")
    public void i_get_a_partial_error() {
        Assert.assertTrue(this.partialException);
        this.partialException = false;
    }

    @When("^I create the documents \\[\'(.*)\', \'(.*)\'\\]$")
    public void i_create_the_documents(String doc1, String doc2) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        String docs = "{\"documents\":[{\"_id\":\""+doc1+"\", \"body\":{}}, {\"_id\":\""+doc2+"\", \"body\":{}}]}";
        try {
            k.getDocument().mCreate(world.index, world.collection, docs, o);
            this.partialException = false;
        } catch (PartialException e) {
            this.partialException = true;
        }
    }

    @Then("^I should have no partial error$")
    public void i_should_have_no_partial_error() {
        Assert.assertFalse(this.partialException);
    }

    @When("^I replace the documents \\[\'(.*)\', \'(.*)\'\\]$")
    public void i_replace_the_documents(String doc1, String doc2) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        String docs = "{\"documents\":[{\"_id\":\""+doc1+"\", \"body\":{\"foo\":\"barz\"}}, {\"_id\":\""+doc2+"\", \"body\":{\"foo\":\"barz\"}}]}";
        try {
            k.getDocument().mReplace(world.index, world.collection, docs, o);
            this.partialException = false;
        } catch (PartialException e) {
            this.partialException = true;
        }
    }

    @Then("^the document \'([^\"]*)\' should be replaced")
    public void the_documents_should_be_replaced(String id1) throws Exception {
        String doc = k.getDocument().get(world.index, world.collection, id1);
        Assert.assertNull(this.errorMessage);

        Gson gson = new Gson();
        Source s = gson.fromJson(doc, Source.class);
        Assert.assertEquals("\"barz\"", s._source.get("foo").toString());
    }

    @When("^I update the documents \\[\'(.*)\', \'(.*)\'\\]$")
    public void i_update_the_documents(String doc1, String doc2) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        String docs = "{\"documents\":[{\"_id\":\""+doc1+"\", \"body\":{\"foo\":\"barz\"}}, {\"_id\":\""+doc2+"\", \"body\":{\"foo\":\"barz\"}}]}";
        try {
            k.getDocument().mUpdate(world.index, world.collection, docs, o);
            this.partialException = false;
        } catch (PartialException e) {
            this.partialException = true;
        } catch (KuzzleException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document \'([^\"]*)\' should be updated")
    public void the_documents_should_be_updated(String id1) throws Exception {
        Assert.assertNull(this.errorMessage);
        String doc = k.getDocument().get(world.index, world.collection, id1);
        Assert.assertNull(this.errorMessage);

        Gson gson = new Gson();
        Source s = gson.fromJson(doc, Source.class);
        Assert.assertEquals("\"barz\"", s._source.get("foo").toString());
    }

    @When("^I createOrReplace the documents \\[\'(.*)\', \'(.*)\'\\]$")
    public void i_createOrReplace_the_documents(String doc1, String doc2) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        String docs = "{\"documents\":[{\"_id\":\""+doc1+"\", \"body\":{\"foo\":\"barz\"}}, {\"_id\":\""+doc2+"\", \"body\":{\"foo\":\"barz\"}}]}";
        try {
            k.getDocument().mCreateOrReplace(world.index, world.collection, docs, o);
            this.partialException = false;
        } catch (PartialException e) {
            this.partialException = true;
        } catch (KuzzleException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document \'([^\"]*)\' should be created")
    public void the_documents_should_be_created(String id1) throws Exception {
        Assert.assertNull(this.errorMessage);
        String doc = k.getDocument().get(world.index, world.collection, id1);
        Assert.assertNull(this.errorMessage);

        Gson gson = new Gson();
        Source s = gson.fromJson(doc, Source.class);
        Assert.assertEquals("\"barz\"", s._source.get("foo").toString());
    }

    @When("^I check if \'([^\"]*)\' exists$")
    public void i_createOrReplace_the_documents(String doc) throws Exception {
        this.documentExists = false;
        try {
            this.documentExists = k.getDocument().exists(world.index, world.collection, doc);
        } catch (KuzzleException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the document should exist$")
    public void the_document_should_exists() throws Exception {
        Assert.assertNull(this.errorMessage);
        Assert.assertTrue(this.documentExists);
    }

    @Then("^the document should not exist$")
    public void the_document_should_not_exists() throws Exception {
        Assert.assertNull(this.errorMessage);
        Assert.assertFalse(this.documentExists);
    }

    @When("^I get documents \\[\'(.*)\', \'(.*)\'\\]$")
    public void i_get_document_mget_my_document_id_and_mget_my_document_id(String id1, String id2) throws Exception {
        try {
            StringVector v = new StringVector();
            v.add(id1);
            v.add(id2);
            jsonDocuments = k.getDocument().mGet(world.index, world.collection, v, false);
        } catch (KuzzleException e) {
            this.errorMessage = e.getMessage();
        }
    }

    @Then("^the documents should be retrieved$")
    public void the_documents_should_be_retrieved() throws Exception {
        Assert.assertNull(this.errorMessage);
        Assert.assertNotNull(this.jsonDocuments);
    }
}
