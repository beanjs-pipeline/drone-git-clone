# drone-git-clone

## usage

```yaml
kind: pipeline
type: docker
name: default

steps:
  - name: clone <xxx>
    image: beanjs/drone-git-clone
    settings:
      remote: repository
      username: <xxx>
      password: <bbb>
      cache: ./cache
      branch: master
```
