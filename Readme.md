
# Notes for the Lexer Implementation

## Lexer Struct

- The `Lexer` struct holds the following fields:
- `input` (string): The source code input to be lexed.
- `position` (int): The current position in the input string being examined.
- `readPosition` (int): The position of the next character to be read from the input string.
- `ch` (byte): The current character being examined.

## New(input string) *Lexer

- This function creates a new `Lexer` instance with the given `input` string.
- It initializes the `position`, `readPosition`, and `ch` fields by calling `readChar()`.

## readChar()

- This method updates the `ch` field with the next character from the input string.
- If the end of the input is reached, it sets `ch` to 0 (NUL character).
- After updating `ch`, it updates `position` to the previous `readPosition` and increments `readPosition` to point to the next character.

## NextToken() token.Token

- This is the main method of the lexer. It returns the next token from the input string.
- It uses a switch statement to check the current character (`ch`) and create the corresponding token.
- If the character is one of the recognized token types (e.g., `=`, `;`, `(`, `)`, `,`, `+`, `{`, `}`), it creates a new token using `newToken()` with the appropriate token type and literal value.
- If the current character is 0 (NUL), it creates an `EOF` (End-Of-File) token.
- After creating the token, it calls `readChar()` to advance to the next character in the input.

## newToken(tokenType token.TokenType, ch byte) token.Token

- This helper function creates a new `token.Token` with the given `tokenType` and `ch` (literal value).

## Future Improvements

- The current implementation only recognizes specific characters as tokens.
- In the future, you'll need to extend the lexer to handle more complex token types, such as:
- Identifiers (variable names)
- Integers
- Keywords (e.g., `let`, `fn`)
- The existing code provides a foundation for building a more comprehensive lexer for the Monkey language.
