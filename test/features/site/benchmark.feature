@benchmark
Feature: Benchmark Web
  Make sure the root page performs at its best.

  Scenario: Visit the root page in a good time frame and memory.
    When I visit the site which performs in 15 ms
    And the process 'server' should consume less than '70mb' of memory
