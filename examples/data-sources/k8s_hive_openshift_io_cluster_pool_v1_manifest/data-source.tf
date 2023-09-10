data "k8s_hive_openshift_io_cluster_pool_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
