# Concourse SSH Resource [![CircleCI](https://circleci.com/gh/henry40408/concourse-ssh-resource.svg?style=shield)](https://circleci.com/gh/henry40408/concourse-ssh-resource) [![goreportcard](https://goreportcard.com/badge/github.com/henry40408/concourse-ssh-resource)](https://goreportcard.com/report/github.com/henry40408/concourse-ssh-resource) [![Docker Repository on Quay](https://quay.io/repository/henry40408/concourse-ssh-resource/status "Docker Repository on Quay")](https://quay.io/repository/henry40408/concourse-ssh-resource) [![GitHub release](https://img.shields.io/github/release/henry40408/concourse-ssh-resource.svg)](https://github.com/henry40408/concourse-ssh-resource) [![license](https://img.shields.io/github/license/henry40408/concourse-ssh-resource.svg)](https://github.com/henry40408/concourse-ssh-resource)

> SSH resource for Concourse CI

## Source Configuration

- `host`: host name of remote machine
- `port`: port of SSH server on remote machine, `22` by default
- `user`: user for executing shell script on remote machine
- `password`: plain password for user on remote machine
- `private_key`: private SSH key for user on remote machine

According to [appleboy/easyssh-proxy](https://github.com/appleboy/easyssh-proxy/blob/b777a323265704a7015f3526c3fe31b4f0daa722/easyssh.go#L69-L105), if `password` and `private_key` both exist, `password` would be used first, then `private_key`.

## Behavior

This is a `put`-only resource, `check` and `in` does nothing.

### `out`: Run commands via SSH

Execute shell script on remote machine via SSH.

#### Parameters

- `interpreter`: string, path to interpreter on remote machine, e.g. `/usr/bin/python3`, `/bin/sh` by default
- `script`: string, shell script to run on remote machine

## Examples

```yaml
---
resource_types:
- name: ssh
  type: docker-image
  source:
    repository: quay.io/henry40408/concourse-ssh-resource

resources:
- name: staging-server
  type: ssh
  source:
    host: 127.0.0.1
    user: root
    password: ((ssh_password))

jobs:
- name: echo
  plan:
  - put: staging-server
    params:
      interpreter: /usr/bin/env python3
      script: |
        print("Hello, world!")
```

## License

MIT
