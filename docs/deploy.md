
# Deploying functions

To deploy one or more functions all you need to do is run `apex deploy`. Apex deploys are idempotent; a build is created
for each function, and Apex and performs a checksum to see if the deployed function matches the local build, if so
it's not deployed.

If you prefer to be explicit you can pass one or more function names to `apex deploy`.

## Examples

Deploy all functions in the current directory:

```
$ apex deploy
```

Deploy all functions in the directory "~/dev/myapp":

```
$ apex deploy -C ~/dev/myapp
```

Deploy specific functions:

```
$ apex deploy auth
$ apex deploy auth api
```