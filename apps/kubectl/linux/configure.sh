#/bin/sh

shell=$(readlink /proc/$$/exe)
echo $HOME

if test -f "$HOME/.bashrc"; then # if bashrc is used
  if ! grep "source /etc/bash_completion" "$HOME/.bashrc"; then # if autocompletion is disabled
    sed -i '1i source /etc/bash_completion' "$HOME/.bashrc" # enable bash completion (insert at first line)
  fi
  if ! grep "source <(kubectl completion bash)" "$HOME/.bashrc"; then # if kubectl completion is not set up
    echo "source <(kubectl completion bash)" >> "$HOME/.bashrc" # # set kubectl completion up
  fi
  printf "%s\n" "Configured bash autocompletion for 'kubectl'."
else 
  printf "%s\n" "No '~/.bashrc' found."
fi

if test -f "$HOME/.zshrc"; then # if zshrc is used
  if ! grep "source <(kubectl completion zsh)" "$HOME/.zshrc"; then # if kubectl completion is not set up
    echo "source <(kubectl completion zsh)" >> "$HOME/.zshrc" # # set kubectl completion up
  fi
else
  printf "%s\n" "No '~/.zshrc' found."
fi
