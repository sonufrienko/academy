# Academy

Academy - demo project.

Built with:

- AWS
  - EKS
  - NLB
- Kubernetes (1.18)
  - NGINX Ingress Controller
- NodeJS (v14)
  - Express
- Python (3.8)
  - FastAPI
- Go (1.16)
  - Fiber

## Application Structure

```
.
├── api               # All our microservices
│   ├── account       # Account microservice with NodeJS
│   ├── comments      # Comments microservice with Go
│   └── courses       # Courses microservice with Python
├── kubernetes        # Kubernetes deployments, services, etc.
├── web
├── LICENSE
└── README.md
```

## Getting Started

## :rocket: Deploy to AWS EKS

1. Setup AWS EKS with [eksctl](https://github.com/weaveworks/eksctl)

   ```shell
   eksctl create cluster -f kubernetes/cluster.yaml
   ```

2. Setup AWS NLB to expose the NGINX Ingress controller

   ```shell
   kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.44.0/deploy/static/provider/aws/deploy.yaml
   ```

   Fix https issue.

   ```
   kubectl edit svc ingress-nginx-controller -n ingress-nginx
   ```

   Replace `targetPort: https` -> `targetPort: http`

3. Configure AWS

   1. Create Hosted Zone on AWS Route53

      Example: `academy.onufrienko.com`

   2. Create Certificate on AWS ACM

      Example: `onufrienko.com` and `*.onufrienko.com`

   3. Configure AWS NLB to use TLS

      Console > EC2 > Load Balancers > your NLB > Listeners > Edit (TCP: 443)

      > Protocol: TLS

      > Port: 443

      > Default SSL certificate: From ACM + select your certificate

4. Deploy

   ```shell
   cd kubernetes
   kubectl apply -f ingress.yaml
   kubectl apply -f account-deployment.yaml
   kubectl apply -f account-service.yaml
   kubectl apply -f comments-deployment.yaml
   kubectl apply -f comments-service.yaml
   kubectl apply -f courses-deployment.yaml
   kubectl apply -f courses-service.yaml
   ```
