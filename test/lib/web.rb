# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

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
  end
end
