# OpenAI Evolution Strategies
[![Build Status](https://travis-ci.org/cshenton/evolution.svg?branch=master)](https://travis-ci.org/cshenton/evolution)
[![Coverage Status](https://coveralls.io/repos/github/cshenton/evolution/badge.svg?branch=master)](https://coveralls.io/github/cshenton/evolution?branch=master)

Replication of the OpenAI [Evolution Strategies](https://blog.openai.com/evolution-strategies/) paper in Golang.


## Intra-cluster communication

Uses a slightly different, homogenous network architecture, rather than a redis
cluster. Each member of the cluster uses gossip to maintain a memberlist, then
directly communicates over gRPC.

Protocol buffers are generated from `/protos` into their respective packages so
from this folder:

```bash
protoc -I protos/ protos/server.proto --go_out=plugins=grpc:server/pb
protoc -I protos/ protos/mlp.proto --go_out=agent/mlp/pb
# etc... protoc -I protos/ protos/<agent_name>.proto --go_out=agent/<agent_name>/pb
```
