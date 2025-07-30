<p align="center">
    <img src="/static/logo.png?v=0.1.2" width="300" />
    <h3 align="center">Moose</h3>
    <p align="center">MCP Server Boilerplate In Go</p>
    <p align="center">
        <a href="https://github.com/clivern/moose/actions/workflows/build.yml">
            <img src="https://github.com/clivern/moose/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/clivern/moose/actions">
            <img src="https://github.com/clivern/moose/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/clivern/moose/releases">
            <img src="https://img.shields.io/badge/Version-0.1.2-9B59B6.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/clivern/moose">
            <img src="https://goreportcard.com/badge/github.com/clivern/moose?v=0.1.2">
        </a>
        <a href="https://godoc.org/github.com/clivern/moose">
            <img src="https://godoc.org/github.com/clivern/moose?status.svg">
        </a>
        <a href="https://github.com/clivern/moose/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-E74C3C.svg">
        </a>
    </p>
</p>
<br/>

`Moose` is an MCP (Model Context Protocol) server boilerplate written in Go. It provides both `STDIO` and `SSE` (Server-Sent Events) modes of operation.


### Cursor Integration

Moose supports both `STDIO` and `SSE` (Server-Sent Events) modes for Cursor AI IDE integration.


#### STDIO Mode (Recommended)

To run `moose` in `STDIO` mode in Cursor AI IDE, use the following configuration:

```json
{
  "mcpServers": {
    "moose": {
      "command": "moose server"
    }
  }
}
```

```json
{
  "mcpServers": {
    "moose": {
      "command": "moose server --log-format json --log-level debug --log-output /var/log/moose.log"
    }
  }
}
```


#### SSE Mode

To run `moose` in `SSE` mode, first start the server:

```bash
moose server --sse --url http://127.0.0.1:8080
```

Then configure `Cursor` to connect to the `SSE` endpoint:

```json
{
  "mcpServers": {
    "moose": {
      "url": "http://127.0.0.1:8080/sse"
    }
  }
}
```

**Note:** When using `SSE` mode, make sure the server is running before starting `Cursor`, as `Cursor` will attempt to connect to the specified URL.


### Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Moose is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/moose/releases) for changelogs for each release version of Moose. It contains summaries of the most noteworthy changes made in each release.


### Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/moose/issues


### Security Issues

If you discover a security vulnerability within Moose, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


### Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


### License

© 2025, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Moose** is authored and maintained by [@clivern](http://github.com/clivern).
