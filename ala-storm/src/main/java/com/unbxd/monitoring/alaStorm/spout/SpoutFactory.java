package com.unbxd.monitoring.alaStorm.spout;

import com.unbxd.monitoring.alaStorm.util.ConfKeys;
import org.apache.storm.kafka.*;
import org.apache.storm.spout.SchemeAsMultiScheme;

import java.util.Properties;

/**
 * Created by prashantsachan on 26/04/17.
 */
public class SpoutFactory {

    public Properties configs = null;

    public SpoutFactory(Properties configs) {
        this.configs = configs;
    }
    public KafkaSpout buildKafkaSpout() {
        BrokerHosts hosts = new ZkHosts(configs.getProperty(ConfKeys.KAFKA_ZOOKEEPER));
        String topic = configs.getProperty(ConfKeys.KAFKA_TOPIC);
        String zkRoot = configs.getProperty(ConfKeys.KAFKA_ZKROOT);
        String groupId = configs.getProperty(ConfKeys.KAFKA_CONSUMERGROUP);
        SpoutConfig spoutConfig = new SpoutConfig(hosts, topic, zkRoot+topic, groupId);
        spoutConfig.scheme = new SchemeAsMultiScheme(new StringScheme());
        KafkaSpout kafkaSpout = new KafkaSpout(spoutConfig);
        return kafkaSpout;
    }
}
