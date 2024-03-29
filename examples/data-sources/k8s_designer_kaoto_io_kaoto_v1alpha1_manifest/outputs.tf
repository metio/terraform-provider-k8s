output "manifests" {
  value = {
    "example" = data.k8s_designer_kaoto_io_kaoto_v1alpha1_manifest.example.yaml
  }
}
