FROM public.ecr.aws/lambda/provided:al2 as build

# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct
# cache dependencies
ADD go.mod go.sum ./
ADD . .
RUN go mod download
# build
RUN go build -o /bin/main
RUN ls
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2

COPY --from=build /bin/main /main
ENTRYPOINT [ "/main" ]
