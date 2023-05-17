Linting is proposed to be done using spectral It is a
[CLI](https://github.com/stoplightio/spectral) that originates from Spotlight.

The idea is to create a shell script and a GitHub action.

Usage of the spectral linter is as follows:

```bash
spectral lint myapifile.yaml
```

To have custom rulesets do the following:

```bash
spectral lint myapifile.yaml --ruleset myruleset.yaml
```
