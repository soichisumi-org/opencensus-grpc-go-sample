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

![image](https://user-images.githubusercontent.com/30210641/93374377-96131200-f891-11ea-9765-08a162120dbf.png)

