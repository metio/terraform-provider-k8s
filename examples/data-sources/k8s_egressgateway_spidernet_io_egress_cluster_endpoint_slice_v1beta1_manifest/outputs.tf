output "manifests" {
  value = {
    "example" = data.k8s_egressgateway_spidernet_io_egress_cluster_endpoint_slice_v1beta1_manifest.example.yaml
  }
}
