module ecos

go 1.16

require (
	github.com/SUMStudio/grocksdb v1.6.47
	github.com/aws/aws-sdk-go-v2 v1.16.2
	github.com/aws/aws-sdk-go-v2/config v1.15.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.26.4
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/cockroachdb/errors v1.9.0 // indirect
	github.com/deathowl/go-metrics-prometheus v0.0.0-20200518174047-74482eab5bfb
	github.com/deckarep/golang-set v1.8.0
	github.com/elliotchance/orderedmap v1.4.0
	github.com/getsentry/sentry-go v0.13.0 // indirect
	github.com/gin-contrib/timeout v0.0.3
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jedib0t/go-pretty/v6 v6.3.0
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/minio/minio-go/v7 v7.0.43
	github.com/minio/sha256-simd v1.0.0
	github.com/mitchellh/copystructure v1.2.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/orcaman/concurrent-map v1.0.0
	github.com/prometheus/client_golang v1.12.1
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/serialx/hashring v0.0.0-20200727003509-22c0c7ab6b1b
	github.com/shirou/gopsutil/v3 v3.22.3
	github.com/sincerexia/gocrush v0.0.0-20220213012034-a24f00a6eb8e
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.8.1
	github.com/twmb/murmur3 v1.1.6
	github.com/wxnacy/wgo v1.0.4
	go.etcd.io/etcd/client/pkg/v3 v3.5.2
	go.etcd.io/etcd/pkg/v3 v3.5.2
	go.etcd.io/etcd/raft/v3 v3.5.2
	go.etcd.io/etcd/server/v3 v3.5.2
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261
	google.golang.org/api v0.58.0 // indirect
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4 // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
