build_bin:
	go mod download
	go build -i /bin/main

build_image:
	docker build -t scaffolder-sample .
	docker run -p 9000:8080 scaffolder-sample

invoke:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"event":"incoming message!"}'


