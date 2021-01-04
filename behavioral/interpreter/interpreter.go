package pattern

/*

Pattern:
interpreter

Plan:
Visitors are approached by a small being which gives them a book and says
'We the maze denizens have developed the language written on the
walls for efficiency and private communication about maze routes.
If you're willing to put in the time to learn this language passed down by many
generations you will be rewarded with deft navigation of the maze'

Premise:
- Expression: interface for expression that can be interpreted
- Concrete expressions: parts of the language that can be interpreted

expression = action - frequency - location
action := do thing
location := numberpaces direction

do literals := '🃕' | '❖' | '✿' (push, pull, twist)
thing literals := '♕' | '🁡' | '❐' (wall, lever, button)
frequency literals := '☽' | '✂︎' | '⚭' (do once, do reversed, do many times)
direction literals := '⁍' | '◎' | '🀫' | '⌘' (north, south, east, west)
numberpaces literals: octal digit

valid expression example:
✿ ❐ - ✂︎ - 412 ⌘

^^ not sure how useful this all is so saving this one for later/never

*/
