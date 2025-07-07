Namefinder
=================

- Namefinder is an Authoritative DNS server backed by MongoDB as a data store for DNS records.
- It provides programmatic access via a REST API.

### Setup

-------------

```shell
git clone https://github.com/lucidstacklabs/namefinder
cd namefinder
```

```shell
go install ./...
```

```shell
$GOPATH/bin/namefinder
```

#### Environment Variables

| Environment Variable | Default Value                 | Description                   |
|----------------------|-------------------------------|-------------------------------|
| `DNS_HOST`           | `0.0.0.0`                     | DNS server bind address       |
| `DNS_PORT`           | `5300`                        | DNS server port               |
| `ADMIN_HOST`         | `0.0.0.0`                     | Admin API bind address        |
| `ADMIN_PORT`         | `5301`                        | Admin API port                |
| `MONGO_ENDPOINT`     | `mongodb://localhost:27017`   | MongoDB connection string     |
| `MONGO_DB`           | `namefinder`                  | MongoDB database name         |
| `JWT_SIGNING_KEY`    | `secret`                      | JWT signing key               |
| `JWT_ISSUER`         | `namefinder`                  | JWT token issuer              |
| `JWT_AUDIENCE`       | `namefinder`                  | JWT token audience            |

