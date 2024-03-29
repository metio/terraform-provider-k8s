output "manifests" {
  value = {
    "example" = data.k8s_postgresql_cnpg_io_pooler_v1_manifest.example.yaml
  }
}
