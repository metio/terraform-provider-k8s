output "manifests" {
  value = {
    "example" = data.k8s_networking_k8s_io_ingress_class_v1_manifest.example.yaml
  }
}
