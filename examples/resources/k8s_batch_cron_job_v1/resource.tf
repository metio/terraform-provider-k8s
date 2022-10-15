resource "k8s_batch_cron_job_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_batch_cron_job_v1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    concurrency_policy            = "Replace"
    failed_jobs_history_limit     = 5
    schedule                      = "1 0 * * *"
    starting_deadline_seconds     = 10
    successful_jobs_history_limit = 10

    job_template = {
      metadata = {}

      spec = {
        backoff_limit              = 2
        ttl_seconds_after_finished = 10

        template = {
          metadata = {}

          spec = {
            containers = [
              {
                name    = "hello"
                image   = "busybox"
                command = ["/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"]
              }
            ]
          }
        }
      }
    }
  }
}
