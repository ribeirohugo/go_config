# Go Config

**Go Config** is a simple package that allows to load basic data for a Golang application.
It supports `toml`, `json`, `yaml` or `xml` file or content to instantiate environment variables.

## 1. Config

Before to run application you may setup a supported``config.toml`` with ``Config`` values.
Check out the following tables to know all ``Config`` parameters detailed.

| Parameter                     | Description                                                                           | Type                | Default | Required |
|:------------------------------|:--------------------------------------------------------------------------------------|:--------------------|:--------|:---------|
| ``environment``               | Website environment.                                                                  | `string`            | ` `     | **NO**   |
| ``service``                   | Website service identifier as string.                                                 | `string`            | ` `     | **NO**   |
| ``[server]``                  | Http server config data.                                                              | `Server`            | ` `     | **YES**  |
| ``[token]``                   | Token data config data.                                                               | `Token`             | ` `     | **NO**   |
| ``[mongodb]``                 | Postgres database config data.                                                        | `Database`          | ` `     | **NO**   |
| ``[mysql]``                   | MySql database config data.                                                           | `Database`          | ` `     | **NO**   |
| ``[postgres]``                | Postgres database config data.                                                        | `Database`          | ` `     | **NO**   |
| ``[audit]``                   | Auditing options config data.                                                         | `ExternalService`   | ` `     | **NO**   |
| ``[jaeger]``                  | Jaeger tracing options config data.                                                   | `ExternalService`   | ` `     | **NO**   |
| ``[loki]``                    | Grafana Loki options config data.                                                     | `ExternalService`   | ` `     | **NO**   |
| ``[tempo]``                   | Tempo tracer config data.                                                             | `ExternalService`   | ` `     | **NO**   |
| ``[prometheus]``              | Prometheus monitoring config data.                                                    | `ExternalService`   | ` `     | **NO**   |
| ``[redis]``                   | Redis cache config data.                                                              | `ExternalService`   | ` `     | **NO**   |
| ``[settings]`` <sup>(1)</sup> | Holds mapped key-value attributes.<br/> Allows an undefined number of new attributes. | `map[string]string` | ` `     | **NO**   |

> <sup>(1)</sup> Settings is not available to `xml` loader.

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
| ``port``            | Database port.                                      | `int`    | `3306`, `5432`, `27017` <sup>(2)</sup>       | **YES**  |
| ``user``            | Database user with needed privileges over database. | `string` | ` `                                          | **YES**  |

> <sup>(2)</sup> `3306` for MySQL, `5432` for Postgres and `27017` for MongoDB. 

### 1.3. Token type

| Parameter      | Description                               | Type     | Default  | Required |
|:---------------|:------------------------------------------|:---------|:---------|:---------|
| ``secret``     | Website token secret string.              | `string` | ` `      | **YES**  |
| ``max_age``    | Maximum duration of a token (in seconds). | `int`    | `86400`  | **NO**   |

### 1.4. External Service type

| Parameter  | Description                      | Type     | Default            | Required |
|:-----------|:---------------------------------|:---------|:-------------------|:---------|
| ``Enable`` | Enable flag to activate tracing. | `bool`   | `FALSE`            | **NO**   |
| ``Host``   | Service host address.            | `string` | ` ` <sup>(3)</sup> | **NO**   |
| ``Token``  | Service token string.            | `string` | ` `                | **NO**   |

> <sup>(3)</sup> Host default values are specified in `External Host Default Values` table.

### 1.4.1. External Host Default Values

For `ExternalService` main config, the `Host` default value depends on the main service,
and those values are described in the next table.

| Service        | Default Host Value                       |
|:---------------|:-----------------------------------------|
| ``Audit``      | ` `                                      |
| ``Jaeger``     | `http://localhost:14268/api/traces`      |
| ``Loki``       | `http://localhost:3100/loki/api/v1/push` |
| ``Prometheus`` | ` `                                      |
| ``Tempo``      | `http://localhost:4318/v1/traces`        |
| ``Redis``      | `localhost:6379`                         |

## 2. Load data

First you need to get dependency `go_config` dependency by calling `go get`, with the wanted release.

``
go get github.com/ribeirohugo/go_config@latest
``

### 2.1. Environment

Then, data can be loaded by calling `Load` method.
It loads config variables from environment.

```
os.SetVar("SERVER_HOST", "localhost")

cfg, err := toml.Load()
if err != nil {
    log.Fatal(err)
}
```

Variables and subvariables uses the following logic:

```
ENVIRONMENT = "TEST"

SERVER_HOST = "localhost"

MONGODB_USER = "username"
```

It will return a `config.Config` struct variable or an error, if anything unexpected occurs.

### 2.2. Toml

Then, data can be loaded by calling `Load` method.
It supports a `config.toml` file properly fulfilled.

```
cfg, err := env.Load(configFile)
if err != nil {
    log.Fatal(err)
}
```

It will return a `config.Config` struct variable or an error, if anything unexpected occurs.

### 2.3. YAML

Then, data can be loaded by calling `Load` method, with a `config.yaml` file properly fulfilled.

```
cfg, err := yaml.Load(configFile)
if err != nil {
    log.Fatal(err)
}
```

It will return a `config.Config` struct variable or an error, if anything unexpected occurs.

### 2.4. JSON

Then, data can be loaded by calling `Load` method, with a `config.json` file properly fulfilled.

```
cfg, err := json.Load(configFile)
if err != nil {
    log.Fatal(err)
}
```

It will return a `config.Config` struct variable or an error, if anything unexpected occurs.

### 2.5. XML

Then, data can be loaded by calling `Load` method, with a `config.json` file properly fulfilled.

```
cfg, err := xml.Load(configFile)
if err != nil {
    log.Fatal(err)
}
```

It will return a `config.Config` struct variable or an error, if anything unexpected occurs.
