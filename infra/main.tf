terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.5"
    }
  }
}

provider "google" {
  credentials = file("../config/krouly-be18aada8f9a.json")
  project     = "krouly"
  region      = "us-central1"
  zone        = "us-central1-a"
}

resource "google_container_cluster" "primary" {
  name     = "krouly-gke-cluster"
  location = "us-central1"

  remove_default_node_pool = true

  node_pool {
    name       = "krouly-node-pool"
    node_count = 1

    node_config {
      machine_type = "e2-medium" // other: f1-micro
      disk_size_gb = 100
    }
  }
}

/*
TODO:

Logging and Monitoring: Integrate Google Cloud Operations 
(formerly Stackdriver) for logging and monitoring to ensure you
can troubleshoot effectively and maintain system health.

Networking: Define VPC and subnet settings if you need custom 
network configurations or isolation between different parts 
of your infrastructure.

Auto-scaling: Implement an auto-scaling policy for your node 
pool to automatically adjust the number of nodes based on load.

Persistent Storage: If your application requires state or 
data persistence, define persistent volume claims or use 
Google Cloud Storage.

Security: Set up network policies, and consider using 
Google-managed keys for encrypting Kubernetes secrets.
*/