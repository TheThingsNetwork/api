<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/api.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>api.SystemStats.MemoryStats</code>
 */
class SystemStats_MemoryStats extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>uint64 total = 1;</code>
     */
    private $total = 0;
    /**
     * <code>uint64 available = 2;</code>
     */
    private $available = 0;
    /**
     * <code>uint64 used = 3;</code>
     */
    private $used = 0;

    public function __construct() {
        \GPBMetadata\GithubCom\TheThingsNetwork\Api\Api::initOnce();
        parent::__construct();
    }

    /**
     * <code>uint64 total = 1;</code>
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * <code>uint64 total = 1;</code>
     */
    public function setTotal($var)
    {
        GPBUtil::checkUint64($var);
        $this->total = $var;
    }

    /**
     * <code>uint64 available = 2;</code>
     */
    public function getAvailable()
    {
        return $this->available;
    }

    /**
     * <code>uint64 available = 2;</code>
     */
    public function setAvailable($var)
    {
        GPBUtil::checkUint64($var);
        $this->available = $var;
    }

    /**
     * <code>uint64 used = 3;</code>
     */
    public function getUsed()
    {
        return $this->used;
    }

    /**
     * <code>uint64 used = 3;</code>
     */
    public function setUsed($var)
    {
        GPBUtil::checkUint64($var);
        $this->used = $var;
    }

}

