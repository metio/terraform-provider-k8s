output "manifests" {
  value = {
    "example" = data.k8s_network_openshift_io_net_namespace_v1_manifest.example.yaml
  }
}
