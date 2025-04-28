Feature: Web
  Web is a website for lean-thoughts.com

  Scenario Outline: Visit sections
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>"

    Examples:
      | layout  | section |
      | full    | root    |
      | full    | home    |
      | full    | books   |
      | partial | root    |
      | partial | home    |
      | partial | books   |
