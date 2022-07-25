#!/bin/bash

prompt="Choose a Directory"
options="$(ls)"

./serve --prompt "$prompt" --options "${options}"


