# frozen_string_literal: true

Given('my preferred languages are {string}') do |preferences|
  @preferred_languages = preferences
end

Then('I should see {string} in {string}') do |section, language|
  if language == 'en'
    if section == 'missing'
      step("I should see the page is missing with layout \"#{@layout}\"")
    else
      step("I should see \"#{section}\"")
    end
    next
  end

  expect(@response.code).to eq(section == 'missing' ? 404 : 200)
  expect(@response.headers[:content_type]).to eq('text/html; charset=utf-8')
  expect_security_headers(@response)

  html = Nokogiri::HTML.parse(@response.body)
  page = Web::Locale::PAGES.fetch(language).fetch(section)

  expect_page_title(html, page)
  expect_page_description(html, page, @layout)
  expect(html.at_css('h1').text).to eq(page.fetch(:heading))
  expect(html.at_css('.introduction').text.strip).to eq(page.fetch(:introduction))
  expect(html.at_css('blockquote').text).to include(page.fetch(:quote)) if page.key?(:quote)
  expect(html.at_css('a[href^="mailto:"]').text).to eq(Web::Locale::LAYOUTS.fetch(language).fetch(:contact)) if section == 'root'
  expect_localized_books(html, language) if section == 'books'
  expect_localized_layout(html, language)
end

Then('the response language should be {string}') do |language|
  expect(@response.headers[:content_language]).to eq(language)
  expect(@response.headers[:vary].split(',').map(&:strip)).to include('Accept-Language')

  html = Nokogiri::HTML.parse(@response.body)
  expect(html.at_css('html')[:lang]).to eq(language) if @layout == 'full'
end
