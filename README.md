# ASCII ART
## Project Overview
This is a system that logs printable ASCII characters in a formatted style to the terminal console. **How does it work?** User pass in an argument(s) of strings, then system is to check the string and output it formatted in respect to the style selcted. The styles are ready made, no need for generation of the stylero use. The each style is stored in a file named according to the style name, the styles to use are *shadow*, *standard* and *thinkertoy*.

## Basic NO-CODE approach of acheiving the desired output
- collect the string to format from the user
- Isolate each  characters of the sttring.
- Match each character with the formated character. But how do i acheive that if the formated character spans across 8 lines in the file. 
- Note that the standard printable ASCII characters are from decimal point of 32 and stops at 127, making a total of 95 characters. So, if each character takes up 8 lines in the file, that mean there'll be a total of 760 line. Mind you, after every character in the file, there is always a new line at the end. That makes it a total of 855 lines.
- 