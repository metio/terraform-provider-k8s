output "manifests" {
  value = {
    "example" = data.k8s_externaldns_k8s_io_dns_endpoint_v1alpha1_manifest.example.yaml
  }
}
