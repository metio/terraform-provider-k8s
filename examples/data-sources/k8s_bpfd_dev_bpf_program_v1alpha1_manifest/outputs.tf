output "manifests" {
  value = {
    "example" = data.k8s_bpfd_dev_bpf_program_v1alpha1_manifest.example.yaml
  }
}
