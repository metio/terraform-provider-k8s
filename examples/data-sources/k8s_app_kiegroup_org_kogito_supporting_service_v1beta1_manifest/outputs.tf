output "manifests" {
  value = {
    "example" = data.k8s_app_kiegroup_org_kogito_supporting_service_v1beta1_manifest.example.yaml
  }
}
