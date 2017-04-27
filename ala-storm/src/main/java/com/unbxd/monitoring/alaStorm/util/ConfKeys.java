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

    //sink bolt
    //public static final String SINK_TYPE_BOLT_ID = "sink-type-bolt";
    //public static final String SINK_BOLT_COUNT = "sinkbolt.count";

    public static final String INFLUX_BOLT_ID = "influx-bolt";
    public static final String INFLUX_BOLT_COUNT = "influxbolt.count";
}
