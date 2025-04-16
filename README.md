# DataStore

A simple thread-safe, key-value store implementation in Go with support for multiple storage backends and data types.

> **Note**: This sample project serves as a demonstration of implementing a simple data store using `Go generics`, `thread-safety with mutexes, and the repository pattern for extensibility`. It is intended as an educational resource showcasing different implementation approaches rather than a production-ready library.

## Implementation Files

- [`/basic/data-store.go`](https://github.com/svelama/datastore/blob/main/basic/data-store.go) - Basic mutex-based implementation
- [`/multiple-storage-backends/data-store.go`](https://github.com/svelama/datastore/blob/main/multiple-storage-backends/data-store.go) - Repository interface and DataStore implementation
- [`/multiple-storage-backends/map-repository.go`](https://github.com/svelama/datastore/blob/main/multiple-storage-backends/map-repository.go) - Map-based repository implementation
