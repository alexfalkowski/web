# frozen_string_literal: true

module Web
  # Expected site content and response policy for the HTTP acceptance scenarios.
  module Site
    SECURITY_HEADERS = {
      content_security_policy: "default-src 'self'; script-src 'self' https://cdn.jsdelivr.net " \
                               "https://static.cloudflareinsights.com; style-src 'self' https://cdn.jsdelivr.net " \
                               "'unsafe-inline'; img-src 'self'; " \
                               "connect-src 'self' https://cloudflareinsights.com; " \
                               "font-src 'self'; object-src 'none'; base-uri 'self'; frame-ancestors 'none'",
      x_content_type_options: 'nosniff',
      referrer_policy: 'strict-origin-when-cross-origin',
      x_frame_options: 'DENY',
      permissions_policy: 'camera=(), geolocation=(), microphone=()',
      strict_transport_security: 'max-age=86400'
    }.freeze

    EXPECTED_BOOKS = [
      {
        title: 'Kanban: Successful Evolutionary Change for Your Technology Business',
        author: 'David J. Anderson',
        topics: 'Kanban, evolutionary change, flow',
        note: 'A practical introduction to improving delivery with visual work management and evolutionary change.',
        link: 'https://www.amazon.de/-/en/David-J-Anderson/dp/0984521402'
      },
      {
        title: 'Modern Software Engineering: Doing What Works to Build Better Software Faster',
        author: 'David Farley',
        topics: 'software engineering, continuous delivery, quality',
        note: 'Connects engineering discipline with fast feedback, learning, and reliable software delivery.',
        link: 'https://www.amazon.de/-/en/Modern-Software-Engineering-Better-Faster/dp/0137314914'
      },
      {
        title: 'Team Topologies: Organizing Business and Technology Teams for Fast Flow',
        author: 'Matthew Skelton and Manuel Pais',
        topics: 'team design, cognitive load, flow',
        note: 'Explains how team structures and interaction modes shape software delivery outcomes.',
        link: 'https://www.amazon.de/-/en/Team-Topologies-Organizing-Business-Technology/dp/1942788819'
      }
    ].freeze

    EXPECTED_BOOK_HEADERS = %w[Title Author Topics Why Link].freeze

    EXPECTED_PAGE_TEXT = {
      'root' => 'Vince Lombardi',
      'books' => 'Margaret Fuller'
    }.freeze

    EXPECTED_PAGE_METADATA = {
      'root' => {
        title: 'Lean Thoughts',
        description: 'Lean Thoughts shares ideas about lean thinking and continuous improvement.'
      },
      'books' => {
        title: 'Books | Lean Thoughts',
        description: 'Recommended books about lean thinking, software delivery, and team flow.'
      },
      'missing' => {
        title: 'Page not found | Lean Thoughts',
        description: 'The requested Lean Thoughts page could not be found.'
      }
    }.freeze

    EXPECTED_SITEMAP_URLS = [
      'https://web.lean-thoughts.com/',
      'https://web.lean-thoughts.com/books'
    ].freeze

    SITEMAP_URL = 'https://web.lean-thoughts.com/sitemap.xml'

    EXPECTED_NAV_LINKS = [
      { selector: 'a[href="/"][hx-put="/"]', text: 'Home' },
      { selector: 'a[href="/books"][hx-put="/books"]', text: 'Books' }
    ].freeze

    FULL_LAYOUT_MARKERS = [
      /<!doctype/i,
      /<html\b/i,
      /<body\b/i,
      /<nav\b/i
    ].freeze
  end
end
