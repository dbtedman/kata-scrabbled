# [Scrabbled](https://github.com/dbtedman/kata-scrabbled)

> **⚠️ WARNING:** Not production ready code, instead a [Code Kata](https://github.com/dbtedman#code-kata) intended to
> hone my programming skills through practice and repetition.

[![CI GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/kata-scrabbled/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-scrabbled/actions/workflows/ci.yml)
[![sast workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-scrabbled/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/kata-scrabbled/actions/workflows/sast.yml)
![language: go](https://img.shields.io/badge/language-go-blue.svg?style=for-the-badge&logo=go)
[![MIT License](https://img.shields.io/github/license/dbtedman/kata-scrabbled?color=orange&style=for-the-badge)](https://github.com/dbtedman/kata-scrabbled/blob/main/LICENSE.md)

A solving program that suggests words using your current letters and the layout of the board.

-   [Getting Started](#getting-started)
-   [Verification](#verification)
-   [Design](#design)
-   [Attribution](#attribution)
-   [References](#references)
-   [License](#license)

## Getting Started

```shell
nvm use && make && ./scrabbled
```

Visit [http://localhost:8080](http://localhost:8080)

## Verification

### Linting

-   [Prettier](https://prettier.io)

```shell
make lint
```

These rules can then be automatically applied:

```shell
make format
```

### Unit Testing

```shell
make test
```

## Design

### Domain Entities

#### [Boards](./internal/domain/board/boards.go)

_Placeholder_

#### [BoardSquare](./internal/domain/board_square/board_square.go)

_Placeholder_

#### [Tiles](./internal/domain/tile/tiles.go)

_Placeholder_

### Domain Use Cases

#### [Render Squares](./internal/domain/square/render_squares.go)

_Placeholder_

### Gateways

#### Words

Words list used to populate internal dictionary.

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

#### [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)

[Github Security](https://github.com/features/security) detects secrets incorrectly committed into the repository.

#### [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)

_Placeholder_

#### [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)

_Placeholder_

#### [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)

_Placeholder_

#### [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)

_Placeholder_

#### [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)

[Snyk](https://snyk.io) and [Github Security](https://github.com/features/security) scan Gradle and NPM dependencies for know vulnerabilities and create pull requests to resolve the vulnerabilities when available.

#### [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/)

_Placeholder_

#### [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)

_Placeholder_

#### [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)

_Placeholder_

#### [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

_Placeholder_

## Attribution

### Words List

`./support/words.txt` accessed from [github.com/dwyl/english-words](https://github.com/dwyl/english-words), accessed
under [public domain](https://github.com/dwyl/english-words/blob/master/LICENSE.md).

## References

_Placeholder_

## License

The [MIT License](./LICENSE.md) is used by this project.
