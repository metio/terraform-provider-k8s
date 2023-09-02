output "manifests" {
  value = {
    "example" = data.k8s_external_secrets_io_cluster_external_secret_v1beta1_manifest.example.yaml
  }
}
