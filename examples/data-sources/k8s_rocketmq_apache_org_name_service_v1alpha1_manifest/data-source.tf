data "k8s_rocketmq_apache_org_name_service_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
