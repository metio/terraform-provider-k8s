data "k8s_work_karmada_io_cluster_resource_binding_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    resource = {
      api_version = "v1"
      kind        = "Service"
      name        = "some-service"
    }
  }
}
