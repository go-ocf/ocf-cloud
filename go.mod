module github.com/plgd-dev/cloud

go 1.16

require (
	github.com/fsnotify/fsnotify v1.5.1
	github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/golang/snappy v0.0.4
	github.com/google/go-querystring v1.1.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/json-iterator/go v1.1.12
	github.com/jtacoma/uritemplates v1.0.0
	github.com/karrick/tparse/v2 v2.8.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lestrrat-go/jwx v1.2.6
	github.com/nats-io/nats.go v1.12.1
	github.com/panjf2000/ants/v2 v2.4.6
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/plgd-dev/go-coap/v2 v2.4.1-0.20210917095946-231b8cdcfcc3
	github.com/plgd-dev/kit v0.0.0-20210614190235-99984a49de48
	github.com/plgd-dev/sdk v0.0.0-20210830094304-889b5b108b02
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.7.2
	go.uber.org/atomic v1.9.0
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.0.0-20210908191846-a5e095526f91
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20210909211513-a8c4777a87af
	google.golang.org/grpc v1.40.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

replace gopkg.in/yaml.v2 v2.4.0 => github.com/cizmazia/yaml v0.0.0-20200220134304-2008791f5454
