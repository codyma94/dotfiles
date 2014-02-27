if [ -f ~/.profile ]; then
  . ~/.profile
fi

if [ -f ~/.aliases ]; then
  . ~/.aliases
fi

### Better Completion ###
if [ -f $(brew --prefix)/etc/bash_completion ]; then
  . $(brew --prefix)/etc/bash_completion
fi
