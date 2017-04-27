package com.unbxd.monitoring.alaStorm.bolt;

import java.util.Properties;

/**
 * Created by prashantsachan on 26/04/17.
 */
public class BoltFactory {
    Properties properties;
    public BoltFactory(Properties props){
        this.properties = props;
    }

    public InfluxBolt buildInfluxBolt() {
        return new InfluxBolt();
    }
}
