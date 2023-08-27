data "k8s_apps_daemon_set_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
