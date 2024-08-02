output "manifests" {
  value = {
    "example" = data.k8s_networking_gke_io_health_check_policy_v1_manifest.example.yaml
  }
}
