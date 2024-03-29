output "manifests" {
  value = {
    "example" = data.k8s_forklift_konveyor_io_forklift_controller_v1beta1_manifest.example.yaml
  }
}
