output "manifests" {
  value = {
    "example" = data.k8s_devops_kubesphere_io_releaser_controller_v1alpha1_manifest.example.yaml
  }
}
