package export

import "github.com/prometheus/client_golang/prometheus"

// CoordinationExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type CoordinationExporter struct {
	TierTotalCapacity           *prometheus.GaugeVec `description:"Total capacity in bytes available in each tier."`
	TierRequiredCapacity        *prometheus.GaugeVec `description:"Total capacity in bytes required in each tier."`
	TierReplicationFactor       *prometheus.GaugeVec `description:"Configured maximum replication factor in each tier."`
	TierHistoricalCount         *prometheus.GaugeVec `description:"Number of available historical nodes in each tier."`
	SegmentUnderReplicatedCount *prometheus.GaugeVec `description:"Number of segments (including replicas) left to load until segments that should be loaded in the cluster are available for queries."`
	SegmentUnavailableCount     *prometheus.GaugeVec `description:"Number of segments (not including replicas) left to load until segments that should be loaded in the cluster are available for queries."`
	SegmentOvershadowedCount    *prometheus.GaugeVec `description:"Number of overshadowed segments."`
	SegmentCount                *prometheus.GaugeVec `description:"Number of used segments belonging to a data source. Emitted only for data sources to which at least one used segment belongs."`
	SegmentSize                 *prometheus.GaugeVec `description:"Total size of used segments in a data source. Emitted only for data sources to which at least one used segment belongs."`
	SegmentDropQueueCount       *prometheus.GaugeVec `description:"Number of segments to drop."`
	SegmentLoadQueueCount       *prometheus.GaugeVec `description:"Number of segments to load."`
	SegmentLoadQueueFailed      *prometheus.GaugeVec `description:"Number of segments that failed to load."`
	SegmentLoadQueueSize        *prometheus.GaugeVec `description:"Size in bytes of segments to load.	"`
	SegmentCostNormalized       *prometheus.GaugeVec `description:"Used in cost balancing. The normalized cost of hosting segments."`
	SegmentCostNormalization    *prometheus.GaugeVec `description:"Used in cost balancing. The normalization of hosting segments."`
	SegmentCostRaw              *prometheus.GaugeVec `description:"Used in cost balancing. The raw cost of hosting segments."`
	SegmentUnneededCount        *prometheus.GaugeVec `description:"Number of segments dropped due to being marked as unused."`
	SegmentDeletedCount         *prometheus.GaugeVec `description:"Number of segments dropped due to rules."`
	SegmentDroppedCount         *prometheus.GaugeVec `description:"Number of segments dropped due to being overshadowed."`
	SegmentMovedCount           *prometheus.GaugeVec `description:"Number of segments moved in the cluster."`
	SegmentAssignedCount        *prometheus.GaugeVec `description:"number of segments assigned to be loaded in the cluster"`
}

// NewCoordinationExporter returns a new Jetty exporter object
func NewCoordinationExporter() *CoordinationExporter {
	ce := &CoordinationExporter{
		TierTotalCapacity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_total_capacity",
			Help:      "Total capacity in bytes available in each tier.",
		}, []string{"tier"}),
		TierRequiredCapacity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_required_capacity",
			Help:      "Total capacity in bytes required in each tier.",
		}, []string{"tier"}),
		TierHistoricalCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_historical_count",
			Help:      "Number of available historical nodes in each tier.",
		}, []string{"tier"}),
		TierReplicationFactor: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_replication_factor",
			Help:      "Configured maximum replication factor in each tier.",
		}, []string{"tier"}),
		SegmentCostNormalization: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_cost_normalization",
			Help:      "Used in cost balancing. The normalization of hosting segments.",
		}, []string{"tier"}),
		SegmentCostNormalized: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_cost_normalized",
			Help:      "Used in cost balancing. The normalized cost of hosting segments.",
		}, []string{"tier"}),
		SegmentCostRaw: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_cost_raw",
			Help:      "Used in cost balancing. The raw cost of hosting segments.",
		}, []string{"tier"}),
		SegmentCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_count",
			Help:      "Number of used segments belonging to a data source. Emitted only for data sources to which at least one used segment belongs.",
		}, []string{"dataSource"}),
		SegmentDeletedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_deleted_count",
			Help:      "Number of segments dropped due to rules.",
		}, []string{"tier"}),
		SegmentDropQueueCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_drop_queue_count",
			Help:      "Number of segments to drop.",
		}, []string{"server"}),
		SegmentDroppedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_dropped_count",
			Help:      "Number of segments dropped due to being overshadowed.",
		}, []string{"tier"}),
		SegmentLoadQueueCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_load_queue_count",
			Help:      "Number of segments to load.",
		}, []string{"server"}),
		SegmentLoadQueueFailed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_load_queue_failed",
			Help:      "Number of segments that failed to load.",
		}, []string{"server"}),
		SegmentLoadQueueSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_load_queue_size",
			Help:      "Size in bytes of segments to load",
		}, []string{"server"}),
		SegmentMovedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_moved_count",
			Help:      "Number of segments moved in the cluster.",
		}, []string{"tier"}),
		SegmentOvershadowedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_overshadowed_count",
			Help:      "Number of overshadowed segments.",
		}, []string{}),
		SegmentSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_size",
			Help:      "Total size of used segments in a data source. Emitted only for data sources to which at least one used segment belongs.",
		}, []string{"dataSource"}),
		SegmentUnavailableCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_unavailable_count",
			Help:      "Number of segments (not including replicas) left to load until segments that should be loaded in the cluster are available for queries.",
		}, []string{"dataSource"}),
		SegmentUnderReplicatedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_underreplicated_count",
			Help:      "Number of segments (including replicas) left to load until segments that should be loaded in the cluster are available for queries.",
		}, []string{"tier", "dataSource"}),
		SegmentUnneededCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_unneeded_count",
			Help:      "Number of segments dropped due to being marked as unused.",
		}, []string{"tier"}),
		SegmentAssignedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_assigned_count",
			Help:      "number of segments assigned to be loaded in the cluster",
		}, []string{"tier"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(ce.TierHistoricalCount)
	prometheus.MustRegister(ce.TierReplicationFactor)
	prometheus.MustRegister(ce.TierRequiredCapacity)
	prometheus.MustRegister(ce.TierTotalCapacity)
	prometheus.MustRegister(ce.SegmentCostNormalization)
	prometheus.MustRegister(ce.SegmentCostNormalized)
	prometheus.MustRegister(ce.SegmentCostRaw)
	prometheus.MustRegister(ce.SegmentCount)
	prometheus.MustRegister(ce.SegmentDeletedCount)
	prometheus.MustRegister(ce.SegmentDropQueueCount)
	prometheus.MustRegister(ce.SegmentDroppedCount)
	prometheus.MustRegister(ce.SegmentLoadQueueCount)
	prometheus.MustRegister(ce.SegmentLoadQueueFailed)
	prometheus.MustRegister(ce.SegmentLoadQueueSize)
	prometheus.MustRegister(ce.SegmentMovedCount)
	prometheus.MustRegister(ce.SegmentAssignedCount)
	prometheus.MustRegister(ce.SegmentOvershadowedCount)
	prometheus.MustRegister(ce.SegmentSize)
	prometheus.MustRegister(ce.SegmentUnavailableCount)
	prometheus.MustRegister(ce.SegmentUnderReplicatedCount)
	prometheus.MustRegister(ce.SegmentUnneededCount)

	return ce
}

// SetTierTotalCapacity .
func (ce *CoordinationExporter) SetTierTotalCapacity(tier string, val float64) {
	ce.TierTotalCapacity.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetTierRequiredCapacity .
func (ce *CoordinationExporter) SetTierRequiredCapacity(tier string, val float64) {
	ce.TierRequiredCapacity.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetTierReplicationFactor .
func (ce *CoordinationExporter) SetTierReplicationFactor(tier string, val float64) {
	ce.TierReplicationFactor.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetTierHistoricalCount .
func (ce *CoordinationExporter) SetTierHistoricalCount(tier string, val float64) {
	ce.TierHistoricalCount.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentUnderReplicatedCount .
func (ce *CoordinationExporter) SetSegmentUnderReplicatedCount(labels map[string]string, val float64) {
	ce.SegmentUnderReplicatedCount.With(labels).Add(val)
}

// SetSegmentUnavailableCount .
func (ce *CoordinationExporter) SetSegmentUnavailableCount(source string, val float64) {
	ce.SegmentUnavailableCount.With(prometheus.Labels{"dataSource": source}).Set(val)
}

// SetSegmentOvershadowedCount .
func (ce *CoordinationExporter) SetSegmentOvershadowedCount(val float64) {
	ce.SegmentOvershadowedCount.WithLabelValues().Add(val)
}

// SetSegmentCount .
func (ce *CoordinationExporter) SetSegmentCount(source string, val float64) {
	ce.SegmentCount.With(prometheus.Labels{"dataSource": source}).Set(val)
}

// SetSegmentSize .
func (ce *CoordinationExporter) SetSegmentSize(source string, val float64) {
	ce.SegmentSize.With(prometheus.Labels{"dataSource": source}).Set(val)
}

// SetSegmentDropQueueCount .
func (ce *CoordinationExporter) SetSegmentDropQueueCount(server string, val float64) {
	ce.SegmentDropQueueCount.With(prometheus.Labels{"server": server}).Set(val)
}

// SetSegmentLoadQueueCount .
func (ce *CoordinationExporter) SetSegmentLoadQueueCount(server string, val float64) {
	ce.SegmentLoadQueueCount.With(prometheus.Labels{"server": server}).Set(val)
}

// SetSegmentLoadQueueFailed .
func (ce *CoordinationExporter) SetSegmentLoadQueueFailed(server string, val float64) {
	ce.SegmentLoadQueueFailed.With(prometheus.Labels{"server": server}).Set(val)
}

// SetSegmentLoadQueueSize .
func (ce *CoordinationExporter) SetSegmentLoadQueueSize(server string, val float64) {
	ce.SegmentLoadQueueSize.With(prometheus.Labels{"server": server}).Set(val)
}

// SetSegmentCostNormalized .
func (ce *CoordinationExporter) SetSegmentCostNormalized(tier string, val float64) {
	ce.SegmentCostNormalized.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentCostNormalization .
func (ce *CoordinationExporter) SetSegmentCostNormalization(tier string, val float64) {
	ce.SegmentCostNormalization.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentCostRaw .
func (ce *CoordinationExporter) SetSegmentCostRaw(tier string, val float64) {
	ce.SegmentCostRaw.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentUnneededCount .
func (ce *CoordinationExporter) SetSegmentUnneededCount(tier string, val float64) {
	ce.SegmentUnneededCount.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentDeletedCount .
func (ce *CoordinationExporter) SetSegmentDeletedCount(tier string, val float64) {
	ce.SegmentDeletedCount.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentDroppedCount .
func (ce *CoordinationExporter) SetSegmentDroppedCount(tier string, val float64) {
	ce.SegmentDroppedCount.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentMovedCount .
func (ce *CoordinationExporter) SetSegmentMovedCount(tier string, val float64) {
	ce.SegmentMovedCount.With(prometheus.Labels{"tier": tier}).Set(val)
}

// SetSegmentAssignedCount .
func (ce *CoordinationExporter) SetSegmentAssignedCount(tier string, val float64) {
	ce.SegmentAssignedCount.With(prometheus.Labels{"tier": tier}).Set(val)
}
