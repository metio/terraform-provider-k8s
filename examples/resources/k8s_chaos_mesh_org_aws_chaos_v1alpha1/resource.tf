resource "k8s_chaos_mesh_org_aws_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    action       = "ec2-restart"
    aws_region   = "eu-central1"
    ec2_instance = "t3.xlarge"
  }
}
