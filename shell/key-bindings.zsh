{
  [[ -o interactive ]] || return 0

  tome-show-widget() {
    selected=$(tome show | fzf --query=${LBUFFER} --with-nth=5..100 --delimiter=';' | cut -d ';' -f 1)
    local ret=$?
    BUFFER=$(tome get --id "$selected")
    zle reset-prompt
    return $ret
  }
  zle     -N     tome-show-widget
  bindkey '^R^R' tome-show-widget
}
