// Package security provides browser security headers for the site.
//
// The exported middleware sets the site's Content-Security-Policy,
// X-Content-Type-Options, Referrer-Policy, X-Frame-Options, Permissions-Policy,
// and Strict-Transport-Security headers. The CSP allowlist is coupled to the
// external assets used by the full-page templates, including jsDelivr-hosted
// HTMX/Pico CSS assets and Cloudflare Insights browser analytics endpoints.
package security
