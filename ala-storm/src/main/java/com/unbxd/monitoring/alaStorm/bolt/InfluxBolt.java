package com.unbxd.monitoring.alaStorm.bolt;

import backtype.storm.task.OutputCollector;
import backtype.storm.task.TopologyContext;
import backtype.storm.topology.OutputFieldsDeclarer;
import backtype.storm.topology.base.BaseRichBolt;
import backtype.storm.tuple.Tuple;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.unbxd.monitoring.alaStorm.util.ConfKeys;
import okhttp3.OkHttpClient;
import org.apache.commons.lang.exception.ExceptionUtils;
import org.influxdb.InfluxDB;
import org.influxdb.InfluxDBFactory;
import org.influxdb.dto.Point;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Map;
import java.util.concurrent.TimeUnit;

/**
 * Created by prashantsachan on 26/04/17.
 */
public class InfluxBolt extends BaseRichBolt {
    Logger logger = LoggerFactory.getLogger(InfluxBolt.class);
    public static final String measureName = "availability";
    public static final String RET_POLICY = "default";
    private OutputCollector collector;
    String dbName;
    InfluxDB influxDB;
    ObjectMapper jsonMapper;

    public InfluxBolt() {
        jsonMapper = new ObjectMapper();
    }

    public void prepare(Map map, TopologyContext topologyContext, OutputCollector outputCollector) {
        this.collector = outputCollector;
        OkHttpClient.Builder builder = new OkHttpClient.Builder().readTimeout(10,
                TimeUnit.SECONDS).connectTimeout(10, TimeUnit.SECONDS);
        String hostPort = (String) map.get(ConfKeys.INFLUX_HOSTPORT);
        String user = (String) map.get(ConfKeys.INFLUX_USER);
        String pass = (String) map.get(ConfKeys.INFLUX_PASS);
        this.influxDB = InfluxDBFactory.connect(hostPort, user, pass, builder);
        this.dbName = (String) map.get(ConfKeys.INFLUX_DBNAME);
    }

    public void execute(Tuple tuple) {
        String val = tuple.getString(0);
        System.out.println("received : " + val);
        try {
            Map<String, Object> data = jsonMapper.readValue(val, Map.class);
//            Map<String, String> apiData = (Map<String, String>) data.get("Api");
            long t = toLong((String) data.get("Timestamp"));
            Point p = Point.measurement(measureName)
                    .time(t, TimeUnit.MILLISECONDS)
                    .addField("url", (String) ((Map) data.get("Api")).get("url"))
                    .addField("status", toInt((String) data.get("Status")))
                    .build();
            influxDB.write(dbName, RET_POLICY, p);
            collector.ack(tuple);
        } catch (IOException e) {
            e.printStackTrace();
            logger.error(ExceptionUtils.getStackTrace(e));
        }

    }

    private long toLong(String nanoTimeStamp) {
        //"2017-04-24T17:36:44.657839706+05:30"
        int idx = nanoTimeStamp.indexOf('.')+3;
        int plusIdx = nanoTimeStamp.indexOf('+');
        int minusIdx = nanoTimeStamp.indexOf('-');
        int max = plusIdx>minusIdx?plusIdx:minusIdx;
        String timeStamp = nanoTimeStamp.substring(0,idx+1) +
                nanoTimeStamp.substring(max);
        System.out.println("conveting timeStamp: "+timeStamp+" to date");
        DateFormat df = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSSX");
        try {
            Date d = df.parse(timeStamp);
            return d.getTime();
        } catch (ParseException e) {
            e.printStackTrace();
            logger.error(ExceptionUtils.getStackTrace(e));
            return 0;
        }

    }
    private int toInt(String status){
        if(status.equals("RED"))
            return 0;
        else
            return 1;
    }

    public void declareOutputFields(OutputFieldsDeclarer outputFieldsDeclarer) {

    }
    @Override
    public void cleanup(){
        influxDB.close();
        super.cleanup();
    }
}
