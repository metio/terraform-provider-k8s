data "k8s_batch_job_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    template = {
      spec = {
        restart_policy = "Never"
        containers = [
          {
            name    = "pi"
            image   = "perl"
            command = ["perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"]
          }
        ]
      }
    }
    backoff_limit = 4
  }
}
