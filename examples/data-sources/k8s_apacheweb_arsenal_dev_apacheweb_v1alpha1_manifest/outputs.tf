output "manifests" {
  value = {
    "example" = data.k8s_apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest.example.yaml
  }
}
