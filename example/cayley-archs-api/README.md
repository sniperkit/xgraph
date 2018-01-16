# DOC API

## CLI

```
cd cmd/cli
go build
./cli
```

## Web

```
cd cmd/web
go build
./server
```

## Setup
```
mkdir -p $HOME/.config/doc-api
cp cmd/cli/config.json $HOME/.config/doc-api/
sudo mkdir /var/lib/doc-api
sudo chown $(whoami) /var/lib/doc-api
```

## Documenting Architecture Decisions

We keep architecture decisions in doc/adr and follow the ideas behind [this](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions) blog post. The tool we use is [adr-tools](https://github.com/npryce/adr-tools).
