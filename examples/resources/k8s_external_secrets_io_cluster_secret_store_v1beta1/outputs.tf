output "resources" {
  value = {
    "minimal" = k8s_external_secrets_io_cluster_secret_store_v1beta1.minimal.yaml
    "issue_110" = k8s_external_secrets_io_cluster_secret_store_v1beta1.issue_110.yaml
  }
}
