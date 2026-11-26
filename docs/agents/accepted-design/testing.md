# Accepted design for testing

This entry is part of the [accepted-design index](../accepted-design.md).

Ruby/Cucumber scenarios are the authoritative product behavior checks. Use
`make features` for acceptance tests and `make benchmarks` for performance
scenarios. `make specs` supports Go build and tooling checks.
[Feature suite](../../../test/features) · [CI checks](../../../.circleci/config.yml)
