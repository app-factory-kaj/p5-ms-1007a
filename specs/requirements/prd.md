# Greeter

## Problem Statement

Teams that need a tiny, dependable HTTP endpoint to greet a caller by name
currently have nothing lightweight to point at — they either skip the need or
reach for a heavier service. Greeter exists to be that small, predictable
building block.

## Solution

Greeter is a small Go HTTP service with a single job: given a name on a `GET /hello` request, it returns a JSON greeting. It follows the engineering
conventions of the organization's `app-factory-kaj/e2e-reference` project.
*assumed*

## Actors

- API Client — any caller (a person, script, or another service) that sends
an HTTP request to the greeter endpoint and reads the JSON response.

## Features

- F1 [Greeting API](features/F1-greeting-api.md)

## Product-wide

See [Product-wide](product-wide.md).

## Out of Scope

- Authentication or authorization on the endpoint. *assumed*
- Persisting greetings or any other data. *assumed*
- A user interface of any kind.