data "k8s_boskos_k8s_io_resource_object_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
