//Package strand implements a solution to the rna-transcription prompt from exercism.io
package strand

var transcribe = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

/*
ToRNA returns the RNA strand that compliments the input DNA strand using the following conversion rules:
 DNA -> RNA
 `G` -> `C`
 `C` -> `G`
 `T` -> `A`
 `A` -> `U`
*/
func ToRNA(dna string) (rna string) {
	if len(dna) == 0 {
		return
	}

	for _, nucleotide := range dna {
		rna += transcribe[string(nucleotide)]
	}

	return
}
