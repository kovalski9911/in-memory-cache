# in-memory cache

A simple package for working with in-memory data caching. The package provides a cache structure that supports storing and retrieving various types of data through generic methods.

## Installation

To install the package, run the following command:

```Bash
go get github.com/kovalski9911/inMemoryCache
```


## Usage

Here is a basic example of using the package:

```go
package main

import (
    "fmt"
    "github.com/kovalski9911/inmemorycache"
)

func main() {
    // Create a new cache to store strings
    strCache := inmemorycache.New[string]()
    
    // Set a value in the cache
    strCache.Set("key", "value")
    
    // Retrieve a value from the cache
    val := strCache.Get("key")
    fmt.Println(val) // Output: value
    
    // Remove a value from the cache
    strCache.Delete("key")
}
```


## Methods

The package provides the following methods for working with the cache:

- ```Set(k string, v T)``` – sets the value v by key k.
- ```Get(k string) T``` – retrieves the value by key k. If the key does not exist, it returns the zero value of type T.
- ```Delete(k string)``` – removes the value associated with key k from the cache.