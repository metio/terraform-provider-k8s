output "manifests" {
  value = {
    "example" = data.k8s_upgrade_cattle_io_plan_v1_manifest.example.yaml
  }
}
