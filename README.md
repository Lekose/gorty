# Introduction
This project is a bit of a play to see what we can do with Go to interface with the WinAPI on Windows 10. 

# Installing Go
So there are two parts to this, first is easy enough just head to https://golang.org/doc/install?download=go1.14.6.windows-amd64.msi and follow the instructions.

Unforunately the next step was a pain in the ass, so we want to run this against x64 builds and ideally be able to build shared libraries using GO. To do this we need G++ but not the x86 version, basically download this:

https://sourceforge.net/projects/mingw-w64/

IMPORTANT when installing this make sure you change the target to x86_64 otherwise you're gonna have a bad time. You also need to add mingw to your system environmental path (the bin folder) let me know if you need a hand with this.

# building a DLL

see DllSample.go for the test code i made to create a DLL using go. Some notes:
* we need an empty main function because go screams if we don't have one, don't think this will cause any issues with duplicate entry points etc but its a bit annoying.
* Go forces its own formating and gives errors if you don't use it, sorta hate that.

to build the same naviagte to the folder DllSample.go is and run the following command:

> go build -o libtest.dll -buildmode=c-shared dllsample.go

Havent had chance to play with this much but the DLL looks valid. As part of the build process it also auto gens a c header file, from what i can tell most of go under the hood is just converting its own code to C before compiling when building shared libraries. 