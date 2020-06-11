//Package strand implements a solution to the rna-transcription prompt from exercism.io
package strand

import "strings"

var transcriber = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U")

/*
ToRNA returns the RNA strand that compliments the input DNA strand using the following conversion rules:
 DNA -> RNA
 `G` -> `C`
 `C` -> `G`
 `T` -> `A`
 `A` -> `U`
*/
func ToRNA(dna string) string {
	return transcriber.Replace(dna)
}
