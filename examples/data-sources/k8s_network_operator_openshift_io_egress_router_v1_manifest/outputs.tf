output "manifests" {
  value = {
    "example" = data.k8s_network_operator_openshift_io_egress_router_v1_manifest.example.yaml
  }
}
