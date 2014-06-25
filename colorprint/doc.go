/*
Package colorprint print ANSI coloured text to the standard output.

Colours are passed to the functions by strings that define a colour code.

Colour code syntax:

	r ... red
	g ... green
 	y ... yellow
	b ... blue
	x ... black
	m ... magenta
	c ... cyan
	w ... white
	d ... default
	+ ... bold
	* ... italic
	~ ... reverse
	_ ... underlined
	# ... blink
	? ... concealed

	background coloring is affected by using upper case letters

Some examples:
	+r ... bold red
	rY ... red text on yellow background

*/
package colorprint
