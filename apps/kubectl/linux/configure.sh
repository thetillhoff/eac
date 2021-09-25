#!/bin/sh

if test -f "$HOME/.bashrc"; then # if bashrc is used
  if ! grep -q "source /etc/bash_completion" "$HOME/.bashrc"; then # if autocompletion is disabled
    sed -i '1i source /etc/bash_completion' "$HOME/.bashrc" # enable bash completion (insert at first line)
  fi
  if ! grep -q "source <(kubectl completion bash)" "$HOME/.bashrc"; then # if kubectl completion is not set up
    echo "source <(kubectl completion bash)" >> "$HOME/.bashrc" # # set kubectl completion up
    printf "%s\n" "Configured bash autocompletion for 'kubectl'."
  else
    printf "%s\n" "Bash autocompletion for 'kubectl' already configured."
  fi
else 
  printf "%s\n" "No '~/.bashrc' found."
fi

if test -f "$HOME/.zshrc"; then # if zshrc is used
  if ! grep -q "source <(kubectl completion zsh)" "$HOME/.zshrc"; then # if kubectl completion is not set up
    echo "source <(kubectl completion zsh)" >> "$HOME/.zshrc" # # set kubectl completion up
    printf "%s\n" "Configured zsh autocompletion for 'kubectl'."
  else
    printf "%s\n" "Zsh autocompletion for 'kubectl' already configured."
  fi
else
  printf "%s\n" "No '~/.zshrc' found."
fi
