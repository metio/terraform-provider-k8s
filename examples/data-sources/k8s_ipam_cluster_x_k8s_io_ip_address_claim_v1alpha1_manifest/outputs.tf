output "manifests" {
  value = {
    "example" = data.k8s_ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1_manifest.example.yaml
  }
}
