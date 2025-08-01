output "manifests" {
  value = {
    "example" = data.k8s_hyperspike_io_valkey_v1_manifest.example.yaml
  }
}
