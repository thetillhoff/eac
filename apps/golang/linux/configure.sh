#/bin/sh

if test -f "$HOME/.profile"; then # if bashrc is used
  if ! grep -q "export PATH=$$PATH:/usr/local/go/bin" "$HOME/.profile"; then # if autocompletion is disabled
    sed -i '1i export PATH=$PATH:/usr/local/go/bin' "$HOME/.profile" # enable bash completion (insert at first line)
  else
    printf "%s\n" "Go already in PATH."
  fi
else 
  printf "%s\n" "No '~/.profile' found."
fi
