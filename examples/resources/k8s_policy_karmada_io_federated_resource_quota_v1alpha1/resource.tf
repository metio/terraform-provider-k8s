resource "k8s_policy_karmada_io_federated_resource_quota_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
