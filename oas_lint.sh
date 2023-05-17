#!/bin/bash

file="$1"

vacuum lint -r tyk_oas_lint_ruleset.yml -d $file

