#!/bin/sh
echo "Formating snippets with goimports"
for filename in /var/snippets/go/*.go; do
        a=$(goimports $filename);
        echo "$a" > $filename;
done
