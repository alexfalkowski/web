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
    link: 'https://www.amazon.de/-/en/David-J-Anderson/dp/0984521402'
  },
  {
    title: 'Modern Software Engineering: Doing What Works to Build Better Software Faster',
    link: 'https://www.amazon.de/-/en/Modern-Software-Engineering-Better-Faster/dp/0137314914'
  },
  {
    title: 'Team Topologies: Organizing Business and Technology Teams for Fast Flow',
    link: 'https://www.amazon.de/-/en/Team-Topologies-Organizing-Business-Technology/dp/1942788819'
  }
].freeze

EXPECTED_PAGE_TEXT = {
  'root' => 'Vince Lombardi',
  'books' => 'Margaret Fuller'
}.freeze

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
  opts = {
    headers: { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0', accept: 'text/html' },
    read_timeout: 10, open_timeout: 10
  }
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
  expect_books(html) if section == 'books'
  expect_layout(html)
end

When('I visit the robots file') do
  opts = {
    headers: { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0' },
    read_timeout: 10, open_timeout: 10
  }

  @response = Web::V1.http.get_robots(opts)
end

When('I visit the favicon') do
  opts = {
    headers: { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0' },
    read_timeout: 10, open_timeout: 10
  }

  @response = Web::V1.http.get_favicon(opts)
end

Then('I should see the robots file') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/plain; charset=utf-8')
  expect_security_headers(@response)
  expect(@response.body).to include('User-agent: *')
end

Then('I should see the favicon') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('image/png')
  expect_security_headers(@response)
  expect(@response.body.bytes.first(8)).to eq([137, 80, 78, 71, 13, 10, 26, 10])
end

When('I visit a missing section with layout {string}') do |layout|
  headers = { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0' }
  layout_headers = {
    'full' => { accept: 'text/html' },
    'partial' => { hx_request: 'true' }
  }
  opts = {
    headers: headers.merge(layout_headers[layout]),
    read_timeout: 10, open_timeout: 10
  }
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

def expect_books(html)
  expect(book_rows(html)).to eq(EXPECTED_BOOKS)
end

def book_rows(html)
  html.css('tbody tr').map { |row| book_row(row) }
end

def book_row(row)
  cells = row.css('td')
  link = cells[1].at_css('a')

  { title: cells[0].text, link: link[:href] }
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
