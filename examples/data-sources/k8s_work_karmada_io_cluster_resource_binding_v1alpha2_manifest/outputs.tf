output "manifests" {
  value = {
    "example" = data.k8s_work_karmada_io_cluster_resource_binding_v1alpha2_manifest.example.yaml
  }
}
