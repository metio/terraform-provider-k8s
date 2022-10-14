output "resources" {
  value = {
    "minimal" = k8s_iot_eclipse_org_ditto_v1alpha1.minimal.yaml
    "example" = k8s_iot_eclipse_org_ditto_v1alpha1.example.yaml
  }
}
