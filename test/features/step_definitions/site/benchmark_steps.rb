# frozen_string_literal: true

When('I visit the site which performs in {int} ms') do |time|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Status-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  expect { Web::V1.http.get_root(opts) }.to perform_under(time).ms
end
