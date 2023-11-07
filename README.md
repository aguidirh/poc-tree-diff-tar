# Diff Tar POC

This is a POC for the diff mechanism required to extract only images that were not mirrored previously in oc-mirror

## Using the command line

In the root folder, there are two folders:

``first-mirroring-albo-only`` - this folder simulates the first mirroring where only aws-load-balancer (albo) was mirrored.

``second-mirroring-albo-and-noo`` - this folder simulates the second mirroring where node-observability-operator (noo) was mirrored and aws-load-balancer (albo) was already in the registry

Run the command below
```
go run main.go
```

After running the command above 3 hidden files are going to be generated:

``.first-mirroring-albo-only-tree`` - the content of this file is the tree structure of first-mirroring-albo-only-tree/docker/registry/v2/repositories - this is the history file

``.second-mirroring-albo-and-noo-tree`` - the content of this file is the tree structure of second-mirroring-albo-and-noo/docker/registry/v2/repositories - this simulates what was generated after running the latest batch worker

``.diff-tree`` - the content of this file is the diff between the two trees (history and current). Based on this file it is possible to get only the required blobs under blobs/sha256 and generating the tar file only with the diff between the current and last mirroring.

