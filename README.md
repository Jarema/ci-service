# ci-service

Early development stage.

First start ci-executor, then ci-service to test it.
By default it expects ci-executor to be run on localhost. If you want to change it,
change the grpc.Dial address and recompile.

The "pipeline" is configured in test.yaml, so you can play with it.
