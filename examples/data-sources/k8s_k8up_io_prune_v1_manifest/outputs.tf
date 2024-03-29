output "manifests" {
  value = {
    "example" = data.k8s_k8up_io_prune_v1_manifest.example.yaml
  }
}
