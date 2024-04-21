output "manifests" {
  value = {
    "example" = data.k8s_bpfman_io_fentry_program_v1alpha1_manifest.example.yaml
  }
}
