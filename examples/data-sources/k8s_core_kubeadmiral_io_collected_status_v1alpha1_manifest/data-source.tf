data "k8s_core_kubeadmiral_io_collected_status_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
