data "k8s_rocketmq_apache_org_broker_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    allow_restart          = true
    broker_image           = "some-image"
    host_path              = "some-path"
    image_pull_policy      = "some-policy"
    replica_per_group      = 5
    scale_pod_name         = "some-name"
    size                   = 123
    storage_mode           = "some-mode"
    volume_claim_templates = []
    volumes                = []
    resources              = {}
    env                    = []
  }
}
