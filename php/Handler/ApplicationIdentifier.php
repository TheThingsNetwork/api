<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/handler/handler.proto

namespace Handler;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>handler.ApplicationIdentifier</code>
 */
class ApplicationIdentifier extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string app_id = 1 [(.gogoproto.customname) = "AppID"];</code>
     */
    private $app_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $app_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\GithubCom\TheThingsNetwork\Api\Handler\Handler::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string app_id = 1 [(.gogoproto.customname) = "AppID"];</code>
     * @return string
     */
    public function getAppId()
    {
        return $this->app_id;
    }

    /**
     * Generated from protobuf field <code>string app_id = 1 [(.gogoproto.customname) = "AppID"];</code>
     * @param string $var
     * @return $this
     */
    public function setAppId($var)
    {
        GPBUtil::checkString($var, True);
        $this->app_id = $var;

        return $this;
    }

}

