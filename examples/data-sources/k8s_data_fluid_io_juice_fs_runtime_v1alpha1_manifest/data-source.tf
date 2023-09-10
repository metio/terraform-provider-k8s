data "k8s_data_fluid_io_juice_fs_runtime_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
