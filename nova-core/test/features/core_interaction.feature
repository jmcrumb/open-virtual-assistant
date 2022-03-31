Feature: User interacts with NOVA core

  Scenario: run an unrecognized command
     Given nova is running
      When we ask "sdfghetyjtyers"
      And wait "1" seconds
      Then output will include "I'm sorry, I don't understand."

  Scenario: run a basic command
    Given nova is running
    When we ask "hello there"
    And wait "1" seconds
    Then output will include "Hello! My name is Nova."

  Scenario: As an end user, so that the virtual assistant will understand me, [I want text to speech.]
    Given nova is running
    When we ask "hello there"
    And wait "1" seconds
    Then the output is successfully converted to audio

  Scenario: As an end user, so that I can interact with the virtual assistant within one click or voice invocation, [I want ease of use.]
    Given nova is running
    When we ask "nova"
    Then "nova" recognizes its name

  Scenario: As an end user, so that the virtual assistant will understand me, [I want speech to text.]
    Given nlp is running
    When we input to the microphone
    Then we recieve a text transcription

  Scenario: As an end user, so that I can feel more engaged with my assistant, [I want to be responded to by name.]
    Given nova is running
    When we ask "hello there"
    And wait "1" seconds
    Then output will include "john doe"

  Scenario: As an end user, so that I can have a conversational experience less distinguishable from a conversation with another person (as measured by user testing), [I want the virtual assistant to understand the secondary command (a follow up command relation to a primary, keyword-based command).]
    Given nova is running
    When we ask "hello"
    And wait "1" seconds
    And we ask "howdy"
    And wait "1" seconds
    Then output will include "To teach me more fun things to do, go to the plugin store."

  # Scenario: As an end user, so that I don’t become frustrated with my virtual assistant, [I want to be able to stop the execution of any command immediately.]

  Scenario: As an end user, so that I can add functionality to my virtual assistant, [I want to be able to download others' plugins.]
    Given main is running
    When we request the plugin "placeholder" download via CLI
    And wait "5" seconds
    Then the plugin "placeholder" is in the plugin registry

  Scenario: As a community developer, so that I can provide a good experience to users, [I want the platform to remember previous invocations of a specific command.]
    Given nova is running
    When we ask "what is the surf like at ocean beach"
    And we ask "hello"
    And we ask "what is the surf"
    And wait "2" seconds
    Then output will include "ocean beach"



  