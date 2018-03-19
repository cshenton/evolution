# OpenAI Evolution Strategies
[![Build Status](https://travis-ci.org/cshenton/evolution.svg?branch=master)](https://travis-ci.org/cshenton/evolution)
[![Coverage Status](https://coveralls.io/repos/github/cshenton/evolution/badge.svg?branch=master)](https://coveralls.io/github/cshenton/evolution?branch=master)

Replication of the OpenAI [Evolution Strategies](https://blog.openai.com/evolution-strategies/) paper in Golang.

## Layout

Some packages focus on a high level interface, to allow the distributed components to be reused across different
agents and environments.

- `agent` defines the `Agent` interface, which is a serialisable, randomly perturbable agent.
    - subpackages define concrete implementations, like a multi-layer perceptron.
- `environment` defines the `Environment` interface, which provides a method for evaluating the fitness of an agent.
    - subpackages define concrete implementations, like a simple MNIST evaluation.

The remaining packages use those interface types to implement a generic distributed evolution algorithm.

- `runner` defines a message passing algorithm for running multiple environments on a single machine.
- `server` implements an rpc server to send and listen for results from other members of the cluster.
- `learner` acts as a sink for data from the server and the runner, and is also responsible for dispatching updates to other members of the cluster.
- `protos` holds all the protocol buffers used for intra-cluster communication.


## Intra-cluster communication

I use a homogenous network architecture, rather than a redis
cluster. Each member of the cluster uses gossip to maintain a memberlist, then
directly communicates over gRPC.

Protocol buffers are generated from `/protos` into their respective packages so
from this folder:

```bash
protoc -I protos/ protos/server.proto --go_out=plugins=grpc:server/pb
protoc -I protos/ protos/mlp.proto --go_out=agent/mlp/pb
# etc... protoc -I protos/ protos/<agent_name>.proto --go_out=agent/<agent_name>/pb
```
