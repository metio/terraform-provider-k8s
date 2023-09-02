output "manifests" {
  value = {
    "example" = data.k8s_networking_istio_io_destination_rule_v1beta1_manifest.example.yaml
  }
}
