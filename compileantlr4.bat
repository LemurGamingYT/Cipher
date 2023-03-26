@echo off
antlr4 -Dlanguage=Go -o Core/parser/ -visitor -no-listener Core/Cipher.g4
