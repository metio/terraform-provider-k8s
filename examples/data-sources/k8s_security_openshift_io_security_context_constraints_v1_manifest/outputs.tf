output "manifests" {
  value = {
    "example" = data.k8s_security_openshift_io_security_context_constraints_v1_manifest.example.yaml
  }
}
