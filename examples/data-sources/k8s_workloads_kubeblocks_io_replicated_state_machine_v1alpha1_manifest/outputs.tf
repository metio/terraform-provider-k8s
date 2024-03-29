output "manifests" {
  value = {
    "example" = data.k8s_workloads_kubeblocks_io_replicated_state_machine_v1alpha1_manifest.example.yaml
  }
}
