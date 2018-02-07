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

%javamethodmodifiers kuzzleio::Room::count() "
  /**
   * Returns the number of other subscriptions on that room.
   *
   * @returns the number of subscription
   */
  public";

%javamethodmodifiers kuzzleio::Room::onDone(SubscribeListener* listener) "
  /**
   * Calls the provided callback when the subscription finishes.
   *
   * @params SubscribeListener
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Room::subscribe(NotificationListener* listener) "
  /**
   * Subscribes using the filters provided at the object creation.
   *
   * @params NotificationListener
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Room::unsubscribe() "
  /**
   * Subscribes using the filters provided at the object creation.
   *
   * @params Cancels the current subscription.
   * @return this
   */
  public";