output "manifests" {
  value = {
    "example" = data.k8s_work_karmada_io_cluster_resource_binding_v1alpha1_manifest.example.yaml
  }
}
