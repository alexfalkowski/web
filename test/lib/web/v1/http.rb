# frozen_string_literal: true

module Web
  module V1
    # HTTP client for the web service (API v1) used by acceptance tests.
    #
    # This class provides small, intention-revealing wrappers around
    # {Nonnative::HTTPClient} for the endpoints exercised by the test suite.
    #
    # Each helper delegates to the underlying HTTP verb method (`get`, `put`, …)
    # and forwards an optional options hash to the base client.
    #
    # @example Fetch the root page
    #   client = Web::V1::HTTP.new("http://localhost:8080")
    #   response = client.get_root
    #
    # @example Request the partial version of a page (PUT)
    #   response = client.put_books
    #
    # @example Pass options through to the underlying client (headers, query, etc.)
    #   response = client.get_books(headers: { "Accept" => "text/html" })
    #
    # @see Web::V1.http Convenience constructor for a memoized instance
    class HTTP < Nonnative::HTTPClient
      # GET the root page.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_root(opts = {})
        get('/', opts)
      end

      # PUT the root page, typically used to request a partial/fragment render.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#put}
      # @return [Object] the response returned by the underlying HTTP client
      def put_root(opts = {})
        put('/', opts)
      end

      # GET the home page.
      #
      # Note: whether `/home` is distinct from `/` is determined by the running
      # service; this helper exists because the test suite exercises it.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_home(opts = {})
        get('/home', opts)
      end

      # PUT the home page, typically used to request a partial/fragment render.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#put}
      # @return [Object] the response returned by the underlying HTTP client
      def put_home(opts = {})
        put('/home', opts)
      end

      # GET the books page.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_books(opts = {})
        get('/books', opts)
      end

      # PUT the books page, typically used to request a partial/fragment render.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#put}
      # @return [Object] the response returned by the underlying HTTP client
      def put_books(opts = {})
        put('/books', opts)
      end
    end
  end
end
