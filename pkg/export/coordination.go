package export

import "github.com/prometheus/client_golang/prometheus"

// CoordinationExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type CoordinationExporter struct {
	TierTotalCapacity           prometheus.Counter `description:"Total capacity in bytes available in each tier."`
	TierRequiredCapacity        prometheus.Counter `description:"Total capacity in bytes required in each tier."`
	TierReplicationFactor       prometheus.Counter `description:"Configured maximum replication factor in each tier."`
	TierHistoricalCount         prometheus.Counter `description:"Number of available historical nodes in each tier."`
	SegmentUnderReplicatedCount prometheus.Counter `description:"Number of segments (including replicas) left to load until segments that should be loaded in the cluster are available for queries."`
	SegmentUnavailableCount     prometheus.Counter `description:"Number of segments (not including replicas) left to load until segments that should be loaded in the cluster are available for queries."`
	SegmentOvershadowedCount    prometheus.Counter `description:"Number of overshadowed segments."`
	SegmentCount                prometheus.Counter `description:"Number of used segments belonging to a data source. Emitted only for data sources to which at least one used segment belongs."`
	SegmentSize                 prometheus.Counter `description:"Total size of used segments in a data source. Emitted only for data sources to which at least one used segment belongs."`
	SegmentDropQueueCount       prometheus.Counter `description:"Number of segments to drop."`
	SegmentLoadQueueCount       prometheus.Counter `description:"Number of segments to load."`
	SegmentLoadQueueFailed      prometheus.Counter `description:"Number of segments that failed to load."`
	SegmentLoadQueueSize        prometheus.Counter `description:"Size in bytes of segments to load.	"`
	SegmentCostNormalized       prometheus.Counter `description:"Used in cost balancing. The normalized cost of hosting segments."`
	SegmentCostNormalization    prometheus.Counter `description:"Used in cost balancing. The normalization of hosting segments."`
	SegmentCostRaw              prometheus.Counter `description:"Used in cost balancing. The raw cost of hosting segments."`
	SegmentUnneededCount        prometheus.Counter `description:"Number of segments dropped due to being marked as unused."`
	SegmentDeletedCount         prometheus.Counter `description:"Number of segments dropped due to rules."`
	SegmentDroppedCount         prometheus.Counter `description:"Number of segments dropped due to being overshadowed."`
	SegmentMovedCount           prometheus.Counter `description:"Number of segments moved in the cluster."`
	SegmentAssignedCount        prometheus.Counter `description:"number of segments assigned to be loaded in the cluster"`
}

// NewCoordinationExporter returns a new Jetty exporter object
func NewCoordinationExporter() *CoordinationExporter {
	ce := &CoordinationExporter{
		TierTotalCapacity: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_total_capacity",
			Help:      "Total capacity in bytes available in each tier.",
			ConstLabels: prometheus.Labels{
				"coordinator": "tier-total-capacity",
			},
		}),
		TierRequiredCapacity: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_required_capacity",
			Help:      "Total capacity in bytes required in each tier.",
			ConstLabels: prometheus.Labels{
				"coordinator": "tier-required-capacity",
			},
		}),
		TierHistoricalCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_historical_count",
			Help:      "Number of available historical nodes in each tier.",
			ConstLabels: prometheus.Labels{
				"coordinator": "tier-historical-count",
			},
		}),
		TierReplicationFactor: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "tier_replication_factor",
			Help:      "Configured maximum replication factor in each tier.",
			ConstLabels: prometheus.Labels{
				"coordinator": "tier-replication-factor",
			},
		}),
		SegmentCostNormalization: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_cost_normalization",
			Help:      "Used in cost balancing. The normalization of hosting segments.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-cost-normalization",
			},
		}),
		SegmentCostNormalized: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_cost_normalized",
			Help:      "Used in cost balancing. The normalized cost of hosting segments.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-cost-normalized",
			},
		}),
		SegmentCostRaw: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_cost_raw",
			Help:      "Used in cost balancing. The raw cost of hosting segments.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-cost-raw",
			},
		}),
		SegmentCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_count",
			Help:      "Number of used segments belonging to a data source. Emitted only for data sources to which at least one used segment belongs.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-count",
			},
		}),
		SegmentDeletedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_deleted_count",
			Help:      "Number of segments dropped due to rules.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-deleted-count",
			},
		}),
		SegmentDropQueueCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_drop_queue_count",
			Help:      "Number of segments to drop.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-drop-queue-count",
			},
		}),
		SegmentDroppedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_dropped_count",
			Help:      "Number of segments dropped due to being overshadowed.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-dropped-count",
			},
		}),
		SegmentLoadQueueCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_load_queue_count",
			Help:      "Number of segments to load.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-load-queue-count",
			},
		}),
		SegmentLoadQueueFailed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_load_queue_failed",
			Help:      "Number of segments that failed to load.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-load-queue-failed",
			},
		}),
		SegmentLoadQueueSize: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_load_queue_size",
			Help:      "Size in bytes of segments to load",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-load-queue-size",
			},
		}),
		SegmentMovedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_moved_count",
			Help:      "Number of segments moved in the cluster.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-moved-count",
			},
		}),
		SegmentOvershadowedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_overshadowed_count",
			Help:      "Number of overshadowed segments.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-overshadowed-count",
			},
		}),
		SegmentSize: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_size",
			Help:      "Total size of used segments in a data source. Emitted only for data sources to which at least one used segment belongs.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-size",
			},
		}),
		SegmentUnavailableCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_unavailable_count",
			Help:      "Number of segments (not including replicas) left to load until segments that should be loaded in the cluster are available for queries.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-unavailable-count",
			},
		}),
		SegmentUnderReplicatedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_underreplicated_count",
			Help:      "Number of segments (including replicas) left to load until segments that should be loaded in the cluster are available for queries.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-underreplicated-count",
			},
		}),
		SegmentUnneededCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_unneeded_count",
			Help:      "Number of segments dropped due to being marked as unused.",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-unneeded-count",
			},
		}),
		SegmentAssignedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_assigned_count",
			Help:      "number of segments assigned to be loaded in the cluster",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-assigned-count",
			},
		}),
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
func (ce *CoordinationExporter) SetTierTotalCapacity(val float64) {}

// SetTierRequiredCapacity .
func (ce *CoordinationExporter) SetTierRequiredCapacity(val float64) {}

// SetTierReplicationFactor .
func (ce *CoordinationExporter) SetTierReplicationFactor(val float64) {}

// SetTierHistoricalCount .
func (ce *CoordinationExporter) SetTierHistoricalCount(val float64) {}

// SetSegmentUnderReplicatedCount .
func (ce *CoordinationExporter) SetSegmentUnderReplicatedCount(val float64) {}

// SetSegmentUnavailableCount .
func (ce *CoordinationExporter) SetSegmentUnavailableCount(val float64) {}

// SetSegmentOvershadowedCount .
func (ce *CoordinationExporter) SetSegmentOvershadowedCount(val float64) {}

// SetSegmentCount .
func (ce *CoordinationExporter) SetSegmentCount(val float64) {}

// SetSegmentSize .
func (ce *CoordinationExporter) SetSegmentSize(val float64) {}

// SetSegmentDropQueueCount .
func (ce *CoordinationExporter) SetSegmentDropQueueCount(val float64) {}

// SetSegmentLoadQueueCount .
func (ce *CoordinationExporter) SetSegmentLoadQueueCount(val float64) {}

// SetSegmentLoadQueueFailed .
func (ce *CoordinationExporter) SetSegmentLoadQueueFailed(val float64) {}

// SetSegmentLoadQueueSize .
func (ce *CoordinationExporter) SetSegmentLoadQueueSize(val float64) {}

// SetSegmentCostNormalized .
func (ce *CoordinationExporter) SetSegmentCostNormalized(val float64) {}

// SetSegmentCostNormalization .
func (ce *CoordinationExporter) SetSegmentCostNormalization(val float64) {}

// SetSegmentCostRaw .
func (ce *CoordinationExporter) SetSegmentCostRaw(val float64) {}

// SetSegmentUnneededCount .
func (ce *CoordinationExporter) SetSegmentUnneededCount(val float64) {}

// SetSegmentDeletedCount .
func (ce *CoordinationExporter) SetSegmentDeletedCount(val float64) {}

// SetSegmentDroppedCount .
func (ce *CoordinationExporter) SetSegmentDroppedCount(val float64) {}

// SetSegmentMovedCount .
func (ce *CoordinationExporter) SetSegmentMovedCount(val float64) {}

// SetSegmentAssignedCount .
func (ce *CoordinationExporter) SetSegmentAssignedCount(val float64) {}
