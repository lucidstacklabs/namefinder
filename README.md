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


### QuickStart

---------------

#### Set up an admin account

Set up the first admin account after bringing up the server

```shell
curl localhost:5301/api/v1/admins/init -XPOST -d '{"username": "admin", "password": "123"}'
```

Fetch the token for the admin account

```shell
curl localhost:5301/api/v1/admins/token -XPOST -d '{"username": "admin", "password": "123"}'
```

Set the admin token in a variable for further use

```shell
ADMIN_TOKEN=eyJhbGciOiJ...
```

#### Set up a namespace

Create a namespace
```shell
curl localhost:5301/api/v1/namespaces -H "Authorization: Bearer $ADMIN_TOKEN"  -XPOST -d '{"name": "test"}'
```

List all namespaces

```shell
curl localhost:5301/api/v1/namespaces -H "Authorization: Bearer $ADMIN_TOKEN"
```

Set namespace id for further use

```shell
NAMESPACE_ID=namespaceid 
```
#### Set up DNS records

Set A records

```shell
curl localhost:5301/admin/api/v1/namespaces/$NAMESPACE_ID/records -H "Authorization: Bearer $ADMIN_TOKEN" -XPOST -d '{"name": "example.com", "type": "A", "value": "192.168.0.105", "ttl": 60, "class": "IN"}'
curl localhost:5301/admin/api/v1/namespaces/$NAMESPACE_ID/records -H "Authorization: Bearer $ADMIN_TOKEN" -XPOST -d '{"name": "example.com", "type": "A", "value": "192.168.0.106", "ttl": 60, "class": "IN"}'
```

```shell
curl localhost:5301/admin/api/v1/namespaces/$NAMESPACE_ID/records -H "Authorization: Bearer $ADMIN_TOKEN" -XPOST -d '{"name": "www.example.com", "type": "CNAME", "value": "example.github.io", "ttl": 60, "class": "IN"}'
```

Set MX records

```shell
curl localhost:5301/admin/api/v1/namespaces/$NAMESPACE_ID/records -H "Authorization: Bearer $ADMIN_TOKEN" -XPOST -d '{"name": "example.com", "type": "MX", "value": "10 mail1.example.com", "ttl": 60, "class": "IN"}'
curl localhost:5301/admin/api/v1/namespaces/$NAMESPACE_ID/records -H "Authorization: Bearer $ADMIN_TOKEN" -XPOST -d '{"name": "example.com", "type": "MX", "value": "20 mail2.example.com", "ttl": 60, "class": "IN"}'
```

#### Testing it out 

```shell
dig @localhost -p 5300 example.com

; <<>> DiG 9.20.10 <<>> @localhost -p 5300 example.com
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 55909
;; flags: qr aa rd; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;example.com.			IN	A

;; ANSWER SECTION:
example.com.		60	IN	A	192.168.0.105
example.com.		60	IN	A	192.168.0.106

;; Query time: 1 msec
;; SERVER: ::1#5300(localhost) (UDP)
;; WHEN: Mon Jul 07 23:35:13 IST 2025
;; MSG SIZE  rcvd: 83
```

```shell
dig @localhost -p 5300 www.example.com CNAME

; <<>> DiG 9.20.10 <<>> @localhost -p 5300 www.example.com CNAME
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 6702
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;www.example.com.		IN	CNAME

;; ANSWER SECTION:
www.example.com.	60	IN	CNAME	example.github.io.

;; Query time: 1 msec
;; SERVER: ::1#5300(localhost) (UDP)
;; WHEN: Mon Jul 07 23:34:36 IST 2025
;; MSG SIZE  rcvd: 79
```

```shell
dig @localhost -p 5300 example.com MX

; <<>> DiG 9.20.10 <<>> @localhost -p 5300 example.com MX
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 63220
;; flags: qr aa rd; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;example.com.			IN	MX

;; ANSWER SECTION:
example.com.		60	IN	MX	10 mail1.example.com.
example.com.		60	IN	MX	20 mail2.example.com.

;; Query time: 1 msec
;; SERVER: ::1#5300(localhost) (UDP)
;; WHEN: Mon Jul 07 23:27:41 IST 2025
;; MSG SIZE  rcvd: 115
```

### API Documentation

-------------------

Full API documentation is available [here](docs/api.md)