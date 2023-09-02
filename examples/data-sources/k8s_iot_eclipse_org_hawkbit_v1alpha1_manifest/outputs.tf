output "manifests" {
  value = {
    "example" = data.k8s_iot_eclipse_org_hawkbit_v1alpha1_manifest.example.yaml
  }
}
