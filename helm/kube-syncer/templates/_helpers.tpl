{{- define "kube-syncer.namespace" -}}
{{- default .Release.Namespace .Values.namespaceOverride -}}
{{- end -}}

{{- define "kube-syncer.labels" -}}
app.kubernetes.io/name: kube-syncer
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: Helm
{{- end -}}

{{- define "kube-syncer.selectorLabels" -}}
app.kubernetes.io/name: kube-syncer
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "kube-syncer.fullname" -}}
{{- printf "%s-%s" .Release.Name "kube-syncer" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
