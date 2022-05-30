### Tool to interpolate variables in text files

The tool evaluates expressions like _${VAR}_ and tries to replace them with the value of the corresponding environment variable _VAR_.
If none is found, the expression is omitted.

## Usage:
```
interpolator $source_file ($target_file)
```

If no target file is specified, an in-place edit will be performed.

## Build
The project is built using the Makefile:
 
 ```bash
 make
 ```
 
The image can be either pushed to the registry via

```bash
# execute `docker login` before
TAG=1.0.0 make docker-push
```

or in a CI pipeline with

```bash
docker build . -t interpolator:${TAG}
```

The project can be cleaned via ```make clean```.