data "k8s_data_fluid_io_jindo_runtime_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}