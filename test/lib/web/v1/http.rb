# frozen_string_literal: true

module Web
  module V1
    class HTTP < Nonnative::HTTPClient
      def home(opts = {})
        get('/', opts)
      end
    end
  end
end
