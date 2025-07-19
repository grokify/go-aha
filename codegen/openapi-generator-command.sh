java -jar openapi-generator-cli.jar generate -i openapi_spec.json -g go -o aha --package-name=aha --additional-properties=disallowAdditionalPropertiesIfNotPresent=false
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> aha/client.go
gofmt -s -w aha/*.go