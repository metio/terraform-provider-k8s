output "manifests" {
  value = {
    "example" = data.k8s_acid_zalan_do_postgresql_v1_manifest.example.yaml
  }
}
