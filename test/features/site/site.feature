Feature: Web
  Web is a website for lean-thoughts.com
  Site rendering scenarios assert successful HTML responses, security headers,
  expected page content, and the requested full or partial layout.

  Scenario Outline: Render sections with expected layout and headers
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>"
    And the response language should be "en"

    Examples:
      | layout  | section |
      | full    | root    |
      | full    | books   |
      | partial | root    |
      | partial | books   |

  Scenario Outline: Render sections in the preferred language
    Given my preferred languages are "<preferences>"
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>" in "<language>"
    And the response language should be "<language>"

    Examples:
      | layout  | section | preferences                | language |
      | full    | root    | de                         | de       |
      | partial | root    | de                         | de       |
      | full    | books   | de                         | de       |
      | partial | books   | de                         | de       |
      | full    | missing | de                         | de       |
      | partial | missing | de                         | de       |
      | full    | root    | es                         | es       |
      | partial | root    | es                         | es       |
      | full    | books   | es                         | es       |
      | partial | books   | es                         | es       |
      | full    | missing | es                         | es       |
      | partial | missing | es                         | es       |
      | full    | root    | es-MX                      | es       |
      | partial | books   | de-CH                      | de       |
      | full    | books   | es;q=0.4,de;q=0.9,en;q=0.1 | de       |
      | partial | root    | de;q=0.4,es;q=0.9,en;q=0.1 | es       |

  Scenario Outline: Fall back to English for unsupported or invalid preferences
    Given my preferred languages are "<preferences>"
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>" in "en"
    And the response language should be "en"

    Examples:
      | layout  | section | preferences |
      | full    | root    | ja          |
      | partial | books   | ja          |
      | full    | missing | ja          |
      | partial | missing | ja          |
      | partial | root    | invalid_!   |
      | full    | books   | invalid_!   |

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
    And the response language should be "en"

    Examples:
      | layout  |
      | full    |
      | partial |
