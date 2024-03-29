output "manifests" {
  value = {
    "example" = data.k8s_network_openshift_io_egress_network_policy_v1_manifest.example.yaml
  }
}
