output "manifests" {
  value = {
    "example" = data.k8s_core_strimzi_io_strimzi_pod_set_v1beta2_manifest.example.yaml
  }
}
