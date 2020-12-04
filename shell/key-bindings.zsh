{
  [[ -o interactive ]] || return 0

  tome-show-widget() {
    local selected=$(tome show | tac |
      FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-40%} $FZF_DEFAULT_OPTS --tiebreak=index --bind=ctrl-r:toggle-sort $FZF_CTRL_R_OPTS --query=${LBUFFER} +m --with-nth=5..100 --delimiter=';'" $(__fzfcmd) |
      cut -d ';' -f 1)
    local ret=$?
    BUFFER=$(tome get --id "$selected")
    zle reset-prompt
    return $ret
  }
  zle     -N     tome-show-widget
  bindkey '^R^R' tome-show-widget
}
