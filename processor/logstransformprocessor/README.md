# Logs Transform Processor

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [development]: logs   |
| Distributions | [observiq], [splunk], [sumo] |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aprocessor%2Flogstransform%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aprocessor%2Flogstransform%20&label=closed&color=blue&logo=opentelemetry) |

[development]: https://github.com/open-telemetry/opentelemetry-collector#development
[observiq]: https://github.com/observIQ/observiq-otel-collector
[splunk]: https://github.com/signalfx/splunk-otel-collector
[sumo]: https://github.com/SumoLogic/sumologic-otel-collector
<!-- end autogenerated section -->

NOTE - This processor is experimental, with the intention that its functionality will be reimplemented in the [transform processor](../transformprocessor/README.md) in the future.

The logs transform processor can be used to apply [log operators](../../pkg/stanza/docs/operators) to logs coming from any receiver.
Please refer to [config.go](./config.go) for the config spec.

Examples:

```yaml
processors:
  logstransform:
    operators:
      - type: regex_parser
        regex: '^(?P<time>\d{4}-\d{2}-\d{2}) (?P<sev>[A-Z]*) (?P<msg>.*)$'
        timestamp:
          parse_from: body.time
          layout: "%Y-%m-%d"
        severity:
          parse_from: body.sev
```

Refer to [config.yaml](./testdata/config.yaml) for detailed
examples on using the processor.