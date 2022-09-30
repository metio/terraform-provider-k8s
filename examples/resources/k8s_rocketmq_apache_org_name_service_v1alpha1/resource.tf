resource "k8s_rocketmq_apache_org_name_service_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    dns_policy             = "some-policy"
    host_network           = true
    host_path              = "some-path"
    image_pull_policy      = "some-policy"
    name_service_image     = "some-image"
    resources              = {}
    size                   = 123
    storage_mode           = "some-mode"
    volume_claim_templates = []
  }
}
