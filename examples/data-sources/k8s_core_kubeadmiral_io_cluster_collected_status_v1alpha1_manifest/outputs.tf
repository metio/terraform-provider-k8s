output "manifests" {
  value = {
    "example" = data.k8s_core_kubeadmiral_io_cluster_collected_status_v1alpha1_manifest.example.yaml
  }
}
