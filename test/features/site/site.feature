Feature: Web
  Web is a website for lean-thoughts.com
  Site rendering scenarios assert successful HTML responses, security headers,
  expected page content, and the requested full or partial layout.

  Scenario Outline: Render sections with expected layout and headers
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>"

    Examples:
      | layout  | section |
      | full    | root    |
      | full    | books   |
      | partial | root    |
      | partial | books   |

  Scenario: Visit robots
    When I visit the robots file
    Then I should see the robots file

  Scenario: Visit sitemap
    When I visit the sitemap file
    Then I should see the sitemap file

  Scenario: Visit favicon
    When I visit the favicon
    Then I should see the favicon

  Scenario Outline: Missing section
    When I visit a missing section with layout "<layout>"
    Then I should see the page is missing with layout "<layout>"

    Examples:
      | layout  |
      | full    |
      | partial |
