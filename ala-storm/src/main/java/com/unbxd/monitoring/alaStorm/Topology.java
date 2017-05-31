package com.unbxd.monitoring.alaStorm;

import backtype.storm.Config;
import backtype.storm.topology.TopologyBuilder;
import backtype.storm.utils.Utils;
import com.unbxd.monitoring.alaStorm.bolt.BoltFactory;
import com.unbxd.monitoring.alaStorm.bolt.InfluxBolt;
import com.unbxd.monitoring.alaStorm.spout.SpoutFactory;
import com.unbxd.monitoring.alaStorm.util.ConfKeys;
import storm.kafka.KafkaSpout;

import java.io.FileInputStream;
import java.net.URL;
import java.util.Properties;

/**
 * Created by prashantsachan on 26/04/17.
 */
public class Topology {

    public Properties properties;
    public BoltFactory boltBuilder;
    public SpoutFactory spoutBuilder;
    public Config config;


    public Topology(String configFile, Config conf) throws Exception {
        properties = new Properties();
        config = conf;
        if(configFile == null){
            ClassLoader classLoader = getClass().getClassLoader();
            URL filename = classLoader.getResource("application.properties");
            System.out.print("props file: "+filename);
            properties.load(classLoader.getResourceAsStream("application.properties"));
        }else {
            properties.load(new FileInputStream(configFile));
        }

        try {

            boltBuilder = new BoltFactory(properties, config);
            spoutBuilder = new SpoutFactory(properties);
        } catch (Exception ex) {
            ex.printStackTrace();
            System.exit(0);
        }
    }

    public TopologyBuilder buildTopology() throws Exception {
        TopologyBuilder builder = new TopologyBuilder();

        KafkaSpout kafkaSpout = spoutBuilder.buildKafkaSpout();
        InfluxBolt influxBolt = boltBuilder.buildInfluxBolt();


        //set the kafkaSpout to topology
        //parallelism-hint for kafkaSpout - defines number of executors/threads to be spawn per container
        int kafkaSpoutCount = Integer.parseInt(properties.getProperty(ConfKeys
                .KAFKA_SPOUT_COUNT));
        builder.setSpout(ConfKeys.KAFKA_SPOUT_ID, kafkaSpout, kafkaSpoutCount);

        int influxBoltCount = Integer.parseInt(properties.getProperty(ConfKeys
                .INFLUX_BOLT_COUNT));
        builder.setBolt(ConfKeys.INFLUX_BOLT_ID,influxBolt,influxBoltCount)
                .shuffleGrouping(ConfKeys.KAFKA_SPOUT_ID);
        String topologyName = properties.getProperty(ConfKeys.TOPOLOGY_NAME);
        return builder;

    }

}