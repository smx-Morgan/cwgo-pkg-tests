module cwgo-test

go 1.21

replace github.com/cloudwego-contrib/cwgo-pkg/registry/consul => ../../cwgo-pkg-registry/registry/consul

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

replace github.com/cloudwego-contrib/cwgo-pkg/config/consul => ../../cwgo-pkg-registry/config/consul

replace github.com/cloudwego-contrib/cwgo-pkg/config/common => ../../cwgo-pkg-registry/config/common

replace github.com/cloudwego-contrib/cwgo-pkg/registry/nacos => ../../cwgo-pkg-registry/registry/nacos

replace github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoshertz/v2 => ../../cwgo-pkg-registry/registry/nacos/nacoshertz/v2

replace github.com/cloudwego-contrib/cwgo-pkg/config/nacos => ../../cwgo-pkg-registry/config/nacos

replace github.com/cloudwego-contrib/cwgo-pkg/config/nacos/v2 => ../../cwgo-pkg-registry/config/nacos/v2

replace github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoskitex/v2 => ../../cwgo-pkg-registry/registry/nacos/nacoskitex/v2

replace github.com/cloudwego-contrib/cwgo-pkg/registry/etcd => ../../cwgo-pkg-registry/registry/etcd

replace github.com/cloudwego-contrib/cwgo-pkg/config/etcd => ../../cwgo-pkg-registry/config/etcd

replace github.com/cloudwego-contrib/cwgo-pkg/registry/eureka => ../../cwgo-pkg-registry/registry/eureka

replace github.com/cloudwego-contrib/cwgo-pkg/registry/polaris => ../../cwgo-pkg-registry/registry/polaris

replace github.com/cloudwego-contrib/cwgo-pkg/registry/redis => ../../cwgo-pkg-registry/registry/redis

replace github.com/cloudwego-contrib/cwgo-pkg/registry/servicecomb => ../../cwgo-pkg-registry/registry/servicecomb

replace github.com/cloudwego-contrib/cwgo-pkg/registry/zookeeper => ../../cwgo-pkg-registry/registry/zookeeper

replace github.com/cloudwego-contrib/cwgo-pkg/config/zookeeper => ../../cwgo-pkg-registry/config/zookeeper

replace github.com/cloudwego-contrib/cwgo-pkg/config/apollo => ../../cwgo-pkg-registry/config/apollo

replace github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/options => ../../cwgo-pkg-registry/registry/nacos/options

require (
	github.com/cloudwego-contrib/cwgo-pkg/config/apollo v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/config/consul v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/config/etcd v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/config/nacos v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/config/nacos/v2 v2.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/config/zookeeper v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/consul v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/etcd v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/eureka v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/nacos v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoshertz/v2 v2.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoskitex/v2 v2.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/polaris v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/redis v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/servicecomb v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/registry/zookeeper v0.0.0-00010101000000-000000000000
	github.com/cloudwego-contrib/cwgo-pkg/telemetry v0.0.0-20241014044734-80a98dbe0b6a
	github.com/cloudwego/gopkg v0.1.2
	github.com/cloudwego/hertz v0.9.3
	github.com/cloudwego/kitex v0.11.3
	github.com/cloudwego/kitex-examples v0.3.3
	github.com/go-chassis/sc-client v0.7.1-0.20220829010936-e0ff6c891c04
	github.com/hashicorp/consul/api v1.29.5
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus v0.1.1
	github.com/hudl/fargo v1.4.0
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20240205032422-93b4c82b7dcd
	github.com/nacos-group/nacos-sdk-go v1.1.5
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.7
	github.com/prometheus/client_golang v1.19.1
	go.opentelemetry.io/otel v1.28.0
)

require (
	github.com/alibabacloud-go/debug v1.0.0 // indirect
	github.com/alibabacloud-go/tea v1.1.17 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.4 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.63.27 // indirect
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.2.2 // indirect
	github.com/aliyun/alibabacloud-dkms-transfer-go-sdk v0.1.7 // indirect
	github.com/apache/thrift v0.20.0 // indirect
	github.com/apolloconfig/agollo/v4 v4.3.1 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/bytedance/sonic v1.12.2 // indirect
	github.com/bytedance/sonic/loader v0.2.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/cloudwego-contrib/cwgo-pkg/config/common v0.0.0-00010101000000-000000000000 // indirect
	github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/options v0.0.0-00010101000000-000000000000 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/configmanager v0.2.2 // indirect
	github.com/cloudwego/dynamicgo v0.4.0 // indirect
	github.com/cloudwego/fastpb v0.0.5 // indirect
	github.com/cloudwego/frugal v0.2.0 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cloudwego/localsession v0.0.2 // indirect
	github.com/cloudwego/netpoll v0.6.4 // indirect
	github.com/cloudwego/runtimex v0.1.0 // indirect
	github.com/cloudwego/thriftgo v0.3.17 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/franela/goreq v0.0.0-20171204163338-bcd34c9993f8 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-chassis/cari v0.9.0 // indirect
	github.com/go-chassis/foundation v0.4.0 // indirect
	github.com/go-chassis/openlog v1.1.3 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-zookeeper/zk v1.0.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.4.3-0.20210424162022-e8629af678b7 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karlseguin/ccache/v2 v2.0.8 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/polarismesh/polaris-go v1.3.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/redis/go-redis/v9 v9.4.0 // indirect
	github.com/shima-park/agollo v1.2.14 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/thoas/go-funk v0.9.3 // indirect
	github.com/tidwall/gjson v1.17.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/v3 v3.5.12 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.45.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.20.0 // indirect
	go.opentelemetry.io/contrib/propagators/ot v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/gcfg.v1 v1.2.3 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
