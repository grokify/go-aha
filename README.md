# Go API client for Aha! (aha.io)

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/grokify/go-aha/actions/workflows/ci.yaml/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/go-aha/actions/workflows/ci.yaml
 [build-status-svg]: https://api.travis-ci.org/grokify/go-aha.svg?branch=master
 [build-status-url]: https://travis-ci.org/grokify/go-aha
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-aha
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-aha
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-aha
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-aha/v2
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-aha/blob/master/LICENSE

Go / Golang API Client for [Aha! Roadmap Service](https://www.aha.io/)

## Overview

This API client was generated by the [openapi-generator](https://github.com/OpenAPITools/openapi-generator) project.

It also comes with a Postman Collection at [`codegen/postman_spec.json`](codegen/postman_spec.json).

## Installation

```bash
$ go get github.com/grokify/go-aha/...
```

## Reference

See docs in:

* [`client/README.md`](client/README.md)

## Example Usage

See:

* [`examples/get_features`](examples/get_features)
* [`examples/get_products`](examples/get_products)

# Helpful Hints

To use curl to retreive an API response, use the following:

curl -XGET https://company.aha.io/api/v1/features/FEAT-1?access_token=<my_access_token>
