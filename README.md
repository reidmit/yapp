# yapp

A framework for apps written entirely in YAML, powered by [ytt](https://github.com/vmware-tanzu/carvel-ytt).

Highly experimental! Do not use!

```sh
# build & start the server...
make build
./yapp run examples/hello-world/yapp.yml
# and then...
curl localhost:7000/hello -d '{"name": "reid"}'
```

## how it works

You configure your app's routes in a `yapp.yml` file. You can use ytt to template this file! And for POST requests, the request body will be parsed as YAML and passed to the template as [data values](https://carvel.dev/ytt/docs/latest/how-to-use-data-values/) under `data.values.request`.

## examples

A simple yapp is just a YAML file like this:

```yaml
routes:
  GET /hello:
    status: 200
    body:
      message: "hello!"
```

Run with `yapp run` and try `curl localhost:7000/hello`:

```sh
$ curl localhost:7000/hello
message: hello!
```

But that's just plain YAML. It gets really fun when you start using ytt!

```yaml
#@ load("@ytt:data", "data")
---
routes:
  POST /hello:
    status: 200
    body:
      message: "hello!"
      name: #@ data.values.request.name
```

Run with `yapp run` and try `curl localhost:7000/hello -d '{"name": "reid"}'`:

```sh
$ curl localhost:7000/hello -d '{"name": "reid"}'
message: hello!
name: reid
```

See `examples` directory for more!
