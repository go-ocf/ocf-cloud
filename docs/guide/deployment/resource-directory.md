# Resource Directory

Resource Directory keeps and updates in-memory cache of devices which were requested by clients through the plgd gateways.

## Docker Image

```bash
docker pull plgd/resource-directory:latest
```

## Docker Run

### How to make certificates

Before you run docker image of plgd/resource-directory, you make sure certificates exists on `.tmp/certs` folder.
If not exists, you can create certificates from plgd/bundle image by following step only once.

```bash
# Create local folder for certificates and run plgd/bundle image to execute shell.
mkdir -p $(pwd).tmp/certs
docker run -it \
  --network=host \
  -v $(pwd)/.tmp/certs:/certs \
  -e CLOUD_SID=00000000-0000-0000-0000-000000000001 \
  --entrypoint /bin/bash \
  plgd/bundle:latest

# Copy & paste below commands on the bash shell of plgd/bundle container.
certificate-generator --cmd.generateRootCA --outCert=/certs/root_ca.crt --outKey=/certs/root_ca.key --cert.subject.cn=RootCA
certificate-generator --cmd.generateCertificate --outCert=/certs/http.crt --outKey=/certs/http.key --cert.subject.cn=localhost --cert.san.domain=localhost --signerCert=/certs/root_ca.crt --signerKey=/certs/root_ca.key

# Exit shell.
exit
```

```bash
# See common certificates for plgd cloud services.
ls .tmp/certs
http.crt  http.key  root_ca.crt  root_ca.key
```

### How to get configuration file

A configuration template is available on [resource-directory/config.yaml](https://github.com/plgd-dev/cloud/blob/v2/resource-directory/config.yaml).
You can also see `config.yaml` configuration file on the `resource-directory` folder by downloading `git clone https://github.com/plgd-dev/cloud.git`.

```bash
# Copy & paste configuration template from the link and save the file named `resource-directory.yaml` on the local folder.
vi resource-directory.yaml

# Or download configuration template.
curl https://github.com/plgd-dev/cloud/blob/v2/resource-directory/config.yaml --output resource-directory.yaml
```

### Edit configuration file

You can edit configuration file such as server port, certificates, OAuth provider and so on.
Read more detail about how to configure OAuth Provider [here](https://github.com/plgd-dev/cloud/blob/v2/docs/guide/developing/authorization.md#how-to-configure-auth0).

See an example of address, tls, event bus/store and OAuth config on the followings.

```yaml
...
apis:
  grpc:
    address: "0.0.0.0:9082"
    tls:
      caPool: "/data/certs/root_ca.crt"
      keyFile: "/data/certs/http.key"
      certFile: "/data/certs/http.crt"
    authorization:
      authority: "https://auth.example.com/authorize"
      audience: "https://api.example.com"
      http:
        tls:
          caPool: "/data/certs/root_ca.crt"
          keyFile: "/data/certs/http.key"
          certFile: "/data/certs/http.crt"
...
clients:
  eventBus:
    nats:
      url: "nats://localhost:4222"
      tls:
        caPool: "/data/certs/root_ca.crt"
        keyFile: "/data/certs/http.key"
        certFile: "/data/certs/http.crt"
...
  eventStore:
    mongoDB:
      uri: "mongodb://localhost:27017"
      database: "eventStore"
      tls:
        caPool: "/data/certs/root_ca.crt"
        keyFile: "/data/certs/http.key"
        certFile: "/data/certs/http.crt"
...
  authorizationServer:
    grpc:
      address: "localhost:9081"
      tls:
        caPool: "/data/certs/root_ca.crt"
        keyFile: "/data/certs/http.key"
        certFile: "/data/certs/http.crt"
    oauth:
      clientID: "412dsFf53Sj6$"
      clientSecret: "235Jgdf65jsd4Shls"
      scopes: "openid"
      tokenURL: "https://auth.example.com/oauth/token"
      audience: "https://api.example.com"
...
```

### Run docker image

You can run plgd/resource-directory image using certificates and configuration file on the folder you made certificates.

```bash
docker run -d --network=host \
  --name=resource-directory \
  -v $(pwd)/.tmp/certs:/data/certs \
  -v $(pwd)/resource-directory.yaml:/data/resource-directory.yaml \
  plgd/resource-directory:latest --config=/data/resource-directory.yaml
```

## YAML Configuration

### Logging

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `log.debug` | bool | `Set to true if you would like to see extra information on logs.` | `false` |

### gRPC API

gRPC API of the Resource Directory service.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `api.grpc.address` | string | `Listen specification <host>:<port> for grpc client connection.` | `"0.0.0.0:9100"` |
| `api.grpc.ownerCacheExpiration` | string | `Time limit of how long to keep subscribed to device updates after last use of the given cache item.` | `1m` |
| `api.grpc.enforcementPolicy.minTime` | string | `The minimum amount of time a client should wait before sending a keepalive ping. Otherwise the server close connection.` | `5s`|
| `api.grpc.enforcementPolicy.permitWithoutStream` | bool |  `If true, server allows keepalive pings even when there are no active streams(RPCs). Otherwise the server close connection.`  | `true` |
| `api.grpc.keepAlive.maxConnectionIdle` | string | `A duration for the amount of time after which an idle connection would be closed by sending a GoAway. 0s means infinity.` | `0s` |
| `api.grpc.keepAlive.maxConnectionAge` | string | `A duration for the maximum amount of time a connection may exist before it will be closed by sending a GoAway. 0s means infinity.` | `0s` |
| `api.grpc.keepAlive.maxConnectionAgeGrace` | string | `An additive period after MaxConnectionAge after which the connection will be forcibly closed. 0s means infinity.` | `0s` |
| `api.grpc.keepAlive.time` | string | `After a duration of this time if the server doesn't see any activity it pings the client to see if the transport is still alive.` | `2h` |
| `api.grpc.keepAlive.timeout` | string | `After having pinged for keepalive check, the client waits for a duration of Timeout and if no activity is seen even after that the connection is closed.` | `20s` |
| `api.grpc.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `api.grpc.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `api.grpc.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `api.grpc.tls.clientCertificateRequired` | bool | `If true, require client certificate.` | `true` |
| `api.grpc.authorization.authority` | string | `Authority is the address of the token-issuing authentication server. Services will use this URI to find and retrieve the public key that can be used to validate the token’s signature.` | `""` |
| `api.grpc.authorization.audience` | string | `Identifier of the API configured in your OAuth provider.` | `""` |
| `api.grpc.authorization.http.maxIdleConns` | int | `It controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.` | `16` |
| `api.grpc.authorization.http.maxConnsPerHost` | int | `It optionally limits the total number of connections per host, including connections in the dialing, active, and idle states. On limit violation, dials will block. Zero means no limit.` | `32` |
| `api.grpc.authorization.http.maxIdleConnsPerHost` | int | `If non-zero, controls the maximum idle (keep-alive) connections to keep per-host. If zero, DefaultMaxIdleConnsPerHost is used.` | `16` |
| `api.grpc.authorization.http.idleConnTimeout` | string | `The maximum amount of time an idle (keep-alive) connection will remain idle before closing itself. Zero means no limit.` | `30s` |
| `api.grpc.authorization.http.timeout` | string | `A time limit for requests made by this Client. A Timeout of zero means no timeout.` | `10s` |
| `api.grpc.authorization.http.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `api.grpc.authorization.http.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `api.grpc.authorization.http.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `api.grpc.authorization.http.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

### Event Bus

Plgd cloud uses NATS messaging system as a event bus.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `clients.eventBus.goPoolSize` | int | `Number of routines to process events in projection.` | `16` |
| `clients.eventBus.nats.url` | string | `URL to nats messaging system.` | `"nats://localhost:4222"` |
| `clients.eventBus.nats.pendingLimits.msgLimit` | int | `Limit number of messages in queue. -1 means unlimited` | `524288` |
| `clients.eventBus.nats.pendingLimits.bytesLimit` | int | `Limit buffer size of queue. -1 means unlimited` | `67108864` |
| `clients.eventBus.nats.tls.caPool` | string | `root certificate the root certificate in PEM format.` |  `""` |
| `clients.eventBus.nats.tls.keyFile` | string | `File name of private key in PEM format.` | `""` |
| `clients.eventBus.nats.tls.certFile` | string | `File name of certificate in PEM format.` | `""` |
| `clients.eventBus.nats.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

### Event Store

Plgd cloud uses MongoDB database as a event store.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `clients.eventStore.cacheExpiration` | string | `Expiration time of cached resource in projection.` | `20m` |
| `clients.eventStore.mongoDB.uri` | string | `URI to mongo database.` | `"mongodb://localhost:27017"` |
| `clients.eventStore.mongoDB.database` | string | `Name of database` | `"eventStore"` |
| `clients.eventStore.mongoDB.batchSize` | int | `Limits number of queries in one find request.` | `16` |
| `clients.eventStore.mongoDB.maxPoolSize` | int | `Limits number of connections.` | `16` |
| `clients.eventStore.mongoDB.maxConnIdleTime` | string | `Close connection when idle time reach the value.` | `240s` |
| `clients.eventStore.mongoDB.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `clients.eventStore.mongoDB.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `clients.eventStore.mongoDB.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `clients.eventStore.mongoDB.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

### Authorization Server Client

Client configurations to internally connect to Authorization Server service.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `clients.authorizationServer.pullFrequency` | string | `Frequency to pull changed user device.` | `15s` |
| `clients.authorizationServer.cacheExpiration` | string | `Expiration time of cached user device.` | `1m` |
| `clients.authorizationServer.ownerClaim` | string | `Owner claim of OAuth provider.` | `"sub"` |
| `clients.authorizationServer.grpc.address` | string | `Authorization service address.` | `"127.0.0.1:9100"` |
| `clients.authorizationServer.grpc.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `clients.authorizationServer.grpc.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `clients.authorizationServer.grpc.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `clients.authorizationServer.grpc.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |
| `clients.authorizationServer.grpc.keepAlive.time` | string | `After a duration of this time if the client doesn't see any activity it pings the server to see if the transport is still alive.` | `10s` |
| `clients.authorizationServer.grpc.keepAlive.timeout` | string | `After having pinged for keepalive check, the client waits for a duration of Timeout and if no activity is seen even after that the connection is closed.` | `20s` |
| `clients.authorizationServer.grpc.keepAlive.permitWithoutStream` | bool | `If true, client sends keepalive pings even with no active RPCs. If false, when there are no active RPCs, Time and Timeout will be ignored and no keepalive pings will be sent.` | `false` |

### OAuth2.0 Service Client

>Configured OAuth2.0 client is used by internal service to request a token used to authorize all calls they execute against other plgd APIs.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `clients.authorizationServer.oauth.clientID` | string | `Client ID to exchange an authorization code for an access token.` | `""` |
| `clients.authorizationServer.oauth.clientSecret` | string | `Client secret to exchange an authorization code for an access token.` |  `""` |
| `clients.authorizationServer.oauth.scopes` | string array | `List of required scopes.` | `""` |
| `clients.authorizationServer.oauth.tokenURL` | string | `Token endpoint of OAuth provider.` | `""` |
| `clients.authorizationServer.oauth.audience` | string | `Identifier of the API configured in your OAuth provider.` | `""` |
| `clients.authorizationServer.oauth.verifyServiceTokenFrequency` | string | `Frequency to verify service token.` | `10s` |
| `clients.authorizationServer.oauth.http.maxIdleConns` | int | `It controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.` | `16` |
| `clients.authorizationServer.oauth.http.maxConnsPerHost` | int | `It optionally limits the total number of connections per host, including connections in the dialing, active, and idle states. On limit violation, dials will block. Zero means no limit.` | `32` |
| `clients.authorizationServer.oauth.http.maxIdleConnsPerHost` | int | `If non-zero, controls the maximum idle (keep-alive) connections to keep per-host. If zero, DefaultMaxIdleConnsPerHost is used.` | `16` |
| `clients.authorizationServer.oauth.http.idleConnTimeout` | string | `The maximum amount of time an idle (keep-alive) connection will remain idle before closing itself. Zero means no limit.` | `30s` |
| `clients.authorizationServer.oauth.http.timeout` | string | `A time limit for requests made by this Client. A Timeout of zero means no timeout.` | `10s` |
| `clients.authorizationServer.oauth.http.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `clients.authorizationServer.oauth.http.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `clients.authorizationServer.oauth.http.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `clients.authorizationServer.oauth.http.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

::: tip Audience
You might have one client, but multiple APIs in the OAuth system. What you want to prevent is to be able to contact all the APIs of your system with one token. This audience allows you to request the token for a specific API. If you configure it to myplgdc2c.api in the Auth0, you have to set it here if you want to also validate it.
:::

### Public Configuration

These configurations are `Coap Cloud Conf` information for device registration to plgd cloud as well as root CA certificate, certificate authority address to get identity certificate for ssl connection to plgd cloud before device registration.
This will be served by HTTP Gateway API as defined [here](https://github.com/plgd-dev/cloud/blob/v2/http-gateway/uri/uri.go#L14) and also see [cloud-configuration](https://try.plgd.cloud/.well-known/cloud-configuration).

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `publicConfiguration.caPool` | string | `File path to root CA which was used to sign coap-gw certificate.` | `""` |
| `publicConfiguration.ownerClaim` | string | `Claim used to identify owner of the device.` | `"sub"` |
| `publicConfiguration.deviceIdClaim` | string | `Claim used to identify device id of the device. Empty means that JWT doesn't contain it.` | `""` |
| `publicConfiguration.signingServerAddress` | string | `Address of ceritificate authority for plgd-dev/sdk.` | `""` |
| `publicConfiguration.cloudID` | string | `Cloud ID which is stored in coap-gw certificate.` | `""` |
| `publicConfiguration.cloudURL` | string | `Cloud URL for onboard device.` | `""` |
| `publicConfiguration.cloudAuthorizationProvider` | string | `OAuth authorization provider for onboard device.` | `""` |
| `publicConfiguration.defaultCommandTimeToLive` | string | `Exposes default command time to live for CreateResource, RetrieveResource, UpdateResource, DeleteResource, and UpdateDeviceMetadata commands when it is not set in the request. 0s - means forever.` | `"0s"` |
| `publicConfiguration.authorizationServer` | string | `Authority is the address of the token-issuing authentication server. Services will use this URI to find and retrieve the public key that can be used to validate the token’s signature. Must be same as in coap-gateway configuration.` | `""` |

> Note that the string type related to time (i.e. timeout, idleConnTimeout, expirationTime) is decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units are "ns", "us", "ms", "s", "m", "h".
