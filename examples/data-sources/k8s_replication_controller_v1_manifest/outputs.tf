output "manifests" {
  value = {
    "example" = data.k8s_replication_controller_v1_manifest.example.yaml
  }
}
