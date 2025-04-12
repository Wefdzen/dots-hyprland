return {
  { "catppuccin/nvim" },
  { "Mofiqul/dracula.nvim" },
  { "ellisonleao/gruvbox.nvim" },

  {
    "LazyVim/LazyVim",
    opts = {
      colorscheme = "catppuccin-frappe", --dracula --gruvbox
      priority = 1000,
    },
  },

  {
    "catppuccin",
    opts = {
      transparent_background = true,
    },
  },
  {
    "dracula.nvim",
    opts = {
      --transparent_bg = true,
    },
  },
  {
    "gruvbox.nvim",
    opts = {
      transparent_mode = false
    },
  },
}
