resource "k8s_apps_clusternet_io_manifest_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  template = {}
}
