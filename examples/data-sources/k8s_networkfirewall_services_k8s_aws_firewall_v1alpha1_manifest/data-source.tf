data "k8s_networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
