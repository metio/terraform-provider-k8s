output "manifests" {
  value = {
    "example" = data.k8s_bpfman_io_cluster_bpf_application_v1alpha1_manifest.example.yaml
  }
}
