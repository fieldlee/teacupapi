#!/bin/bash

allPassed=true

while IFS= read -r -d '' file
do 
    name=$(basename "$file")
    if [[ $name =~ [A-Z]+ ]]; then
        allPassed=false
        echo "$file"
    fi
done< <(find . -name "*.go" -print0)

if [ "$allPassed" == false ]; then
    echo "Go source file should use lowercase or underscore"
    exit 1
fi
