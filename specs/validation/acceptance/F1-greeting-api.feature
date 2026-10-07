Feature: F1 Greeting API

  @story-F1.1
  Rule: A named caller is greeted by name

    Scenario: Greeting a given name
      When an API Client sends GET /hello?name=Ada
      Then the response is a JSON greeting "Hello, Ada"

  @story-F1.2
  Rule: A caller who gives no name still gets a greeting

    Scenario: Greeting with no name given
      When an API Client sends GET /hello with no name parameter
      Then the response is a JSON greeting "Hello, World!"

  @story-F1.3
  Rule: A name over 40 characters is refused

    @negative
    Scenario: An overlong name is rejected
      When an API Client sends GET /hello with a name 41 characters long
      Then the response is a 400 error and no greeting is returned
