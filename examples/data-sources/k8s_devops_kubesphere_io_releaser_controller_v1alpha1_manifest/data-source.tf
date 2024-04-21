data "k8s_devops_kubesphere_io_releaser_controller_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
