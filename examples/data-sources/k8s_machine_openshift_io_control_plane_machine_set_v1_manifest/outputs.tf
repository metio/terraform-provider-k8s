output "manifests" {
  value = {
    "example" = data.k8s_machine_openshift_io_control_plane_machine_set_v1_manifest.example.yaml
  }
}
