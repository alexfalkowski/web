@benchmark
Feature: Benchmark Web
  Make sure these endpoints perform at their best.

  Scenario: Visit site in a good time frame and memory.
    When I visit the site which performs in 15 ms
    And the process 'server' should consume less than '70mb' of memory
