# Path to your oh-my-zsh installation.
export ZSH=$HOME/.oh-my-zsh

# Look in ~/.oh-my-zsh/themes/
ZSH_THEME="bureau"

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# Load aliases
if [[ -f ~/.aliases ]]; then
  source ~/.aliases
fi

# Load env vars
if [[ -f ~/.env ]]; then
  source ~/.env
fi

setopt hist_ignore_all_dups
setopt hist_ignore_space

DISABLE_UNTRACKED_FILES_DIRTY="true"

# plugins can be found in ~/.oh-my-zsh/plugins/
plugins=(brew git osx python scala terminalapp tmux rails)

source $ZSH/oh-my-zsh.sh
