#!/usr/bin/env bash
set -e

# Configure services
export PATH="/usr/local/bin:$PATH"

export CERTIFICATES_PATH="/data/certs"
export OAUTH_KEYS_PATH="/data/oauth/keys"
export LOGS_PATH="/data/log"
export MONGO_PATH="/data/db"
export NGINX_PATH="/data/nginx"
export JETSTREAM_PATH="/data/jetstream"

export SERVICE_OWNER_CLAIM=${OWNER_CLAIM}
export DEVICE_OWNER_CLAIM=${OWNER_CLAIM}

export CERTIFICATE_AUTHORITY_ADDRESS="localhost:${CERTIFICATE_AUTHORITY_PORT}"
export MOCKED_OAUTH_SERVER_ADDRESS="localhost:${MOCKED_OAUTH_SERVER_PORT}"
export RESOURCE_AGGREGATE_ADDRESS="localhost:${RESOURCE_AGGREGATE_PORT}"
export RESOURCE_DIRECTORY_ADDRESS="localhost:${RESOURCE_DIRECTORY_PORT}"
export AUTHORIZATION_ADDRESS="localhost:${AUTHORIZATION_PORT}"
export AUTHORIZATION_HTTP_ADDRESS="localhost:${AUTHORIZATION_HTTP_PORT}"
export GRPC_GATEWAY_ADDRESS="localhost:${GRPC_GATEWAY_PORT}"
export HTTP_GATEWAY_ADDRESS="localhost:${HTTP_GATEWAY_PORT}"

export INTERNAL_CERT_DIR_PATH="$CERTIFICATES_PATH/internal"
export GRPC_INTERNAL_CERT_NAME="endpoint.crt"
export GRPC_INTERNAL_CERT_KEY_NAME="endpoint.key"

export EXTERNAL_CERT_DIR_PATH="$CERTIFICATES_PATH/external"
export COAP_GATEWAY_FILE_CERT_NAME="coap-gateway.crt"
export COAP_GATEWAY_FILE_CERT_KEY_NAME="coap-gateway.key"

# ROOT CERTS
export CA_POOL_DIR="$CERTIFICATES_PATH"
export CA_POOL_NAME_PREFIX="root_ca"
export CA_POOL_CERT_PATH="$CA_POOL_DIR/$CA_POOL_NAME_PREFIX.crt"
export CA_POOL_CERT_KEY_PATH="$CA_POOL_DIR/$CA_POOL_NAME_PREFIX.key"

# DIAL CERTS
export DIAL_TYPE="file"
export DIAL_FILE_CA_POOL="$CA_POOL_CERT_PATH"
export DIAL_FILE_CERT_DIR_PATH="$INTERNAL_CERT_DIR_PATH"
export DIAL_FILE_CERT_NAME="$GRPC_INTERNAL_CERT_NAME"
export DIAL_FILE_CERT_KEY_NAME="$GRPC_INTERNAL_CERT_KEY_NAME"
export DIAL_FILE_USE_SYSTEM_CERTIFICATION_POOL="true"

#LISTEN CERTS
export LISTEN_TYPE="file"
export LISTEN_FILE_CA_POOL="$CA_POOL_CERT_PATH"
export LISTEN_FILE_CERT_DIR_PATH="$INTERNAL_CERT_DIR_PATH"
export LISTEN_FILE_CERT_NAME="$GRPC_INTERNAL_CERT_NAME"
export LISTEN_FILE_CERT_KEY_NAME="$GRPC_INTERNAL_CERT_KEY_NAME"

#SECRETS
export SECRETS_DIRECTORY=/data/secrets

#OAUTH-SEVER KEYS
export OAUTH_ID_TOKEN_KEY_PATH=${OAUTH_KEYS_PATH}/id-token.pem
export OAUTH_ACCESS_TOKEN_KEY_PATH=${OAUTH_KEYS_PATH}/access-token.pem

#ENDPOINTS
export MONGODB_HOST="localhost:$MONGO_PORT"
export MONGODB_URL="mongodb://$MONGODB_HOST"
export MONGODB_URI="mongodb://$MONGODB_HOST"
export MONGO_URI="mongodb://$MONGODB_HOST"
export NATS_HOST="localhost:$NATS_PORT"
export NATS_URL="nats://${NATS_HOST}"
export SERVICE_NATS_URL=${NATS_URL}

export AUTH_SERVER_ADDRESS=${AUTHORIZATION_ADDRESS}
export FQDN_NGINX_HTTPS=${FQDN}:${NGINX_PORT}
export DOMAIN=${FQDN_NGINX_HTTPS}
if [ "$NGINX_PORT" = "443" ]; then
  export DOMAIN=${FQDN}
fi

#OAUTH SERVER
if [ -z "${OAUTH_AUDIENCE}" ]
then
  export OAUTH_AUDIENCE=test
  export SERVICE_OAUTH_AUDIENCE=${OAUTH_AUDIENCE}
  export SDK_OAUTH_AUDIENCE=${FQDN}:${NGINX_PORT}
else
  export OAUTH_AUDIENCE=${OAUTH_AUDIENCE}
  export SERVICE_OAUTH_AUDIENCE=${OAUTH_AUDIENCE}
  export SDK_OAUTH_AUDIENCE=${OAUTH_AUDIENCE}
fi

if [ -z "${OAUTH_ENDPOINT_AUTH_URL}" ]
then
  export DEVICE_OAUTH_ENDPOINT_AUTH_URL="https://localhost:${MOCKED_OAUTH_SERVER_PORT}/authorize"
  export OAUTH_ENDPOINT_CODE_URL="https://localhost:${MOCKED_OAUTH_SERVER_PORT}/authorize"
  export SERVICE_CLIENT_CONFIGURATION_AUTHCODEURL="https://${FQDN}:${NGINX_PORT}/authorize?client_id=test"
else
  export DEVICE_OAUTH_REDIRECT_URL=https://${FQDN}:${NGINX_PORT}/api/authz/callback
  export SDK_OAUTH_REDIRECT_URL=https://${FQDN}:${NGINX_PORT}/api/authz/callback
  export DEVICE_OAUTH_ENDPOINT_AUTH_URL=${OAUTH_ENDPOINT_AUTH_URL}
  export OAUTH_ENDPOINT_CODE_URL=${OAUTH_ENDPOINT_AUTH_URL}
  export SERVICE_CLIENT_CONFIGURATION_AUTHCODEURL=https://${FQDN}:${NGINX_PORT}/api/authz/code
fi

if [ -z "${OAUTH_ENDPOINT_TOKEN_URL}" ]
then
  export DEVICE_OAUTH_ENDPOINT_TOKEN_URL=https://localhost:${MOCKED_OAUTH_SERVER_PORT}/oauth/token
  export SERVICE_CLIENT_CONFIGURATION_ACCESSTOKENURL="https://${FQDN}:${NGINX_PORT}/oauth/token?client_id=test&audience=test"
  export SERVICE_OAUTH_ENDPOINT_TOKEN_URL=https://localhost:${MOCKED_OAUTH_SERVER_PORT}/oauth/token
  export SDK_OAUTH_ENDPOINT_AUTH_URL="https://${FQDN}:${NGINX_PORT}/oauth/token"
else
  export DEVICE_OAUTH_REDIRECT_URL=https://${FQDN}:${NGINX_PORT}/api/authz/callback
  export SDK_OAUTH_REDIRECT_URL=https://${FQDN}:${NGINX_PORT}/api/authz/callback
  export DEVICE_OAUTH_ENDPOINT_TOKEN_URL=${OAUTH_ENDPOINT_TOKEN_URL}
  export SERVICE_CLIENT_CONFIGURATION_ACCESSTOKENURL=https://${FQDN}:${NGINX_PORT}/api/authz/token
  export SERVICE_OAUTH_ENDPOINT_TOKEN_URL=${OAUTH_ENDPOINT_TOKEN_URL}
  export SDK_OAUTH_ENDPOINT_AUTH_URL=${OAUTH_ENDPOINT_AUTH_URL}
fi

if [ -z "${OAUTH_ENDPOINT}" ]
then
  export OAUTH_ENDPOINT=${DOMAIN}
fi

if [ -z "${JWKS_URL}" ]
then
  export JWKS_URL=https://localhost:${MOCKED_OAUTH_SERVER_PORT}/.well-known/jwks.json
fi

if [ -z "${OAUTH_CLIENT_ID}" ]
then
  export DEVICE_OAUTH_CLIENT_ID=test
  export SDK_OAUTH_CLIENT_ID=test
  export OAUTH_CLIENT_ID=test
else
  export DEVICE_OAUTH_CLIENT_ID=${OAUTH_CLIENT_ID}
  export SDK_OAUTH_CLIENT_ID=${OAUTH_CLIENT_ID}
  export OAUTH_CLIENT_ID=${OAUTH_CLIENT_ID}
fi
export DEVICE_OAUTH_CLIENT_SECRET=${OAUTH_CLIENT_SECRET}

export COAP_GATEWAY_UNSECURE_FQDN=$FQDN
export COAP_GATEWAY_FQDN=$FQDN

mkdir -p $CA_POOL_DIR
mkdir -p $INTERNAL_CERT_DIR_PATH
mkdir -p $EXTERNAL_CERT_DIR_PATH
mkdir -p ${SECRETS_DIRECTORY}
ln -s ${SECRETS_DIRECTORY} /secrets

export CA_POOL=$CA_POOL_CERT_PATH
export CERT_FILE=$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_NAME
export KEY_FILE=$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_KEY_NAME

fqdnSAN="--cert.san.domain=$FQDN"
if ip route get $FQDN 2>/dev/null >/dev/null; then
  fqdnSAN="--cert.san.ip=$FQDN"
fi
echo "generating CA cert"
certificate-generator --cmd.generateRootCA --outCert=$CA_POOL_CERT_PATH --outKey=$CA_POOL_CERT_KEY_PATH --cert.subject.cn="Root CA"
echo "generating GRPC internal cert"
certificate-generator --cmd.generateCertificate --outCert=$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_NAME --outKey=$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_KEY_NAME --cert.subject.cn="localhost" --cert.san.domain="localhost" --cert.san.ip="0.0.0.0" --cert.san.ip="127.0.0.1" $fqdnSAN --signerCert=$CA_POOL_CERT_PATH --signerKey=$CA_POOL_CERT_KEY_PATH
echo "generating COAP-GW cert"
certificate-generator --cmd.generateIdentityCertificate=$COAP_GATEWAY_CLOUD_ID --outCert=$EXTERNAL_CERT_DIR_PATH/$COAP_GATEWAY_FILE_CERT_NAME --outKey=$EXTERNAL_CERT_DIR_PATH/$COAP_GATEWAY_FILE_CERT_KEY_NAME --cert.san.domain=$COAP_GATEWAY_FQDN --signerCert=$CA_POOL_CERT_PATH --signerKey=$CA_POOL_CERT_KEY_PATH
echo "generating NGINX cert"
certificate-generator --cmd.generateCertificate --outCert=$EXTERNAL_CERT_DIR_PATH/$DIAL_FILE_CERT_NAME --outKey=$EXTERNAL_CERT_DIR_PATH/$DIAL_FILE_CERT_KEY_NAME --cert.subject.cn="localhost" --cert.san.domain="localhost" --cert.san.ip="0.0.0.0" --cert.san.ip="127.0.0.1" $fqdnSAN --signerCert=$CA_POOL_CERT_PATH --signerKey=$CA_POOL_CERT_KEY_PATH

# copy ceritficates to paths
## oauth-server
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/oauth-server.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/oauth-server.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/oauth-server.yaml | sort | uniq)

## authorization
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/authorization.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/authorization.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/authorization.yaml | sort | uniq)

## resource-aggregate
### setup root-cas
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/resource-aggregate.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/resource-aggregate.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/resource-aggregate.yaml | sort | uniq)

## resource-directory
### setup root-cas
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/resource-directory.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/resource-directory.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/resource-directory.yaml | sort | uniq)

## coap-gateway
### setup root-cas
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/coap-gateway.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/coap-gateway.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/coap-gateway.yaml | sort | uniq)

## grpc-gateway
### setup root-cas
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/grpc-gateway.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/grpc-gateway.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/grpc-gateway.yaml | sort | uniq)

## http-gateway
### setup root-cas
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/http-gateway.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/http-gateway.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/http-gateway.yaml | sort | uniq)

## certificate-authority
### setup root-cas
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CA_POOL ${file}
done < <(yq e '[.. | select(has("caPool")) | .caPool]' /configs/certificate-authority.yaml | sort | uniq)
### setup certificates
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $CERT_FILE ${file}
done < <(yq e '[.. | select(has("certFile")) | .certFile]' /configs/certificate-authority.yaml | sort | uniq)
### setup private keys
while read -r line; do
  file=`echo $line | yq e '.[0]' - `
  mkdir -p `dirname ${file}`
  cp $KEY_FILE ${file}
done < <(yq e '[.. | select(has("keyFile")) | .keyFile]' /configs/certificate-authority.yaml | sort | uniq)


mkdir -p ${OAUTH_KEYS_PATH}
openssl genrsa -out ${OAUTH_ID_TOKEN_KEY_PATH} 4096
openssl ecparam -name prime256v1 -genkey -noout -out ${OAUTH_ACCESS_TOKEN_KEY_PATH}

mkdir -p $MONGO_PATH
mkdir -p $CERTIFICATES_PATH
mkdir -p $LOGS_PATH
mkdir -p ${NGINX_PATH}
cp /nginx/nginx.conf.template ${NGINX_PATH}/nginx.conf
sed -i "s/REPLACE_NGINX_PORT/$NGINX_PORT/g" ${NGINX_PATH}/nginx.conf
sed -i "s/REPLACE_HTTP_GATEWAY_PORT/$HTTP_GATEWAY_PORT/g" ${NGINX_PATH}/nginx.conf
sed -i "s/REPLACE_GRPC_GATEWAY_PORT/$GRPC_GATEWAY_PORT/g" ${NGINX_PATH}/nginx.conf
sed -i "s/REPLACE_MOCKED_OAUTH_SERVER_PORT/$MOCKED_OAUTH_SERVER_PORT/g" ${NGINX_PATH}/nginx.conf
sed -i "s/REPLACE_CERTIFICATE_AUTHORITY_PORT/$CERTIFICATE_AUTHORITY_PORT/g" ${NGINX_PATH}/nginx.conf
sed -i "s/REPLACE_AUTHORIZATION_HTTP_PORT/$AUTHORIZATION_HTTP_PORT/g" ${NGINX_PATH}/nginx.conf

# nats
echo "starting nats-server"
cat > /data/nats.config <<EOF
port: $NATS_PORT
max_pending: 128Mb
write_deadline: 10s
tls: {
  verify: true
  cert_file: "$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_NAME"
  key_file: "$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_KEY_NAME"
  ca_file: "$CA_POOL_CERT_PATH"
}
EOF
# Enable jetstream
if [ "${JETSTREAM}" = "true" ]; then
  cat >> /data/nats.config <<EOF
jetstream: {
  store_dir: "$JETSTREAM_PATH"
  // 1GB
    max_memory_store: 1073741824

  // 10GB
  max_file_store: 10737418240
}
EOF
fi
nats-server -c /data/nats.config >$LOGS_PATH/nats-server.log 2>&1 &
status=$?
nats_server_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start nats-server: $status"
  sync
  cat $LOGS_PATH/nats-server.log
  exit $status
fi

# waiting for nats. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if nc -z localhost $NATS_PORT; then
    break
  fi
  echo "Try to reconnect to nats(${NATS_HOST}) $i"
  sleep 1
done

if [ "${JETSTREAM}" = "true" ]; then
  echo "Setup streaming at nats"
  nats --tlscert="$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_NAME" --tlskey="$DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_KEY_NAME" --tlsca="$CA_POOL_CERT_PATH" str add EVENTS --config /configs/jetstream.json
fi

# mongo
echo "starting mongod"
cat $DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_NAME > $DIAL_FILE_CERT_DIR_PATH/mongo.key
cat $DIAL_FILE_CERT_DIR_PATH/$DIAL_FILE_CERT_KEY_NAME >> $DIAL_FILE_CERT_DIR_PATH/mongo.key
mongod --setParameter maxNumActiveUserIndexBuilds=64 --port $MONGO_PORT --dbpath $MONGO_PATH --sslMode requireSSL --sslCAFile $CA_POOL_CERT_PATH --sslPEMKeyFile $DIAL_FILE_CERT_DIR_PATH/mongo.key >$LOGS_PATH/mongod.log 2>&1 &
status=$?
mongo_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start mongod: $status"
  sync
  cat $LOGS_PATH/mongod.log
  exit $status
fi

# waiting for mongo DB. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${MONGODB_HOST} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to mongodb(${MONGODB_HOST}) $i"
  sleep 1
done

# starting nginx
echo "starting nginx"
nginx -c $NGINX_PATH/nginx.conf >$LOGS_PATH/nginx.log 2>&1
status=$?
nginx_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start nginx: $status"
  sync
  cat $LOGS_PATH/nginx.log
  exit $status
fi

i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${FQDN_NGINX_HTTPS} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to nginx(${FQDN_NGINX_HTTPS}) $i"
  sleep 1
done

# oauth-server
echo "starting oauth-server"

## setup cfg
cat /configs/oauth-server.yaml | yq e "\
  .apis.http.address = \"${MOCKED_OAUTH_SERVER_ADDRESS}\" |
  .oauthSigner.idTokenKeyFile = \"${OAUTH_ID_TOKEN_KEY_PATH}\" |
  .oauthSigner.accessTokenKeyFile = \"${OAUTH_ACCESS_TOKEN_KEY_PATH}\" |
  .oauthSigner.domain = \"${DOMAIN}\"
" - > /data/oauth-server.yaml

oauth-server --config /data/oauth-server.yaml >$LOGS_PATH/oauth-server.log 2>&1 &
status=$?
oauth_server_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start oauth-server: $status"
  sync
  cat $LOGS_PATH/oauth-server.log
  exit $status
fi

i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${MOCKED_OAUTH_SERVER_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to oauth-server(${MOCKED_OAUTH_SERVER_ADDRESS}) $i"
  sleep 1
done
    
# authorization
## configuration
cat /configs/authorization.yaml | yq e "\
  .apis.grpc.address = \"${AUTHORIZATION_ADDRESS}\" |
  .apis.grpc.authorization.http.tls.useSystemCAPool = true |
  .apis.grpc.authorization.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .apis.grpc.authorization.authority = \"https://${OAUTH_ENDPOINT}\" |
  .apis.http.address = \"${AUTHORIZATION_HTTP_ADDRESS}\" |
  .clients.storage.ownerClaim = \"${SERVICE_OWNER_CLAIM}\" |
  .clients.storage.mongoDB.uri = \"${MONGODB_URI}\" |
  .oauthClients.device.provider = \"${DEVICE_PROVIDER}\" |
  .oauthClients.device.clientID = \"${DEVICE_OAUTH_CLIENT_ID}\" |
  .oauthClients.device.clientSecret = \"${DEVICE_OAUTH_CLIENT_SECRET}\" |
  .oauthClients.device.authorizationURL = \"${DEVICE_OAUTH_ENDPOINT_AUTH_URL}\" |
  .oauthClients.device.redirectURL = \"${DEVICE_OAUTH_REDIRECT_URL}\" |
  .oauthClients.device.tokenURL = \"${DEVICE_OAUTH_ENDPOINT_TOKEN_URL}\" |
  .oauthClients.device.scopes = [ \"${DEVICE_OAUTH_SCOPES}\" ] |
  .oauthClients.device.http.tls.useSystemCAPool = true |
  .oauthClients.client.clientID = \"${SDK_OAUTH_CLIENT_ID}\" |
  .oauthClients.client.authorizationURL = \"${SDK_OAUTH_ENDPOINT_AUTH_URL}\" |
  .oauthClients.client.redirectURL = \"${SDK_OAUTH_REDIRECT_URL}\" |
  .oauthClients.client.audience = \"${SDK_OAUTH_AUDIENCE}\" |
  .oauthClients.client.http.tls.useSystemCAPool = true |
  .oauthClients.client.scopes=[ \"${SDK_OAUTH_SCOPES}\" ]
" - > /data/authorization.yaml

echo "starting authorization"
authorization --config /data/authorization.yaml >$LOGS_PATH/authorization.log 2>&1 &
status=$?
authorization_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start authorization: $status"
  sync
  cat $LOGS_PATH/authorization.log
  exit $status
fi

i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${AUTHORIZATION_HTTP_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to authorization service(${AUTHORIZATION_HTTP_ADDRESS}) $i"
  sleep 1
done

# resource-aggregate
## configuration
cat /configs/resource-aggregate.yaml | yq e "\
  .apis.grpc.address = \"${RESOURCE_AGGREGATE_ADDRESS}\" |
  .apis.grpc.authorization.http.tls.useSystemCAPool = true |
  .apis.grpc.authorization.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .apis.grpc.authorization.authority = \"https://${OAUTH_ENDPOINT}\" |
  .clients.eventStore.mongoDB.uri = \"${MONGODB_URI}\" |
  .clients.eventBus.nats.url = \"${NATS_URL}\" |
  .clients.eventBus.nats.jetstream = ${JETSTREAM} |
  .clients.authorizationServer.ownerClaim = \"${SERVICE_OWNER_CLAIM}\" |
  .clients.authorizationServer.grpc.address = \"${AUTHORIZATION_ADDRESS}\" |
  .clients.authorizationServer.oauth.clientID = \"${SERVICE_OAUTH_CLIENT_ID}\" |
  .clients.authorizationServer.oauth.clientSecret = \"${SERVICE_OAUTH_CLIENT_SECRET}\" |
  .clients.authorizationServer.oauth.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .clients.authorizationServer.oauth.http.tls.useSystemCAPool = true |
  .clients.authorizationServer.oauth.tokenURL = \"${SERVICE_OAUTH_ENDPOINT_TOKEN_URL}\"
" - > /data/resource-aggregate.yaml

echo "starting resource-aggregate"
resource-aggregate --config /data/resource-aggregate.yaml >$LOGS_PATH/resource-aggregate.log 2>&1 &
status=$?
resource_aggregate_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start resource-aggregate: $status"
  sync
  cat $LOGS_PATH/resource-aggregate.log
  exit $status
fi

# waiting for resource-aggregate. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${RESOURCE_AGGREGATE_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to resource-aggregate(${RESOURCE_AGGREGATE_ADDRESS}) $i"
  sleep 1
done

# resource-directory
## configuration
cat /configs/resource-directory.yaml | yq e "\
  .apis.grpc.address = \"${RESOURCE_DIRECTORY_ADDRESS}\" |
  .apis.grpc.authorization.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .apis.grpc.authorization.http.tls.useSystemCAPool = true |
  .apis.grpc.authorization.authority = \"https://${OAUTH_ENDPOINT}\" |
  .clients.eventStore.mongoDB.uri = \"${MONGODB_URI}\" |
  .clients.eventBus.nats.url = \"${NATS_URL}\" |
  .clients.authorizationServer.ownerClaim = \"${SERVICE_OWNER_CLAIM}\" |
  .clients.authorizationServer.grpc.address = \"${AUTHORIZATION_ADDRESS}\" |
  .clients.authorizationServer.oauth.clientID = \"${SERVICE_OAUTH_CLIENT_ID}\" |
  .clients.authorizationServer.oauth.clientSecret = \"${SERVICE_OAUTH_CLIENT_SECRET}\" |
  .clients.authorizationServer.oauth.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .clients.authorizationServer.oauth.http.tls.useSystemCAPool = true |
  .clients.authorizationServer.oauth.tokenURL = \"${SERVICE_OAUTH_ENDPOINT_TOKEN_URL}\" |
  .publicConfiguration.tokenURL = \"${SERVICE_CLIENT_CONFIGURATION_ACCESSTOKENURL}\" |
  .publicConfiguration.authorizationURL = \"${SERVICE_CLIENT_CONFIGURATION_AUTHCODEURL}\" |
  .publicConfiguration.cloudID = \"${COAP_GATEWAY_CLOUD_ID}\" |
  .publicConfiguration.cloudURL = \"coaps+tcp://${COAP_GATEWAY_FQDN}:${COAP_GATEWAY_PORT}\" |
  .publicConfiguration.signingServerAddress = \"${FQDN_NGINX_HTTPS}\" |
  .publicConfiguration.ownerClaim = \"${OWNER_CLAIM}\"
" - > /data/resource-directory.yaml

echo "starting resource-directory"
resource-directory --config /data/resource-directory.yaml >$LOGS_PATH/resource-directory.log 2>&1 &
status=$?
resource_directory_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start resource-directory: $status"
  sync
  cat $LOGS_PATH/resource-directory.log
  exit $status
fi

# waiting for resource-directory. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${RESOURCE_DIRECTORY_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to resource-directory(${RESOURCE_DIRECTORY_ADDRESS}) $i"
  sleep 1
done

# coap-gateway-unsecure
echo "starting coap-gateway-unsecure"
## configuration
cat /configs/coap-gateway.yaml | yq e "\
  .log.debug = ${LOG_DEBUG} |
  .log.dumpCoapMessages = ${COAP_GATEWAY_LOG_MESSAGES} |
  .apis.coap.address = \"${COAP_GATEWAY_UNSECURE_ADDRESS}\" |
  .apis.coap.externalAddress = \"${FQDN}:${COAP_GATEWAY_UNSECURE_PORT}\" |
  .apis.coap.tls.enabled = false |
  .clients.eventBus.nats.url = \"${NATS_URL}\" |
  .clients.resourceAggregate.grpc.address = \"${RESOURCE_AGGREGATE_ADDRESS}\" |
  .clients.resourceDirectory.grpc.address = \"${RESOURCE_DIRECTORY_ADDRESS}\" |
  .clients.authorizationServer.grpc.address = \"${AUTHORIZATION_ADDRESS}\" |
  .clients.authorizationServer.oauth.clientID = \"${SERVICE_OAUTH_CLIENT_ID}\" |
  .clients.authorizationServer.oauth.clientSecret = \"${SERVICE_OAUTH_CLIENT_SECRET}\" |
  .clients.authorizationServer.oauth.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .clients.authorizationServer.oauth.http.tls.useSystemCAPool = true |
  .clients.authorizationServer.oauth.tokenURL = \"${SERVICE_OAUTH_ENDPOINT_TOKEN_URL}\"
" - > /data/coap-gateway-unsecure.yaml

coap-gateway --config /data/coap-gateway-unsecure.yaml >$LOGS_PATH/coap-gateway-unsecure.log 2>&1 &
status=$?
coap_gw_unsecure_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start coap-gateway-unsecure: $status"
  sync
  cat $LOGS_PATH/coap-gateway-unsecure.log
  exit $status
fi

# coap-gateway-secure
echo "starting coap-gateway-secure"
### setup cfgs from env
cat /configs/coap-gateway.yaml | yq e "\
  .log.debug = ${LOG_DEBUG} |
  .log.dumpCoapMessages =  ${COAP_GATEWAY_LOG_MESSAGES} |
  .apis.coap.address = \"${COAP_GATEWAY_ADDRESS}\" |
  .apis.coap.externalAddress = \"${FQDN}:${COAP_GATEWAY_PORT}\" |
  .apis.coap.tls.enabled = true |
  .apis.coap.tls.keyFile = \"${EXTERNAL_CERT_DIR_PATH}/${COAP_GATEWAY_FILE_CERT_KEY_NAME}\" |
  .apis.coap.tls.certFile = \"${EXTERNAL_CERT_DIR_PATH}/${COAP_GATEWAY_FILE_CERT_NAME}\" |
  .clients.eventBus.nats.url = \"${NATS_URL}\" |
  .clients.resourceAggregate.grpc.address = \"${RESOURCE_AGGREGATE_ADDRESS}\" |
  .clients.resourceDirectory.grpc.address = \"${RESOURCE_DIRECTORY_ADDRESS}\" |
  .clients.authorizationServer.grpc.address = \"${AUTHORIZATION_ADDRESS}\" |
  .clients.authorizationServer.oauth.clientID = \"${SERVICE_OAUTH_CLIENT_ID}\" |
  .clients.authorizationServer.oauth.clientSecret = \"${SERVICE_OAUTH_CLIENT_SECRET}\" |
  .clients.authorizationServer.oauth.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .clients.authorizationServer.oauth.http.tls.useSystemCAPool = true |
  .clients.authorizationServer.oauth.tokenURL = \"${SERVICE_OAUTH_ENDPOINT_TOKEN_URL}\"
" - > /data/coap-gateway-secure.yaml

coap-gateway --config /data/coap-gateway-secure.yaml >$LOGS_PATH/coap-gateway.log 2>&1 &
status=$?
coap_gw_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start coap-gateway: $status"
  sync
  cat $LOGS_PATH/coap-gateway.log
  exit $status
fi

# grpc-gateway
echo "starting grpc-gateway"
## configuration
cat /configs/grpc-gateway.yaml | yq e "\
  .apis.grpc.address = \"${GRPC_GATEWAY_ADDRESS}\" |
  .apis.grpc.authorization.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .apis.grpc.authorization.http.tls.useSystemCAPool = true |
  .apis.grpc.authorization.authority = \"https://${OAUTH_ENDPOINT}\" |
  .clients.eventBus.nats.url = \"${NATS_URL}\" |
  .clients.resourceAggregate.grpc.address = \"${RESOURCE_AGGREGATE_ADDRESS}\" |
  .clients.resourceDirectory.grpc.address = \"${RESOURCE_DIRECTORY_ADDRESS}\"
" - > /data/grpc-gateway.yaml
grpc-gateway --config=/data/grpc-gateway.yaml >$LOGS_PATH/grpc-gateway.log 2>&1 &
status=$?
grpc_gw_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start grpc-gateway: $status"
  sync
  cat $LOGS_PATH/grpc-gateway.log
  exit $status
fi

# waiting for grpc-gateway. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${GRPC_GATEWAY_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to grpc-gateway(${GRPC_GATEWAY_ADDRESS}) $i"
  sleep 1
done


# http-gateway
## configuration
cat /configs/http-gateway.yaml | yq e "\
  .apis.http.address = \"${HTTP_GATEWAY_ADDRESS}\" |
  .apis.http.authorization.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .apis.http.authorization.http.tls.useSystemCAPool = true |
  .apis.http.authorization.authority = \"https://${OAUTH_ENDPOINT}\" |
  .clients.grpcGateway.grpc.address = \"${GRPC_GATEWAY_ADDRESS}\" |
  .ui.enabled = true |
  .ui.oauthClient.domain = \"${OAUTH_ENDPOINT}\" |
  .ui.oauthClient.clientID = \"${OAUTH_CLIENT_ID}\" |
  .ui.oauthClient.audience = \"${OAUTH_AUDIENCE}\" |
  .ui.oauthClient.scope = \"openid offline_access\" |
  .ui.oauthClient.httpGatewayAddress =\"https://${FQDN_NGINX_HTTPS}\"
" - > /data/http-gateway.yaml

echo "starting http-gateway"
http-gateway --config=/data/http-gateway.yaml >$LOGS_PATH/http-gateway.log 2>&1 &
status=$?
http_gw_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start http-gateway: $status"
  sync
  cat $LOGS_PATH/http-gateway.log
  exit $status
fi

# waiting for http-gateway. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${HTTP_GATEWAY_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to http-gateway(${HTTP_GATEWAY_ADDRESS}) $i"
  sleep 1
done

# certificate-authority
echo "starting certificate-authority"
## configuration
cat /configs/certificate-authority.yaml | yq e "\
  .apis.grpc.address = \"${CERTIFICATE_AUTHORITY_ADDRESS}\" |
  .apis.grpc.authorization.audience = \"${SERVICE_OAUTH_AUDIENCE}\" |
  .apis.grpc.authorization.http.tls.useSystemCAPool = true |
  .apis.grpc.authorization.authority = \"https://${OAUTH_ENDPOINT}\" |
  .signer.keyFile = \"${CA_POOL_CERT_KEY_PATH}\" |
  .signer.certFile = \"${CA_POOL_CERT_PATH}\"
" - > /data/certificate-authority.yaml
certificate-authority --config /data/certificate-authority.yaml >$LOGS_PATH/certificate-authority.log 2>&1 &
status=$?
certificate_authority_pid=$!
if [ $status -ne 0 ]; then
  echo "Failed to start certificate-authority: $status"
  sync
  cat $LOGS_PATH/certificate-authority.log
  exit $status
fi

# waiting for grpc-gateway. Without wait, sometimes auth service didn't connect.
i=0
while true; do
  i=$((i+1))
  if openssl s_client -connect ${CERTIFICATE_AUTHORITY_ADDRESS} -cert ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_NAME} -key ${INTERNAL_CERT_DIR_PATH}/${DIAL_FILE_CERT_KEY_NAME} <<< "Q" 2>/dev/null > /dev/null; then
    break
  fi
  echo "Try to reconnect to certificate-authority(${CERTIFICATE_AUTHORITY_ADDRESS}) $i"
  sleep 1
done


echo "Open browser at https://${DOMAIN}"

# Naive check runs checks once a minute to see if either of the processes exited.
# This illustrates part of the heavy lifting you need to do if you want to run
# more than one service in a container. The container exits with an error
# if it detects that either of the processes has exited.
# Otherwise it loops forever, waking up every 60 seconds
while sleep 10; do
  ps aux |grep $nats_server_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "nats-server has already exited."
    sync
    cat $LOGS_PATH/nats-server.log
    exit 1
  fi
  ps aux |grep $mongo_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "mongod has already exited."
    sync
    cat $LOGS_PATH/mongod.log
    exit 1
  fi
  ps aux |grep $authorization_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "authorization has already exited."
    sync
    cat $LOGS_PATH/authorization.log
    exit 1
  fi
  ps aux |grep $resource_aggregate_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "resource-aggregate has already exited."
    sync
    cat $LOGS_PATH/resource-aggregate.log
    exit 1
  fi
  ps aux |grep $resource_directory_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "resource-directory has already exited."
    sync
    cat $LOGS_PATH/resource-directory.log
    exit 1
  fi
  ps aux |grep $coap_gw_unsecure_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "coap-gateway-unsecure has already exited."
    sync
    cat $LOGS_PATH/coap-gateway-unsecure.log
    exit 1
  fi
  ps aux |grep $coap_gw_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "coap-gateway has already exited."
    sync
    cat $LOGS_PATH/coap-gateway.log
    exit 1
  fi
  ps aux |grep $grpc_gw_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "grpc-gateway has already exited."
    sync
    cat $LOGS_PATH/grpc-gateway.log
   exit 1
  fi
  ps aux |grep $http_gw_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "http-gateway has already exited."
    sync
    cat $LOGS_PATH/http-gateway.log
   exit 1
  fi
  ps aux |grep $certificate_authority_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "certificate-authority has already exited."
    sync
    cat $LOGS_PATH/certificate-authority.log
   exit 1
  fi
  ps aux |grep $oauth_server_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "oauth-server has already exited."
    sync
    cat $LOGS_PATH/oauth-server.log
   exit 1
  fi
  ps aux |grep $nginx_pid |grep -q -v grep
  if [ $? -ne 0 ]; then 
    echo "nginx has already exited."
    sync
    cat $LOGS_PATH/nginx.log
   exit 1
  fi
done
