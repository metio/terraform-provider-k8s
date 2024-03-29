output "manifests" {
  value = {
    "example" = data.k8s_machine_openshift_io_machine_set_v1beta1_manifest.example.yaml
  }
}
