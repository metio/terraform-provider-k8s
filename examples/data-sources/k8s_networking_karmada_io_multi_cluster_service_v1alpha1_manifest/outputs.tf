output "manifests" {
  value = {
    "example" = data.k8s_networking_karmada_io_multi_cluster_service_v1alpha1_manifest.example.yaml
  }
}
