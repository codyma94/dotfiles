call plug#begin()

" plugins
Plug 'gmarik/vundle'
Plug 'tpope/vim-surround'
Plug 'rking/ag.vim'
Plug 'Valloric/MatchTagAlways'
Plug 'docunext/closetag.vim'
Plug 'Raimondi/delimitMate'
Plug 'scrooloose/nerdcommenter'
Plug 'scrooloose/nerdtree'
Plug 'jistr/vim-nerdtree-tabs'
Plug 'ervandew/supertab'
Plug 'scrooloose/syntastic'
Plug 'bronson/vim-trailing-whitespace'
Plug 'derekwyatt/vim-scala'
Plug 'christoomey/vim-tmux-navigator'
Plug 'easymotion/vim-easymotion'
Plug 'tpope/vim-repeat'
Plug 'sjl/gundo.vim'
Plug 'nathanaelkane/vim-indent-guides'
Plug 'tpope/vim-fugitive'
Plug 'fatih/vim-go'

" color schemes
Plug 'altercation/vim-colors-solarized'
Plug 'tomasr/molokai'
Plug 'morhetz/gruvbox'
Plug 'chriskempson/vim-tomorrow-theme'
Plug 'w0ng/vim-hybrid'

call plug#end()

" don't need compatibility with vi
set nocompatible

" turn filetype detection on for plugins
filetype plugin on

" filetype indenting
filetype indent on

" set encoding
set encoding=utf-8

" syntax highlighting
syntax enable

" copy to system clipboard
set clipboard=unnamed

" no annoying error sound on errors
set noerrorbells visualbell t_vb=
autocmd GUIEnter * set visualbell t_vb=

" set backup directory
set backupdir=~/.vim/backup
set directory=~/.vim/backup

" put cursor at previous position on file open
autocmd BufReadPost * exe "normal! g`\""

" configure backspace so it acts as it should act
set backspace=indent,eol,start
set whichwrap+=<,>,h,l

" keep 500 lines of command line history
set history=500

" show the cursor position all the time
set ruler

" do incremental searching
set incsearch

" highlight search results
set hlsearch

" always set autoindenting on
set autoindent

" expand tabs to spaces
set expandtab

" be smart when using tabs
set smarttab

" 1 tab = 2 spaces
set shiftwidth=2
set tabstop=2

" round indent to nearest multiple of 2
set shiftround

" time to wait after ESC (default has an annoying delay)
set timeoutlen=200

" ignore case when searching
set ignorecase

" show matching brackets when text indicator is over them
set showmatch

" how many tenths of a second to blink when matching brackets
set mat=2

" set replace all as default
set gdefault

" start scrolling before cursor reaches the edge
set scrolloff=4

" format the status line
set laststatus=2
set statusline=%F "full file path"
set statusline+=%h "help file flag
set statusline+=%m "modified flag
set statusline+=%r "read only flag
set statusline+=%y "filetype
set statusline+=%= "left/right separator
set statusline+=%c, "cursor column
set statusline+=%l/%L "cursor line/total lines
set statusline+=\ %P " percent through file"

" fix typos
:command WQ wq
:command Wq wq
:command W w
:command Q q
:command QW wq
cnoreabbrev ag Ag!

"""""""""""""""""""""""""""""""""""""""""""""""
" Custom mappings                             "
"""""""""""""""""""""""""""""""""""""""""""""""
" set <leader> to comma
let mapleader = ","

" swap ; and :
nnoremap ; :
nnoremap : ;

" exit insert mode
inoremap jj <Esc>

" treat long lines as break lines
map j gj
map k gk

" move around windows easily
nnoremap <C-h> <C-w>h
nnoremap <C-j> <C-w>j
nnoremap <C-k> <C-w>k
nnoremap <C-l> <C-w>l

" map 0 to first nonblank character
map 0 ^

" Fix extra whitespace
nnoremap <leader>f :FixWhitespace<CR>

" Gundo
nnoremap <leader>g :GundoToggle <CR>

" open NERDTree
nnoremap <leader>n :NERDTreeTabsToggle<CR>

" Ag
nnoremap <leader>p :Ag


" paste mode
set pastetoggle=<leader>p

" toggle spellcheck
nnoremap <leader>s :setlocal spell! spelllang=en_us<CR>

" vertical split new window with <leader>v
nnoremap <leader>v :vsp

"clear highlighted search
nmap <silent> ,/ :nohlsearch<CR>

" set line numbering
set number

" mouse scroll
set mouse=a

"""""""""""""""""""""""""""""""""""""""""""""""
" Function Defn's                             "
""""""""""""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""""""""""""""""""""
" Coloring                                    "
"""""""""""""""""""""""""""""""""""""""""""""""
set background=dark

" solarized settings
let g:solarized_termtrans=1
let g:solarized_termcolors=256
let g:solarized_contrast="high"
let g:solarized_visibility="high"

" gruvbox settings
let &t_ZH="\e[3m"
let &t_ZR="\e[23m"

" select colorscheme
colorscheme gruvbox

" syntastic settings
let g:syntastic_html_tidy_ignore_errors = [
    \  'plain text isn''t allowed in <head> elements',
    \  '<base> escaping malformed URI reference',
    \  'discarding unexpected <body>',
    \  '<script> escaping malformed URI reference',
    \  '</head> isn''t allowed in <body> elements'
    \ ]

" vim-go settings
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1
let g:go_highlight_interfaces = 1
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1
let g:go_fmt_command = "goimports"
let g:syntastic_go_checkers = ['golint', 'govet', 'errcheck']
let g:syntastic_mode_map = { 'mode': 'active', 'passive_filetypes': ['go'] }

" autocmds
au FileType tex :NoMatchParen
au FileType tex set norelativenumber
au FileType go nmap <leader>t <Plug>(go-test)
au FileType go nmap <leader>c <Plug>(go-coverage)
au FileType go nmap <Leader>i <Plug>(go-info)
