output "manifests" {
  value = {
    "example" = data.k8s_flows_netobserv_io_flow_collector_v1beta1_manifest.example.yaml
  }
}
