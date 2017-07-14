<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github.com/TheThingsNetwork/api/gateway/gateway.proto

namespace Gateway;

/**
 * Protobuf enum <code>gateway.LocationMetadata.LocationSource</code>
 */
class LocationMetadata_LocationSource
{
    /**
     * <pre>
     * The source of the location is not known or not set
     * </pre>
     *
     * <code>UNKNOWN = 0;</code>
     */
    const UNKNOWN = 0;
    /**
     * <pre>
     * The location is determined by GPS
     * </pre>
     *
     * <code>GPS = 1;</code>
     */
    const GPS = 1;
    /**
     * <pre>
     * The location is fixed by configuration
     * </pre>
     *
     * <code>CONFIG = 2;</code>
     */
    const CONFIG = 2;
    /**
     * <pre>
     * The location is set in and updated from a registry
     * </pre>
     *
     * <code>REGISTRY = 3;</code>
     */
    const REGISTRY = 3;
    /**
     * <pre>
     * The location is estimated with IP Geolocation
     * </pre>
     *
     * <code>IP_GEOLOCATION = 4;</code>
     */
    const IP_GEOLOCATION = 4;
}

