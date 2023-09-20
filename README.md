# Grafana HTTP OpenAPI Client for Go

> :warning: Grafana's OpenAPI Go client is currently under construction. Progress is tracked [here](https://github.com/grafana/grafana/issues/47827).

This HTTP Go client for [Grafana](https://github.com/grafana/grafana) is generated from [Grafana's OpenAPI specification](https://github.com/grafana/grafana/blob/main/public/api-merged.json) using [swagger for Go](https://github.com/go-swagger/go-swagger).

## Dependencies

This project uses [bingo](https://github.com/bwplotka/bingo) (located in [.bingo/](.bingo/)), a tool to automate the versioning of Go packages. Here, bingo manages verison of [swagger for Go](https://github.com/go-swagger/go-swagger) version so that the client generation is consistent.

**[Required]**  In order to generate the client, you must have installed bingo locally. Then, get swagger at the version specified in bingo's [swagger.mod](.bingo/swagger.mod).
```bash
go install github.com/bwplotka/bingo@latest
bingo get swagger
```

## Generate the Go client

You are ready to generate the client once bingo & swagger are installed (see [Dependencies](#dependencies)).

To generate the client for a specific API, find the name of its tag and model in the [API specification](https://github.com/grafana/grafana/blob/main/public/api-merged.json). Then, set those as environment variables and run the command to generate it:
```bash
export API_TAG=folders
export MODEL=Folder
make generate-client
```

## Custom templates

In order to generate the client, `go-swagger` uses default templates. These templates can be customised to add custom configuration that are applied each time the client is generated.

For more information, check out the `go-swagger` docs on how to [use custom templates](https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md). The default template definitions for the client can be found in [go-swagger/generator/templates/client/](https://github.com/go-swagger/go-swagger/tree/master/generator/templates/client).

In this project, the custom templates can be found in `templates/`. They are provided to the generation command through the flag `--template-dir=templates`.

The custom templates provide added functionality for things such as authentication, TLS/SSL, retries, and custom error handling.