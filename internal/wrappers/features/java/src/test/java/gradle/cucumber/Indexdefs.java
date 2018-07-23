package gradle.cucumber;

import cucumber.api.PendingException;
import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.Kuzzle;
import io.kuzzle.sdk.KuzzleException;
import io.kuzzle.sdk.QueryOptions;
import io.kuzzle.sdk.StringVector;
import org.junit.Assert;

public class Indexdefs {
    private Kuzzle k;
    private World world;
    private String errorMessage;

    @Before
    public void before() {
        k = KuzzleSingleton.getInstance();
    }

    @Given("^there is no index called \'([^\"]*)\'$")
    public void there_is_no_index_called(String index) {
        try {
            QueryOptions o = new QueryOptions();
            o.setRefresh("wait_for");

            k.getIndex().delete(index, o);
        } catch (KuzzleException e) {}
    }

    @When("^I create an index called \'([^\"]*)\'$")
    public void i_create_an_index_called_test_index(String index) throws Exception {
        try {
            k.getIndex().create(index);
        } catch (KuzzleException e) {
            this.errorMessage = e.getMessage();
        }
        world.index = index;
    }

    @Then("^the index should exist$")
    public void the_index_should_exists() throws Exception {
            Assert.assertTrue(k.getIndex().exists(world.index));
    }


    @Then("^I get an error$")
    public void i_get_an_error() throws Exception {
        Assert.assertNotNull(this.errorMessage);
    }

    @Given("^there is the indexes \'([^\"]*)\' and \'([^\"]*)\'$")
    public void there_is_the_indexes_test_index_and_test_index(String index1, String index2) throws Exception {
        try {
            QueryOptions o = new QueryOptions();
            o.setRefresh("wait_for");

            k.getIndex().create(index1, o);
            k.getIndex().create(index2, o);
        } catch(Exception e) {}
    }

    @When("^I delete the indexes \'([^\"]*)\' and \'([^\"]*)\'$")
    public void i_delete_the_indexes_test_index_and_test_index(String index1, String index2) throws Exception {
        QueryOptions o = new QueryOptions();
        o.setRefresh("wait_for");

        StringVector v = new StringVector();
        v.add(index1);
        v.add(index2);

        k.getIndex().mDelete(v, o);
    }

    @Then("^indexes \'([^\"]*)\' and \'([^\"]*)\' don't exist$")
    public void indexes_test_index_and_test_index_don_t_exist(String index1, String index2) throws Exception {
        Assert.assertFalse(k.getIndex().exists(index1));
        Assert.assertFalse(k.getIndex().exists(index2));
    }

    @When("^I list indexes$")
    public void i_list_indexes() throws Exception {
      world.stringArray = k.getIndex().list();
    }

    @Then("^I get \'([^\"]*)\' and \'([^\"]*)\'$")
    public void i_get_test_index_and_test_index(String index1, String index2) throws Exception {
      int match = 0;

      for (int i = 0; i < world.stringArray.size(); ++i) {
        if (world.stringArray.get(i).equals(index1) || world.stringArray.get(i).equals(index2))
          match++;
      }

      Assert.assertTrue(match == 2);
    }
}
