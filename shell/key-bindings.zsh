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
  zle     -N    tome-show-widget
  bindkey '^Tr' tome-show-widget

  tome-write-widget() {
    local selected num
    setopt localoptions noglobsubst noposixbuiltins pipefail no_aliases 2> /dev/null
    selected=( $(fc -rl 1 | perl -ne 'print if !$seen{(/^\s*[0-9]+\**\s+(.*)/, $1)}++' |
    FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-40%} $FZF_DEFAULT_OPTS -n2..,.. --tiebreak=index --bind=ctrl-r:toggle-sort $FZF_CTRL_R_OPTS --query=${(qqq)LBUFFER} +m" $(__fzfcmd)) )
    local ret=$?

    if [ -n "$selected" ]; then
      num=$selected[1]
      if [ -n "$num" ]; then
        zle vi-fetch-history -n $num
        echo $BUFFER | tome write
        BUFFER=''
      fi
    fi
    zle reset-prompt
    return $ret
  }

  zle     -N    tome-write-widget
  bindkey '^Ta' tome-write-widget
}
