data "k8s_chaos_mesh_org_aws_chaos_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    action       = "ec2-restart"
    aws_region   = "eu-central1"
    ec2_instance = "t3.xlarge"
  }
}
