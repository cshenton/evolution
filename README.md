# OpenAI Evolution Strategies
[![Build Status](https://travis-ci.org/cshenton/evolution.svg?branch=master)](https://travis-ci.org/cshenton/evolution)

Replication of the OpenAI [Evolution Strategies](https://blog.openai.com/evolution-strategies/) paper in Golang.


## Roadmap

- Build out [stand-alone API for black box optimisers](https://github.com/cshenton/opt)
- **First Experiment:** Atari RAM (MLP)
- Broaden Agent interface to accept nd-tensor input data
- CNN agent
- **Second Experiment:** Atari Pixels (CNN)
- Agent methods for continuous policies
- MuJoCo RPC server (like the one at [cshenton/atari-rpc](https://github.com/cshenton/atari-rpc))
- **Third Experiment:** MuJoCo control


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

I use a homogenous network architecture. Each member of the cluster uses gossip to maintain a memberlist, then
directly communicates over gRPC.

Protocol buffers are generated from `/protos` into their respective packages so from this folder:

```bash
protoc -I . server/proto/server.proto --go_out=plugins=grpc:.
protoc -I . agent/mlp/proto/mlp.proto --go_out=.
protoc -I ../atari-rpc/proto/ atari.proto --go_out=plugins=grpc:environment/atari/proto

# etc... protoc -I protos/ protos/<agent_name>.proto --go_out=agent/<agent_name>/pb
```
