# Certificate Authority

This service is used to sign identity certificates for devices and [plgd-dev/sdk](https://github.com/plgd-dev/sdk) client.

> Initialization of the SDK requires an Identity Certificate issued by this service after successful user authorization. The SDK is then able to discover and configure the ownership as described [here](https://openconnectivity.org/specs/OCF_Security_Specification_v2.2.3.pdf#page=37). During this process, a CSR (Certificate Signing Request) is sent to the Certificate Authority which returns device's Identity Certificate.
>
> e.g. plgd mobile application uses SDK for device ownership configuration. Available on both [Google Play](https://play.google.com/store/apps/details?id=dev.plgd.client&hl=sk&gl=US), [Apple Store](https://apps.apple.com/sk/app/plgd/id1536315811).

## Docker Image

```bash
docker pull plgd/certificate-authority:latest
```

## Docker Run

### How to make certificates

Before you run docker image of plgd/certificate-authority, you make sure certificates exists on `.tmp/certs` folder.
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

A configuration template is available on [certificate-authority/config.yaml](https://github.com/plgd-dev/cloud/blob/v2/certificate-authority/config.yaml).
You can also see `config.yaml` configuration file on the `certificate-authority` folder by downloading `git clone https://github.com/plgd-dev/cloud.git`.

```bash
# Copy & paste configuration template from the link and save the file named `certificate-authority.yaml` on the local folder.
vi certificate-authority.yaml

# Or download configuration template.
curl https://github.com/plgd-dev/cloud/blob/v2/certificate-authority/config.yaml --output certificate-authority.yaml
```

### Edit configuration file

You can edit configuration file including server port, certificates, OAuth provider and so on.
Read more detail about how to configure OAuth Provider [here](https://github.com/plgd-dev/cloud/blob/v2/docs/guide/developing/authorization.md#how-to-configure-auth0).

See an example of address, tls on the followings.

```yaml
...
apis:
  grpc:
    address: "0.0.0.0:9087"
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
signer:
  keyFile: "/data/certs/root_ca.key"
  certFile: "/data/certs/root_ca.crt"
...
```

### Run docker image

You can run plgd/certificate-authority image using certificates and configuration file on the folder you made certificates.

```bash
docker run -d --network=host \
  --name=certificate-authority \
  -v $(pwd)/.tmp/certs:/data/certs \
  -v $(pwd)/certificate-authority.yaml:/data/certificate-authority.yaml \
  plgd/certificate-authority:latest --config=/data/certificate-authority.yaml
```

## YAML Configuration

### Logging

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `log.debug` | bool | `Set to true if you would like to see extra information on logs.` | `false` |

### gRPC API

gRPC API of the Certificate Authority service as defined [here](https://github.com/plgd-dev/cloud/blob/v2/certificate-authority/pb/service_grpc.pb.go#L19).

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
| `api.grpc.authorization.ownerClaim` | string | `Claim used to identify owner of the device.` | `"sub"` |
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

### Signer

Signer configuration to issue identity certificates for devices or mobile application.

| Property | Type | Description | Default |
| ---------- | -------- | -------------- | ------- |
| `signer.keyFile` | string | `File path to the signer private key in PEM format.` |  `""` |
| `signer.certFile` | string | `File path to the signer certificate in PEM format.` |  `""` |
| `signer.validFrom` | string | `The time from when the certificate is valid. (Format: https://github.com/karrick/tparse)` |  `"now-1h"` |
| `signer.expiresIn` | string | `The time up to which the certificate is valid.` |  `"87600h"` |

> Note that the string type related to time (i.e. timeout, idleConnTimeout, expirationTime) is decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units are "ns", "us", "ms", "s", "m", "h".
