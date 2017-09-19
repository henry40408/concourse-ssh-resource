# SSH Shell Resource [![CircleCI](https://circleci.com/gh/henry40408/ssh-shell-resource.svg?style=shield)](https://circleci.com/gh/henry40408/ssh-shell-resource) [![Docker Repository on Quay](https://quay.io/repository/henry40408/ssh-shell-resource/status "Docker Repository on Quay")](https://quay.io/repository/henry40408/ssh-shell-resource) [![GitHub release](https://img.shields.io/github/release/henry40408/ssh-shell-resource.svg)](https://github.com/henry40408/ssh-shell-resource) [![license](https://img.shields.io/github/license/henry40408/ssh-shell-resource.svg)](https://github.com/henry40408/ssh-shell-resource)

> SSH shell resource for Concourse CI

## Source Configuration

- `host`: host name of target machine
- `port`: port of SSH server on target machine, `22` by default
- `user`: user for executing shell script on target machine
- `password`: plain password for user on target machine
- `private_key`: private SSH key for user on target machine

According to [appleboy/easyssh-proxy](https://github.com/appleboy/easyssh-proxy/blob/b777a323265704a7015f3526c3fe31b4f0daa722/easyssh.go#L69-L105), if `password` and `private_key` both exist, `password` would be used first, then `private_key`.

## Behavior

This is a `put`-only resource, so `check` and `in` do nothing.

### `out`: Run commands via SSH

Execute shell script on target machine via SSH.

#### Parameters

- `script`: string, shell script to run on target machine

## Examples

```yaml
---
resource_types:
- name: ssh-shell
  type: docker-image
  source:
    repository: quay.io/henry40408/ssh-shell-resource

resources:
- name: staging-server
  type: ssh-shell
  source:
    host: 127.0.0.1
    user: root
    password: ((ssh_password))

jobs:
- name: echo
  plan:
  - put: staging-server
    params:
      script: |
        #!/bin/sh
        echo "Hello, world!"
```

## License

MIT
