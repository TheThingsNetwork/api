<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/handler/handler.proto

namespace Handler;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>handler.DeviceIdentifier</code>
 */
class DeviceIdentifier extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>string app_id = 1;</code>
     */
    private $app_id = '';
    /**
     * <code>string dev_id = 2;</code>
     */
    private $dev_id = '';

    public function __construct() {
        \GPBMetadata\GithubCom\TheThingsNetwork\Api\Handler\Handler::initOnce();
        parent::__construct();
    }

    /**
     * <code>string app_id = 1;</code>
     */
    public function getAppId()
    {
        return $this->app_id;
    }

    /**
     * <code>string app_id = 1;</code>
     */
    public function setAppId($var)
    {
        GPBUtil::checkString($var, True);
        $this->app_id = $var;
    }

    /**
     * <code>string dev_id = 2;</code>
     */
    public function getDevId()
    {
        return $this->dev_id;
    }

    /**
     * <code>string dev_id = 2;</code>
     */
    public function setDevId($var)
    {
        GPBUtil::checkString($var, True);
        $this->dev_id = $var;
    }

}

