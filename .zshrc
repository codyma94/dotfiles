PROMPT="%n@%m:%~%# "

# auto cd into dir
setopt AUTO_CD

# pipe multiple outputs
setopt MULTIOS

# spell check commands
setopt CORRECT

# cd always pushd
setopt AUTO_PUSHD

# expand globs
setopt GLOB_COMPLETE
setopt PUSHD_MINUS

# no pushd messages
setopt PUSHD_SILENT

# blank pushd goes home
setopt PUSHD_TO_HOME

# ignore multiple dirs in stack
setopt PUSHD_IGNORE_DUPS

# 10 second delay when trying to delete everything
setopt RM_STAR_WAIT

# vim as default editor
# export EDITOR="vim"

# disable beeps
setopt NO_BEEP

# case insensitive globbing
setopt NO_CASE_GLOB

if [[ -f ~/.aliases ]]; then
  . ~/.aliases
fi
