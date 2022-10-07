package hello_test

import (
    "testing"
    "app/hello"
)

func TestHello(t *testing.T) {
    result := hello.Hello()
    if result != "Hello, World" {
       t.Errorf("Response incorrect, got: %s, want: %s.", result, "Hello, World")
    }
}
