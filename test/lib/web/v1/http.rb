# frozen_string_literal: true

module Web
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_root(opts = {})
        get('/', opts)
      end

      def put_root(opts = {})
        put('/', opts)
      end

      def get_home(opts = {})
        get('/home', opts)
      end

      def put_home(opts = {})
        put('/home', opts)
      end

      def get_books(opts = {})
        get('/books', opts)
      end

      def put_books(opts = {})
        put('/books', opts)
      end
    end
  end
end
