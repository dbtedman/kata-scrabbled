# [Kata](https://github.com/dbtedman/kata) // [Scrabbled](https://github.com/dbtedman/kata-scrabbled)

> ⚠️ WARNING: Not production ready code.

[![CI GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/kata-scrabbled/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-scrabbled/actions/workflows/ci.yml)

A solving program that suggests words using your current letters and the layout of the board.

![Screenshot](./Screenshot.png)

-   [Getting Started](#getting-started)
-   [Design](#design)
-   [Attribution](#attribution)
-   [License](#license)

## Getting Started

```shell
nvm use && make && ./scrabbled
```

Visit [http://localhost:8080](http://localhost:8080)

## Design

### Domain Entities

| Entity                                                        | Notes |
| :------------------------------------------------------------ | :---- |
| [Boards](./internal/domain/board/boards.go)                   |       |
| [BoardSquare](./internal/domain/board_square/board_square.go) |       |
| [Tiles](./internal/domain/tile/tiles.go)                      |       |

### Domain Use Cases

| Use Case                                                    | Notes |
| :---------------------------------------------------------- | :---- |
| [RenderSquares](./internal/domain/square/render_squares.go) |       |

### Gateways

| Gateway | Notes                                            |
| :------ | :----------------------------------------------- |
| Words   | Words list used to populate internal dictionary. |

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

| Security Risk                                                                                                                       | Mitigation |
| :---------------------------------------------------------------------------------------------------------------------------------- | :--------- |
| [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)                                           |            |
| [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)                                         |            |
| [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)                                                                   |            |
| [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)                                                       |            |
| [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)                                   |            |
| [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)                 |            |
| [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/) |            |
| [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)             |            |
| [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)     |            |
| [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)                    |            |

## Attribution

### Words List

`./support/words.txt` accessed from [github.com/dwyl/english-words](https://github.com/dwyl/english-words), accessed
under [public domain](https://github.com/dwyl/english-words/blob/master/LICENSE.md).

## License

The [MIT License](./LICENSE.md) is used by this project.
