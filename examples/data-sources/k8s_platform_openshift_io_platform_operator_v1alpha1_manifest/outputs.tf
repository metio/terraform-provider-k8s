output "manifests" {
  value = {
    "example" = data.k8s_platform_openshift_io_platform_operator_v1alpha1_manifest.example.yaml
  }
}
