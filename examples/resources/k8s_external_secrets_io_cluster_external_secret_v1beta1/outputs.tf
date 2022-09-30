output "resources" {
  value = {
    "minimal" = k8s_external_secrets_io_cluster_external_secret_v1beta1.minimal.yaml
  }
}
