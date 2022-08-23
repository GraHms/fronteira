# Fronteira Microgateway

The Fronteira Microgateway is a lightweight, cloud-native, open-source, and developer-centric API management proxy that secures and routes requests and responses between API consumers and API providers.

## Usage
Here's an example of a simple Fronteira Policy
```yaml
apiVersion: v1
kind: FronteiraPolicy
metadata:
  name: myapi-policy
  labels:
    app: myapi
spec:
  target: https://www.myapi.com
  operations:
   - method: POST
     path: /api/v1/resources/
     scopes: ["user:create", "admin:create"]

   - method: GET
     path: /api/v1/resources/
     scopes: ["admin:read"]

   - method: GET
     path: /api/v1/resoruces/:id
     scopes: [ "user:read", "admin:read" ]



