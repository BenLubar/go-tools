#!/bin/bash -e

case "`~/.golang/bin/go env GOOS`" in
"windows")
	case "`~/.golang/bin/go env GOARCH`" in
	"386")
		export CC=i686-w64-mingw32-gcc
		export CXX=i686-w64-mingw32-g++
		export CGO_ENABLED=1
		;;
	"amd64")
		export CC=x86_64-w64-mingw32-gcc
		export CXX=x86_64-w64-mingw32-g++
		export CGO_ENABLED=1
		;;
	esac
	;;
esac

~/.golang/bin/go "$@" < /dev/stdin
