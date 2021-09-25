#!/bin/sh

if test -f "$HOME/.profile"; then # if bashrc is used
  if ! grep -q 'export GOROOT=/usr/local/go' "$HOME/.profile"; then # if goroot is not configured
    printf "%s\n" 'export GOROOT=/usr/local/go' >> "$HOME/.profile" # configure goroot
  fi
  if ! grep -q 'export GOPATH=$HOME/go' "$HOME/.profile"; then # if gopath is not configured
    printf "%s\n" 'export GOPATH=$HOME/go' >> "$HOME/.profile" # configure gopath
  fi
  if ! grep -q 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' "$HOME/.profile"; then # if go is not in path
    printf "%s\n" 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> "$HOME/.profile" # add go to path
  fi

  if ! grep -q "export PATH=$$PATH:/usr/local/go/bin" "$HOME/.profile"; then # if autocompletion is disabled
    sed -i '1i export PATH=$PATH:/usr/local/go/bin' "$HOME/.profile" # enable bash completion (insert at first line)
  else
    printf "%s\n" "Go already in PATH."
  fi
else 
  printf "%s\n" "No '~/.profile' found."
fi
