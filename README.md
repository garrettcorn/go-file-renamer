# go-file-renamer
Simple golang program to change the file extension of a file.

## Install
Assuming golang is installed, run the following to clone the repository and install go-file-renamer.
```
go get -u github.com/garrettcorn/go-file-renamer
```
```
go install github.com/garrettcorn/go-file-renamer
```

If $GOPATH/bin is in the PATH, go-file-renamer's help can be viewed running the following.
```
go-file-renamer --help
```

## To Add $GOPATH to .bashrc
Open your .bashrc with your editor of choice.
```
code ~/.bashrc
```

append $GOPATH/bin to your PATH by adding the following line to your .bashrc file
```
export PATH=$PATH:$GOPATH/bin
```

## Usage
By default go-file-renamer will automatically change all of the files with the `rar` extension in the current working directory to `mkv`. To configure go-file-renamer simply provide `-from` and `-to` flags with the desired options. For example, to change all files in the current directory from `foo` to `bar` run the following.
```
go-file-renamer -from foo -to bar
```
