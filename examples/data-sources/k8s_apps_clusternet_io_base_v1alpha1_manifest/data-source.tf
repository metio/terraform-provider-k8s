data "k8s_apps_clusternet_io_base_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    feeds = []
  }
}
