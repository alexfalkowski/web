# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'nokogiri'

require 'web/v1/http'

module Web
  module V1
    class << self
      def http
        @http ||= Web::V1::HTTP.new(Nonnative.configuration.url)
      end
    end
  end
end
