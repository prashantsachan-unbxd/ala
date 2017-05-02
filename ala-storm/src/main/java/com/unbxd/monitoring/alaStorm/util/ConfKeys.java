package com.unbxd.monitoring.alaStorm.util;

/**
 * Created by prashantsachan on 26/04/17.
 */
public class ConfKeys {


    public static final String TOPOLOGY_NAME = "ala-storm-topology";

    //kafka spout
    public static final String KAFKA_SPOUT_ID = "kafka-spout";
    public static final String KAFKA_ZOOKEEPER = "kafka.zookeeper";
    public static final String KAFKA_TOPIC = "kafa.topic";
    public static final String KAFKA_ZKROOT = "kafka.zkRoot";
    public static final String KAFKA_CONSUMERGROUP = "kafka.consumer.group";
    public static final String KAFKA_SPOUT_COUNT = "kafkaspout.count";

    public static final String INFLUX_BOLT_ID = "influx-bolt";
    public static final String INFLUX_BOLT_COUNT = "influxbolt.count";

    // influx db configuration
    public static final String INFLUX_HOSTPORT = "influx.hostport";
    public static final String INFLUX_DBNAME = "influx.dbname";
    public static final String INFLUX_USER  = "influx.user";
    public static final String INFLUX_PASS  = "influx.pass";

}
