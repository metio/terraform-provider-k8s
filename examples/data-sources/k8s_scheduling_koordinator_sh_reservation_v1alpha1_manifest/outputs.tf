output "manifests" {
  value = {
    "example" = data.k8s_scheduling_koordinator_sh_reservation_v1alpha1_manifest.example.yaml
  }
}
