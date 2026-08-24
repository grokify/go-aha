# Regenerating the OpenAPI Client

The `oag7/aha` package is **generated** by [OpenAPI Generator](https://openapi-generator.tech/)
(currently **7.24.0**) from the spec at `codegen/openapi_spec.yaml`. Do not edit
the generated files by hand — regenerate instead.

## Prerequisites

- Java (to run `openapi-generator-cli.jar`)
- `codegen/openapi-generator-cli.jar` (the generator; see `codegen/openapi-generator-version.txt`)
- The updated spec at `codegen/openapi_spec.yaml`

## Workflow

```bash
cd codegen

# 1. Generate the client into codegen/aha from openapi_spec.yaml
sh openapi-generator-command.sh

# 2. Drop the standalone module files the generator emits — the client is part
#    of the parent go-aha module, not its own module.
rm aha/go.mod aha/go.sum

cd ..

# 3. Swap the old client for the new one
rm -rf oag7/aha
mv codegen/aha oag7/aha

# 4. Resolve dependencies and verify
go mod tidy
go build ./...
go test ./...
```

### What the generator command does

`codegen/openapi-generator-command.sh`:

```bash
java -jar openapi-generator-cli.jar generate \
  -i openapi_spec.yaml -g go -o aha \
  --package-name=aha \
  --git-repo-id go-aha/v3/oag7/aha --git-user-id grokify \
  --additional-properties=disallowAdditionalPropertiesIfNotPresent=false

# Append a helper the generator does not provide, then format:
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> aha/client.go
gofmt -s -w aha/*.go
```

## Important: optional fields become pointers

OpenAPI Generator's Go generator renders schema properties based on whether they
are in the schema's `required` list:

| In `required`? | Generated Go type |
|----------------|-------------------|
| Yes            | value, e.g. `string` |
| No (optional)  | **pointer**, e.g. `*string` with `json:",omitempty"` |

So a property that is not marked `required` in `codegen/openapi_spec.yaml`
generates as a pointer. This is by design (it distinguishes "absent" from
"empty"), and the generated model docs say so explicitly, e.g.:

```
**ServiceName** | Pointer to **string** | | [optional]
```

### Consequences for hand-written wrappers

Hand-written packages that wrap the generated client (`oag7/features`,
`oag7/ideas`, …) must account for pointer fields. Two options:

1. **Use the generated safe getters** (recommended). Every optional field `X`
   has a `GetX() T` method that returns the zero value when the pointer is nil:

   ```go
   // instead of: intf.ServiceName == "jira"   (fails: *string vs string)
   if intf.GetServiceName() == "jira" && intf.GetName() == "key" {
       return strings.TrimSpace(intf.GetValue())
   }
   ```

2. **Mark the property `required` in the spec**, then regenerate. The field
   becomes a value type (`string`). Only do this when the API genuinely always
   returns the field, and be aware it changes the generated type for every
   consumer of that field.

## Post-regeneration checklist

1. `go build ./...` — fix any wrapper breakage (usually pointer/value mismatches;
   see above).
2. `go test ./...` — the generated `oag7/aha/test` suite should pass (its cases
   are skipped stubs).
3. Review renamed/removed schemas — the generator drops models the new spec no
   longer defines (e.g. the old `*Wrap` envelopes were superseded by `*Response`).
4. Update the README "Supported APIs" table if the surface changed.
5. `golangci-lint run` and commit the regenerated client and wrapper fixes.
