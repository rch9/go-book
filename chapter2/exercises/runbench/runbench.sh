#!/usr/bin/env bash

# must run from this folder!

echo "tests started"
echo

echo "example:"
cd ../../6/popcount/
go test -bench=.

echo
echo "ex2-3:"
cd ../../exercises/ex2-3/popcount/
go test -bench=.

echo
echo "ex2-4:"
cd ../../ex2-4/popcount/
go test -bench=.

echo
echo "ex2-5:"
cd ../../ex2-5/popcount/
go test -bench=.
