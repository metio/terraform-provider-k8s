output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_mesh_fault_injection_v1alpha1_manifest.example.yaml
  }
}
