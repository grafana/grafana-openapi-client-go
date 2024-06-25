# Grafana HTTP OpenAPI Client for Go

This HTTP Go client for [Grafana](https://github.com/grafana/grafana) is generated from [Grafana's OpenAPI specification](https://github.com/grafana/grafana/blob/main/public/api-merged.json) using [swagger for Go](https://github.com/go-swagger/go-swagger).

Both code and non-code contributions are very welcome - please feel free to open an issue or PR in this repository if you spot a possible improvement!

## Dependencies

This project uses [bingo](https://github.com/bwplotka/bingo) (located in [.bingo/](.bingo/)), a tool to automate the versioning of Go packages. Here, bingo manages verison of [swagger for Go](https://github.com/go-swagger/go-swagger) version so that the client generation is consistent.

**[Required]**  In order to generate the client, you must have installed bingo locally. Then, get swagger at the version specified in bingo's [swagger.mod](.bingo/swagger.mod).
```bash
go install github.com/bwplotka/bingo@latest
bingo get swagger
```

## Generate the client

Once bingo & swagger are installed (see [Dependencies](#dependencies)), generate the client for _all Grafana APIs_, with the following `make` command:

```bash
make generate-client
```

This runs the Swagger generation command.

To generate the client for a _specific Grafana API_, find the name of its tag and model in the [Grafana OpenAPI specification](https://github.com/grafana/grafana/blob/main/public/api-merged.json). Then, set those as environment variables and run the command to generate it:
```bash
export API_TAG=folders
export MODEL=Folder
make generate-client
```

### How to use custom templates

In order to generate the client, `go-swagger` uses default templates. These templates can be customised to add custom configuration that are applied each time the client is generated.

For more information, check out the `go-swagger` docs on how to [use custom templates](https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md). The default template definitions for the client can be found in [go-swagger/generator/templates/client/](https://github.com/go-swagger/go-swagger/tree/master/generator/templates/client).

In this project, the custom templates can be found in `templates/`. They are provided to the generation command through the flag `--template-dir=templates`.

The custom templates provide added functionality for things such as authentication, TLS/SSL, retries, and custom error handling.

## Build the client

### Configuration

The client has the following friendly configuration options:

```go
import goapi "github.com/grafana/grafana-openapi-client-go/client"

cfg := &goapi.TransportConfig{
    // Host is the doman name or IP address of the host that serves the API.
    Host:       "localhost:3000",
    // BasePath is the URL prefix for all API paths, relative to the host root.
    BasePath:   "/api",
    // Schemes are the transfer protocols used by the API (http or https).
    Schemes:    []string{"http"},
    // APIKey is an optional API key or service account token.
    APIKey:     os.Getenv("API_ACCESS_TOKEN"),
    // BasicAuth is optional basic auth credentials.
    BasicAuth:  url.UserPassword("admin", "admin"),
    // OrgID provides an optional organization ID.
    // OrgID is only supported with BasicAuth since API keys are already org-scoped.
    OrgID:      1,
    // TLSConfig provides an optional configuration for a TLS client
    TLSConfig:  &tls.Config{},
    // NumRetries contains the optional number of attempted retries
    NumRetries: 3,
    // RetryTimeout sets an optional time to wait before retrying a request
    RetryTimeout: 0,
    // RetryStatusCodes contains the optional list of status codes to retry
    // Use "x" as a wildcard for a single digit (default: [429, 5xx])
    RetryStatusCodes: []string{"420", "5xx"},
    // HTTPHeaders contains an optional map of HTTP headers to add to each request
    HTTPHeaders: map[string]string{},
}

client := goapi.NewHTTPClientWithConfig(strfmt.Default, cfg)
```

### Examples

Checkout how the Grafana Terraform Provider initialises and uses the client [here](https://github.com/grafana/terraform-provider-grafana/blob/2988bb3560acc55f3e686532f44109a224825568/internal/provider/provider.go#L419-L446).

The `goswagger` documentation have more information about how to [build a client](https://goswagger.io/go-swagger/generate/client/).

### Roadmap

We are planning a few improvements around processes such as automation, testing, release, and integration into other dependencies. Some of this work is tracked [here](https://github.com/grafana/grafana/issues/47827).
