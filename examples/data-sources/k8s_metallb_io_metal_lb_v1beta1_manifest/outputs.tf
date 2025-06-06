output "manifests" {
  value = {
    "example" = data.k8s_metallb_io_metal_lb_v1beta1_manifest.example.yaml
  }
}
