# parca-profiling
Continuous Profiling of Go Programs with parca; integration with pprof. 


# pprof-demo
This is a demo Go application that generates different workloads for profiling. It exposes [pprof](https://pkg.go.dev/net/http/pprof) endpoints on port 6060 and is fully compatible with [Parca](https://www.parca.dev) for continuous profiling.

---

## Features
- CPU-intensive workloads: Fibonacci, prime calculations, matrix multiplication, bubble sort, complex math.
- Memory-intensive workloads for heap profiling.
- Goroutine, mutex, and channel contention workloads.
- Pprof endpoints at `/debug/pprof/`.

---

## Requirements
- Kubernetes cluster with Istio installed.
- Parca (optional) for profiling visualization.
- Go >= 1.20 (for building locally if needed).

---

## Steps
1. Create a namespace (if not exists):
   
      _kubectl create namespace parca_

3. Apply Deployment and Service:
   
     _ kubectl apply -f deployment.yaml
      
      kubectl apply -f service.yaml_

4. Apply Istio Gateway and VirtualService:
   
     _ kubectl apply -f gateway.yaml

      kubectl apply -f virtualservice.yaml_

6. Access pprof and parca via the command line or browser (for visualization).

**Notes**
1. Replace pprof.example.com with your domain or localhost if testing locally.
2. All workloads run automatically in the background to generate meaningful profiling data.
3. While this demo uses Istio to expose pprof endpoints via a Gateway and VirtualService, you can use any ingress or reverse proxy (like NGINX Ingress, Traefik, or even `kubectl port-forward`) to access the endpoints. The key requirement is that port 6060 is reachable for profiling.



