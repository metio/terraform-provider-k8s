resource "k8s_apps_clusternet_io_feed_inventory_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    feeds = []
  }
}
