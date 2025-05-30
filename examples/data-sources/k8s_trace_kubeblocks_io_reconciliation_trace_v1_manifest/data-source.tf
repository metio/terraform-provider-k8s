data "k8s_trace_kubeblocks_io_reconciliation_trace_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
