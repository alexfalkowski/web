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
    #   client = Web::V1::HTTP.new("http://localhost:11000")
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
        put('/', nil, opts)
      end

      # GET the books page.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_books(opts = {})
        get('/books', opts)
      end

      # GET the robots.txt file.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_robots(opts = {})
        get('/robots.txt', opts)
      end

      # GET the browser favicon.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_favicon(opts = {})
        get('/favicon.ico', opts)
      end

      # PUT the books page, typically used to request a partial/fragment render.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#put}
      # @return [Object] the response returned by the underlying HTTP client
      def put_books(opts = {})
        put('/books', nil, opts)
      end

      # GET a route that should not be handled by the site.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#get}
      # @return [Object] the response returned by the underlying HTTP client
      def get_missing(opts = {})
        get('/not-a-real-page', opts)
      end

      # PUT a route that should not be handled by the site.
      #
      # @param opts [Hash] request options forwarded to {Nonnative::HTTPClient#put}
      # @return [Object] the response returned by the underlying HTTP client
      def put_missing(opts = {})
        put('/not-a-real-page', nil, opts)
      end
    end
  end
end
