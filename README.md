# NextGen CI Foundation
The `nextgen-ci-foundation` repository provides foundational components and configurations essential for the next-generation continuous integration (CI) infrastructure. This repository includes tools and scripts for logging configuration, monitoring, server graceful shutdown, and other infrastructure utilities.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
    - [Logging](#logging)
    - [Graceful Shutdown](#graceful-shutdown)
    - [Utilities](#utilities)
      - [File Utilities](#file-utilities)
      - [String Utilities](#string-utilities)
      - [Array Utilities](#array-utilities)
      - [Jitter Utility](#jitter-utility)
- [Contributing](#contributing)
- [License](#license)

## Features
The `nextgen-ci-foundation` repository provides the following features:

- **Logging**: Flexible and configurable logging system with support for multiple log formats (plain text, console JSON, Cloud Logging, etc.).
- **Graceful Shutdown**: Ensure services can shut down gracefully, preserving state and avoiding data loss.
- **Utilities**: Various utility functions for working with files, strings, arrays, and jitter.

## Installation

To use the `nextgen-ci-foundation` library in your Go project, add it as a dependency in your `go.mod` file:

```sh
go get github.com/nextgenci/nextgen-ci-foundation
```

## Usage

### Logging

The `nextgen-ci-foundation` library provides a flexible logging system that can be configured to log to different destinations (e.g., console, file, Cloud Logging). To use the logging system, import the `logging` package and create a new logger:

```go
package main

import (
	"github.com/nextgenci/nextgen-ci-foundation/foundation/domain"
	foundation "github.com/nextgenci/nextgen-ci-foundation/foundation/logging"
)

func main() {
	appInfo := domain.NewApplicationInfo("appGroup", "appName", "1.0.0", "main", "abc123", "2023-01-01")
	foundation.InitLoggingFromEnv(appInfo)
}
```

### Graceful Shutdown

The graceful shutdown package ensures that services can shut down gracefully.
```go
package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/nextgenci/nextgen-ci-foundation/foundation/shutdown"
)

func main() {
	var waitGroup sync.WaitGroup

	// Define functions to execute on shutdown
	functionsOnShutdown := []func(){
		func() {
			// Add your shutdown logic here
		},
	}

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	go foundation.HandleGracefulShutdown(gracefulShutdown, &waitGroup, functionsOnShutdown...)

	// Application code...
	waitGroup.Add(1)
	// Simulate work
	waitGroup.Done()
}
``` 

### Utilities

The `nextgen-ci-foundation` library provides various utility functions for working with files, strings, arrays, and jitter.

#### File Utilities

The file utilities package provides functions for working with files, such as reading and writing files, watching files for changes, and creating temporary files.

#### String Utilities

The string utilities package provides functions for working with strings, such upper snake case, lowercase snake case, and more.

#### Array Utilities

The array utilities package provides functions for working with arrays, such as checking if a array contains a specific string/int, and more.

#### Jitter Utility

The jitter utility package provides functions for adding jitter to time durations.

# Contributing

Contributions are welcome! Please read our [Contributing Guidelines]() for more details on how to contribute.


# License

This project is licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details.
