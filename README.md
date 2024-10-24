
# kube-syncer

**kube-syncer** is a powerful Golang-based tool designed to streamline the synchronization of Kubernetes ConfigMaps between different namespaces. It offers a flexible and reliable solution for maintaining consistent configuration across multiple environments, ensuring seamless configuration management.

# Tools and Libraries
<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=kubernetes,go,mongo" />
  </a>
</p>

## Features

* **Namespace-Based Synchronization:** Seamlessly syncs ConfigMaps from a source namespace to a target namespace.
* **Configurable Synchronization Interval:** Adjust the frequency of synchronization to meet your specific needs.
* **Error Handling and Logging:** Robust error handling and logging capabilities provide transparency and troubleshooting ease.
* **Customizable Filters:** Apply filters to selectively synchronize specific ConfigMaps based on their names or labels.
* **Kubernetes API Client Integration:** Leverages the Kubernetes API client to interact with the cluster efficiently.

## Installation

**Prerequisites:**

- Go 1.22 or later
- A Kubernetes cluster
- MongoDB Server

**Instructions:**

1. **Clone the Repository:**
	  ```bash
	  git clone https://github.com/your-username/kube-syncer.git
	```
2. **Build**
	```bash
	go get -v
	```
3. **Generate Binary**
	```bash
	go build .
	```
4. **Create kube directory**
	```bash
	mkdir -p ~/.kube/
	``` 
5. **Create kubeconfig file and paste kubeconfig**
	```bash
	vi ~/.kube/config
	```
6. **Change binary permission**
	```bash
	chmod +x kubesync
	```
7. **Run application**
	```bash
	./kubesync
	```
