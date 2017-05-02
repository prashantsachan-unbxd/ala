package com.unbxd.monitoring.alaStorm.bolt;

import backtype.storm.Config;
import com.unbxd.monitoring.alaStorm.util.ConfKeys;

import java.util.Properties;

/**
 * Created by prashantsachan on 26/04/17.
 */
public class BoltFactory {
    Properties properties;
    Config config;
    public BoltFactory(Properties props, Config conf){
        this.properties = props;
        config = conf;
    }

    public InfluxBolt buildInfluxBolt() {
        String influxHost = properties.getProperty(ConfKeys.INFLUX_HOSTPORT);
        String influxDbName = properties.getProperty(ConfKeys.INFLUX_DBNAME);
        String influxUser = properties.getProperty(ConfKeys.INFLUX_USER);
        String influxPass = properties.getProperty(ConfKeys.INFLUX_PASS);
        config.put(ConfKeys.INFLUX_HOSTPORT, influxHost.trim());
        config.put(ConfKeys.INFLUX_DBNAME, influxDbName.trim());
        config.put(ConfKeys.INFLUX_USER, influxUser.trim());
        config.put(ConfKeys.INFLUX_PASS, influxPass.trim());
        return new InfluxBolt();
    }
}
