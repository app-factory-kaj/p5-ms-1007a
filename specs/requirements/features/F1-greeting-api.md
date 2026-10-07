# Greeting API

## Purpose

Lets an API Client get a JSON greeting for a given name over plain HTTP.

## User Stories

- F1.1 As an API Client, I send `GET /hello?name=X` and receive a JSON
greeting addressed to X.
- F1.2 As an API Client, I send `GET /hello` without a `name` and receive a
default JSON greeting rather than an error.
- F1.3 As an API Client, I send `GET /hello` with a `name` longer than 40
characters and receive a 400 error rather than a greeting. \[Greeting style guide\]

## Decisions

- When `name` is omitted, the response defaults to a generic greeting (e.g.
`{"message": "Hello, World!"}`) instead of an error.
- The greeting message is the word "Hello" followed by the name (e.g.
`Hello, Ada`), kept short and friendly. \[Greeting style guide\]
- A `name` longer than 40 characters is rejected with a 400 error rather than
being greeted. \[Greeting style guide\]