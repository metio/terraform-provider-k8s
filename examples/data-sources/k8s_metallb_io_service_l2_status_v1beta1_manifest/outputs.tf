output "manifests" {
  value = {
    "example" = data.k8s_metallb_io_service_l2_status_v1beta1_manifest.example.yaml
  }
}
