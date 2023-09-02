output "manifests" {
  value = {
    "example" = data.k8s_sfn_services_k8s_aws_state_machine_v1alpha1_manifest.example.yaml
  }
}
