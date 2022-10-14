resource "k8s_hive_openshift_io_cluster_pool_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    base_domain = "example.com"
    image_set_ref = {
      name = "some-image-set"
    }
    size     = 123
    platform = {}
  }
}
