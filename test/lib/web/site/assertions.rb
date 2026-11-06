# frozen_string_literal: true

module Web
  module Site
    # Assertions mixed into Cucumber worlds, using the scenario response and layout.
    module Assertions
      def expect_security_headers(response)
        Web::Site::SECURITY_HEADERS.each do |header, value|
          expect(response.headers[header]).to eq(value)
        end
      end

      def expect_static_cache(response)
        expect(response.headers[:cache_control]).to eq('public, max-age=3600')
      end

      def expect_page_text(section, html)
        expect(html.text).to include(Web::Site::EXPECTED_PAGE_TEXT.fetch(section))
      end

      def expect_page_metadata(section, html, layout)
        metadata = Web::Site::EXPECTED_PAGE_METADATA.fetch(section)

        expect_page_title(html, metadata)
        expect_page_description(html, metadata, layout)
      end

      def expect_page_title(html, metadata)
        title = html.at_css('title')

        expect(title).not_to be_nil
        expect(title.text).to eq(metadata.fetch(:title))
      end

      def expect_page_description(html, metadata, layout)
        description = html.at_css('meta[name="description"]')

        expect(description).not_to be_nil
        expect(description[:content]).to eq(metadata.fetch(:description))
        expect(description[:'hx-swap-oob']).to eq('true') if layout == 'partial'
      end

      def expect_books(html)
        expect(book_headers(html)).to eq(Web::Site::EXPECTED_BOOK_HEADERS)
        expect(book_rows(html)).to eq(Web::Site::EXPECTED_BOOKS)
      end

      def book_headers(html)
        html.css('thead th').map(&:text)
      end

      def book_rows(html)
        html.css('tbody tr').map { |row| book_row(row) }
      end

      def book_row(row)
        cells = row.css('td')
        link = cells[4].at_css('a')

        {
          title: cells[0].text,
          author: cells[1].text,
          topics: cells[2].text,
          note: cells[3].text,
          link: link[:href]
        }
      end

      def expect_layout(html)
        return expect_full_layout(html) if @layout == 'full'

        expect_partial_layout(html)
      end

      def expect_full_layout(html)
        Web::Site::EXPECTED_NAV_LINKS.each do |link|
          expect(html.at_css(link.fetch(:selector)).text).to eq(link.fetch(:text))
        end
      end

      def expect_partial_layout(html)
        Web::Site::FULL_LAYOUT_MARKERS.each { |marker| expect(@response.body).not_to match(marker) }

        Web::Site::EXPECTED_NAV_LINKS.each do |link|
          expect(html.at_css(link.fetch(:selector))).to be_nil
        end

        expect(html.text).not_to include('Copyright')
      end
    end
  end
end
