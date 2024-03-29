output "manifests" {
  value = {
    "example" = data.k8s_sharedresource_openshift_io_shared_secret_v1alpha1_manifest.example.yaml
  }
}
