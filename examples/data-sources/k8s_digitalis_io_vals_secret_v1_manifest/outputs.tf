output "manifests" {
  value = {
    "example" = data.k8s_digitalis_io_vals_secret_v1_manifest.example.yaml
  }
}
