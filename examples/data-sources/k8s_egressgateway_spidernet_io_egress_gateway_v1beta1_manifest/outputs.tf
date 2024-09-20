output "manifests" {
  value = {
    "example" = data.k8s_egressgateway_spidernet_io_egress_gateway_v1beta1_manifest.example.yaml
  }
}
