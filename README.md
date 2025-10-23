# SRE Take-Home Assignment - Monorepo CFX

Repository ini dibuat sebagai solusi untuk SRE Take-Home Assignment yang diberikan oleh interviewer. Project ini mendemonstrasikan implementasi CI/CD pipeline untuk monorepo yang berisi dua backend services (Go dan Node.js) yang di-deploy ke Kubernetes cluster.

## 📋 Assignment Overview

### Requirements
- Monorepo dengan dua backend services:
  - Service 1: Go (REST API)
  - Service 2: Node.js (Express.js)
- Infrastructure: Kubernetes
- CI/CD Pipeline automation
- Public service exposure via DNS

## 🏗️ Project Structure

```
monorepo-cfx/
├── services/
│   ├── go-service/              # Go application
│   │   ├── main.go
│   │   ├── go.mod
│   │   ├── Dockerfile
│   │   └── k8s/                 # Kubernetes manifests
│   │       ├── base/
│   │       └── overlays/
│   │           ├── development/
│   │           ├── staging/
│   │           └── production/
│   └── node-service/            # Node.js application
│       ├── index.js
│       ├── package.json
│       ├── Dockerfile
│       └── k8s/                 # Kubernetes manifests
│           ├── base/
│           └── overlays/
│               ├── development/
│               ├── staging/
│               └── production/
├── .github/
│   └── workflows/
│       └── ci.yml               # GitHub Actions CI/CD pipeline
├── monitoring/
│   └── grafana-alloy.yaml       # Monitoring configuration
└── README.md
```

## 🚀 Services Description

### Go Service (`/api/go`)
- **Framework**: Standard library `net/http`
- **Endpoints**:
  - `GET /` - Health check
  - `GET /health` - Detailed health check
  - `GET /api/users` - List users
  - `POST /api/users` - Create user
- **Port**: 8080

### Node.js Service (`/api/node`)
- **Framework**: Express.js
- **Endpoints**:
  - `GET /` - Health check
  - `GET /health` - Detailed health check
  - `GET /api/products` - List products
  - `POST /api/products` - Create product
- **Port**: 3000

## 🔄 CI/CD Pipeline Workflow

### GitHub Actions Pipeline (`.github/workflows/ci.yml`)

Pipeline berjalan pada setiap push/PR ke `main` branch:

1. **Build & Test Stage**:
   - Build Go application
   - Test Go service dengan coverage
   - Install dan test Node.js application
   - Build Node.js application

2. **Container Build & Push**:
   - Build Docker image untuk Go service
   - Build Docker image untuk Node.js service
   - Push images ke GitHub Container Registry (ghcr.io)

3. **Kubernetes Deployment**:
   - Generate Kubernetes manifests menggunakan Kustomize
   - Deploy ke Kubernetes cluster
   - Validasi deployment status

### Environment Setup
- **Development**: `dev.domain.com` (menggunakan xip.io)
- **Staging**: `staging.domain.com` (menggunakan xip.io)
- **Production**: `prod.domain.com`

## 🛠️ Technology Stack

### Application Development
- **Go**: v1.21+ (REST API dengan standard library)
- **Node.js**: v18+ (Express.js framework)
- **Containerization**: Docker + Multi-stage builds

### CI/CD Pipeline
- **CI/CD Platform**: GitHub Actions
- **Container Registry**: GitHub Container Registry (ghcr.io)
- **Infrastructure as Code**: Kubernetes YAML + Kustomize
- **Secret Management**: Kubernetes Secrets

### Infrastructure
- **Container Orchestration**: Kubernetes (v1.28+)
- **Ingress Controller**: NGINX Ingress Controller
- **Load Balancer**: Kubernetes Service + Ingress
- **DNS**: xip.io untuk development/testing

### Monitoring (Bonus)
- **Monitoring Stack**: Grafana Alloy (Prometheus + Grafana)
- **Metrics**: Application metrics + Infrastructure metrics
- **Logging**: Structured logging dengan correlation IDs

## 🚀 How to Run/Test

### Prerequisites
```bash
# Install required tools
kubectl install
docker install
helm install (optional for ingress controller)
```

### Local Development
```bash
# Clone repository
git clone https://github.com/username/monorepo-cfx.git
cd monorepo-cfx

# Run Go service
cd services/go-service
go run main.go

# Run Node.js service (separate terminal)
cd services/node-service
npm install
npm start
```

### Kubernetes Deployment

#### 1. Setup Kubernetes Cluster
```bash
# Untuk local development
minikube start
atau
kind create cluster
atau
docker desktop Kubernetes
```

#### 2. Install Ingress Controller
```bash
# Install NGINX Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```

#### 3. Deploy Services
```bash
# Deploy ke development environment
kubectl apply -k services/go-service/k8s/overlays/development/
kubectl apply -k services/node-service/k8s/overlays/development/

# Deploy ke staging environment
kubectl apply -k services/go-service/k8s/overlays/staging/
kubectl apply -k services/node-service/k8s/overlays/staging/

# Deploy ke production environment
kubectl apply -k services/go-service/k8s/overlays/production/
kubectl apply -k services/node-service/k8s/overlays/production/
```

#### 4. Verify Deployment
```bash
# Check deployment status
kubectl get pods -n monorepo-cfx
kubectl get services -n monorepo-cfx
kubectl get ingress -n monorepo-cfx

# Test services
kubectl port-forward svc/go-service 8080:80 -n monorepo-cfx
kubectl port-forward svc/node-service 3000:80 -n monorepo-cfx

# Test via curl
curl http://localhost:8080/
curl http://localhost:3000/
```

### Testing with Public DNS
```bash
# Gunakan xip.io untuk public access
# Update ingress hostname di deployment manifests
# Contoh: go-service.192.168.49.2.nip.io

# Access services via browser
http://go-service.192.168.49.2.nip.io/
http://node-service.192.168.49.2.nip.io/
```

## 📊 Monitoring Setup (Bonus)

### Grafana Alloy Configuration
```bash
# Apply monitoring configuration
kubectl apply -f monitoring/grafana-alloy.yaml

# Access Grafana UI
kubectl port-forward svc/grafana 3000:3000 -n monitoring
```

### Metrics Collection
- **Go Service**: Prometheus metrics via `/metrics` endpoint
- **Node.js Service**: Custom metrics middleware
- **Infrastructure**: Node exporter + kube-state-metrics

## 🔧 Configuration

### Environment Variables
```yaml
# Go Service Environment
- name: PORT
  value: "8080"
- name: LOG_LEVEL
  value: "info"
- name: ENVIRONMENT
  valueFrom:
    fieldRef:
      fieldPath: metadata.namespace

# Node.js Service Environment
- name: PORT
  value: "3000"
- name: NODE_ENV
  valueFrom:
    fieldRef:
      fieldPath: metadata.namespace
```

### Kubernetes Resources
- **Go Service**:
  - CPU request: 100m, limit: 500m
  - Memory request: 128Mi, limit: 512Mi
- **Node.js Service**:
  - CPU request: 100m, limit: 500m
  - Memory request: 128Mi, limit: 512Mi

## 🐛 Troubleshooting

### Common Issues
1. **Pod stuck in Pending**: Check resource requests vs cluster capacity
2. **Image pull errors**: Verify registry credentials and image names
3. **Ingress not working**: Check Ingress Controller status and DNS configuration

### Debug Commands
```bash
# Check pod logs
kubectl logs -f deployment/go-service -n monorepo-cfx
kubectl logs -f deployment/node-service -n monorepo-cfx

# Debug pods
kubectl exec -it deployment/go-service -n monorepo-cfx -- /bin/sh
kubectl describe pod <pod-name> -n monorepo-cfx

# Check events
kubectl get events -n monorepo-cfx --sort-by='.lastTimestamp'
```

## 📈 Performance Considerations

### Optimizations Implemented
- **Multi-stage Docker builds** untuk smaller image sizes
- **Kustomize overlays** untuk environment-specific configurations
- **Resource limits** untuk prevent resource starvation
- **Health checks** untuk proper load balancing
- **Graceful shutdown** untuk zero-downtime deployments

### Scalability
- **Horizontal Pod Autoscaling** dapat ditambahkan
- **Multiple replicas** untuk high availability
- **Separate namespaces** untuk environment isolation

## 🔒 Security Considerations

- **Non-root containers** untuk security
- **Read-only filesystem** dimana memungkinkan
- **Secret management** dengan Kubernetes Secrets
- **Network policies** untuk traffic isolation (optional)
- **Image scanning** dapat diintegrasikan di CI pipeline

## 🚀 Future Enhancements

1. **Advanced Monitoring**: APM integration (New Relic/DataDog)
2. **Security Scanning**: Container image vulnerability scanning
3. **Automated Testing**: Integration + E2E tests
4. **GitOps**: ArgoCD/Flux untuk continuous deployment
5. **Service Mesh**: Istio/Linkerd untuk advanced traffic management
6. **Auto-scaling**: HPA + VPA untuk automatic scaling

## 📞 Support

Untuk pertanyaan atau diskusi lebih lanjut tentang implementasi ini:
- **GitHub Issues**: [Create Issue](https://github.com/username/monorepo-cfx/issues)
- **Documentation**: Check inline code comments
- **Deployment Guide**: Follow step-by-step instructions above

---

**Note**: Repository ini dibuat khusus untuk memenuhi requirements SRE assignment dan mendemonstrasikan best practices dalam CI/CD, Kubernetes deployment, dan microservices architecture.