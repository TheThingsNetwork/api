<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/gateway/gateway.proto

namespace GPBMetadata\GithubCom\TheThingsNetwork\Api\Gateway;

class Gateway
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\GithubCom\Gogo\Protobuf\Gogoproto\Gogo::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0aea0e0a356769746875622e636f6d2f5468655468696e67734e6574776f" .
            "726b2f6170692f676174657761792f676174657761792e70726f746f1207" .
            "6761746577617922f9010a104c6f636174696f6e4d65746164617461120c" .
            "0a0474696d6518012001280312100a086c61746974756465180220012802" .
            "12110a096c6f6e67697475646518032001280212100a08616c7469747564" .
            "6518042001280512100a08616363757261637918052001280512380a0673" .
            "6f7572636518062001280e32282e676174657761792e4c6f636174696f6e" .
            "4d657461646174612e4c6f636174696f6e536f7572636522540a0e4c6f63" .
            "6174696f6e536f75726365120b0a07554e4b4e4f574e100012070a034750" .
            "531001120a0a06434f4e4649471002120c0a085245474953545259100312" .
            "120a0e49505f47454f4c4f434154494f4e100422c5040a0a52784d657461" .
            "6461746112210a0a676174657761795f6964180120012809420de2de1f09" .
            "47617465776179494412170a0f676174657761795f747275737465641802" .
            "2001280812110a0974696d657374616d70180b2001280d120c0a0474696d" .
            "65180c2001280312160a0e656e637279707465645f74696d65180d200128" .
            "0c12100a0872665f636861696e18152001280d120f0a076368616e6e656c" .
            "18162001280d122d0a08616e74656e6e6173181e2003280b321b2e676174" .
            "657761792e52784d657461646174612e416e74656e6e6112110a09667265" .
            "7175656e6379181f2001280412160a04727373691820200128024208e2de" .
            "1f045253534912140a03736e721821200128024207e2de1f03534e52122b" .
            "0a086c6f636174696f6e18292001280b32192e676174657761792e4c6f63" .
            "6174696f6e4d657461646174611a81020a07416e74656e6e61120f0a0761" .
            "6e74656e6e6118012001280d120f0a076368616e6e656c18022001280d12" .
            "160a04727373691803200128024208e2de1f045253534912250a0c636861" .
            "6e6e656c5f72737369180520012802420fe2de1f0b4368616e6e656c5253" .
            "5349123a0a17727373695f7374616e646172645f646576696174696f6e18" .
            "06200128024219e2de1f15525353495374616e6461726444657669617469" .
            "6f6e12180a106672657175656e63795f6f66667365741807200128031214" .
            "0a03736e721804200128024207e2de1f03534e5212160a0e656e63727970" .
            "7465645f74696d65180a2001280c12110a0966696e655f74696d65180b20" .
            "0128032295010a0f5478436f6e66696775726174696f6e12110a0974696d" .
            "657374616d70180b2001280d12100a0872665f636861696e18152001280d" .
            "12110a096672657175656e6379181620012804120d0a05706f7765721817" .
            "20012805121e0a16706f6c6172697a6174696f6e5f696e76657273696f6e" .
            "181f20012808121b0a136672657175656e63795f646576696174696f6e18" .
            "202001280d22c3050a0653746174757312110a0974696d657374616d7018" .
            "012001280d120c0a0474696d6518022001280312170a0f67617465776179" .
            "5f7472757374656418032001280812110a09626f6f745f74696d65180420" .
            "01280312120a026970180b200328094206e2de1f02495012100a08706c61" .
            "74666f726d180c2001280912150a0d636f6e746163745f656d61696c180d" .
            "2001280912130a0b6465736372697074696f6e180e2001280912160a0e66" .
            "72657175656e63795f706c616e180f20012809120e0a0662726964676518" .
            "1020012809120e0a06726f7574657218112001280912160a046670676118" .
            "122001280d4208e2de1f044650474112140a0364737018132001280d4207" .
            "e2de1f0344535012140a0368616c1814200128094207e2de1f0348414c12" .
            "2b0a086c6f636174696f6e18152001280b32192e676174657761792e4c6f" .
            "636174696f6e4d6574616461746112140a03727474181f2001280d4207e2" .
            "de1f03525454120d0a0572785f696e18292001280d120d0a0572785f6f6b" .
            "182a2001280d120d0a0574785f696e182b2001280d120d0a0574785f6f6b" .
            "182c2001280d120d0a056c6d5f6f6b182d2001280d120d0a056c6d5f7374" .
            "182e2001280d120d0a056c6d5f6e77182f2001280d12170a056c5f707073" .
            "18302001280d4208e2de1f044c505053122d0a026f7318332001280b3219" .
            "2e676174657761792e5374617475732e4f534d6574726963734206e2de1f" .
            "024f5312100a086d657373616765731834200328091a97010a094f534d65" .
            "7472696373120e0a066c6f61645f31180120012802120e0a066c6f61645f" .
            "35180220012802120f0a076c6f61645f313518032001280212290a0e6370" .
            "755f70657263656e74616765180b200128024211e2de1f0d435055506572" .
            "63656e7461676512190a116d656d6f72795f70657263656e746167651815" .
            "2001280212130a0b74656d7065726174757265181f20012802427e0a206f" .
            "72672e7468657468696e67736e6574776f726b2e6170692e676174657761" .
            "79420c4761746577617950726f746f50015a276769746875622e636f6d2f" .
            "5468655468696e67734e6574776f726b2f6170692f67617465776179aa02" .
            "1c5468655468696e67734e6574776f726b2e4150492e47617465776179f8" .
            "e11e00620670726f746f33"
        ));

        static::$is_initialized = true;
    }
}

