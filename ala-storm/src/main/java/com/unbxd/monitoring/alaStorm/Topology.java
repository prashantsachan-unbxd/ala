package com.unbxd.monitoring.alaStorm;

import backtype.storm.Config;
import backtype.storm.LocalCluster;
import backtype.storm.StormSubmitter;
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

    public Properties configs;
    public BoltFactory boltBuilder;
    public SpoutFactory spoutBuilder;


    public Topology(String configFile) throws Exception {
        configs = new Properties();
        if(configFile == null){
            ClassLoader classLoader = getClass().getClassLoader();
            URL filename = classLoader.getResource("/application.properties");
            System.out.print("props file: "+filename);
            configs.load(classLoader.getResourceAsStream("/application.properties"));
        }else {
            configs.load(new FileInputStream(configFile));
        }

        try {

            boltBuilder = new BoltFactory(configs);
            spoutBuilder = new SpoutFactory(configs);
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
        int kafkaSpoutCount = Integer.parseInt(configs.getProperty(ConfKeys
                .KAFKA_SPOUT_COUNT));
        builder.setSpout(ConfKeys.KAFKA_SPOUT_ID, kafkaSpout, kafkaSpoutCount);

        int influxBoltCount = Integer.parseInt(configs.getProperty(ConfKeys
                .INFLUX_BOLT_COUNT));
        builder.setBolt(ConfKeys.INFLUX_BOLT_ID,influxBolt,influxBoltCount)
                .shuffleGrouping(ConfKeys.KAFKA_SPOUT_ID);
        String topologyName = configs.getProperty(ConfKeys.TOPOLOGY_NAME);
        return builder;

    }

    public static void main(String[] args) throws Exception {
        String configFile = null;
        if (args.length == 0) {
            System.out.println("Missing input : config file location, using default");

        } else {
            configFile = args[0];
        }

        Topology ingestionTopology = new Topology(configFile);

        Config conf = new Config();
        conf.setDebug(true);
        conf.setNumWorkers(2);
        TopologyBuilder builder = ingestionTopology.buildTopology();
        //StormSubmitter.submitTopology(topologyName, conf, builder.createTopology());
        LocalCluster cluster = new LocalCluster();
        cluster.submitTopology("test", conf, builder.createTopology());
        Utils.sleep(1000000);
        cluster.killTopology("test");
        cluster.shutdown();

    }
}