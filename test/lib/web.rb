# frozen_string_literal: true

# Top-level namespace for the web service test client.
#
# This library is used by the Ruby acceptance tests under `test/` to interact
# with a running instance of the service.
#
# @see Web::V1
module Web
  require 'securerandom'
  require 'yaml'
  require 'base64'

  require 'nokogiri'

  require 'web/v1/http'

  # Versioned namespace for the web service test client.
  #
  # The V1 namespace provides a stable entrypoint for interacting with the
  # service over HTTP. It exposes a single convenience constructor, {.http},
  # which returns an instance of {Web::V1::HTTP}.
  #
  # The returned client is memoized for the life of the Ruby process, so repeated
  # calls to {.http} reuse the same client instance.
  #
  # @example Get a client and fetch the root page
  #   client = Web::V1.http
  #   response = client.get_root
  #
  # @example Pass request options through to the underlying HTTP client
  #   Web::V1.http.get_books(headers: { 'Accept' => 'text/html' })
  #
  # @see Web::V1::HTTP
  module V1
    class << self
      # Return the memoized HTTP client for API v1.
      #
      # The client is configured to talk to the base URL provided by the test
      # harness configuration (see `Nonnative.configuration.url`).
      #
      # @return [Web::V1::HTTP] an HTTP client bound to the configured base URL
      def http
        @http ||= Web::V1::HTTP.new(Nonnative.configuration.url)
      end
    end
  end
end
