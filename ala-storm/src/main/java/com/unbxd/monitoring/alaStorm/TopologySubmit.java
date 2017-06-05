package com.unbxd.monitoring.alaStorm;

import org.apache.storm.Config;
import org.apache.storm.StormSubmitter;
import org.apache.storm.topology.TopologyBuilder;

/**
 * Created by prashantsachan on 31/05/17.
 */
public class TopologySubmit {
    public static String topologyName = "MetricCollectTopology";
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
        StormSubmitter.submitTopology(topologyName, conf, builder.createTopology());

    }
}
