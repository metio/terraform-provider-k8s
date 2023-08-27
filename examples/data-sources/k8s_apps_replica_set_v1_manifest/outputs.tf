output "manifests" {
  value = {
    "example" = data.k8s_apps_replica_set_v1_manifest.example.yaml
  }
}
