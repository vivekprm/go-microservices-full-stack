# Introduction
Todo app using below microservices:
- Todo Service
- User Service
  - Login
  - Registration

# Functionality
Anonymous users can see the todos created but can't modify or delete.
Authenticated Users can modify and delete todos created by them only.

# Deployment
Setup CI/CD

Use nginx or envoy as API gateway and to host static resources.

Deploy all these using docker compose
Deploy in kubernetes env.