output "manifests" {
  value = {
    "example" = data.k8s_bpfman_io_bpf_program_v1alpha1_manifest.example.yaml
  }
}
