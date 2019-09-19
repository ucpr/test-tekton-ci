update-plugins:
	kubectl create configmap plugins --from-file=plugins.yaml=./plugins.yaml --dry-run -o yaml | kubectl replace configmap plugins -f -

update-config:
	kubectl create configmap config --from-file=config.yaml=./config.yaml --dry-run -o yaml | kubectl replace configmap config -f -
