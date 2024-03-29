output "manifests" {
  value = {
    "example" = data.k8s_integration_rock8s_com_plug_v1beta1_manifest.example.yaml
  }
}
