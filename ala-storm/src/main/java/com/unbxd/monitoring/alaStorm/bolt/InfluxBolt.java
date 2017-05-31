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
    public static final String MSG_FIELD_SERVICE = "service";
    public static final String MSG_FIELD_ID = "id";
    public static final String MSG_FIELD_METRIC_NAME = "metricName";
    public static final String MSG_FIELD_METRIC_VALUE= "value";


    public static final String FIELD_SERVICE_ID = "serviceId";
    private OutputCollector collector;
    String dbName;
    InfluxDB influxDB;
    ObjectMapper jsonMapper;
    String RET_POLICY;

    public InfluxBolt() {jsonMapper = new ObjectMapper();}

    public void prepare(Map map, TopologyContext topologyContext, OutputCollector outputCollector) {
        this.collector = outputCollector;
        OkHttpClient.Builder builder = new OkHttpClient.Builder().readTimeout(10,
                TimeUnit.SECONDS).connectTimeout(10, TimeUnit.SECONDS);
        String hostPort = (String) map.get(ConfKeys.INFLUX_HOSTPORT);
        String user = (String) map.get(ConfKeys.INFLUX_USER);
        String pass = (String) map.get(ConfKeys.INFLUX_PASS);
        this.influxDB = InfluxDBFactory.connect(hostPort, user, pass, builder);
        this.dbName = (String) map.get(ConfKeys.INFLUX_DBNAME);
        RET_POLICY = (String)map.get(ConfKeys.INFLUX_RET_POLICY);
    }

    public void execute(Tuple tuple) {
        String val = tuple.getString(0);
        System.out.println("received : " + val);
        try {
            Map<String, Object> data = jsonMapper.readValue(val, Map.class);
            String nanoTS  =(String) data.get("timestamp");
            String milliTS = milliTimeStamp(nanoTS);
            long t = toLong(milliTS);
            String metricName = (String) data.get(MSG_FIELD_METRIC_NAME);
            Point p = Point.measurement(metricName)
                    .time(t, TimeUnit.MILLISECONDS)
                    .addField(FIELD_SERVICE_ID, (String) ((Map) data.get
                            (MSG_FIELD_SERVICE)).get(MSG_FIELD_ID))
                    .addField(MSG_FIELD_METRIC_NAME, metricName)
                    .addField(MSG_FIELD_METRIC_VALUE,(Number) data.get
                            (MSG_FIELD_METRIC_VALUE))
                    .build();
            influxDB.write(dbName, RET_POLICY, p);
            collector.ack(tuple);
        } catch (IOException e) {
            e.printStackTrace();
            logger.error(ExceptionUtils.getStackTrace(e));
        }

    }

    private long toLong(String milliTimeStamp) {
        DateFormat df = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSSXXX");
        try {
            Date d = df.parse(milliTimeStamp);
            return d.getTime();
        } catch (ParseException e) {
            e.printStackTrace();
            logger.error(ExceptionUtils.getStackTrace(e));
            return 0;
        }

    }
    private String milliTimeStamp(String nanoTimeStamp){
        int idx = nanoTimeStamp.indexOf('.')+3;
        int plusIdx = nanoTimeStamp.indexOf('+');
        int minusIdx = nanoTimeStamp.indexOf('-');
        int max = plusIdx>minusIdx?plusIdx:minusIdx;
        String timeStamp = nanoTimeStamp.substring(0,idx+1) +
                nanoTimeStamp.substring(max);
        return timeStamp;
    }

    public void declareOutputFields(OutputFieldsDeclarer outputFieldsDeclarer) {

    }
    @Override
    public void cleanup(){
        influxDB.close();
        super.cleanup();
    }
}
