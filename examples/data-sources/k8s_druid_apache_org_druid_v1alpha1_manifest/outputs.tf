output "manifests" {
  value = {
    "example" = data.k8s_druid_apache_org_druid_v1alpha1_manifest.example.yaml
  }
}
