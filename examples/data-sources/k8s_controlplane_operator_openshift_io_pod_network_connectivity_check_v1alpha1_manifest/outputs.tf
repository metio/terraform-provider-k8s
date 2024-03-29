output "manifests" {
  value = {
    "example" = data.k8s_controlplane_operator_openshift_io_pod_network_connectivity_check_v1alpha1_manifest.example.yaml
  }
}
