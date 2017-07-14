<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/networkserver/networkserver.proto

namespace Networkserver;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>networkserver.DevicesResponse</code>
 */
class DevicesResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>repeated .lorawan.Device results = 1;</code>
     */
    private $results;

    public function __construct() {
        \GPBMetadata\GithubCom\TheThingsNetwork\Api\Networkserver\Networkserver::initOnce();
        parent::__construct();
    }

    /**
     * <code>repeated .lorawan.Device results = 1;</code>
     */
    public function getResults()
    {
        return $this->results;
    }

    /**
     * <code>repeated .lorawan.Device results = 1;</code>
     */
    public function setResults(&$var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Lorawan\Device::class);
        $this->results = $arr;
    }

}

