<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/protocol/protocol.proto

namespace Protocol;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>protocol.RxMetadata</code>
 */
class RxMetadata extends \Google\Protobuf\Internal\Message
{
    protected $protocol;

    public function __construct() {
        \GPBMetadata\GithubCom\TheThingsNetwork\Api\Protocol\Protocol::initOnce();
        parent::__construct();
    }

    /**
     * <code>.lorawan.Metadata lorawan = 1;</code>
     */
    public function getLorawan()
    {
        return $this->readOneof(1);
    }

    /**
     * <code>.lorawan.Metadata lorawan = 1;</code>
     */
    public function setLorawan(&$var)
    {
        GPBUtil::checkMessage($var, \Lorawan\Metadata::class);
        $this->writeOneof(1, $var);
    }

    public function getProtocol()
    {
        return $this->whichOneof("protocol");
    }

}

