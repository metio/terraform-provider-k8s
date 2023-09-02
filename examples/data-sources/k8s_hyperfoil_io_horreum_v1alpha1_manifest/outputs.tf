output "manifests" {
  value = {
    "example" = data.k8s_hyperfoil_io_horreum_v1alpha1_manifest.example.yaml
  }
}
