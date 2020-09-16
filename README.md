# Opencensus-gRPC-Go-Sample

A sample repo of OpenCensus trace instrumentation for gRPC-go server

# Installation

* [Evans](https://github.com/ktr0731/evans)
  * brew tap ktr0731/evans
  * brew install evans
  
# Run

0. Get Cloud Trace permission. Do A or B
    * A: [Download service account key and add to PATH](https://cloud.google.com/trace/docs/setup/go?hl=ja#running_locally_and_elsewhere)
    * B: `gcloud auth application-default login`
1. run servers
    * make GCP_PROJECT=xxxx run-child
    * make GCP_PROJECT=xxxx run-parent
2. load
    * make load
3. see result

![image](https://user-images.githubusercontent.com/30210641/93375436-1d14ba00-f893-11ea-80b2-510e1f67b5f8.png)


