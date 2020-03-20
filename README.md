# Casbin Svc

Rewrite casbin service by go-kit to provide http and grpc interface

# Configuration

`viper` is used to load configuration

| label | meaning | default |
| ------| ------- | ------- |
| `address.debug` | debug transport address | `:8080` |
| `address.http`| http transport address | `:8081` |
| `address.grpc`| grpc transport address | `:8080` |
| `model.default` | default model file | `./model/model.conf` |
