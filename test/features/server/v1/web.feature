Feature: Web
  Web is a website for lean-thoughts.com

  Scenario Outline: Visit sections
    When I visit "<section>"
    Then I should see "<section>"

    Examples:
      | section |
      | root    |
      | home    |
      | books   |
