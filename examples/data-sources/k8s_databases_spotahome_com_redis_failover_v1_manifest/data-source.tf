data "k8s_databases_spotahome_com_redis_failover_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
