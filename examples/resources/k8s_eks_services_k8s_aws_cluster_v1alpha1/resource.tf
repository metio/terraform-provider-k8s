resource "k8s_eks_services_k8s_aws_cluster_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
