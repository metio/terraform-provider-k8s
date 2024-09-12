data "k8s_core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
