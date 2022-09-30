output "resources" {
  value = {
    "minimal" = k8s_cert_manager_io_cluster_issuer_v1.minimal.yaml
  }
}
