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

## Generate the Go client

Once bingo & swagger are installed (see [Dependencies](#dependencies)), generate the client:
```bash
make generate-client
```