# yapp

a framework for apps written entirely in YAML, powered by [ytt](https://github.com/vmware-tanzu/carvel-ytt)

highly experimental! do not use!

```
# start the server...
go run . -f examples/hello-world/yapp.yml
# and then...
curl localhost:7000/hello -d '{"name": "reid"}'
```
