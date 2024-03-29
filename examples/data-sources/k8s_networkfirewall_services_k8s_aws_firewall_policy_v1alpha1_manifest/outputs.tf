output "manifests" {
  value = {
    "example" = data.k8s_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest.example.yaml
  }
}
