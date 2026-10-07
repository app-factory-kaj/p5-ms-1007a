# Greeting API

## Purpose

Lets an API Client get a JSON greeting for a given name over plain HTTP.

## User Stories

- F1.1 As an API Client, I send `GET /hello?name=X` and receive a JSON
greeting addressed to X.
- F1.2 As an API Client, I send `GET /hello` without a `name` and receive a
default JSON greeting rather than an error.

## Decisions

- When `name` is omitted, the response defaults to a generic greeting (e.g.
`{"message": "Hello, World!"}`) instead of an error.