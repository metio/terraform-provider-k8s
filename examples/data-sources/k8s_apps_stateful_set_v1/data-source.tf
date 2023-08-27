data "k8s_apps_stateful_set_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
