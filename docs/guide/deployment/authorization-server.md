# Authorization Server
Authorization Server authorizes users and devices interacting with the plgd cloud.

## Docker Image

```bash
docker pull plgd/authorization:latest
```

## Docker Run
### How to make certificates
Before you run docker image of plgd/authorization, you make sure certificates exists on `.tmp/certs` folder. 
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
http.crt	http.key	root_ca.crt	root_ca.key
```
### How to get configuration file
A configuration template is available on [authorization/config.yaml](https://github.com/plgd-dev/cloud/blob/v2/authorization/config.yaml). 
You can also see `config.yaml` configuration file on the `authorization` folder by downloading `git clone https://github.com/plgd-dev/cloud.git`. 
```bash
# Copy & paste configuration template from the link and save the file named `authorization.yaml` on the local folder.
vi authorization.yaml

# Or download configuration template.
curl https://github.com/plgd-dev/cloud/blob/v2/authorization/config.yaml --output authorization.yaml 
```

### Edit configuration file 
You can edit configuration file including server port, certificates, OAuth provider and so on.
Read more detail about how to configure OAuth Provider [here](https://github.com/plgd-dev/cloud/blob/v2/docs/guide/developing/authorization.md#how-to-configure-auth0). 

See an example of address, tls and OAuth config on the followings.
```yaml
...
apis:
  grpc:
    address: "0.0.0.0:9081"
    tls:
      caPool: "/data/certs/root_ca.crt"
      keyFile: "/data/certs/http.key"
      certFile: "/data/certs/http.crt"
...
  http:
    address: "0.0.0.0:9085"
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
oauthClients:
  device:
    provider: "plgd"
    clientID: "ij12OJj2J23K8KJs"
    clientSecret: "654hkja12asd123d"
    scopes: "profile,openid,offline_access"
    authorizationURL: "https://auth.example.com/authorize"
    tokenURL: "https://auth.example.com/oauth/token"
    audience: "https://api.example.com"
    redirectURL: "https://localhost:9085/api/authz/callback"
...
  client:
    clientID: "412dsFf53Sj6$"
    clientSecret: "235Jgdf65jsd4Shls"
    scopes: "openid"
    authorizationURL: "https://auth.example.com/authorize"
    audience: "https://api.example.com"
    redirectURL: "https://localhost:9085/api/authz/callback"
...
```

### Run docker image 
You can run plgd/authorization image using certificates and configuration file on the folder you made certificates.
```bash
docker run -d --network=host \
	--name=authorization \
	-v $(pwd)/.tmp/certs:/data/certs \
	-v $(pwd)/authorization.yaml:/data/authorization.yaml \
	plgd/authorization:latest --config=/data/authorization.yaml
```

## YAML Configuration
### Logging

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `log.debug` | bool | `Set to true if you would like to see extra information on logs.` | `false` |

### gRPC API
gRPC API of the Authorization Server service as defined [here](https://github.com/plgd-dev/cloud/blob/v2/authorization/pb/service_grpc.pb.go#L19).

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `api.grpc.address` | string | `Listen specification <host>:<port> for grpc client connection.` | `"0.0.0.0:9100"` |
| `api.grpc.enforcementPolicy.minTime` | string | `The minimum amount of time a client should wait before sending a keepalive ping. Otherwise the server close connection.` | `5s`|
| `api.grpc.enforcementPolicy.permitWithoutStream` | bool |  `If true, server allows keepalive pings even when there are no active streams(RPCs). Otherwise the server close connection.`  | `true` |
| `api.grpc.keepAlive.maxConnectionIdle` | string | `A duration for the amount of time after which an idle connection would be closed by sending a GoAway. 0s means infinity.` | `0s` |
| `api.grpc.keepAlive.maxConnectionAge` | string | `A duration for the maximum amount of time a connection may exist before it will be closed by sending a GoAway. 0s means infinity.` | `0s` |
| `api.grpc.keepAlive.maxConnectionAgeGrace` | string | `An additive period after MaxConnectionAge after which the connection will be forcibly closed. 0s means infinity.` | `0s` |
| `api.grpc.keepAlive.time` | string | `After a duration of this time if the server doesn't see any activity it pings the client to see if the transport is still alive.` | `2h` |
| `api.grpc.keepAlive.timeout` | string | `After having pinged for keepalive check, the client waits for a duration of Timeout and if no activity is seen even after that the connection is closed.` | `20s` |
| `api.grpc.tls.caPool` | string | `File path to the root certificate in PEM format.` |  `""` |
| `api.grpc.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `api.grpc.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `api.grpc.tls.clientCertificateRequired` | bool | `If true, require client certificate.` | `true` |
| `api.grpc.authorization.authority` | string | `Endpoint of OAuth provider.` | `""` |
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

### HTTP API
HTTP API of the Authorization Server service as defined [here](https://github.com/plgd-dev/cloud/blob/v2/authorization/uri/uri.go)

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `api.http.address` | string | `Listen specification <host>:<port> for http client connection.` | `"0.0.0.0:9100"` |
| `api.http.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `api.http.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `api.http.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `api.http.tls.clientCertificateRequired` | bool | `If true, require client certificate.` | `true` |

### OAuth2.0 Client for Device
>Configured OAuth2.0 client is used to request an authorization code used for onboarding and exchange it for the token during the [device registration](https://plgd.dev/guide/architecture/component-overview.html#coap-gateway).

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `oauthClients.device.provider` | string | `Value which comes from the device during the sign-up ("apn").` | `"generic"` |
| `oauthClients.device.clientID` | string | `Client ID to exchange an authorization code for an access token.` | `""` |
| `oauthClients.device.clientSecret` | string | `Client secret to exchange an authorization code for an access token.` |  `""` |
| `oauthClients.device.scopes` | string | `Comma separated list of required scopes.` | `""` |
| `oauthClients.device.authorizationURL` | string | `Authorization endpoint of OAuth provider.` | `""` |
| `oauthClients.device.tokenURL` | string | `Token endpoint of OAuth provider.` | `""` |
| `oauthClients.device.audience` | string | `Identifier of the API configured in your OAuth provider.` | `""` |
| `oauthClients.device.redirectURL` | string | `Redirect url used to obtain device access token.` | `""` |
| `oauthClients.device.responseType` | string | `One of "code/token".` | `"code"` |
| `oauthClients.device.responseMode` | string | `One of "query/post_form".` | `"post_form"` |
| `oauthClients.device.http.maxIdleConns` | int | `It controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.` | `16` |
| `oauthClients.device.http.maxConnsPerHost` | int | `It optionally limits the total number of connections per host, including connections in the dialing, active, and idle states. On limit violation, dials will block. Zero means no limit.` | `32` |
| `oauthClients.device.http.maxIdleConnsPerHost` | int | `If non-zero, controls the maximum idle (keep-alive) connections to keep per-host. If zero, DefaultMaxIdleConnsPerHost is used.` | `16` |
| `oauthClients.device.http.idleConnTimeout` | string | `The maximum amount of time an idle (keep-alive) connection will remain idle before closing itself. Zero means no limit.` | `30s` |
| `oauthClients.device.http.timeout` | string | `A time limit for requests made by this Client. A Timeout of zero means no timeout.` | `10s` |
| `oauthClients.device.http.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `oauthClients.device.http.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `oauthClients.device.http.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `oauthClients.device.http.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

::: tip Audience 
You might have one client, but multiple APIs in the OAuth system. What you want to prevent is to be able to contact all the APIs of your system with one token. This audience allows you to request the token for a specific API. If you configure it to myplgdc2c.api in the Auth0, you have to set it here if you want to also validate it.
:::

### OAuth2.0 Client for UI and SDK
>Configured OAuth2.0 client is used by the mobile application or SDK to request a token used to authorize all calls they execute against other plgd APIs.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `oauthClients.client.clientID` | string | `Client ID to exchange an authorization code for an access token.` | `""` |
| `oauthClients.client.clientSecret` | string | `Client secret to exchange an authorization code for an access token.` |  `""` |
| `oauthClients.client.scopes` | string | `Comma separated list of required scopes.` | `""` |
| `oauthClients.client.authorizationURL` | string | `Authorization endpoint of OAuth provider.` | `""` |
| `oauthClients.client.audience` | string | `Identifier of the API configured in your OAuth provider.` | `""` |
| `oauthClients.client.redirectURL` | string | `Redirect url used to obtain device access token.` | `""` |
| `oauthClients.client.responseMode` | string | `One of "query/post_form".` | `"post_form"` |
| `oauthClients.client.http.maxIdleConns` | int | `It controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.` | `16` |
| `oauthClients.client.http.maxConnsPerHost` | int | `It optionally limits the total number of connections per host, including connections in the dialing, active, and idle states. On limit violation, dials will block. Zero means no limit.` | `32` |
| `oauthClients.client.http.maxIdleConnsPerHost` | int | `If non-zero, controls the maximum idle (keep-alive) connections to keep per-host. If zero, DefaultMaxIdleConnsPerHost is used.` | `16` |
| `oauthClients.client.http.idleConnTimeout` | string | `The maximum amount of time an idle (keep-alive) connection will remain idle before closing itself. Zero means no limit.` | `30s` |
| `oauthClients.client.http.timeout` | string | `A time limit for requests made by this Client. A Timeout of zero means no timeout.` | `10s` |
| `oauthClients.client.http.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `oauthClients.client.http.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `oauthClients.client.http.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `oauthClients.client.http.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

::: tip Audience 
You might have one client, but multiple APIs in the OAuth system. What you want to prevent is to be able to contact all the APIs of your system with one token. This audience allows you to request the token for a specific API. If you configure it to myplgdc2c.api in the Auth0, you have to set it here if you want to also validate it.
:::

### Event Bus
Plgd cloud uses NATS messaging system as a event bus.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `clients.eventBus.nats.url` | string | `URL to nats messaging system.` | `"nats://localhost:4222"` |
| `clients.eventBus.nats.jetstream `| bool | `If true, events will be published to jetstream.` | `false` |
| `clients.eventBus.nats.tls.caPool` | string | `root certificate the root certificate in PEM format.` |  `""` |
| `clients.eventBus.nats.tls.keyFile` | string | `File name of private key in PEM format.` | `""` |
| `clients.eventBus.nats.tls.certFile` | string | `File name of certificate in PEM format.` | `""` |
| `clients.eventBus.nats.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

### Storage
Plgd cloud uses MongoDB database as owner's device store.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `clients.storage.ownerClaim` | string | `Claim used to identify owner of the device.` | `"sub"` |
| `clients.storage.mongoDB.uri` | string | `URI to mongo database.` | `"mongodb://localhost:27017"` |
| `clients.storage.mongoDB.database` | string | `Name of database.` | `"ownersDevices"` |
| `clients.storage.mongoDB.tls.caPool` | string | `File path to the root certificate in PEM format which might contain multiple certificates in a single file.` |  `""` |
| `clients.storage.mongoDB.tls.keyFile` | string | `File path to private key in PEM format.` | `""` |
| `clients.storage.mongoDB.tls.certFile` | string | `File path to certificate in PEM format.` | `""` |
| `clients.storage.mongoDB.tls.useSystemCAPool` | bool | `If true, use system certification pool.` | `false` |

> Note that the string type related to time (i.e. timeout, idleConnTimeout, expirationTime) is decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units are "ns", "us", "ms", "s", "m", "h".

