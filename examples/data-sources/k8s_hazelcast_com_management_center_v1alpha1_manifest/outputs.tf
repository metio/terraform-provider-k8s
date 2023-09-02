output "manifests" {
  value = {
    "example" = data.k8s_hazelcast_com_management_center_v1alpha1_manifest.example.yaml
  }
}
