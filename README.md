# tf-pd-import

[![Build Status](https://travis-ci.org/darkowlzz/tf-pd-import.svg?branch=master)](https://travis-ci.org/darkowlzz/tf-pd-import)

Terraform Pagerduty Resource Importer and Config Generator.


## Setup

1. Clone the repo & `cd` into it.
2. Ensure that [`govendor`](https://github.com/kardianos/govendor) is installed
and install all the vendor dependencies with `govendor sync`.

## Development

* Add new dependencies with `govendor fetch <packagename>`. This would install
the dependencies under `vendor/` and add them to `vendor/vendor.json`, which
should be checked-in.

## Tests

`go test`
