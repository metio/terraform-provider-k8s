output "manifests" {
  value = {
    "example" = data.k8s_camel_apache_org_kamelet_v1alpha1_manifest.example.yaml
  }
}
