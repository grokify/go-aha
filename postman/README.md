# Using Go-Aha with Postman

`go-aha` uses [`grokify/swaggman`](https://github.com/grokify/swaggman) to generate a usable Postman spec from the OpenAPI / Swagger 2.0 spec used to auto-generate the SDK.

Two environment variables are used:

* `AHA_ACCOUNT` which is the hostname subdomain for your account.
* `AHA_API_KEY` which is your API key.

Set these to values in your Postman file.

These are the same names and values used by [`grokify/oauth2more/aha`](https://github.com/grokify/oauth2more).