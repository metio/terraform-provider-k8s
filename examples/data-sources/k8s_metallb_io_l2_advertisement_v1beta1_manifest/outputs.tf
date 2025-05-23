output "manifests" {
  value = {
    "example" = data.k8s_metallb_io_l2_advertisement_v1beta1_manifest.example.yaml
  }
}
