# frozen_string_literal: true

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

When('I visit {string} with layout {string}') do |section, layout|
  @layout = layout
  opts = Web.http_options(headers: { user_agent: 'Web-client/1.0 HTTP/1.0', accept: 'text/html' })
  methods = {
    'full' => 'get',
    'partial' => 'put'
  }
  method = methods[layout]

  @response = Web::V1.http.send("#{method}_#{section}", opts)
end

Then('I should see {string}') do |section|
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/html; charset=utf-8')
  expect_security_headers(@response)

  html = Nokogiri::HTML.parse(@response.body)

  expect_page_text(section, html)
  expect_page_metadata(section, html, @layout)
  expect_books(html) if section == 'books'
  expect_layout(html)
end

When('I visit the robots file') do
  opts = Web.http_options(headers: { user_agent: 'Web-client/1.0 HTTP/1.0' })

  @response = Web::V1.http.get_robots(opts)
end

When('I visit the sitemap file') do
  opts = Web.http_options(headers: { user_agent: 'Web-client/1.0 HTTP/1.0' })

  @response = Web::V1.http.get_sitemap(opts)
end

When('I visit the favicon') do
  opts = Web.http_options(headers: { user_agent: 'Web-client/1.0 HTTP/1.0' })

  @response = Web::V1.http.get_favicon(opts)
end

Then('I should see the robots file') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/plain; charset=utf-8')
  expect_security_headers(@response)
  expect(@response.body).to include('User-agent: *')
  expect(@response.body).to include("Sitemap: #{SITEMAP_URL}")
end

Then('I should see the sitemap file') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/xml; charset=utf-8')
  expect_security_headers(@response)

  xml = Nokogiri::XML.parse(@response.body)
  urls = xml.css('url loc').map(&:text)

  expect(urls).to eq(EXPECTED_SITEMAP_URLS)
end

Then('I should see the favicon') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('image/png')
  expect_security_headers(@response)
  expect(@response.body.bytes.first(8)).to eq([137, 80, 78, 71, 13, 10, 26, 10])
end

When('I visit a missing section with layout {string}') do |layout|
  headers = { user_agent: 'Web-client/1.0 HTTP/1.0' }
  layout_headers = {
    'full' => { accept: 'text/html' },
    'partial' => { hx_request: 'true' }
  }
  opts = Web.http_options(headers: headers.merge(layout_headers[layout]))
  methods = {
    'full' => 'get',
    'partial' => 'put'
  }
  method = methods[layout]

  @response = Web::V1.http.send("#{method}_missing", opts)
end

Then('I should see the page is missing with layout {string}') do |layout|
  expect(@response.code).to eq(404)
  expect(@response.headers[:content_type]).to eq('text/html; charset=utf-8')
  expect_security_headers(@response)

  html = Nokogiri::HTML.parse(@response.body)

  expect(html.text).to include('Page not found')
  expect_page_metadata('missing', html, layout)
  expect(html.text.include?('Copyright')).to eq(layout == 'full')
end

def expect_security_headers(response)
  SECURITY_HEADERS.each do |header, value|
    expect(response.headers[header]).to eq(value)
  end
end

def expect_page_text(section, html)
  expect(html.text).to include(EXPECTED_PAGE_TEXT.fetch(section))
end

def expect_page_metadata(section, html, layout)
  metadata = EXPECTED_PAGE_METADATA.fetch(section)

  expect_page_title(html, metadata)
  expect_page_description(html, metadata, layout)
end

def expect_page_title(html, metadata)
  title = html.at_css('title')

  expect(title).not_to be_nil
  expect(title.text).to eq(metadata.fetch(:title))
end

def expect_page_description(html, metadata, layout)
  description = html.at_css('meta[name="description"]')

  expect(description).not_to be_nil
  expect(description[:content]).to eq(metadata.fetch(:description))
  expect(description[:'hx-swap-oob']).to eq('true') if layout == 'partial'
end

def expect_books(html)
  expect(book_headers(html)).to eq(EXPECTED_BOOK_HEADERS)
  expect(book_rows(html)).to eq(EXPECTED_BOOKS)
end

def book_headers(html)
  html.css('thead th').map(&:text)
end

def book_rows(html)
  html.css('tbody tr').map { |row| book_row(row) }
end

def book_row(row)
  cells = row.css('td')
  link = cells[4].at_css('a')

  {
    title: cells[0].text,
    author: cells[1].text,
    topics: cells[2].text,
    note: cells[3].text,
    link: link[:href]
  }
end

def expect_layout(html)
  return expect_full_layout(html) if @layout == 'full'

  expect_partial_layout(html)
end

def expect_full_layout(html)
  EXPECTED_NAV_LINKS.each do |link|
    expect(html.at_css(link.fetch(:selector)).text).to eq(link.fetch(:text))
  end
end

def expect_partial_layout(html)
  FULL_LAYOUT_MARKERS.each { |marker| expect(@response.body).not_to match(marker) }

  EXPECTED_NAV_LINKS.each do |link|
    expect(html.at_css(link.fetch(:selector))).to be_nil
  end

  expect(html.text).not_to include('Copyright')
end
