data "k8s_work_karmada_io_resource_binding_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    resource = {
      api_version = "v1"
      kind        = "Service"
      name        = "some-service"
    }
  }
}
