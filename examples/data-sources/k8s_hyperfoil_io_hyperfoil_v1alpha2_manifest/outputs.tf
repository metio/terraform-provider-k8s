output "manifests" {
  value = {
    "example" = data.k8s_hyperfoil_io_hyperfoil_v1alpha2_manifest.example.yaml
  }
}
