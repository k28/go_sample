# エディタの設定

## Vim

### Install Language Server
```
$ go get -u golang.org/x/tools/gopls
$ go get -u github.com/sourcegraph/go-langserver
```
### vimrc
```
" for golang
NeoBundle 'mattn/vim-lsp-settings'
NeoBundle 'mattn/vim-goimports'

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

