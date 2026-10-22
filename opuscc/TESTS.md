# Test organization

Small kernel/pointer tests are grouped by subsystem rather than one file per
function. Test and helper names are unchanged, so `go test ./opuscc -run Name`
continues to work. `// From ...` comments identify the original source sections.

| Area | Grouped test files |
| --- | --- |
| CELT entropy | `celt_entropy_{decode,encode}_pointers_test.go`, `celt_laplace_pointers_test.go`, `celt_cwrs_pointers_test.go` |
| CELT signal processing | `celt_{energy,filter,pitch,stereo,vector_layout}_pointers_test.go` |
| PCM conversion, clipping and fades | `pcm_channel_pointers_test.go` |
| Packet parsing and duration | `opus_packet_pointers_test.go` |
| Projection | `projection_{decode,matrix_io}_pointers_test.go` |
| SILK resampling and filters | `silk_resampler_{fir,iir}_pointers_test.go` |
| SILK prediction and concealment | `silk_{gain,lpc,nlsf,plc,stereo}_pointers_test.go` |
| SILK entropy/pulses | `silk_pulse_coding_pointers_test.go` |

Large codec, ownership, FFT, decoder validation and C-reference suites remain
separate. In particular, `lpc_pointers_test.go` is the existing large CELT
LPC/FIR/IIR/pitch suite, not the small fixed-point SILK LPC group. The external
package test `interpolate_pointers_test.go` and build-tagged
`decoder_image_compare_test.go` also remain separate.

Consolidation changes file placement/import lists only—not declarations, fixture
paths, test names, assertions, helpers or build tags. Native reference fixtures
remain in `testdata/` and must still run with the package working directory.
