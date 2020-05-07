# Apache Druid Prometheus Exporter

By setting up Apache Druid to push the metrics to a HTTP service, you are able to collect metrics for monitoring. This service exposes only two endpoints:

1. `/collect`
2. `/metrics`

## Collect endpoint

This endpoint is a `POST` where you can instruct Druid to send the metrics to. The endpoint will parse each `feed` and update the corresponded prometheus metric.

## Metrics endpoint

The prometheus client exposes all the metrics on this endpoint. In this way, your prometheus server can read the exposed metrics.
