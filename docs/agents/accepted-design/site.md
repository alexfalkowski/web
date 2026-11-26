# Accepted design for the site

These entries are part of the [accepted-design index](../accepted-design.md).

## Shared metadata

Compute the footer year at startup from the local clock and share it through
dependency injection. A running instance keeps that value through a calendar
rollover. Revisit this behavior when a task explicitly changes metadata
freshness. [Metadata constructor](../../../internal/site/meta/meta.go) ·
[Module wiring](../../../internal/site/meta/module.go)

## Embedded asset caching

Use the shared MVC static-file handler for `/robots.txt`, `/sitemap.xml`, and
`/favicon.ico`. Web sets `Cache-Control: public, max-age=3600` on all three.
The handler sends the full asset on each successful GET request. Revisit
conditional responses if shared MVC gains a reliable content validator for
embedded assets.
[Static registration](../../../internal/site/static/static.go) ·
[go-service version](../../../go.mod)

## Browser assets and content security policy

The full layout loads pinned HTMX and individual daisyUI stylesheets from
jsDelivr with integrity hashes. Keep the stylesheets separate to avoid source
map requests outside the `connect-src` allowlist. The CSP permits jsDelivr for
these resources, Cloudflare Insights for production analytics, and inline
styles for the layout's embedded CSS. Review the layout and CSP together when
changing these assets. [Full layout](../../../internal/site/root/layout/full.tmpl) ·
[Security headers](../../../internal/site/security/headers.go)
