data "k8s_data_fluid_io_goose_fs_runtime_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
