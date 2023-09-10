data "k8s_cache_kubedl_io_cache_backend_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
