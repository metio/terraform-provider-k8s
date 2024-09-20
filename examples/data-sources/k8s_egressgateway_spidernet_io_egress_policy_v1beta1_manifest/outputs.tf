output "manifests" {
  value = {
    "example" = data.k8s_egressgateway_spidernet_io_egress_policy_v1beta1_manifest.example.yaml
  }
}
