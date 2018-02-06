%typemap(javaimports) kuzzleio::Room "
/* The type Room. */"

%javamethodmodifiers kuzzleio::Room::Room(Collection *collection, json_object* filters, room_options* options) "
  /**
   * The Room object is the result of a subscription request, allowing to manipulate the subscription itself.
   *
   * @param collection           - An instantiated Collection object
   * @param filters              - Subscription filters
   * @param options              - Subscription options
   */
  public";

%javamethodmodifiers kuzzleio::Room::Room(Collection *collection, json_object* filters) "
  /**
   * {@link #Room(Collection collection, JsonObject filters, RoomOptions options)}
   */
  public";

  %javamethodmodifiers kuzzleio::Room::Room(Collection *collection) "
  /**
   * {@link #Room(Collection collection, JsonObject filters, RoomOptions options)}
   */
  public";