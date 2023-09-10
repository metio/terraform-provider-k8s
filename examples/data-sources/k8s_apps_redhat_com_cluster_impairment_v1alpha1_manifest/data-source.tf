data "k8s_apps_redhat_com_cluster_impairment_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    duration = 480
    egress = {
      bandwidth = 10000
      latency   = 50
      loss      = 0.02
    }
    ingress = {
      bandwidth = 10000
      latency   = 50
      loss      = 0.02
    }
    interfaces = ["ens2f0"]
    link_flapping = {
      down_time = 30
      enable    = true
      up_time   = 120
    }
    start_delay = 5
  }
}
