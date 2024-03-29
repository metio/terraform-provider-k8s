data "k8s_core_kubeadmiral_io_federated_object_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
