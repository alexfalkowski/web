# frozen_string_literal: true

When('I visit {string} with layout {string}') do |section, layout|
  opts = {
    headers: { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0' },
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
    'home' => 'Vince Lombardi',
    'books' => 'Margaret Fuller'
  }
  html = Nokogiri::HTML.parse(@response.body)

  expect(html.text).to include(expected[section])
end
