# エディタの設定

## Vim

### Install Language Server and tools
```
$ go get -u golang.org/x/tools/gopls
$ go get -u github.com/sourcegraph/go-langserver
$ go get golang.org/x/tools/cmd/goimports
```

### dein.toml
```
# for golang
[[plugins]]
repo = 'mattn/vim-lsp-settings'

[[plugins]]
repo = 'mattn/vim-goimports'
```
### vimrc
```

" go lsp-setting
if executable('gopls')
    au User lsp_setup call lsp#register_server({
        \ 'name': 'gopls',
        \ 'cmd': {server_info->['gopls']},
        \ 'whitelist': ['go'],
        \ })
    autocmd BufWritePre *.go LspDocumentFormatSync

    augroup go_lsp_omnifunf
        autocmd!
        autocmd FileType go setlocal omnifunc=lsp#complete
    augroup END
endif

if executable('go-langserver')
    au User lsp_setup call lsp#register_server({
        \ 'name': 'go-langserver',
        \ 'cmd': {server_info->['go-langserver', '-gocodecompletion']},
        \ 'whitelist': ['go'],
        \ })
    autocmd BufWritePre *.go LspDocumentFormatSync
endif
```

