package gradle.cucumber;

import io.kuzzle.sdk.Kuzzle;

public class KuzzleSingleton {
    private static Kuzzle kuzzle = null;

    public static Kuzzle getInstance() {
        if (kuzzle != null) {
            return kuzzle;
        }

        kuzzle = new Kuzzle((System.getenv().get("KUZZLE_HOST") != null) ? (System.getenv().get("KUZZLE_HOST")) : "localhost");
        return kuzzle;
    }
}
