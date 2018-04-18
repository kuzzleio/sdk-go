package gradle.cucumber;

import cucumber.api.java.After;
import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import io.kuzzle.sdk.Kuzzle;
import io.kuzzle.sdk.KuzzleRequest;
import org.junit.Assert;

public class Subscribedefs {
    private Kuzzle k;

    @Before
    public void before() {
        k = new Kuzzle("localhost", null);
    }

    @After
    public void after() {
        k.disconnect();
    }

}
