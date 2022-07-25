#!/bin/bash

prompt="Choose a Directory"
options="$(ls)"

./project --prompt "$prompt" --options "${options}"


