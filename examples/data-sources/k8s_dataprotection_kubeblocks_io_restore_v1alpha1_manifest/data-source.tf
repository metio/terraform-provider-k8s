data "k8s_dataprotection_kubeblocks_io_restore_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
