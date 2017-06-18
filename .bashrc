if [ -f ~/.env ]; then
  . ~/.env
fi

if [ -f ~/.aliases ]; then
  . ~/.aliases
fi

[ -f ~/.fzf.bash ] && source ~/.fzf.bash
