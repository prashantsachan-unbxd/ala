package com.unbxd.monitoring.alaStorm;

import backtype.storm.Config;
import backtype.storm.topology.TopologyBuilder;
import backtype.storm.utils.Utils;
import backtype.storm.LocalCluster;
import com.unbxd.monitoring.alaStorm.bolt.BoltFactory;
import com.unbxd.monitoring.alaStorm.bolt.InfluxBolt;
import com.unbxd.monitoring.alaStorm.spout.SpoutFactory;
import com.unbxd.monitoring.alaStorm.util.ConfKeys;
import storm.kafka.KafkaSpout;
/**
 * Created by prashantsachan on 31/05/17.
 */
public class TopologyRun {

    public static void main(String[] args) throws Exception {
        String configFile = null;
        if (args.length == 0) {
            System.out.println("Missing input : config file location, using default");

        } else {
            configFile = args[0];
        }

        Config conf = new Config();
        //conf.setDebug(true);
        conf.setNumWorkers(2);
        Topology ingestionTopology = new Topology(configFile, conf);

        TopologyBuilder builder = ingestionTopology.buildTopology();
        //StormSubmitter.submitTopology(topologyName, conf, builder.createTopology());
        LocalCluster cluster = new LocalCluster();
        cluster.submitTopology("test", conf, builder.createTopology());
        Utils.sleep(1000*1000*1000);
        cluster.killTopology("test");
        cluster.shutdown();

    }
}
