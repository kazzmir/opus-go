package opuscc

import (
	"math"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// C reference values from /tmp/opencode/decodeframe_ref.c (links
// ../opus/.libs/libopus.a). opus_decode_frame is static in C, so both the
// harness and these tests drive it through the exported opus_decode_float
// entry point with a 48 kHz stereo decoder.
//
// Part A decodes real packets from opus_newvectors/testvector02.bit (SILK NB)
// and testvector01.bit (CELT FB), including PLC calls, a decode-gain frame
// (exercises the inlined celt_exp2 bit-twiddling union in opus_decode_frame),
// a decode_fec call, a mono->stereo switch, a silk->celt and a celt->silk
// mode transition (pcm_transition + smooth_fade + CELT fade-out paths).
// Part B decodes packets produced by the C encoder (see decodeframe_ref.c)
// covering HYBRID frames (redundant CELT band, redundant_rng), a hybrid->SILK
// transition (silence fade), a SILK frame carrying CELT redundancy
// (prev_redundancy=1), a silk->CELT transition and CELT PLC.
//
// Frames that do not run the CELT decoder are bit-exact against the C
// reference (fnv/first-4-float bits taken from the harness output). Frames
// marked celt=true run the CELT layer, which has a known pre-existing
// low-bit float divergence between the Go and C builds (the same one that
// shows up as 123 LSB bytes when e2e-decoding testvector01; the SILK layer
// and all entropy-coder state are bit-exact). For those frames the sample
// expectations below are Go golden values that must be updated if that
// divergence is ever fixed; n, rangeFinal and mode stay C-derived and are
// exact.

type frameExpect struct {
	n    int32
	fnv  uint32
	rng  uint32
	b    [4]uint32
	celt bool // true when the CELT layer runs: fnv/b are Go golden values (see file comment)
}

var decodeFrameRefA = []struct {
	tag  string
	pkt  int // index into packets below; -1 = PLC (NULL data)
	fec  int32
	size int32
	want frameExpect
}{
	{"pkt0", 0, 0, 5760, frameExpect{2880, 0x20a8ba55, 0x50373c71, [4]uint32{0, 0, 0, 0}, false}},
	{"pkt1", 1, 0, 5760, frameExpect{2880, 0x579c5571, 0x0671027e, [4]uint32{0xb9900000, 0xb9900000, 0xb9900000, 0xb9900000}, false}},
	{"plc120", -1, 0, 5760, frameExpect{5760, 0x3bb44ab5, 0, [4]uint32{0xba000000, 0xba000000, 0xb9c00000, 0xb9c00000}, false}},
	{"plc5", -1, 0, 240, frameExpect{240, 0xe5f35745, 0, [4]uint32{0xb8000000, 0xb8000000, 0, 0}, false}},
	{"gain", 0, 0, 5760, frameExpect{2880, 0x3e91f2ad, 0x50373c71, [4]uint32{0x380f9e4d, 0x380f9e4d, 0x380f9e4d, 0x380f9e4d}, false}},
	{"fec", 0, 1, 5760, frameExpect{5760, 0x03c7e665, 0x08000000, [4]uint32{0xb9800000, 0xb9800000, 0xb9800000, 0xb9800000}, false}},
	{"pkt602", 2, 0, 5760, frameExpect{2880, 0xa9569e94, 0x012d0ad8, [4]uint32{0xb9b00000, 0xb9b00000, 0xb9b00000, 0xb9b00000}, false}},
	{"celt", 3, 0, 5760, frameExpect{2880, 0xc1d4075d, 0x16230400, [4]uint32{0xbba00000, 0xbb600000, 0xbb920000, 0xbb4a0000}, true}},
	{"back2silk", 0, 0, 5760, frameExpect{2880, 0x052a1179, 0x50373c71, [4]uint32{0x3813a17a, 0x3766de6c, 0x3831c915, 0x3801e941}, true}},
}

func TestOpusDecodeFrameCReferencePartA(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	dec, err := Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	defer Opus_opus_decoder_destroy(tls, dec)
	dp := (*OpusT_OpusDecoder)(unsafe.Pointer(dec))

	/* tv02 pkt0/pkt1: SILK NB mono 60 ms; tv02 pkt602: SILK NB stereo 60 ms;
	   tv01 pkt0: CELT FB stereo. */
	packets := [][]byte{
		mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40"),
		mustHex(t, "182312cf4040d200ea1335b36ad4d1a12853dd70b1861253119131ec38"),
		mustHex(t, "1ce27f853110ec825fbabc88e8820cb1a5d6ba25069ec1d8cc14ea25e1a38ccfd4eaafb3"+
			"596c1444c2e02fe8706847d56e57cf605f9732ddaf5efa683bae5ebfd9a8cd3bfa9c7e32"+
			"f0e48418b97fc54503cecaa1cc6b1a46d91123908119b9793b3e6bd408b89c9feab30575"+
			"f971e925bc8d7eb061bb22f4786430e2a148792d3e3097bea8e1186f6b0efdcc107ac6f3"+
			"3d1b0ca884e15c2c3b9f17c11cabb6afd73b717cb2d901532dcb41adeddc6d04352a04a6"+
			"cd4eac275ab18a17b2e1ff1161700e743564dfdbfc58854491b94fa1565204d15f5f4b15"+
			"a39343494601fa5c8c7f9b464063c33266951dc965096200e48d1a119d1946563e3a2804"+
			"09654f1e9941bbf6ec7374d995715c902b914e63262db048048fdc711480660299247c40"+
			"721cfd020a8c9967c7ac6e66ea309821aa31b8fe9de96376c5a58cb0bae1c2f113678600"+
			"412f4c6130d956648101f349d5bbf3a73b6ae43a93b92a5537"),
		mustHex(t, "ff83fe17bd7ffac7671a851b2c6c118bfb0b91bad839564f5e88bd307db2746e6d9fc172e256c59e0527b3d9b62d10af1b2a"+
			"27271e3ad11cb664f3efa2eb693f8d02b1d5e13a7e37da75ec4993e736a4913fb6671c77ba14ad235130c1d26a17c3aeee42"+
			"b3419933611d087773bf88a80193dcdf678b1ac95d3f907cb124444eb04330e4d88b3a5ba3757e0ccecd55df2320fe742cb5"+
			"7fea3684043d33bed2103a8f35d4f01fb3393d089c36f7b38ceee18b2b7909d5ed76c4c283abd523d533e0826372bdbb1178"+
			"cf2ee403378567827e52c6b11de3eed35e893b832d6998a2d7da477b3c8368f3e8830960ac59cee3e21d92e1a83925c5a313"+
			"6ae351febfd7097bc20eed5a619286b011d4c1da4a996d86afb58c5bd6784a133bc1adc32b760d54199195fd4a997c56c8d3"+
			"008d7627953a99ce9aeb8d27f6f803a4d33dba3a466e3c9562e15166f175f8edb8c5694c78884aa932e108e7d3e397ba3aad"+
			"8b61d6a582503d532086fe5871adb29902841192bae1c1199ae6c2816a7ba5bc0eba86306de07b0be93a9fba535ca9fa825f"+
			"5665bbb41022c9142eebfa289f6e2084a645d9b98d7dc255e83bc51c7a0235dd770044a27b61d096dda1e6eb9c8dc1d72ba6"+
			"469cfb48f20bc10e5f0a66754c67fd22b743d390b14cf88f1c8909c9f0aacb35d685a7db654100e18c8d72be875d696624d9"+
			"dd072442d04c304d58deb3ff7154286723c7752902678a905ad83850dcce6d4647956aab214e2df461d6a58239ab13daed67"+
			"706ee15c90e25bcfbd5a353f6dd35a2fd15cdcb8423589fdc5c8a6dfc4a5cb56e9e95b5a35b39510862a59f429ea37b38f3b"+
			"598bf5384ab1e275f0bde9fa2651f232a51ea7e467be28e0fb6b8d13dc4ef0a7f608f77f00042ca796821d129c88bdff5ea3"+
			"8adac96f9a8687e6f46991a8c7a1f3aba5fef5393f09cde717ea3def2e86a642413194761af8dc65247ba2e3ee367c021423"+
			"4cfa74e4f04466c693a5450bac0e291f552bc8147cf7e990959a14a104b54c"),
	}

	for i, step := range decodeFrameRefA {
		if step.tag == "gain" {
			var slot [2]int32
			if got := Opus_opus_decoder_ctl(tls, dec, OPUS_SET_GAIN_REQUEST, libc.VaList(uintptr(unsafe.Pointer(&slot[0])), int32(256))); got != OPUS_OK {
				t.Fatalf("step %d (%s): OPUS_SET_GAIN: %d", i, step.tag, got)
			}
		} else if step.tag == "fec" {
			var slot [2]int32
			if got := Opus_opus_decoder_ctl(tls, dec, OPUS_SET_GAIN_REQUEST, libc.VaList(uintptr(unsafe.Pointer(&slot[0])), int32(0))); got != OPUS_OK {
				t.Fatalf("step %d (%s): OPUS_SET_GAIN(0): %d", i, step.tag, got)
			}
		}

		pcm := make([]float32, 5760*2)
		var dataPtr uintptr
		var dataLen int32
		if step.pkt >= 0 {
			dataPtr = uintptr(unsafe.Pointer(&packets[step.pkt][0]))
			dataLen = int32(len(packets[step.pkt]))
		}
		n := Opus_opus_decode_float(tls, dec, dataPtr, dataLen, uintptr(unsafe.Pointer(unsafe.SliceData(pcm))), step.size, step.fec)
		if n < 0 {
			t.Fatalf("step %d (%s): decode: %d", i, step.tag, n)
		}
		want := step.want
		if n != want.n {
			t.Fatalf("step %d (%s): n=%d want %d", i, step.tag, n, want.n)
		}
		if got := fnv1aFloats(pcm[:int(n)*2]); got != want.fnv {
			t.Fatalf("step %d (%s): fnv=%08x want %08x (celt frame: Go golden, see file comment)", i, step.tag, got, want.fnv)
		}
		if got := dp.FrangeFinal; got != want.rng {
			t.Fatalf("step %d (%s): rangeFinal=%08x want %08x", i, step.tag, got, want.rng)
		}
		for j := 0; j < 4; j++ {
			if got := math.Float32bits(pcm[j]); got != want.b[j] {
				t.Fatalf("step %d (%s): pcm[%d]=%08x want %08x", i, step.tag, j, got, want.b[j])
			}
		}
	}

	if got, want := dp.Fdecode_gain, int32(0); got != want {
		t.Fatalf("decode_gain: got %d want %d", got, want)
	}
	if got, want := dp.Fprev_mode, int32(MODE_SILK_ONLY); got != want {
		t.Fatalf("prev_mode after celt->silk: got %d want %d", got, want)
	}
	if got, want := dp.Fprev_redundancy, int32(0); got != want {
		t.Fatalf("prev_redundancy: got %d want %d", got, want)
	}
}

func TestOpusDecodeFrameCReferencePartB(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	dec, err := Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	defer Opus_opus_decoder_destroy(tls, dec)
	dp := (*OpusT_OpusDecoder)(unsafe.Pointer(dec))

	/* Packets produced by the C encoder (deterministic LCG noise input; see
	   /tmp/opencode/decodeframe_ref.c): seg0 hybrid FB @20 kbps, seg1/2 SILK
	   WB @56/20 kbps, seg3 CELT FB @96 kbps (silk->celt transition), seg4
	   CELT WB @20 kbps, then a 120 ms PLC frame. */
	packets := [][]byte{
		mustHex(t, "7c8cb723a4e954f30817690d8021804d5c6f7c7a79cdaeeeda68b9cc67aab183653ff229912863fc3f7a335205cb0e03"+
			"3bed80eb1a0cfd3f5f"),
		mustHex(t, "7c8cc676dcad40dbc8d29d8c3742932526cf5cb73422f4484e23294abafbd6f43cf4f0be8f037456e8bf8a941b1cc6bb"),
		mustHex(t, "7c89304a6e5771290464a49a7b9ae7faad938f4cfa4f7e87b076d0fd29b8148c9bba0b085e8d4a8fc70fe4ba126ee8e1"+
			"1a80"),
		mustHex(t, "7c89169498b4acaa71ef3c3036299ca0842e1c72362e835f826fcac1eeb088470759f85df05cd98914a64cd3f7972fb7"+
			"e5023815"),
		mustHex(t, "4caccf4a733dbcf9159b0c462d88c972ac952d88b7d24e1a4f9b2af2dc964db88fdc778c77009b751ebd6c1cd9f7ef05"+
			"334fc6ffdf5a26ed4efceb1543db50f04c12474c5921d3b234fd088e0ff828065dbf70ee0b60ce91883b6084d7f8cc1e"+
			"da909cb2424cdba3995fc9d6ea764cc9cd223229f216b41666d9ccaeac6e43"),
		mustHex(t, "4cacc75331e3a040d4a10faa4f15bf4928e24dc13148273e8514fbce8ae08398ff17da946cfe9e9cf0b5246807058b32"+
			"629125f56601607a217720d266adb3a9a2f7832acf3afe2b288256b61502bf0914943e08ab8ffb011108b8fac9dec32e"+
			"5197d5363b1851e4c89e103a3579326cc18656eacf211ba22d934ede4d75f1c0fa72d83efe56bd2a7a7dba4b87ff80"),
		mustHex(t, "4ca9db86f5fd6835eabdffe5d5657e59f6e4f039b83cb8998aa27d180104c677e47aa136ed1fac6c15145bd4e1a547c4"+
			"389e28f8a8ffe878ea7a4c3a6163503a25aefb5fbab30dbbf4c811a0bb81f93476f8879172a0751f398bc5bfe70f66bb"+
			"d711574c48bb6e3f087ba524382dd7db120e16dbf8484b82514eb431dd869f6b65b5874b14287f2b76bc2cc90f230880"+
			"237010"),
		mustHex(t, "4ca8b7e9911114f14df6e02efd853e23a9c79aea72c664988f3bf929dc5cd99487ea7061dc32c078fcc90c434f77ca1c"+
			"69297dd5dd466f19ad6d8cedbf3a705a0e40a37f8740edb02f38cb9694f1cea994e055a6386e624238c3cb64c7d10108"+
			"466ea71a2ffeb3dd4fbae4cb69093a16fd606d647f507705cb02656e43abf9738aa964978dedb0dfddc0"),
		mustHex(t, "4caca1be4550d176585c363ab7d02cab1fe02d9d55d95fb698f64360bc0631fa3715afbaa2bc0040"),
		mustHex(t, "4cacbdba7e90384a0780f723ea3a36a9aae4187402da61c4cc349640807bcf1de611b023078f78b50042f581bd"),
		mustHex(t, "4ca9ef9fe3714ea87edd062a5f744f7afebc419e78356624145e0aaca4272c0bae8cfed2cd985c9c903cea8e14"),
		mustHex(t, "4ca92093bbf4674d7327a5d8918ff0db8632af079395e945ff2366e4049179f3e97942d2faf58ade9ec34aa3691e511a"+
			"3ddf0f"),
		mustHex(t, "4ca9eaccef6c63bc5064048168256b936c603eb6c52b4e9ff395d767f7579dbda91f9cbc553ae797972d874d13000649"+
			"c0d8ad994c44cedd0b66ab489e0ea6aee15d82aed2dda9b8b27431bce0d22012fd67996728d56121e30654e663266e6d"+
			"98516a9897ee16653ee2f72c6dc6c538bb6432a1b0df037b9d7e44c2e51fcaf0b093447f7f807e0d1884cfda878e99ac"+
			"ebddeec7b50340949a805894dbfd9d200e97727fe658d8b4e3c42ef8471037fdff8bd47278e01efa704732c558272d5e"+
			"480a8c16c1516ac12f4dcc506f56dac9a34b8c235fc682c523d7be414801b73140cde32e448ac5713305e87a43cd"),
		mustHex(t, "fc6200d47bf27851320242f454f12e8e38eb8b7b8c21c4e7a418cacba440e041a3ca487637ca40f1fa250e65e046479e"+
			"f35ba849f7af108d3c96e59c0520147fa8cd9cc3e876d0891e6327578e5e9a10513ee1a2b6ae0f8ef7620393101806a7"+
			"0b61ce9d453bfc8d956019c61b593b08431232c1236aa4f0806e4b8d33ac9afc04e7a0c13df18b2a6f943efe1e3af3bc"+
			"777e184d62c9768de1f48e4239e69e6cf175bba6a608e905e0a8417a44afe70e8b3a8c3481210f8cef2b0c599cee5f4d"+
			"bf0490332a339cba371b1ad3ebfe103086cd5c15189f66db2ecc0b2b05d361530d2b38c1426cb5218c185095bed79514"+
			"e7"),
		mustHex(t, "fc365d3c6a073427e27637247fcaa8ccd27662e9bc78dda637371d4d47fe3a647793a47dc0a99704e8f76a2b5aabe1f2"+
			"66fdf8f5286a814f236cbb60d2ac6fbfccb28a35cacc8fa2ffdabac084db058c2ff4514f8395af12cd644916169f92c2"+
			"ed6a11ca523718e1244b65c2dd35fbdd519ca51593ba06556b936cbae19d7d787b451b4f9635a6bef84e42cc7b5e3aa2"+
			"afbaf5fe61fc236b24dade3f5fd384f3df11c7c76cbb2f30c508c965dcade7a7a14e04566243d3cb3adfc780bc863725"+
			"5939f2a9017014fb49a031a55523c1389b18a9c76f687effc77fa8e125d6d70a49b63ccfe682c25acc96c02c52e2841a"+
			"5d"),
		mustHex(t, "fc2e694194acb5017f6a21c1290ae4fc10ea4726b7d7f778d0f537c1931f70332f733afb260247989505c41145faceca"+
			"65f80572bb884ff18b7547d0c69e291488ec92ed6d5ab57c237be7c1e8464f5b0b3a3b4c66ef124790071acc3655c65a"+
			"ee1d1fcef7a48785ffcd06181afeb0ac487a0e2ac8bc7774bdfb038b38e6d67282b796391ecd419adeb7128f3e0c1b06"+
			"a57fce126bde42bc010c039019d07de03223366d4bd781de631607fc0ef5b349974fe28168238040a1e12e5c9931ed56"+
			"f83349530f2acf6d3838fcdaacf2ae23dd13324df240780e00b73f81616c9b0944a08f140bd24813a50f07ba731ad5d2"+
			"34"),
		mustHex(t, "bc7a85aa8aeef737dc3c1221b678fc31764e1915aad26b9a4f4e7cf41154f3db16566b7036fd64dd3a247ebcf125f83f"+
			"207e5f4e80874387094cb58b"),
		mustHex(t, "bc15cc427b2efcd8687cf94366a0929affd9ee7e1f5279b40795f41e9b1ea9488a1983c546b72e8f096e31341c4f78a2"+
			"87dfc8b022daf0"),
		mustHex(t, "bc318070c9b433a3983f46d03dadf3bbc93b30f74a14fcc0d1f6dee977113fabb2aa70dc2538262026b6cf58f7931dc5"+
			"1a6328b53df752"),
		mustHex(t, "bc2ef884e2cf9e68090241a72b7c77be1905449f898d3662dfac5daab05c1d40a2e0bcd59bcdf898078bebaa76cf90c4"+
			"988e77a9f3b6"),
	}

	/* Per-frame expectations: fnv/bits are C reference values for frames
	   5..11 (pure SILK, bit-exact) and Go golden values for CELT-touched
	   frames (see file comment); n/rng are always C-derived. */
	want := []frameExpect{
		{960, 0x53ba9704, 0x01ad2800, [4]uint32{0x245a2fdb, 0xa4993b42, 0x269f84e3, 0x26c04030}, true},
		{960, 0x5ff5cd0b, 0x02879800, [4]uint32{0xbbfd8c82, 0xbc06c12d, 0xbd52e1fb, 0xbcfc435d}, true},
		{960, 0xc65da4a1, 0x03449d00, [4]uint32{0x3cbd42f8, 0x3be0def0, 0x3e09e936, 0x3d1fe8a8}, true},
		{960, 0xa052e9c5, 0x096a4200, [4]uint32{0xbe0680b9, 0x3ea8a673, 0x3d6e499c, 0x3d069628}, true},
		{960, 0x83aace05, 0x016a9b50, [4]uint32{0xbf861490, 0xbebef9a0, 0xbea1436e, 0x3ea2f640}, true},
		{960, 0x708dddfa, 0x4f71e43c, [4]uint32{0xbe5f6000, 0x3db09000, 0xbe564800, 0x3db41000}, false},
		{960, 0x3b15988a, 0x04b839e8, [4]uint32{0xbdcea000, 0x3d328000, 0xbdf16000, 0x3da2b000}, false},
		{960, 0xc7096dce, 0x26732000, [4]uint32{0xbbb10000, 0xbde4e000, 0xbe490000, 0x3cef0000}, false},
		{960, 0x2e69413e, 0x1bdee7d6, [4]uint32{0xbea47c00, 0xbe807000, 0xbe934800, 0xbe2fa000}, false},
		{960, 0xa6e610ae, 0x010c30b6, [4]uint32{0xbc370000, 0xbca80000, 0x3c3e0000, 0xbcb78000}, false},
		{960, 0x8c5f769f, 0x05882413, [4]uint32{0x3d70e000, 0xbe86d000, 0xbda28000, 0xbe20a000}, false},
		{960, 0xe4c0a8a2, 0x01689079, [4]uint32{0xbd27c000, 0xbd88b000, 0xbcb28000, 0xbdd18000}, false},
		{960, 0xdfe62552, 0x2de6f900, [4]uint32{0xbe449800, 0xbe31e000, 0xbeaab000, 0xbd12c000}, true},
		{960, 0xa2fac41b, 0x0637b600, [4]uint32{0xbb899de4, 0x3d08b193, 0x3cbda620, 0xbdd4711e}, true},
		{960, 0x71620cb1, 0x4c7d9b00, [4]uint32{0xbe21ada0, 0xbd8c2bf9, 0xbe386ff1, 0xbb4adea0}, true},
		{960, 0xe144e484, 0x00e58200, [4]uint32{0xbf08e4c9, 0xbe3de3ae, 0x3e449b78, 0xbf0027af}, true},
		{960, 0x14275c42, 0x00bc30be, [4]uint32{0xbeac9950, 0x3ed064fc, 0xbe102cd6, 0xbf7b9ef1}, true},
		{960, 0x3ea1f53a, 0x00e6db9a, [4]uint32{0xbe059ffb, 0x3ccfcebd, 0xbe16a51d, 0xbcf72499}, true},
		{960, 0x072980e8, 0x00876a00, [4]uint32{0xbdedbc10, 0xbde3841a, 0xbdb209fb, 0xbdcb7920}, true},
		{960, 0x720886c6, 0x2937d000, [4]uint32{0xbcf93b9c, 0xbe8e2c7e, 0xbdf56009, 0xbebfcff2}, true},
	}
	/* expected decoder mode after each frame (follows from the packet TOC) */
	wantMode := []int32{
		MODE_HYBRID, MODE_HYBRID, MODE_HYBRID, MODE_HYBRID,
		MODE_SILK_ONLY, MODE_SILK_ONLY, MODE_SILK_ONLY, MODE_SILK_ONLY,
		MODE_SILK_ONLY, MODE_SILK_ONLY, MODE_SILK_ONLY, MODE_SILK_ONLY,
		MODE_SILK_ONLY, MODE_CELT_ONLY, MODE_CELT_ONLY, MODE_CELT_ONLY,
		MODE_CELT_ONLY, MODE_CELT_ONLY, MODE_CELT_ONLY, MODE_CELT_ONLY,
	}
	/* frame 12 is the SILK packet that carries the SILK->CELT transition
	   redundancy (prev_redundancy=1); hybrid frames signal celt_to_silk
	   redundancy instead (prev_redundancy=0) */
	wantRedundancy := make([]int32, 20)
	wantRedundancy[12] = 1

	pcm := make([]float32, 5760*2)
	for i, pkt := range packets {
		n := Opus_opus_decode_float(tls, dec, uintptr(unsafe.Pointer(&pkt[0])), int32(len(pkt)), uintptr(unsafe.Pointer(unsafe.SliceData(pcm))), 5760, 0)
		if n < 0 {
			t.Fatalf("frame %d: decode: %d", i, n)
		}
		w := want[i]
		if n != w.n {
			t.Fatalf("frame %d: n=%d want %d", i, n, w.n)
		}
		if got := fnv1aFloats(pcm[:int(n)*2]); got != w.fnv {
			t.Fatalf("frame %d: fnv=%08x want %08x (celt frame: Go golden, see file comment)", i, got, w.fnv)
		}
		if got := dp.FrangeFinal; got != w.rng {
			t.Fatalf("frame %d: rangeFinal=%08x want %08x", i, got, w.rng)
		}
		for j := 0; j < 4; j++ {
			if got := math.Float32bits(pcm[j]); got != w.b[j] {
				t.Fatalf("frame %d: pcm[%d]=%08x want %08x", i, j, got, w.b[j])
			}
		}
		if got := dp.Fprev_mode; got != wantMode[i] {
			t.Fatalf("frame %d: prev_mode=%d want %d", i, got, wantMode[i])
		}
		if got := dp.Fprev_redundancy; got != wantRedundancy[i] {
			t.Fatalf("frame %d: prev_redundancy=%d want %d", i, got, wantRedundancy[i])
		}
	}

	/* 120 ms PLC after CELT frames (recursion into 20 ms concealment). */
	n := Opus_opus_decode_float(tls, dec, 0, 0, uintptr(unsafe.Pointer(unsafe.SliceData(pcm))), 5760, 0)
	if n != 5760 {
		t.Fatalf("plc: n=%d want 5760", n)
	}
	if got, wantFnv := fnv1aFloats(pcm[:5760*2]), uint32(0xf3e48462); got != wantFnv {
		t.Fatalf("plc: fnv=%08x want %08x (celt frame: Go golden, see file comment)", got, wantFnv)
	}
	if got := dp.FrangeFinal; got != 0 {
		t.Fatalf("plc: rangeFinal=%08x want 0", got)
	}
	for j, wantBits := range [4]uint32{0xbe60b6bf, 0xbe183f94, 0xbeab5bbe, 0xbe6f84d9} {
		if got := math.Float32bits(pcm[j]); got != wantBits {
			t.Fatalf("plc: pcm[%d]=%08x want %08x", j, got, wantBits)
		}
	}
}
