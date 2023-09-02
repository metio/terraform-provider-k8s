output "manifests" {
  value = {
    "example" = data.k8s_acid_zalan_do_operator_configuration_v1_manifest.example.yaml
  }
}
