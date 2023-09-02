output "manifests" {
  value = {
    "example" = data.k8s_digitalis_io_db_secret_v1beta1_manifest.example.yaml
  }
}
