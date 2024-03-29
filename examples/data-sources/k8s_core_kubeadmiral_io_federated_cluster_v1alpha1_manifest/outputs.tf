output "manifests" {
  value = {
    "example" = data.k8s_core_kubeadmiral_io_federated_cluster_v1alpha1_manifest.example.yaml
  }
}
