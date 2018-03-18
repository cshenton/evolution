# OpenAI Evolution Strategies
[![Build Status](https://travis-ci.org/cshenton/openai-evolution.svg?branch=master)](https://travis-ci.org/cshenton/openai-evolution)
[![Coverage Status](https://coveralls.io/repos/github/cshenton/openai-evolution/badge.svg?branch=master)](https://coveralls.io/github/cshenton/openai-evolution?branch=master)

Replication of the OpenAI [Evolution Strategies](https://blog.openai.com/evolution-strategies/) paper in Golang.

Uses a slightly different, homogenous network architecture, rather than a redis
cluster. Each member of the cluster uses gossip to maintain a memberlist, then
directly communicates over gRPC (namespaced swarm).