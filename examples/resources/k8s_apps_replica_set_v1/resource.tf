resource "k8s_apps_replica_set_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
