<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/discovery/discovery.proto

namespace GPBMetadata\GithubCom\TheThingsNetwork\Api\Discovery;

class Discovery
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\GPBEmpty::initOnce();
        \GPBMetadata\Google\Api\Annotations::initOnce();
        \GPBMetadata\GithubCom\Gogo\Protobuf\Gogoproto\Gogo::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0ad50e0a396769746875622e636f6d2f5468655468696e67734e6574776f" .
            "726b2f6170692f646973636f766572792f646973636f766572792e70726f" .
            "746f1209646973636f766572791a1c676f6f676c652f6170692f616e6e6f" .
            "746174696f6e732e70726f746f1a2d6769746875622e636f6d2f676f676f" .
            "2f70726f746f6275662f676f676f70726f746f2f676f676f2e70726f746f" .
            "2292010a084d6574616461746112230a0a676174657761795f6964180a20" .
            "012809420de2de1f09476174657761794944480012190a0f6465765f6164" .
            "64725f70726566697818142001280c4800121b0a066170705f6964181e20" .
            "0128094209e2de1f0541707049444800121d0a076170705f657569181f20" .
            "01280c420ae2de1f064170704555494800420a0a086d6574616461746122" .
            "a9020a0c416e6e6f756e63656d656e7412120a0269641801200128094206" .
            "e2de1f02494412140a0c736572766963655f6e616d651802200128091217" .
            "0a0f736572766963655f76657273696f6e18032001280912130a0b646573" .
            "6372697074696f6e180420012809120b0a0375726c180520012809120e0a" .
            "067075626c696318062001280812130a0b6e65745f61646472657373180b" .
            "2001280912120a0a7075626c69635f6b6579180c2001280912130a0b6365" .
            "727469666963617465180d2001280912130a0b6170695f61646472657373" .
            "180e2001280912140a0c6d7174745f61646472657373180f200128091214" .
            "0a0c616d71705f6164647265737318102001280912250a086d6574616461" .
            "746118162003280b32132e646973636f766572792e4d6574616461746122" .
            "290a11476574536572766963655265717565737412140a0c736572766963" .
            "655f6e616d6518012001280922360a0a4765745265717565737412120a02" .
            "69641801200128094206e2de1f02494412140a0c736572766963655f6e61" .
            "6d6518022001280922620a0f4d657461646174615265717565737412120a" .
            "0269641801200128094206e2de1f02494412140a0c736572766963655f6e" .
            "616d6518022001280912250a086d65746164617461180c2001280b32132e" .
            "646973636f766572792e4d6574616461746122420a15416e6e6f756e6365" .
            "6d656e7473526573706f6e736512290a0873657276696365731801200328" .
            "0b32172e646973636f766572792e416e6e6f756e63656d656e74222e0a11" .
            "476574427941707049445265717565737412190a066170705f6964181e20" .
            "0128094209e2de1f054170704944223a0a15476574427947617465776179" .
            "49445265717565737412210a0a676174657761795f6964181e2001280942" .
            "0de2de1f0947617465776179494422660a12476574427941707045554952" .
            "65717565737412500a076170705f657569181f2001280c423fe2de1f0641" .
            "7070455549dade1f316769746875622e636f6d2f5468655468696e67734e" .
            "6574776f726b2f74746e2f636f72652f74797065732e41707045554932fe" .
            "040a09446973636f76657279123b0a08416e6e6f756e636512172e646973" .
            "636f766572792e416e6e6f756e63656d656e741a162e676f6f676c652e70" .
            "726f746f6275662e456d707479126f0a06476574416c6c121c2e64697363" .
            "6f766572792e47657453657276696365526571756573741a202e64697363" .
            "6f766572792e416e6e6f756e63656d656e7473526573706f6e7365222582" .
            "d3e493021f121d2f616e6e6f756e63656d656e74732f7b73657276696365" .
            "5f6e616d657d12610a0347657412152e646973636f766572792e47657452" .
            "6571756573741a172e646973636f766572792e416e6e6f756e63656d656e" .
            "74222a82d3e493022412222f616e6e6f756e63656d656e74732f7b736572" .
            "766963655f6e616d657d2f7b69647d12410a0b4164644d65746164617461" .
            "121a2e646973636f766572792e4d65746164617461526571756573741a16" .
            "2e676f6f676c652e70726f746f6275662e456d70747912440a0e44656c65" .
            "74654d65746164617461121a2e646973636f766572792e4d657461646174" .
            "61526571756573741a162e676f6f676c652e70726f746f6275662e456d70" .
            "747912430a0a47657442794170704944121c2e646973636f766572792e47" .
            "657442794170704944526571756573741a172e646973636f766572792e41" .
            "6e6e6f756e63656d656e74124b0a0e476574427947617465776179494412" .
            "202e646973636f766572792e476574427947617465776179494452657175" .
            "6573741a172e646973636f766572792e416e6e6f756e63656d656e741245" .
            "0a0b4765744279417070455549121d2e646973636f766572792e47657442" .
            "79417070455549526571756573741a172e646973636f766572792e416e6e" .
            "6f756e63656d656e7432120a10446973636f766572794d616e6167657242" .
            "82010a226f72672e7468657468696e67736e6574776f726b2e6170692e64" .
            "6973636f76657279420e446973636f7665727950726f746f50015a296769" .
            "746875622e636f6d2f5468655468696e67734e6574776f726b2f6170692f" .
            "646973636f76657279aa021e5468655468696e67734e6574776f726b2e41" .
            "50492e446973636f76657279620670726f746f33"
        ));

        static::$is_initialized = true;
    }
}

