data "k8s_work_karmada_io_resource_binding_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
