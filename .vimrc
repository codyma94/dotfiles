" load plugins via pathogen
call pathogen#infect()
call pathogen#interpose('bundle/{}')

" use Vim settings
set nocompatible

" turn filetype detection on for plugins
filetype plugin on

" filetype indenting
filetype indent on

" set backup directory 
set backupdir=~/.vim/backup
set directory=~/.vim/backup

" time to wait after ESC (default has an annoying delay)
set timeoutlen=250

" yanks to system clipboard
set clipboard+=unnamed

" configure backspace so it acts as it should act
set backspace=indent,eol,start
set whichwrap+=<,>,h,l

" mouse stuff
set mouse=a
set ttymouse=xterm

" set line width
set wrap
set textwidth=79
set formatoptions=qrn1

" ignore case when searching
set ignorecase

" don't redraw while executing macros
set lazyredraw

" show matching brackets when text indicator is over them
set showmatch

" how many tenths of a second to blink when matching brackets
set mat=2

" set replace all as default
set gdefault

" set relative numbers as default
set relativenumber

" no annoying error sound on errors
set noerrorbells visualbell t_vb=
autocmd GUIEnter * set visualbell t_vb=

" use spaces instead of tabs
set expandtab

" be smart when using tabs
set smarttab

" 1 tab = 2 spaces
set shiftwidth=2
set tabstop=2

" set encoding
set encoding=utf-8

" always set autoindenting on
set autoindent

" always show the status line
set laststatus=2

" format the status line
set statusline=%F "full file path"
"set statusline+=[%{&ff}] "file format
set statusline+=%h "help file flag
set statusline+=%m "modified flag
set statusline+=%r "read only flag
set statusline+=%y "filetype
set statusline+=%= "left/right separator
set statusline+=%c, "cursor column
set statusline+=%l/%L "cursor line/total lines
set statusline+=\ %P " percent through file"

" remove the windows ^M when encodings gets messed up
noremap <Leader>m mmHmt:%s/<C-V><cr>//ge<cr>'tzt'm

" solarized coloring
 syntax enable
 set background=light
 let g:solarized_termtrans=1
 let g:solarized_termcolors=256
 let g:solarized_contrast="high"
 let g:solarized_visibility="high"
 colorscheme solarized


" only load closetag on html/xml like files
autocmd FileType html,htmldjango,jinjahtml,eruby,mako let b:closetag_html_style=1
autocmd FileType html,xhtml,xml,htmldjango,jinjahtml,eruby,mako source ~/.vim/bundle/closetag/plugin/closetag.vim

" fix typos
:command WQ wq
:command Wq wq
:command W w
:command Q q
:command QW wq

" map 0 to first nonblank character
map 0 ^

" treat long lines as break lines
map j gj
map k gk

" create backup files
" set backup

set history=50		" keep 50 lines of command line history
set ruler		" show the cursor position all the time
set showcmd		" display incomplete commands
set incsearch		" do incremental searching
set hlsearch

" relative and absolute numbering stuff 
function! NumberToggle()
  if(&relativenumber==1)
    set number
  else
    set relativenumber
  endif
endfunc

" toggle normal line numbers with relative line numbers
nnoremap<C-n> :call NumberToggle()<CR>

" set <leader> to comma
let mapleader = ","

"set ; to do :
nnoremap ; :

" vertical split new window with <leader>v
nnoremap <leader>v :vsp 

" start ack with <leader>a
nnoremap <leader>a :Ack

" toggle spellcheck
nnoremap <leader>s :setlocal spell! spelllang=en_us<CR>

" Gundo keybindings
nnoremap <leader>g :GundoToggle<CR>
"
" exit insert mode with jj 
:imap jj <Esc>

" put cursor at previous position
autocmd BufReadPost * exe "normal! g`\""

" Rainbow parens always on
au VimEnter * RainbowParenthesesToggle
au Syntax * RainbowParenthesesLoadRound
au Syntax * RainbowParenthesesLoadSquare
au Syntax * RainbowParenthesesLoadBraces

" open NERDTree
nnoremap <Leader>n :NERDTreeTabsToggle<CR>

" move around windows easily
nnoremap <C-h> <C-w>h
nnoremap <C-j> <C-w>j
nnoremap <C-k> <C-w>k
nnoremap <C-l> <C-w>l
