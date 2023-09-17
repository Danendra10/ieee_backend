#!/bin/bash

output_name="main"
os="linux"

if [[ $1 == "--help" || $1 == "-h" ]]; then
    echo "Usage: ./script.sh [output_name]"
    echo "-o: OS"
    echo "-n: output name"
    exit 0
fi

while getopts ":o:n:" opt; do
  case ${opt} in
    o )
      os=$OPTARG
      ;;
    n )
      output_name=$OPTARG
      ;;
    \? )
      echo "Invalid option: $OPTARG" 1>&2
      exit 1
      ;;
    : )
      echo "Invalid option: $OPTARG requires an argument" 1>&2
      exit 1
      ;;
  esac
done

echo "Building for OS: $os"
echo "Output name: $output_name"

GOOS = $os
echo "BUILDING ${output_name}"
go build -o ${output_name} main.go
echo "BUILD DONE"