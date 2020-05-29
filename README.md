# Writing an Interpreter in Go

## Parsing
### Parsers
According to Wikipedia
>A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process. [...] The parser is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters;

A parser turns it into a data structure that represents the input

```> var input = '{"name": "Thorsten", "age": 28}'; > var output = JSON.parse(input);
> output
{ name: 'Thorsten', age: 28 }
> output.name 'Thorsten'
> output.age 28
```

How is it related to a Interpreter? A JSON parser takes text as input and builds the data structure that represents the input. That's exactly what a parser for a programming language does.