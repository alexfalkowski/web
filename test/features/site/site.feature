Feature: Web
  Web is a website for lean-thoughts.com

  Scenario Outline: Visit sections
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>"

    Examples:
      | layout  | section |
      | full    | root    |
      | full    | books   |
      | partial | root    |
      | partial | books   |

  Scenario: Missing section
    When I visit a missing section
    Then I should see the page is missing
