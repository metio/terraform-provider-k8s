output "manifests" {
  value = {
    "example" = data.k8s_certman_managed_openshift_io_certificate_request_v1alpha1_manifest.example.yaml
  }
}
