# IPFS-Lite

[![Go Reference](https://pkg.go.dev/badge/github.com/topport/magic.svg)](https://pkg.go.dev/github.com/topport/magic)
[![Go Report Card](https://goreportcard.com/badge/github.com/topport/magic)](https://goreportcard.com/report/github.com/topport/magic)
[![Actions Status](https://github.com/topport/magic/workflows/Go/badge.svg)](https://github.com/topport/magic/actions)
[![codecov](https://codecov.io/gh/datahop/ipfs-lite/branch/master/graph/badge.svg)](https://codecov.io/gh/datahop/ipfs-lite)

IPFS-Lite is an embeddable, lightweight IPFS peer. This fork started
from [ipfs-lite](https://github.com/hsanjuan/ipfs-lite).

It offers all the features of main [ipfs-lite](https://github.com/hsanjuan/ipfs-lite). For certain requirements of
datahop it adds some more features of full ipfs, such as config, repo, leveldb etc.

## Examples

### cli client
```
 go run ./examples/litepeer/litepeer.go
```

### mobile client
```
  go run ./examples/mobilepeer/mobilepeer.go
```

## Objectives

* [x] create cache repo as ipfs
* [x] have persistent config information (Id, keys, ports, bootstraps etc)
* [x] use leveldb as datastore to set up peer
* [x] generate gomobile binding for android
* [x] adding content
* [x] replicating content
* [x] remove content
* [x] remove respective replication info for removed content
* [x] bootstrap mobile client
* [x] cli for bootstrap peer


## Documentation

[Go pkg docs](https://pkg.go.dev/github.com/topport/magic)

## License

Copyright 2021 Datahop

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
License. You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "
AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific
language governing permissions and limitations under the License.

## Acknowledgment

This software is part of the NGI Pointer project "Incentivised Content Dissemination at the Network Edge" that has
received funding from the European Union???s Horizon 2020 research and innovation programme under grant agreement No
871528




go run main.go server --config ~/.magic.yml