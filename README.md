
# SIS Project — Checkpoint 1

## Overview

A synchronous in-memory Go application that analyzes text.

Each task stores its input, lifecycle state, and processing result.

## Features

- Create tasks with unique IDs
- Validate input and reject empty text
- Store and retrieve tasks in memory
- Count words, Unicode characters, and lines
- Track task lifecycle states
- Handle missing tasks and invalid state transitions
- Run automated unit tests

## Task Lifecycle

Pending -> Processing -> Completed

A task can only be processed while its state is Pending.
Processing a completed task returns an error.

## Requirements

- Go 1.27

## Run the Application

```bash
go run .
```

## Run Tests

```bash
go test ./...
```

## Code Quality Checks

```bash
gofmt -w .
go vet ./...
go test ./...
```
