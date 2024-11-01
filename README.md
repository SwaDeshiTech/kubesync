
# kube-syncer

**kube-syncer** is a powerful Golang-based tool designed to streamline the synchronization of Kubernetes ConfigMaps between different namespaces. It offers a flexible and reliable solution for maintaining consistent configuration across multiple environments, ensuring seamless configuration management.

# Tools and Libraries
<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=kubernetes,go,mongo,azure" />
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

**Run with custom config**

To run the kubesync binary with a custom configuration file using the -configFolder argument, you'll need to modify the last step of the instructions:

1. **Clone the config file from below path**
	```bash
	cp kubesync/conf/config.yml <custom_config_path>
	cp kubesync/conf/syncer.yml <custom_config_path>
	```
2. **Run application using custom path**
	```bash
	./kubesync -configFolder "<custom_config_path>"
	for example
	./kubesync -configFolder "/root/custom_config_folder"
	```

	**Explanation:**

	* **configFolder:** This flag tells the kubesync binary to look for the config.yml file in the specified directory.
	* **<custom_config_path>** Replace this with the actual path to the folder containing your custom config.yml and syncer.yml file.

## Configuration File Description

This configuration file defines settings for a Kubernetes synchronization application. Here's a breakdown of the parameters:

### General Settings:

-   **port:** Specifies the port number (8443) on which the application will listen for incoming requests.
-   **disableCronJob:** A boolean flag indicating whether to disable the scheduled job that periodically synchronizes resources.

### Kubernetes Configuration:

-   **kubeConfigPath:** The path to the Kubernetes configuration file (kubeconfig) that contains authentication and authorization information for accessing the cluster.

### MongoDB Configuration:

-   **uri:** The MongoDB connection URI, specifying the hostname (localhost) and port (27017) of the MongoDB server.
-   **username:** The username for authenticating to the MongoDB database.
-   **password:** The password for authenticating to the MongoDB database.
-   **maxPoolSize:** The maximum number of connections in the MongoDB connection pool.
-   **minPoolSize:** The minimum number of idle connections in the MongoDB connection pool.
-   **maxIdleTime:** The maximum idle time (in seconds) for a connection in the pool before it's closed.
-   **maxConnIdleTime:** The maximum time (in seconds) a connection can remain idle in the pool before being removed.
-   **connectTimeout:** The maximum time (in seconds) to wait for a connection to be established to the MongoDB server.

## Configuration File Description: Syncer Configuration

This configuration file defines a specific/multiple synchronization task. Here's a breakdown of the parameters:

### Syncer Configuration:

-   **name:** A unique identifier for this synchronization task.
-   **sourceNamespace:** The Kubernetes namespace from which resources will be synchronized.
-   **destinationNamespace:** A list of Kubernetes namespaces to which resources will be synchronized.
-   **configMapList:** A list of ConfigMaps to be synchronized from the source namespace to the destination namespaces.
-   **secretList:** A list of Secrets to be synchronized from the source namespace to the destination namespaces.
-   **k8sClusterName:** The name of the Kubernetes cluster to which the source & destination namespaces belong.
