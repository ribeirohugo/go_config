# Go Config

**Go Config** is a simple package that allows to load basic data for a Golang application.
It requires a `toml` file to instantiate env variables.

## 1. Config

Before to run application you may setup ``config.toml`` with ``Config`` values.
Check out the following tables to know all ``Config`` parameters detailed.

| Parameter       | Description                           | Type       | Default | Required |
|:----------------|:--------------------------------------|:-----------|:--------|:---------|
| ``environment`` | Website environment.                  | `string`   | ` `     | **YES**  |
| ``service``     | Website service identifier as string. | `string`   | ` `     | **YES**  |
| ``[server]``    | Http server config data.              | `Server`   | ` `     | **YES**  |
| ``[token]``     | Token data config data.               | `Token`    | ` `     | **YES**  |
| ``[mongodb]``   | Postgres database config data.        | `Database` | ` `     | **YES**  |
| ``[mysql]``     | MySql database config data.           | `Database` | ` `     | **YES**  |
| ``[postgres]``  | Postgres database config data.        | `Database` | ` `     | **YES**  |
| ``[tracer]``    | Tracing options config data.          | `Tracer`   | ` `     | **YES**  |

### 1.1. Server type

| Parameter           | Description                      | Type       | Default | Required |
|:--------------------|:---------------------------------|:-----------|:--------|:---------|
| ``host``            | HTTP server host name.           | `string`   | ` `     | **YES**  |
| ``port``            | HTTP server port number.         | `int`      | ` `     | **YES**  |
| ``allowed_origins`` | Allowed hosts to CORS allowance. | `[]string` | ` `     | **NO**   |

### 1.2. Database type

To set up ``[mysql]`` and ``[postgres]`` use the following parameters:

| Parameter           | Description                                         | Type     | Default                                      | Required |
|:--------------------|:----------------------------------------------------|:---------|:---------------------------------------------|:---------|
| ``db``              | Database name.                                      | `string` | ` `                                          | **YES**  |
| ``host``            | Database host.                                      | `string` | ` `                                          | **YES**  |
| ``migrations_path`` | Migrations directory path.                          | `string` | `file://migrations/<mongo><mysql><postgres>` | **NO**   |
| ``password``        | Database password.                                  | `string` | ` `                                          | **YES**  |
| ``port``            | Database port.                                      | `int`    | ` `                                          | **YES**  |
| ``user``            | Database user with needed privileges over database. | `string` | ` `                                          | **YES**  |

### 1.3. Token type

| Parameter      | Description                               | Type     | Default  | Required |
|:---------------|:------------------------------------------|:---------|:---------|:---------|
| ``secret``     | Website token secret string.              | `string` | ` `      | **YES**  |
| ``max_age``    | Maximum duration of a token (in seconds). | `int`    | `86400`  | **NO**   |

### 1.4. Tracer type

| Parameter      | Description                      | Type     | Default                             | Required |
|:---------------|:---------------------------------|:---------|:------------------------------------|:---------|
| ``Enable``     | Enable flag to activate tracing. | `bool`   | `FALSE`                             | **NO**   |
| ``JaegerHost`` | Jaeger host address.             | `string` | `http://localhost:14268/api/traces` | **NO**   |
