# frozen_string_literal: true

module Web
  module V1
    class HTTP < Nonnative::HTTPClient
      def root(opts = {})
        get('/', opts)
      end

      def home(opts = {})
        get('/home', opts)
      end

      def books(opts = {})
        get('/books', opts)
      end
    end
  end
end
