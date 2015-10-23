# yj

Simple tool to read YAML from STDIN and print JSON to STDOUT.

## Synopsis

```sh
$ echo "hello: world" | yj
{
  "hello": "world"
}
```

## Why?

Sometimes I have config files in yaml and I want to access them easily via some bash script. `yj` in conjunction with [`jq`](https://stedolan.github.io/jq/) allows me to do this. `yj` is written in golang which allow me to easily distribute/download static binaries instead of fussing with trying to install ruby/python/perl dependencies.


```sh
$ cat <<EOM > config.yml
key: val
# comment
setting: stuff
list:
  - abc
  - def
map:
  a: 1
  b: 2
EOM

$ yj < config.yml | jq -r .setting
stuff

$ yj < config.yml | jq -r .map.b
2

$ yj < config.yml | jq -r .list[1]
def
```