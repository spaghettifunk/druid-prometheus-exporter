# Apache Druid Prometheus Exporter

By setting up Apache Druid to push the metrics to a HTTP service, you are able to collect metrics for monitoring. This service exposes only two endpoints:

1. `/collect`
2. `/metrics`

## Collect endpoint

This endpoint is a `POST` where you can instruct Druid to send the metrics to. The endpoint will parse each `feed` and update the corresponded prometheus metric.

## Metrics endpoint

The prometheus client exposes all the metrics on this endpoint. In this way, your prometheus server can read the exposed metrics.

## Druid configuration

You need to configure Druid to send the metrics to this service. To do so,you need to change the configuration file to add/update it with the following entries

```
druid_emitter_logging_logLevel=debug
druid_emitter_logging_loggerClass=HttpPostEmitter
druid_emitter=http
druid_emitter_http_recipientBaseUrl=http://{exporter_addr}:7000/collect
```

In the `config` folder of this project, you can find an example.