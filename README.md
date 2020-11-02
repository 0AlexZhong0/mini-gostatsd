# Plan

## Project strategy and workflow

- What components are there?

`statsd` aggregates the metrics such as counters and gauges and sends them to the backend (`Graphite` is the only supported backend at the moment). The project runs `Graphite` and its I/O companion Carbon from using the respective docker images.

- How do I structure the project?

There needs no folder structure for a tiny toy project. It is feasible to do everything inside one file for now.

- How should each component interact?

Spin up `statsd` by running the main module, it then establishes a connection with Graphite via the configured ports, and finally it flushes stats to Graphite in a set interval.

- How easy is it to test each component either individually or together?

Refer to the statsd and the gostatsd repository to see what tests they conduct.

## Implementation details

- How do I send metrics continuously to `Graphite` in a set interval? (where would this implementation be)

Using `time.Tick(interval)`, for every specified set interval, post the stat to Graphite.

- How do I write data to Graphite once the connection is established?

We can feed multiple metrics to `Graphite` by writing data such as `metric_one`, `metric_two` and so one where each of them are in the plain text format specified in the [Graphite documentation](https://graphite.readthedocs.io/en/latest/feeding-carbon.html) and separated by a **new line character** `\n`. We can make use of an array to store the metrics.

# Overview

An implementation of [statsd](https://github.com/statsd/statsd) in Go for learning purposes. `mini_gostatsd` currently aggregates the metrics specified below and `Graphite` is the only supported backend.

## Metrics

- `mini_gostatsd.graphiteStats.last_flush`: the time which the stats were last flushed/sent to `Graphite`.

# Get started

## Using `docker`
Explain my troubles of running `Graphite` locally, and state the reasons for why using an existing docker image is the most reliable way. Of course, give the instructions here, too.

# CI/CD

## Testing
`Coveralls` for test coverage reports and make use of a widely used Go testing library/framework. Get inspiration from the statsd project to see what sort of unit tests or integration tests they conduct.

## Pipeline
`Travis CI`. Talk about some house-keeping stuff, such as setting up the pipeline in `Travis CI` if you don't have it already.

# Future work

- Run the process/daemon from a configuration file rather than hard-coded parameters.
- Add different metric types such as gauges, histogram and summary.
- Generate graphs and add them to Grafana without manual addition.
- Add aws-cloudwatch and elasticsearch to supported backends. (this means I have to have a `backends` folder where I need to specify the initialization function of each backend) See [here](https://github.com/statsd/statsd/blob/master/docs/backend.md) for a comprehensive list of backends.
