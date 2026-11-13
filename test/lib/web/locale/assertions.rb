# frozen_string_literal: true

module Web
  module Locale
    # Assertions mixed into Cucumber worlds, using the scenario response and layout.
    module Assertions
      def expect_localized_books(html, language)
        expected = Web::Site::EXPECTED_BOOKS.zip(Web::Locale::BOOK_DETAILS.fetch(language)).map { |book, details| book.merge(details) }

        expect(book_cards(html)).to eq(expected)
        expect_book_links(html, Web::Locale::BOOK_LINK_TEXT.fetch(language))
      end

      def expect_localized_layout(html, language)
        return expect_partial_layout(html) if @layout == 'partial'

        layout = Web::Locale::LAYOUTS.fetch(language)

        {
          'nav li a[href="/"][hx-put="/"]' => layout.fetch(:home),
          'nav li a[href="/books"][hx-put="/books"]' => layout.fetch(:books),
          'a[href="https://github.com/Lean-Thoughts"]' => layout.fetch(:code)
        }.each do |selector, text|
          expect(html.at_css(selector).text).to eq(text)
        end

        expect_localized_footer(html, layout.fetch(:rights))
      end

      def expect_localized_footer(html, rights)
        expect(html.at_css('body > footer').text.strip).to match(
          /\ACopyright © \d{4} Lean Thoughts \(v\S+\)\. #{Regexp.escape(rights)}\z/
        )
      end
    end
  end
end
