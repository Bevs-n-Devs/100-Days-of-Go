# 100 Days of Go

This repository contains Golang data structures, algorithms, and small projects designed to improve my skills and deepen my understanding of Go.

## Initializing Go Project Module
When creating a module, you should stick to the following steps
```
// create the directory which you will work from 

new-item day-1 -type directory      // Windows
mkdir day-1                         // Linux/MacOS

// navigate to the directory and initialize the module 

cd day-1
go mod init day-1                                               // Locally (when not pushing to a repository)
go mod init github.com/Bevs-n-Devs/100-Days-of-Go/day-1         // Using a GitHub repository (when you plan to publish it or share it publicly)
```

Go modules help manage project dependencies and versions, making it easy to track and manage external libraries or packages your project uses.

Once your module is initialized, you can easily add dependencies using `go get` or by importing them, and Go will manage these for you in the `go.mod` file.

To add module requirements and sums: `go mod tidy`