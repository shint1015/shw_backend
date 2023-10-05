function find_proto() {
  find -name '*.proto' -type f -exec basename {} \;
}

cd api

for file in $(find_proto); do
  echo "$file"
  protoc --go_out=../pkg/grpc --go_opt=paths=source_relative --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative "$file"
done