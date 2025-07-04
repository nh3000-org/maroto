# Maroto V2

[![GoDoc](https://godoc.org/github.com/nh3000-org/maroto?status.svg)](https://pkg.go.dev/github.com/nh3000-org/maroto/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/nh3000-org/maroto)](https://goreportcard.com/report/github.com/nh3000-org/maroto)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines)  
[![CI](https://github.com/nh3000-org/maroto/actions/workflows/goci.yml/badge.svg)](https://github.com/nh3000-org/maroto/actions/workflows/goci.yml)
[![Lint](https://github.com/nh3000-org/maroto/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/nh3000-org/maroto/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/nh3000-org/maroto)](https://codecov.io/gh/nh3000-org/maroto)
[![Visits Badge](https://badges.pufler.dev/visits/nh3000-org/maroto)](https://badges.pufler.dev)
[![Stars Badge](https://img.shields.io/github/stars/nh3000-org/maroto.svg?style=social&label=Stars)](https://github.com/nh3000-org/maroto/stargazers)


A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

![sirmaroto](docs/assets/images/logosmall.png)
> Maroto definition: Brazilian expression, means an astute/clever/intelligent person. 
> [Art by **@marinabankr**](https://www.instagram.com/marinabankr/)

You can write your PDFs like you are creating a site using Bootstrap. A Row may have many Cols, and a Col may have many components. 
Besides that, pages will be added when content may extrapolate the useful area. You can define a header which will be added
always when a new page appear, in this case, a header may have many rows, lines or tablelist. 

#### Maroto `v2.3.1` is here! Try out:

* Installation with`go get`:

```bash
go get github.com/nh3000-org/maroto/v2@v2.3.1
```

* You can see the full `v2` documentation [here](https://maroto.io/).
* The `v1` still exists in [this branch](https://github.com/nh3000-org/maroto/tree/v1), and you can see the doc [here](https://maroto.io/#/v1/README?id=deprecated).

![result](docs/assets/images/result.png)

## Contributing

| Command         | Description                                       | Dependencies                                                  |
|-----------------|---------------------------------------------------|---------------------------------------------------------------|
| `make build`    | Build project                                     | `go`                                                          |
| `make test`     | Run unit tests                                    | `go`                                                          |
| `make fmt`      | Format files                                      | `gofmt`, `gofumpt` and `goimports`                            |
| `make lint`     | Check files                                       | `golangci-lint`                                               |
| `make dod`      | (Definition of Done) Format files and check files | Same as `make build`, `make test`, `make fmt` and `make lint` | 
| `make install`  | Install all dependencies                          | `go`, `curl` and `git`                                        |
| `make examples` | Run all examples                                  | `go`                                                          |
| `make mocks`    | Generate mocks                                    | `go` and `mockery`                                            |
| `make docs`     | Run docsify docs server local                     | `docsify`                                                     |
| `make godoc`    | Run godoc server local                            | `godoc`                                                       |

## Stargazers over time
[![Stargazers over time](https://starchart.cc/nh3000-org/maroto.svg?variant=adaptive)](https://starchart.cc/nh3000-org/maroto)
