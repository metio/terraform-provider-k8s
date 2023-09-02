output "manifests" {
  value = {
    "example" = data.k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest.example.yaml
  }
}
