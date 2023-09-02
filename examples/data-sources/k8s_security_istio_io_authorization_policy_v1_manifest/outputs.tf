output "manifests" {
  value = {
    "example" = data.k8s_security_istio_io_authorization_policy_v1_manifest.example.yaml
  }
}
