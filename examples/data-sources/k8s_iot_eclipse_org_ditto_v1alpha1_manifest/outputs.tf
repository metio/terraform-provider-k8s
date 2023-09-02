output "manifests" {
  value = {
    "example" = data.k8s_iot_eclipse_org_ditto_v1alpha1_manifest.example.yaml
  }
}
