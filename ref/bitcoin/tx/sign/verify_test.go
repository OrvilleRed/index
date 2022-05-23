package sign_test

import (
	"encoding/hex"
	"github.com/jchavannes/btcd/wire"
	"github.com/jchavannes/jgo/jerr"
	"github.com/memocash/index/ref/bitcoin/memo"
	"github.com/memocash/index/ref/bitcoin/tx/sign"
	"testing"
)

var (
	VerifyTx1Bytes, _       = hex.DecodeString("010000000275c074802f97438576bb0f7f416f2fd74eef5b19937a067ee5f4b0f5af859065000000006b48304502210082e0e542997a149f05516adf8dd36a44bfd98b3fe80cc1341962fb4b64e9441602203ad75306f4e721d43c3a5b45a6b34ca8deff835dd058e2dd1460ed2ebfbf06af412103b99faaa1dbfdfe9682d6a2584ed41389840177ceba4608c8b1ea8a71facc5a35ffffffff272762c798b6d6ad150dd64b12cebec9875813a9133cd66879897feb5ddab95b010000006a47304402206c9471f86152dab0746d617449ce2eda66e48d23c1baa19bb114a2e7466b1b01022079400dabf66a2620f5e6153a618615ada4cae942242ec0ab3e5bba5179823095c32102159db18d5c36547b0d9f257be56ce3284ca9c9a2a7e667f35feebdac6b86a792ffffffff050000000000000000406a04534c500001010453454e44207f8889682d57369ed0e32336f8b7e0ffec625a35cca183f4e81fde4e71a538a108000000000000000008000000000098968050622500000000001976a914a4371ef3221bb1e39d2021ed7b456ea7bbbd4a1588ac22020000000000001976a914c098c40db5f86b8d9ad3691ffe9060e17b2b496788ac8e8f0000000000001976a9147245bf3d5f5c5e6f370445601a72b3809f798f4688ac9fcfa301000000001976a914c098c40db5f86b8d9ad3691ffe9060e17b2b496788ac00000000")
	VerifyTx1Child0Bytes, _ = hex.DecodeString("020000000132149f2349ec4d7bb2b657d63ac0a357c7bcb68b9af71e09e84584e9b278914f010000006b483045022100d61a1d5af70e8b7c006b668b8573adb2bedb41263329108145f96c178382f95902203a0b00e19945d71ab4162359ce8077cfdfc68a401c8be8d0ca13bc1bcaf53bd8412102d91855d0d36c1baaa71fbdd05e462299a30f8f40565944923f284871964b8a9fffffffff0280c3c901000000001976a914c098c40db5f86b8d9ad3691ffe9060e17b2b496788acc1ed5c00000000001976a914d01d7a2a20968dd68869e6894c771ea9707aeb5088ac00000000")
	VerifyTx1Child1Bytes, _ = hex.DecodeString("01000000027cc273bab50c7a5e030d969e63945395a59fc535ebd4426d4d0294c72546fec6020000006a47304402200df8d6dfda58498d36c0853cc7cda6fe086c1308164141fe5549bbb829dff06d02204ec8929227c8f2b7d1d6f809798394b166712e6d958c976d697a3fd567155746412102159db18d5c36547b0d9f257be56ce3284ca9c9a2a7e667f35feebdac6b86a792ffffffff0c0b4d325574b3377f3abc5f06496b37d8bff194d4d27263b8afd02e337d9531010000006a47304402200b306b735e764a4ed5e3f3183caff31da3124f16a9e9e6db1eb4e655851259d602206135c25fd2692cc46832747e1f950b4911d7f097b5512abc6d08aadb25c75f1a412102159db18d5c36547b0d9f257be56ce3284ca9c9a2a7e667f35feebdac6b86a792ffffffff040000000000000000406a04534c500001010453454e44207f8889682d57369ed0e32336f8b7e0ffec625a35cca183f4e81fde4e71a538a108000000000098968008000000000330586022020000000000001976a914a4371ef3221bb1e39d2021ed7b456ea7bbbd4a1588ac22020000000000001976a914a4371ef3221bb1e39d2021ed7b456ea7bbbd4a1588ac6bbe0000000000001976a914a4371ef3221bb1e39d2021ed7b456ea7bbbd4a1588ac00000000")
	VerifyTx1, _       = memo.GetMsgFromRaw(VerifyTx1Bytes)
	VerifyTx1Child0, _ = memo.GetMsgFromRaw(VerifyTx1Child0Bytes)
	VerifyTx1Child1, _ = memo.GetMsgFromRaw(VerifyTx1Child1Bytes)

	VerifyTx2Bytes, _       = hex.DecodeString("020000000132149f2349ec4d7bb2b657d63ac0a357c7bcb68b9af71e09e84584e9b278914f010000006b483045022100d61a1d5af70e8b7c006b668b8573adb2bedb41263329108145f96c178382f95902203a0b00e19945d71ab4162359ce8077cfdfc68a401c8be8d0ca13bc1bcaf53bd8412102d91855d0d36c1baaa71fbdd05e462299a30f8f40565944923f284871964b8a9fffffffff0280c3c901000000001976a914c098c40db5f86b8d9ad3691ffe9060e17b2b496788acc1ed5c00000000001976a914d01d7a2a20968dd68869e6894c771ea9707aeb5088ac00000000")
	VerifyTx2Child0Bytes, _ = hex.DecodeString("0200000001142dc856ecf19463c340699f796f289a8bea91ba4def150c18acb2e7b5d4404b010000006a47304402204c308a31ee4170fd6189047883a787418513684581b98c75281a8424e899759c02200ad03e984bf8e6e8455389c639f7f24531474ee1b982f0af3ec92e55e31f279d412102d91855d0d36c1baaa71fbdd05e462299a30f8f40565944923f284871964b8a9fffffffff0240420f00000000001976a914c098c40db5f86b8d9ad3691ffe9060e17b2b496788ac23b22602000000001976a914d01d7a2a20968dd68869e6894c771ea9707aeb5088ac00000000")
	VerifyTx2, _            = memo.GetMsgFromRaw(VerifyTx2Bytes)
	VerifyTx2Child0, _      = memo.GetMsgFromRaw(VerifyTx2Child0Bytes)
)

type VerifyTest struct {
	Tx       *wire.MsgTx
	InputTxs []*wire.MsgTx
	Error    string
}

func (tst VerifyTest) Test(t *testing.T) {
	err := sign.Verify(tst.Tx, tst.InputTxs)
	if (err != nil && tst.Error == "") || (err == nil && tst.Error != "") ||
		(tst.Error != "" && !jerr.HasError(err, tst.Error)) {
		t.Error(jerr.Newf("VerifyTest does not match expected: %s - %s", tst.Error, err))
	}
}

func TestVerifySimple(t *testing.T) {
	VerifyTest{
		Tx: VerifyTx2,
		InputTxs: []*wire.MsgTx{
			VerifyTx2Child0,
		},
	}.Test(t)
}

func TestVerifyFail(t *testing.T) {
	VerifyTest{
		Tx: VerifyTx1,
		InputTxs: []*wire.MsgTx{
			VerifyTx1Child0,
			VerifyTx1Child1,
		},
		Error: "signature not empty on failed checksig",
	}.Test(t)
}

func TestVerifySignature(t *testing.T) {
	err := sign.VerifySignature(VerifyTx2Child0.TxOut[1].PkScript, VerifyTx2, 0, VerifyTx2Child0.TxOut[1].Value)
	if err != nil {
		t.Error(jerr.Get("error verifying signature for test", err))
	}
}
