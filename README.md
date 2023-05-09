# What is rpg?
rpg is random password generator runs in your terminal/console.

# How to Use?
First, you need to install rpg from [here](https://github.com/AlperAkca79/rpg/releases/latest).

- Open your terminal and type this:
```
rpg [arguments]
```

![rpg_usage](https://github.com/AlperAkca79/cat/assets/91411319/db469b29-a4ec-48c0-92f1-2cf7d0f312d2)

### Arguments

`-length <integer>`: Chance to add number to your password. You can type from 0 to infinity.

`-uppercase`: Chance to add uppercase letters to your password.

`-lowercase`: Chance to add lowercase letters to your password.
 
`-number`: Chance to add numbers to your password. (0123456789)

`-special`: Chance to add special characters to your password. This argument contains the special characters `!`, `@`, `#`, `$`, `%`, `^`, `&`, `*`, `(`, `)`, `_`, `-`, `+`, `=`, `[`, `]`, `{`, `}`, `|`, `;`, `:`, `,`, `.`, `<`, `>` and `?`.

### Examples
- Command:

  `rpg -length 12 -uppercase -number`

  Output:

  `CRSFGLCQLRA2`

- Command:

  `rpg -length 16 -lowercase -special`

  Output:

  `!ya$^}w.gb[d]iy&`

- Command:

  `rpg -length 18 -uppercase -lowercase -number -special`

  Output:

  `Obi}M[UnmO3@0,EfZw`