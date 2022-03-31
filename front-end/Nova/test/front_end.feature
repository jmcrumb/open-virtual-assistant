Feature: User interacts with the front end graphic user interface

    Scenario: As a community developer, so that other users can download my plugins, [I want to be able to publish my own plugin]
        Given UI is running
        When UI serves plugin page
        And we interact to download the plugin "placeholder"
        Then we recieve confirmation that "placeholder" is downloaded

    Scenario: As a community developer, so that I can easily publish plugins, [I want a low friction way to add custom plugins using github.]
        Given UI is running
        When UI serves plugin page
        And we we fill out the plugin publish request
        Then confirms the plugin is published

    Scenario: As an end user, so that I can fully interact with my virtual assistant's responses, [I want differently abled accommodations reflected in the user interface.]
        Given UI is running
        When UI serves settings page
        Then accessibility settings exist 

    Scenario: As an end user, so that I can provide feedback to other community members regarding a plugin's efficacy, [I want to be able to rate and leave comments about a plugin.]
        Given UI is running
        When UI serves the "placeholder" plugin detail page
        Then we successfully create a review

    Scenario: As an end user, so that I can download other plugins, [I want to be able to browse available plugins.]
        Given UI is running
        When UI serves plugin page
        Then we can search for the "placeholder" plugin
