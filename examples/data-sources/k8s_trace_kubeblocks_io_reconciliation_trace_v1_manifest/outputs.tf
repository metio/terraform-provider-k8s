output "manifests" {
  value = {
    "example" = data.k8s_trace_kubeblocks_io_reconciliation_trace_v1_manifest.example.yaml
  }
}
