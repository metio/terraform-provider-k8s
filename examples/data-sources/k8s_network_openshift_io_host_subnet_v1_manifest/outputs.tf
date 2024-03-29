output "manifests" {
  value = {
    "example" = data.k8s_network_openshift_io_host_subnet_v1_manifest.example.yaml
  }
}
