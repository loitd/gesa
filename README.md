# gesa 
[![Build Status](https://travis-ci.org/loitd/gesa.svg?branch=master)](https://travis-ci.org/loitd/gesa)
[![codecov](https://codecov.io/gh/loitd/gesa/branch/master/graph/badge.svg)](https://codecov.io/gh/loitd/gesa)
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/gesa-dev/Lobby#)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/loitd/gesa/pulls)
[https://ci.appveyor.com/api/projects/status/32r7s2skrgm9ubva?svg=true](https://ci.appveyor.com/api/projects/status/github/loitd/gesa?branch=master&svg=true)
## Introduction
Go E-SQL Agent for Grafana and more written in Golang
## Installation
## Environment
## Testing
Test db:
```go
go test github.com/loitd/gesa/db -v
```
## Notes
* This program built and run in Windows so `Travis-CI` could be failed for sometime.
* Everytime you change/move the location of additional packages (oracle client, pkg-config) you need to remove and `REINSTALL` ora.v4
## License 
[![Logo](https://licensebuttons.net/l/by-nc-sa/4.0/88x31.png)](https://creativecommons.org/licenses/by-nc-sa/4.0/)  
This work is licensed under a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](https://creativecommons.org/licenses/by-nc-sa/4.0/).  

