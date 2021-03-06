apiVersion: v1
kind: ConfigMap
metadata:
  name: spire-server
  namespace: {{ .Release.Namespace }}
data:
  server.conf: |
    server {
      bind_address = "0.0.0.0"
      bind_port = "8081"
      trust_domain = "test.com"
      data_dir = "/run/spire/data"
      log_level = "DEBUG"
      svid_ttl = "1h"
      upstream_bundle = true
      ca_subject = {
        Country = ["US"],
        Organization = ["SPIFFE"],
        CommonName = "",
      }
    }
    plugins {
      DataStore "sql" {
        plugin_data {
          database_type = "sqlite3"
          connection_string = "/run/spire/data/datastore.sqlite3"
        }
      }
      NodeAttestor "k8s_sat" {
        plugin_data {
          clusters = {
            # NOTE: Change this to your cluster name
            "kubernetes" = {
              use_token_review_api_validation = true
              service_account_whitelist = ["{{ .Release.Namespace }}:spire-agent"]
            }
          }
        }
      }
      NodeResolver "noop" {
        plugin_data {}
      }
      KeyManager "disk" {
        plugin_data {
          keys_path = "/run/spire/data/keys.json"
        }
      }
      UpstreamCA "disk" {
        plugin_data {
          ttl = "12h"
          key_file_path = "/run/spire/secrets/bootstrap.key"
          cert_file_path = "/run/spire/config/bootstrap.crt"
        }
      }
    }
  bootstrap.crt: |
    -----BEGIN CERTIFICATE-----
    MIIBzDCCAVOgAwIBAgIJAJM4DhRH0vmuMAoGCCqGSM49BAMEMB4xCzAJBgNVBAYT
    AlVTMQ8wDQYDVQQKDAZTUElGRkUwHhcNMTgwNTEzMTkzMzQ3WhcNMjMwNTEyMTkz
    MzQ3WjAeMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1BJRkZFMHYwEAYHKoZIzj0C
    AQYFK4EEACIDYgAEWjB+nSGSxIYiznb84xu5WGDZj80nL7W1c3zf48Why0ma7Y7m
    CBKzfQkrgDguI4j0Z+0/tDH/r8gtOtLLrIpuMwWHoe4vbVBFte1vj6Xt6WeE8lXw
    cCvLs/mcmvPqVK9jo10wWzAdBgNVHQ4EFgQUh6XzV6LwNazA+GTEVOdu07o5yOgw
    DwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwGQYDVR0RBBIwEIYOc3Bp
    ZmZlOi8vbG9jYWwwCgYIKoZIzj0EAwQDZwAwZAIwE4Me13qMC9i6Fkx0h26y09QZ
    IbuRqA9puLg9AeeAAyo5tBzRl1YL0KNEp02VKSYJAjBdeJvqjJ9wW55OGj1JQwDF
    D7kWeEB6oMlwPbI/5hEY3azJi16I0uN1JSYTSWGSqWc=
    -----END CERTIFICATE-----
