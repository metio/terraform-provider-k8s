output "resources" {
  value = {
    "minimal" = k8s_camel_apache_org_kamelet_binding_v1alpha1.minimal.yaml
    "example" = k8s_camel_apache_org_kamelet_binding_v1alpha1.example.yaml
  }
}
