# frozen_string_literal: true

When('I visit root') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0'
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Web::V1.server_http.root(opts)
end

When('I visit home') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0'
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Web::V1.server_http.home(opts)
end

When('I visit books') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0'
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Web::V1.server_http.books(opts)
end

Then('I should see root') do
  expect(@response.code).to eq(200)

  html = Nokogiri::HTML.parse(@response.body)
  expect(html.title).to eq('Lean Thoughts')
end

Then('I should see home') do
  expect(@response.code).to eq(200)

  html = Nokogiri::HTML.parse(@response.body)
  expect(html.text).to include('Lean Thinking')
end

Then('I should see books') do
  expect(@response.code).to eq(200)

  html = Nokogiri::HTML.parse(@response.body)
  expect(html.text).to include('Kanban')
end
