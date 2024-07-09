# frozen_string_literal: true

When('I visit home') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0'
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Web::V1.server_http.home(opts)
end

Then('I should be at home') do
  expect(@response.code).to eq(200)

  html = Nokogiri::HTML.parse(@response.body)
  expect(html.title).to eq('Lean Thoughts')
end
