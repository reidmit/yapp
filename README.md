# yapp (pronounced [/ʒæp/](http://ipa-reader.xyz/?text=%CA%92%C3%A6p&voice=Brian))

A framework for apps written entirely in YAML, powered by [ytt](https://github.com/vmware-tanzu/carvel-ytt).

Highly experimental! This is a ridiculous idea! Do not use!

```sh
# build & start the server...
make build
./yapp examples/hello-world/yapp.yml
# and then...
curl localhost:7000/hello -d '{"name": "reid"}'
```

## how it works

You configure your app's routes in a `yapp.yml` file. You can use ytt to template this file!

Details about incoming HTTP requests will be passed to the template as [data values](https://carvel.dev/ytt/docs/latest/how-to-use-data-values/) under `data.values.request`.

These are the values that are currently provided to your template:

- `data.values.request.body`: request body, parsed as YAML
- `data.values.request.headers`: request headers (`map[string]string[]`)
- `data.values.request.query`: request query parameters (`map[string]string[]`)

## examples

A simple yapp is just a YAML file like this:

```yaml
routes:
  GET /hello:
    status: 200
    body:
      message: "hello!"
```

Run with `yapp` and try `curl localhost:7000/hello`:

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
      name: #@ data.values.request.body.name
```

Run with `yapp` and try `curl localhost:7000/hello -d '{"name": "reid"}'`:

```sh
$ curl localhost:7000/hello -d '{"name": "reid"}'
message: hello!
name: reid
```

See the [`examples`](/examples) directory for more!
