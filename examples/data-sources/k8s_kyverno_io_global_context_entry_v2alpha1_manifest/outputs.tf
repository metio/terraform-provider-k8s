output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_global_context_entry_v2alpha1_manifest.example.yaml
  }
}
