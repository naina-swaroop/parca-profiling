# Parca Profiling Demo
A sample Go application demonstrating continuous profiling with Go's `pprof` and Parca in Kubernetes.

## pprof-demo
This demo Go application generates different workloads for profiling. It exposes [pprof](https://pkg.go.dev/net/http/pprof) endpoints on port `6060` and is fully compatible with [Parca](https://www.parca.dev) for continuous profiling.
The application continuously generates CPU, memory, mutex, block, and goroutine workloads, making it ideal for exploring profiling data in Parca.

---

## Features
- CPU-intensive workloads: Fibonacci, prime calculations, matrix multiplication, bubble sort, complex math.
- Memory-intensive workloads for heap profiling.
- Goroutine, mutex, and channel contention workloads.
- Pprof endpoints at `/debug/pprof/`.

---
## Repository Structure

- `deployment.yaml` – Deploys the sample Go application.
- `service.yaml` – Exposes the application within the cluster.
- `ingress.yaml` – Configures the Istio Gateway and VirtualService.
- `parca-cm-static.yaml` – Example Parca ConfigMap for static scraping.
- `parca-cm-dynamic.yaml` – Example Parca ConfigMap for dynamic scraping.

---

## Requirements
- Kubernetes cluster with Istio installed.
- Parca (optional) for profiling visualization.
- Go >= 1.20 (for building locally if needed).

---

## Steps

1. Create the namespace (if it doesn't already exist):

```bash
kubectl create namespace parca
```

2. Deploy the sample application:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

3. Configure Istio:

```bash
kubectl apply -f ingress.yaml
```

4. Configure Parca:

- Use `parca-cm-static.yaml` for static scraping.
- Use `parca-cm-dynamic.yaml` for dynamic scraping.

5. Access the `pprof` endpoints or visualize the collected profiles in Parca.

**Notes**
1. Replace pprof.example.com with your domain or localhost if testing locally.
2. All workloads run automatically in the background to generate meaningful profiling data.
3. While this demo uses Istio to expose pprof endpoints via a Gateway and VirtualService, you can use any ingress or reverse proxy (like NGINX Ingress, Traefik, or even `kubectl port-forward`) to access the endpoints. The key requirement is that port 6060 is reachable for profiling.




