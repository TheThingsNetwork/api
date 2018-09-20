<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/broker/broker.proto

namespace Broker;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>broker.ApplicationHandlerRegistration</code>
 */
class ApplicationHandlerRegistration extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string app_id = 1 [(.gogoproto.customname) = "AppID"];</code>
     */
    private $app_id = '';
    /**
     * Generated from protobuf field <code>string handler_id = 2 [(.gogoproto.customname) = "HandlerID"];</code>
     */
    private $handler_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $app_id
     *     @type string $handler_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\GithubCom\TheThingsNetwork\Api\Broker\Broker::initOnce();
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

    /**
     * Generated from protobuf field <code>string handler_id = 2 [(.gogoproto.customname) = "HandlerID"];</code>
     * @return string
     */
    public function getHandlerId()
    {
        return $this->handler_id;
    }

    /**
     * Generated from protobuf field <code>string handler_id = 2 [(.gogoproto.customname) = "HandlerID"];</code>
     * @param string $var
     * @return $this
     */
    public function setHandlerId($var)
    {
        GPBUtil::checkString($var, True);
        $this->handler_id = $var;

        return $this;
    }

}

