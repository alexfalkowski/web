# frozen_string_literal: true

When('I visit {string} with layout {string}') do |section, layout|
  @layout = layout
  headers = { user_agent: 'Web-client/1.0 HTTP/1.0', accept: 'text/html' }
  headers[:accept_language] = @preferred_languages if @preferred_languages
  opts = Web.http_options(headers: headers)
  methods = {
    'full' => 'get',
    'partial' => 'put'
  }
  method = methods[layout]

  @response = Web::V1.http.send("#{method}_#{section}", opts)
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

When('I visit a missing section with layout {string}') do |layout|
  @layout = layout
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

Then('I should see the robots file') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/plain; charset=utf-8')
  expect_security_headers(@response)
  expect_static_cache(@response)
  expect(@response.body).to include('User-agent: *')
  expect(@response.body).to include("Sitemap: #{Web::Site::SITEMAP_URL}")
end

Then('I should see the sitemap file') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/xml; charset=utf-8')
  expect_security_headers(@response)
  expect_static_cache(@response)

  xml = Nokogiri::XML.parse(@response.body)
  urls = xml.css('url loc').map(&:text)

  expect(urls).to eq(Web::Site::EXPECTED_SITEMAP_URLS)
end

Then('I should see the favicon') do
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('image/png')
  expect_security_headers(@response)
  expect_static_cache(@response)
  expect(@response.body.bytes.first(8)).to eq([137, 80, 78, 71, 13, 10, 26, 10])
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
