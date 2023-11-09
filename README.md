# cicote

[![Quality check status](https://github.com/macie/cicote/actions/workflows/check.yml/badge.svg)](https://github.com/macie/cicote/actions/workflows/check.yml)

**WORK IN PROGRESS**

Calculate circadian color temperature for given geographic coordinates.

## Usage

```sh
$ cicote 52.50 13.24
3500
$ cicote -t '2023-05-23T16:43Z' 52.50 13.24
3500
```

## Development

Use `make` (GNU or BSD):

- `make` - install dependencies
- `make test` - runs test
- `make check` - static code analysis
- `make clean` - removes compilation artifacts
- `make info` - print system info (useful for debugging).

## License

[MIT](./LICENSE)
