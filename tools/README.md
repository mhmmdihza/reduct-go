# Simplify Client Integration with Reduct API using go-swagger

This tool generates client-side code for interacting with the Reduct API using [go-swagger](https://github.com/go-swagger/go-swagger), based on OpenAPI YAML definitions located in the `reduct-api-docs` directory.

## How to Use

1. **Install dependencies**  
   Run in the `tools` folder:
   ```
   go mod tidy

This installs code generation dependencies without affecting the SDK's core dependencies.

2. **Generate client code**

Execute from `tools` directory:

```
go generate -x tools.go
```
## Important Notes

tools.go contains the command:

```
go tool yq eval-all --inplace \
  "select(fileIndex == 0) * select(fileIndex == 1) * select(fileIndex == 2) * select(fileIndex == 3)" \
  output.yaml \
  reduct-api-docs/bucket.yaml \
  reduct-api-docs/entry-read.yaml \
  reduct-api-docs/entry-write.yaml
```
This command merges multiple OpenAPI YAML files into a single output.yaml file.

## Considerations:
- The `--inplace` flag is required because go generate cannot use redirection operators like >.

- target file /`output.yaml` must be defined as first param and it will be the target file modified during merging.

- You need to explicitly select each file with `select(fileIndex == N)` for every file you want to merge.

- This method is a bit hacky, but it ensures compatibility with just Go 1.24—no additional dependencies are needed, regardless of whether the machine is running Windows, Linux, or another OS.
