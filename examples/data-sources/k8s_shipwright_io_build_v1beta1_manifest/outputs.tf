output "manifests" {
  value = {
    "example" = data.k8s_shipwright_io_build_v1beta1_manifest.example.yaml
  }
}
