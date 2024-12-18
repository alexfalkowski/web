# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'nokogiri'

require 'web/v1/http'

module Web
  class << self
    def config
      @config ||= Nonnative.configurations('.config/server.yml')
    end
  end

  module V1
    class << self
      def http
        @http ||= Web::V1::HTTP.new(Nonnative.configuration.url)
      end
    end
  end
end
