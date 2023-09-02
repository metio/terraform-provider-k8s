output "manifests" {
  value = {
    "example" = data.k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest.example.yaml
  }
}
