package export

// Exporter is the interface that will format the metric feeds received from Druid
// into a more appropriate struct
type Exporter interface {
	FormatMetrics(chan string) interface{}
}
