
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
4. **(Optional) Prepare kubeconfig**
	> This step is only required if you are running kube-syncer outside of a Kubernetes cluster and need to authenticate via a kubeconfig file.
	```bash
	mkdir -p ~/.kube/
	vi ~/.kube/config
	```
5. **Change binary permission**
	```bash
	chmod +x kubesync
	```
6. **Run application**
	```bash
	./kubesync
	```

**Run with custom config**

To run the kubesync binary with a custom configuration file using the -configFolder argument, you'll need to copy the `conf/config.yml` into the destination directory and point the binary at it:

1. **Copy the config file into a custom folder**
	```bash
	cp kubesync/conf/config.yml <custom_config_path>
	```
2. **Run the application using the custom path**
	```bash
	./kubesync -configFolder "<custom_config_path>"
	# for example
	./kubesync -configFolder "/root/custom_config_folder"
	```

### Docker Image

1. **Build the container**
	```bash
	docker build -t kubesync:latest .
	```
2. **Run the image**
	```bash
	docker run --rm -p 8443:8443 kubesync:latest
	```

The Dockerfile bundles the `conf/` and `profiles/` folders beside the binary, so the container automatically loads the default `conf/config.yml` at `/app/conf/config.yml`. Set `-configFolder` via `docker run --entrypoint ./kubesync kubesync:latest -configFolder /app/conf` if you need to override the bundled config at runtime.

	**Explanation:**

	* **configFolder:** This flag tells the kubesync binary to look for the config.yml file in the specified directory.
	* **<custom_config_path>** Replace this with the actual path to the folder containing your custom config.yml file.

## Configuration File Description

This configuration file defines settings for a Kubernetes synchronization application. Here's a breakdown of the parameters:

### General Settings:

-   **port:** Specifies the port number (8443) on which the application will listen for incoming requests.
-   **disableCronJob:** A boolean flag indicating whether to disable the scheduled job that periodically synchronizes resources.

### Kubernetes Configuration:

-   **useServiceAccount:** When set to `true`, kube-syncer will use the service account token of the pod where it is running (via `rest.InClusterConfig`) and register that client under `serviceAccountName`.
-   **serviceAccountName:** Logical name assigned to the in-cluster client in the watcher registry (default: `in-cluster`). Keep this name in sync with the `k8sClusterName` entries inside the `syncers` block in `conf/config.yml`.
-   **kubeConfigPath:** _(Optional)_, provide a directory containing kubeconfig files when running outside the cluster. Each file is loaded as a separate client and keyed by filename.
-   **cronSchedules:** A list of cron jobs that should run, including their expressions, job types, and the resources they act upon. All scheduling metadata now lives in `conf/config.yml` rather than a database.
-   **k8sClusters:** Describe the clusters that kube-syncer can interact with. Entries are consumed by the namespace watcher and can replace the earlier MongoDB-backed cluster catalog.


### Syncer Definitions

All syncer declarations now live inside `conf/config.yml` under the `syncers` array. Each entry describes how a watcher should move ConfigMaps and Secrets:

-   **name:** Unique identifier displayed in logs and namespace watchers.
-   **sourceNamespace:** The namespace where the ConfigMaps/Secrets originate.
-   **destinationNamespace:** A list of namespaces to which the resources are copied (the watcher only acts when the destination namespace appears in this list).
-   **configMapList / secretList:** Lists of resource names that should be synchronized.
-   **k8sClusterName:** Logical cluster name that matches one of the kube clients (in-cluster or kubeconfig-derived) so the watcher knows which Kubernetes API to use.
-   **skipNamespace:** An optional array of namespace names that should be ignored even if they appear in `destinationNamespace` or `destinationNamespace` is set to `*`. Namespaces listed here will never trigger a sync.

When `destinationNamespace` contains the wildcard `*`, the syncer will process events for every namespace except those specified in `skipNamespace`. Otherwise it only reacts when the namespace name explicitly matches one of the listed destinations.

## Helm Chart

The `helm/kube-syncer` chart packages the binary, the default `conf/` directory and the `profiles/` samples. You can install it locally with:

1. **Build the Docker image** (see above) and push it to your registry or adjust `values.yaml` to point at your image.
2. **Install the chart**
	```bash
	helm install kube-syncer ./helm/kube-syncer --set image.repository=<your-registry>/kubesync --set image.tag=<tag>
	```

If the image is stored in a private registry and requires authentication, create a Kubernetes Docker registry secret first:

```bash
kubectl create secret docker-registry kubesync-regcred \
  --docker-username=<username> \
  --docker-password=<token> \
  --docker-server=https://index.docker.io/v1/ \
  --namespace=<target-namespace>
```

Then pull the secret name into the chart:

```bash
helm install kube-syncer ./helm/kube-syncer \
  --set image.repository=<your-registry>/kubesync \
  --set image.tag=<tag> \
  --set imagePullSecrets={kubesync-regcred}
```

The chart creates a ConfigMap from `conf/config.yml`, mounts it at `/etc/kubesync/conf`, and sets the `CONFIG_FOLDER` environment variable so the application loads the config automatically. Customize `values.yaml` to adjust replica count or resource requests.
