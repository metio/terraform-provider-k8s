output "manifests" {
  value = {
    "example" = data.k8s_core_kubeadmiral_io_cluster_federated_object_v1alpha1_manifest.example.yaml
  }
}
