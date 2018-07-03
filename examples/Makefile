python-nats:
	kubeless function deploy python-nats --runtime python2.7 --handler pubsub.handler --from-file python/hellowithdata.py
	kubeless trigger nats create python-nats --function-selector created-by=kubeless,function=python-nats --trigger-topic test

python-nats-verify:
	$(eval DATA := $(shell mktemp -u -t XXXXXXXX))
	$(eval NODEPORT := $(shell kubectl get svc nats -n nats-io -o jsonpath="{.spec.ports[0].nodePort}"))
	$(eval MINIKUBE_IP := $(shell minikube ip))
	kubeless trigger nats publish --url nats://$(MINIKUBE_IP):$(NODEPORT) --topic test --message '{"payload":"$(DATA)"}'
	number="1"; \
	timeout="60"; \
	found=false; \
	while [ $$number -le $$timeout ] ; do \
		pod=`kubectl get po -oname -l function=python-nats`; \
		logs=`kubectl logs $$pod | grep $(DATA)`; \
		if [ "$$logs" != "" ]; then \
			found=true; \
			break; \
		fi; \
		sleep 1; \
		number=`expr $$number + 1`; \
	done; \
	$$found
	# Verify event context
	logs=`kubectl logs -l function=python-nats`; \
	echo $$logs | grep -q "event-time.*UTC" && \
	echo $$logs | grep -q "event-type.*application/json" && \
	echo $$logs | grep -q "event-namespace.*natstriggers.kubeless.io" && \
	echo $$logs | grep -q "event-id.*"

nats-python-func1-topic-test:
	kubeless function deploy nats-python-func1-topic-test --runtime python2.7 --handler pubsub.handler --from-file python/hellowithdata.py  --label topic=nats-topic-test

nats-python-func2-topic-test:
	kubeless function deploy nats-python-func2-topic-test --runtime python2.7 --handler pubsub.handler --from-file python/hellowithdata.py  --label topic=nats-topic-test

nats-python-func-multi-topic:
	kubeless function deploy nats-python-func-multi-topic --runtime python2.7 --handler pubsub.handler --from-file python/hellowithdata.py  --label func=nats-python-func-multi-topic

nats-python-trigger-topic-test:
	kubeless trigger nats create nats-python-trigger-topic-test --function-selector created-by=kubeless,topic=nats-topic-test --trigger-topic topic-test

nats-python-trigger-topic1:
	kubeless trigger nats create nats-python-trigger-topic1 --function-selector created-by=kubeless,func=nats-python-func-multi-topic --trigger-topic topic1

nats-python-trigger-topic2:
	kubeless trigger nats create nats-python-trigger-topic2 --function-selector created-by=kubeless,func=nats-python-func-multi-topic --trigger-topic topic2

nats-python-func1-topic-test-verify:
	$(eval DATA := $(shell mktemp -u -t XXXXXXXX))
	$(eval NODEPORT := $(shell kubectl get svc nats -n nats-io -o jsonpath="{.spec.ports[0].nodePort}"))
	$(eval MINIKUBE_IP := $(shell minikube ip))
	kubeless trigger nats publish --url nats://$(MINIKUBE_IP):$(NODEPORT) --topic topic-test --message '{"payload":"$(DATA)"}'
	number="1"; \
	timeout="60"; \
	found=false; \
	while [ $$number -le $$timeout ] ; do \
		pod=`kubectl get po -oname -l function=nats-python-func1-topic-test`; \
		logs=`kubectl logs $$pod | grep $(DATA)`; \
		if [ "$$logs" != "" ]; then \
			found=true; \
			break; \
		fi; \
		sleep 1; \
		number=`expr $$number + 1`; \
	done; \
	$$found
	# Verify event context
	logs=`kubectl logs -l function=nats-python-func1-topic-test`; \
	echo $$logs | grep -q "event-time.*UTC" && \
	echo $$logs | grep -q "event-type.*application/json" && \
	echo $$logs | grep -q "event-namespace.*natstriggers.kubeless.io" && \
	echo $$logs | grep -q "event-id.*"

nats-python-func2-topic-test-verify:
	$(eval DATA := $(shell mktemp -u -t XXXXXXXX))
	$(eval NODEPORT := $(shell kubectl get svc nats -n nats-io -o jsonpath="{.spec.ports[0].nodePort}"))
	$(eval MINIKUBE_IP := $(shell minikube ip))
	kubeless trigger nats publish --url nats://$(MINIKUBE_IP):$(NODEPORT) --topic topic-test --message '{"payload":"$(DATA)"}'
	number="1"; \
	timeout="60"; \
	found=false; \
	while [ $$number -le $$timeout ] ; do \
		pod=`kubectl get po -oname -l function=nats-python-func2-topic-test`; \
		logs=`kubectl logs $$pod | grep $(DATA)`; \
		if [ "$$logs" != "" ]; then \
			found=true; \
			break; \
		fi; \
		sleep 1; \
		number=`expr $$number + 1`; \
	done; \
	$$found
	# Verify event context
	logs=`kubectl logs -l function=nats-python-func2-topic-test`; \
	echo $$logs | grep -q "event-time.*UTC" && \
	echo $$logs | grep -q "event-type.*application/json" && \
	echo $$logs | grep -q "event-namespace.*natstriggers.kubeless.io" && \
	echo $$logs | grep -q "event-id.*"

nats-python-func-multi-topic-verify:
	$(eval DATA := $(shell mktemp -u -t XXXXXXXX))
	$(eval NODEPORT := $(shell kubectl get svc nats -n nats-io -o jsonpath="{.spec.ports[0].nodePort}"))
	$(eval MINIKUBE_IP := $(shell minikube ip))
	kubeless trigger nats publish --url nats://$(MINIKUBE_IP):$(NODEPORT) --topic topic1 --message '{"payload":"$(DATA)"}'
	number="1"; \
	timeout="60"; \
	found=false; \
	while [ $$number -le $$timeout ] ; do \
		pod=`kubectl get po -oname -l function=nats-python-func-multi-topic`; \
		logs=`kubectl logs $$pod | grep $(DATA)`; \
		if [ "$$logs" != "" ]; then \
			found=true; \
			break; \
		fi; \
		sleep 1; \
		number=`expr $$number + 1`; \
	done; \
	$$found
	# Verify event context
	logs=`kubectl logs -l function=nats-python-func-multi-topic`; \
	echo $$logs | grep -q "event-time.*UTC" && \
	echo $$logs | grep -q "event-type.*application/json" && \
	echo $$logs | grep -q "event-namespace.*natstriggers.kubeless.io" && \
	echo $$logs | grep -q "event-id.*"

	kubeless trigger nats publish --url nats://$(MINIKUBE_IP):$(NODEPORT) --topic topic2 --message '{"payload":"$(DATA)"}'
	number="1"; \
	timeout="60"; \
	found=false; \
	while [ $$number -le $$timeout ] ; do \
		pod=`kubectl get po -oname -l function=nats-python-func-multi-topic`; \
		logs=`kubectl logs $$pod | grep $(DATA)`; \
		if [ "$$logs" != "" ]; then \
			found=true; \
			break; \
		fi; \
		sleep 1; \
		number=`expr $$number + 1`; \
	done; \
	$$found
	# Verify event context
	logs=`kubectl logs -l function=nats-python-func-multi-topic`; \
	echo $$logs | grep -q "event-time.*UTC" && \
	echo $$logs | grep -q "event-type.*application/json" && \
	echo $$logs | grep -q "event-namespace.*natstriggers.kubeless.io" && \
	echo $$logs | grep -q "event-id.*"

