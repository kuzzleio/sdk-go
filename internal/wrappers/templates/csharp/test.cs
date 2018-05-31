using System;

public class Example
{
    public static void Main()
    {
        Kuzzle k = new Kuzzle("localhost");
        Console.WriteLine(k.server.now());
    }
}