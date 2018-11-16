java -jar openapi-generator-cli.jar generate -i swagger_spec.yaml -g go -o client -D packageName=aha
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> client/client.go
gofmt -s -w client/*.go
