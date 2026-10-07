# Greeting API

## Purpose

Lets an API Client get a JSON greeting for a given name over plain HTTP.

## Open Questions

1. Should `GET /hello` require the `name` query parameter, or return a
 default greeting when it is omitted? *blocking*
   - Require `name`; respond with an error when it is missing
   - Default to a generic greeting (e.g. "Hello, World!") when `name` is missing