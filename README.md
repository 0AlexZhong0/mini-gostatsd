# Plan

## Project strategy and workflow

- What components are there?
- How do I structure the project?
- How should each component interact?
- How easy is it to test each component either individually or together? (refer to statsd to see what tests they conduct)
- Let's implement them!!!

# Overview

An implementation of [statsd](https://github.com/statsd/statsd) in golang for learning purposes. `mini_gostatsd` currently only aggregates X_METRIC and Graphite is the only supported backend.

# Get started

## Using docker
Explain my troubles of running Graphite locally, and state the reasons for why using an existing docker image is the most reliable way. Of course, give the instructions here, too.

# CI/CD

## Testing
Coveralls for test coverage reports and make use of a widely used Go testing library/framework. Get inspiration from the statsd project to see what sort of unit tests or integration tests they conduct.

## Pipeline
Travis CI. Talk about some house-keeping stuff, such as setting up the pipeline in Travis CI if you don't have it already.

# Future work

- Add different metric types such as gauges, histogram and summary.
- Generate graphs and add them to Grafana without manual addition.
- Add aws-cloudwatch and elasticsearch to supported backends. See [here](https://github.com/statsd/statsd/blob/master/docs/backend.md) for a comprehensive list of backends.
