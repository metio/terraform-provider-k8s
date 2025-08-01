data "k8s_workspace_devfile_io_dev_workspace_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
