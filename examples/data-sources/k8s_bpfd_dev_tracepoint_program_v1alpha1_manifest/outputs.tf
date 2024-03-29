output "manifests" {
  value = {
    "example" = data.k8s_bpfd_dev_tracepoint_program_v1alpha1_manifest.example.yaml
  }
}
