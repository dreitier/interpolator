# interpolator
Interpolate variables in text files.

## Description

*interpolator* evaluates expressions like `${VAR}` and tries to replace them with the value of the corresponding environment variable `VAR`.
If none is found, the expression is omitted.

## Getting started

### Usage
```bash
interpolator $source_file ($target_file)
```

If no target file is specified, an in-place edit will be performed.

### Docker container
You can find ready-to-run Docker containers at [dreitier/interpolator](https://hub.docker.com/repository/docker/dreitier/interpolator).

## Development
### Creating new releases
A new release (artifact & Docker container) is automatically created when a new Git tag is pushed:

```bash
git tag x.y.z
git push origin x.y.z
```

## Support
This software is provided as-is. You can open an issue in GitHub's issue tracker at any time. But we can't promise to get it fixed in the near future.
If you need professionally support, consulting or a dedicated feature, please get in contact with us through our [website](https://dreitier.com).

## Contribution
Feel free to provide a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.