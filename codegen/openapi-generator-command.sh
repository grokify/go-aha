java -jar openapi-generator-cli.jar generate -i openapi_spec.yaml -g go -o aha --package-name=aha --git-repo-id go-aha/v3/oag7/aha --git-user-id grokify --additional-properties=disallowAdditionalPropertiesIfNotPresent=false
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> aha/client.go
gofmt -s -w aha/*.go
