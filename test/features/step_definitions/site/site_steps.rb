# frozen_string_literal: true

When('I visit {string} with layout {string}') do |section, layout|
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

  expected = {
    'root' => 'Vince Lombardi',
    'books' => 'Margaret Fuller'
  }
  html = Nokogiri::HTML.parse(@response.body)

  expect(html.text).to include(expected[section])
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

  html = Nokogiri::HTML.parse(@response.body)

  expect(html.text).to include('Page not found')
  expect(html.text.include?('Copyright')).to eq(layout == 'full')
end
