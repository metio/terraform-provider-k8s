resource "k8s_batch_job_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_batch_job_v1" "example" {
  metadata = {
    name = "test"
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
