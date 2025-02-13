package metrics // import "github.com/pomerium/pomerium/internal/telemetry/metrics"

import (
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// The following tags are applied to stats recorded by this package.
var (
	TagKeyHTTPMethod  tag.Key = tag.MustNewKey("http_method")
	TagKeyService     tag.Key = tag.MustNewKey("service")
	TagKeyGRPCService tag.Key = tag.MustNewKey("grpc_service")
	TagKeyGRPCMethod  tag.Key = tag.MustNewKey("grpc_method")
	TagKeyHost        tag.Key = tag.MustNewKey("host")
	TagKeyDestination tag.Key = tag.MustNewKey("destination")
)

// Default distributions used by views in this package.
var (
	DefaulHTTPSizeDistribution = view.Distribution(
		1, 256, 512, 1024, 2048, 8192, 16384, 32768, 65536, 131072, 262144,
		524288, 1048576, 2097152, 4194304, 8388608)
	DefaultHTTPLatencyDistrubtion = view.Distribution(
		1, 2, 5, 7, 10, 25, 500, 750, 100, 250, 500, 750, 1000, 2500, 5000,
		7500, 10000, 25000, 50000, 75000, 100000)
	grpcSizeDistribution = view.Distribution(
		1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024,
		2048, 4096, 8192, 16384,
	)
	DefaultMillisecondsDistribution = ocgrpc.DefaultMillisecondsDistribution
)

// DefaultViews are a set of default views to view HTTP and GRPC metrics.
var (
	DefaultViews = [][]*view.View{
		GRPCServerViews,
		HTTPServerViews,
		GRPCClientViews,
		GRPCServerViews}
)
