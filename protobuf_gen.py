from subprocess import call
call(["protoc", "-I", "executor", "executor/executor.proto", "--go_out=plugins=grpc:executor"])
