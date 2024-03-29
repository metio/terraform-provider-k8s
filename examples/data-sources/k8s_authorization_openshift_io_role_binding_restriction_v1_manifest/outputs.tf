output "manifests" {
  value = {
    "example" = data.k8s_authorization_openshift_io_role_binding_restriction_v1_manifest.example.yaml
  }
}
