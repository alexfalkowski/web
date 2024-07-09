# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'nokogiri'

require 'web/v1/http'

module Web
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:11000')
    end

    def server_config
      @server_config ||= Nonnative.configurations('.config/server.yml')
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Web::V1::HTTP.new('http://localhost:11000')
      end
    end
  end
end
