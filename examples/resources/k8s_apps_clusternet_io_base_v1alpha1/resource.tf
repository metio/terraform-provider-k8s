resource "k8s_apps_clusternet_io_base_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    feeds = []
  }
}
