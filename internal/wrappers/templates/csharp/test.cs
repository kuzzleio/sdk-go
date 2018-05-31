using System;
using Kuzzleio;

public class Example
{
    public static Kuzzle k;

    public static void Main()
    {
        k = new Kuzzle("localhost", null);
        k.server.now();
    }
}