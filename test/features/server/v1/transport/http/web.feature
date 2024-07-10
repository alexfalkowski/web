Feature: Web

  Web is a website for lean-thoughts.com

  Scenario: Root
    When I visit root
    Then I should see root

  Scenario: Home
    When I visit home
    Then I should see home

   Scenario: Books
    When I visit books
    Then I should see books
