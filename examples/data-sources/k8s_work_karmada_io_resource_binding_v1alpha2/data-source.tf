data "k8s_work_karmada_io_resource_binding_v1alpha2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}