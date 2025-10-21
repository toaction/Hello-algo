package com.toaction.helloalgo;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class HelloTest {
    @Test
    public void testHello() {
        String result = Hello.hello();
        assertEquals("Hello, Algo!", result, "Hello.hello() should return 'Hello, Algo!'");
    }
}
